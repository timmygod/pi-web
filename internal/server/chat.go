package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/sessions"
	"pi-web/internal/workers"
)

type ChatSender interface {
	Send(ctx context.Context, sessionID, sessionPath string, chat chat.Request) error
	SetModel(ctx context.Context, sessionID, sessionPath, provider, modelID string) error
	SetThinkingLevel(ctx context.Context, sessionID, sessionPath, level string) error
	Abort(ctx context.Context, sessionID string) error
	GetState(ctx context.Context, sessionID string) (workers.WorkerStatus, error)
	GetCommands(ctx context.Context, sessionID string) ([]workers.SlashCommand, bool, error)
	Status(sessionID string) workers.WorkerStatus
	EnsureWorker(ctx context.Context, sessionID, sessionPath string) error
}

type compactSender interface {
	Compact(ctx context.Context, sessionID, sessionPath string) error
}

type modePreparingSender interface {
	PrepareMode(sessionID string, config workers.WorkerConfig) error
}

type settledWaitingSender interface {
	WaitSettled(ctx context.Context, sessionID string) error
}

type sessionOperationState struct {
	mu      sync.Mutex
	entries map[string]*sessionOperation
}

type sessionOperation struct {
	token chan struct{}
	refs  int
}

func (s *Server) acquireSessionOperation(ctx context.Context, sessionID string) (func(), error) {
	s.sessionOperations.mu.Lock()
	if s.sessionOperations.entries == nil {
		s.sessionOperations.entries = make(map[string]*sessionOperation)
	}
	operation := s.sessionOperations.entries[sessionID]
	if operation == nil {
		operation = &sessionOperation{token: make(chan struct{}, 1)}
		s.sessionOperations.entries[sessionID] = operation
	}
	operation.refs++
	s.sessionOperations.mu.Unlock()

	select {
	case operation.token <- struct{}{}:
		return func() {
			<-operation.token
			s.releaseSessionOperation(sessionID, operation)
		}, nil
	case <-ctx.Done():
		s.releaseSessionOperation(sessionID, operation)
		return nil, ctx.Err()
	}
}

func (s *Server) releaseSessionOperation(sessionID string, operation *sessionOperation) {
	s.sessionOperations.mu.Lock()
	operation.refs--
	if operation.refs == 0 && s.sessionOperations.entries[sessionID] == operation {
		delete(s.sessionOperations.entries, sessionID)
	}
	s.sessionOperations.mu.Unlock()
}

func (s *Server) forceCompact(ctx context.Context, sessionID string, mode *sessionModeState) error {
	release, err := s.acquireSessionOperation(ctx, sessionID)
	if err != nil {
		return err
	}
	defer release()

	return s.forceCompactSessionLocked(ctx, sessionID, mode)
}

func (s *Server) forceCompactSessionLocked(ctx context.Context, sessionID string, mode *sessionModeState) error {
	resolved, err := sessions.ResolveByID(s.sessionsDir, sessionID)
	if err != nil {
		return err
	}
	if len(resolved.Session.Entries) > 0 && resolved.Session.Entries[len(resolved.Session.Entries)-1]["type"] == "compaction" {
		return nil
	}
	return s.forceCompactLocked(ctx, sessionID, resolved.Path, mode)
}

func (s *Server) forceCompactLocked(ctx context.Context, sessionID, sessionPath string, mode *sessionModeState) error {
	sender, ok := s.chatSender.(compactSender)
	if !ok {
		return errors.New("installed Pi worker does not support compaction")
	}
	if s.chatSender.Status(sessionID).State == workers.WorkerStateRunning {
		if err := s.chatSender.Abort(ctx, sessionID); err != nil {
			return fmt.Errorf("interrupt running session before compaction: %w", err)
		}
	}
	if mode != nil {
		if err := s.prepareWorkerMode(sessionID, *mode); err != nil {
			return err
		}
	}
	return sender.Compact(ctx, sessionID, sessionPath)
}

func (s *Server) prepareWorkerMode(sessionID string, mode sessionModeState) error {
	sender, ok := s.chatSender.(modePreparingSender)
	if !ok {
		return nil
	}
	config := workers.WorkerConfig{}
	if mode.EffectiveMode == sessionModeLocal {
		if mode.ContextWindow <= 0 {
			return errors.New("Local Mode requires a model with a known context window")
		}
		config.LocalContextWindow = mode.ContextWindow
	}
	return sender.PrepareMode(sessionID, config)
}

func (s *Server) sendSessionChat(ctx context.Context, resolved sessions.ResolvedSession, request chat.Request) error {
	sessionID := resolved.Session.ID
	release, err := s.acquireSessionOperation(ctx, sessionID)
	if err != nil {
		return err
	}
	defer release()

	resolved, err = sessions.ResolveByID(s.sessionsDir, sessionID)
	if err != nil {
		return err
	}
	mode := s.refreshSessionModeModel(ctx, resolved.Session)
	if mode.EffectiveMode == sessionModeLocal {
		s.markLocalSessionActive(sessionID)
	}
	if err := s.prepareWorkerMode(sessionID, mode); err != nil {
		return fmt.Errorf("prepare worker mode: %w", err)
	}
	if mode.EffectiveMode == sessionModeLocal &&
		s.chatSender.Status(sessionID).State != workers.WorkerStateRunning &&
		shouldCompactLocalSession(resolved.Session.Entries, mode.ContextWindow, request) {
		if err := s.forceCompactLocked(ctx, sessionID, resolved.Path, nil); err != nil {
			return fmt.Errorf("proactive Local Mode compaction: %w", err)
		}
		s.broadcast(sessionID, "reload")
	}
	return s.chatSender.Send(ctx, sessionID, resolved.Path, request)
}

func (s *Server) handleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	if !resolved.Session.ChatAvailable {
		writeJSONError(w, http.StatusConflict, resolved.Session.ChatDisabledReason)
		return
	}
	chatReq, err := chat.ParseRequest(r, chat.DefaultMaxImageBytes, chat.DefaultMaxRequestBytes)
	if err != nil {
		switch {
		case errors.Is(err, chat.ErrEmptyRequest):
			writeJSONError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, chat.ErrImageTooLarge):
			writeJSONError(w, http.StatusRequestEntityTooLarge, err.Error())
		case errors.Is(err, chat.ErrUnsupportedImageType):
			writeJSONError(w, http.StatusUnsupportedMediaType, err.Error())
		case errors.As(err, new(*http.MaxBytesError)):
			writeJSONError(w, http.StatusRequestEntityTooLarge, err.Error())
		default:
			writeJSONError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	if s.chatSender == nil {
		writeJSONError(w, http.StatusServiceUnavailable, "chat unavailable")
		return
	}
	sessionID := resolved.Session.ID
	s.resetLocalRecoveryBudget(sessionID)
	if !s.startTask(func(ctx context.Context) {
		if err := s.sendSessionChat(ctx, resolved, chatReq); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Fprintf(os.Stderr, "chat send failed for %s: %v\n", sessionID, err)
		}
	}) {
		writeJSONError(w, http.StatusServiceUnavailable, "server is shutting down")
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]any{"ok": true, "status": "queued"})
}

func (s *Server) handleForceCompact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	if s.chatSender == nil {
		writeJSONError(w, http.StatusServiceUnavailable, "chat unavailable")
		return
	}
	mode := s.refreshSessionModeModel(r.Context(), resolved.Session)
	if mode.EffectiveMode != sessionModeLocal {
		writeJSONError(w, http.StatusConflict, "force compact is available only in Local Mode")
		return
	}
	if err := s.forceCompact(r.Context(), resolved.Session.ID, &mode); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	s.broadcast(resolved.Session.ID, "reload")
	writeJSON(w, 0, map[string]any{"ok": true})
}

// recentSessionActivityWindow is the grace period after a JSONL write during
// which a session is still reported as "running" even when no in-process
// chat worker and no session-status file claims it. Kept short so the
// "running" status / Cancel button doesn't linger after the assistant
// finishes streaming its final message.
const recentSessionActivityWindow = 800 * time.Millisecond
const sessionStatusTTL = 10 * time.Second

type sessionStatusFile struct {
	State     string `json:"state"`
	UpdatedAt string `json:"updatedAt"`
}

func (s *Server) readSessionStatus(sessionID string) *workers.WorkerStatus {
	if sessionID == "" {
		return nil
	}
	path := filepath.Join(s.sessionStatusDir(), sessionID)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var status sessionStatusFile
	if err := json.Unmarshal(data, &status); err != nil {
		return nil
	}
	if status.State != "running" {
		return nil
	}
	updatedAt, err := time.Parse(time.RFC3339, status.UpdatedAt)
	if err != nil {
		return nil
	}
	if time.Since(updatedAt) > sessionStatusTTL {
		return nil
	}
	return &workers.WorkerStatus{State: workers.WorkerStateRunning}
}

func (s *Server) handleCancelChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	if s.chatSender == nil {
		writeJSONError(w, http.StatusServiceUnavailable, "chat unavailable")
		return
	}
	s.stopLocalRecovery(resolved.Session.ID)
	if err := s.chatSender.Abort(r.Context(), resolved.Session.ID); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	_ = os.Remove(filepath.Join(s.sessionStatusDir(), resolved.Session.ID))
	s.recomputeAndBroadcastStatus(resolved.Session.ID)
	s.broadcast(resolved.Session.ID, "reload")
	writeJSON(w, 0, map[string]any{"ok": true, "status": "cancelled"})
}

func (s *Server) handleWorkerStatus(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("id")

	status := workers.WorkerStatus{State: workers.WorkerStateIdle}
	if s.computeRunningStatus(sessionID) {
		status.State = workers.WorkerStateRunning
	} else if s.chatSender != nil {
		// Do not create/prewarm workers from status polling. A browser can poll
		// many visible sessions at once; if one pi RPC switch_session hangs, eager
		// prewarming accumulates stuck `pi --mode rpc` processes and starves real
		// chat requests. Only report state for an already-created worker here.
		stateCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if state, err := s.chatSender.GetState(stateCtx, sessionID); err == nil {
			status.Model = state.Model
			status.ModelName = state.ModelName
			status.ModelProvider = state.ModelProvider
			status.ThinkingLevel = state.ThinkingLevel
		}
	}
	writeJSON(w, 0, status)
}

// handleCommands serves the slash-command palette for a session's composer.
// By default it peeks at an existing worker and never spawns one; with
// ?load=1 it ensures a worker first (used when the user opens the palette and
// no worker exists yet). Any failure to query commands degrades to an empty
// list rather than an error — the palette is a non-critical affordance and
// must never break the composer.
func (s *Server) handleCommands(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	if s.chatSender == nil {
		writeJSONError(w, http.StatusServiceUnavailable, "chat unavailable")
		return
	}
	sessionID := resolved.Session.ID
	if r.URL.Query().Get("load") == "1" {
		if err := s.chatSender.EnsureWorker(r.Context(), sessionID, resolved.Path); err != nil {
			fmt.Fprintf(os.Stderr, "commands: ensure worker failed for %s: %v\n", sessionID, err)
		}
	}
	cmds, ready, err := s.chatSender.GetCommands(r.Context(), sessionID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "commands: query failed for %s: %v\n", sessionID, err)
		cmds = nil
	}
	if cmds == nil {
		cmds = []workers.SlashCommand{}
	}
	writeJSON(w, 0, map[string]any{"commands": cmds, "workerReady": ready})
}

func (s *Server) hasRecentSessionActivity(sessionID string) bool {
	if sessionID == "" {
		return false
	}
	now := s.now()
	s.fileModMu.RLock()
	mod, ok := s.fileMod[sessionID]
	s.fileModMu.RUnlock()
	if !ok {
		return false
	}
	return !mod.IsZero() && now.Sub(mod) <= recentSessionActivityWindow
}

func (s *Server) handleSetModel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	var body struct {
		Provider string `json:"provider"`
		ModelID  string `json:"modelId"`
	}
	if !decodeJSONBody(w, r, &body) {
		return
	}
	if body.Provider == "" || body.ModelID == "" {
		writeJSONError(w, http.StatusBadRequest, "provider and modelId required")
		return
	}
	if err := s.chatSender.SetModel(r.Context(), resolved.Session.ID, resolved.Path, body.Provider, body.ModelID); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	s.updateAutoSessionMode(r.Context(), resolved.Session.ID, body.Provider, body.ModelID)
	writeJSON(w, 0, map[string]any{"ok": true})
}

func (s *Server) handleSetThinkingLevel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	var body struct {
		Level string `json:"level"`
	}
	if !decodeJSONBody(w, r, &body) {
		return
	}
	if body.Level == "" {
		writeJSONError(w, http.StatusBadRequest, "level required")
		return
	}
	if err := s.chatSender.SetThinkingLevel(r.Context(), resolved.Session.ID, resolved.Path, body.Level); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	status := s.chatSender.Status(resolved.Session.ID)
	if state, err := s.chatSender.GetState(r.Context(), resolved.Session.ID); err == nil {
		status.ThinkingLevel = state.ThinkingLevel
	}
	writeJSON(w, 0, map[string]any{"ok": true, "thinkingLevel": status.ThinkingLevel})
}
