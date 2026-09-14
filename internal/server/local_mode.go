package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/sessions"
	"pi-web/internal/workers"
)

const (
	sessionModeAuto  = "auto"
	sessionModeLocal = "local"
	sessionModeCloud = "cloud"

	localRecoveryHealthyToolResults = 3
	localRecoveryUserStopped        = -1
	localRecoveryHealthyDuration    = 5 * time.Minute
)

const sessionModesSchema = `CREATE TABLE IF NOT EXISTS session_modes (
	session_id       TEXT PRIMARY KEY,
	configured_mode  TEXT NOT NULL,
	effective_mode   TEXT NOT NULL,
	model_provider   TEXT NOT NULL DEFAULT '',
	model_id         TEXT NOT NULL DEFAULT '',
	context_window   INTEGER NOT NULL DEFAULT 0,
	last_active_at   DATETIME,
	updated_at       DATETIME NOT NULL
)`

const localRecoveryIncidentsSchema = `CREATE TABLE IF NOT EXISTS local_recovery_incidents (
	session_id   TEXT NOT NULL,
	incident_id  TEXT NOT NULL,
	status       TEXT NOT NULL,
	attempted_at DATETIME NOT NULL,
	PRIMARY KEY (session_id, incident_id)
)`

const localRecoveryStateSchema = `CREATE TABLE IF NOT EXISTS local_recovery_state (
	session_id                    TEXT PRIMARY KEY,
	automatic_attempts_since_user INTEGER NOT NULL DEFAULT 0,
	updated_at                    DATETIME NOT NULL
)`

type sessionModeState struct {
	ConfiguredMode string `json:"configuredMode"`
	EffectiveMode  string `json:"effectiveMode"`
	ModelProvider  string `json:"modelProvider,omitempty"`
	ModelID        string `json:"modelId,omitempty"`
	ContextWindow  int    `json:"contextWindow,omitempty"`
	LastActiveAt   string `json:"lastActiveAt,omitempty"`
}

type modelModeMetadata struct {
	Provider      string
	ID            string
	BaseURL       string
	ContextWindow int
}

func (s *Server) localModeNow() time.Time {
	if s.now != nil {
		return s.now()
	}
	return time.Now()
}

func normalizeSessionMode(mode string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "", sessionModeAuto:
		return sessionModeAuto, nil
	case sessionModeLocal:
		return sessionModeLocal, nil
	case sessionModeCloud:
		return sessionModeCloud, nil
	default:
		return "", errors.New("mode must be auto, local, or cloud")
	}
}

func effectiveSessionMode(configured string, model modelModeMetadata) string {
	if configured == sessionModeLocal || configured == sessionModeCloud {
		return configured
	}
	if endpointIsLocal(model.BaseURL) || providerIsLocal(model.Provider) {
		return sessionModeLocal
	}
	return sessionModeCloud
}

func providerIsLocal(provider string) bool {
	switch strings.ToLower(strings.TrimSpace(provider)) {
	case "llama", "ollama", "lmstudio", "lm-studio", "local":
		return true
	default:
		return false
	}
}

func endpointIsLocal(raw string) bool {
	if strings.TrimSpace(raw) == "" {
		return false
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return false
	}
	host := strings.TrimSuffix(strings.ToLower(parsed.Hostname()), ".")
	if host == "localhost" || strings.HasSuffix(host, ".localhost") || strings.HasSuffix(host, ".local") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && (ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast())
}

func modelMetadataFromPayload(data json.RawMessage, provider, modelID string) modelModeMetadata {
	metadata := modelModeMetadata{Provider: provider, ID: modelID}
	if len(data) == 0 || modelID == "" {
		return metadata
	}
	var payload struct {
		Models []struct {
			Provider      string `json:"provider"`
			ID            string `json:"id"`
			ModelID       string `json:"modelId"`
			BaseURL       string `json:"baseUrl"`
			ContextWindow int    `json:"contextWindow"`
		} `json:"models"`
	}
	if json.Unmarshal(data, &payload) != nil {
		return metadata
	}
	for _, model := range payload.Models {
		id := model.ID
		if id == "" {
			id = model.ModelID
		}
		if id == modelID && (provider == "" || model.Provider == provider) {
			metadata.Provider = model.Provider
			metadata.ID = id
			metadata.BaseURL = model.BaseURL
			metadata.ContextWindow = model.ContextWindow
			return metadata
		}
	}
	return metadata
}

func (s *Server) resolveModelMetadata(ctx context.Context, provider, modelID string) modelModeMetadata {
	metadata := modelModeMetadata{Provider: provider, ID: modelID}
	if s.models == nil || modelID == "" {
		return metadata
	}
	data, err := s.models(ctx)
	if err != nil {
		return metadata
	}
	return modelMetadataFromPayload(data, provider, modelID)
}

func (s *Server) saveSessionMode(sessionID, configured string, model modelModeMetadata) (sessionModeState, error) {
	configured, err := normalizeSessionMode(configured)
	if err != nil {
		return sessionModeState{}, err
	}
	effective := effectiveSessionMode(configured, model)
	state := sessionModeState{
		ConfiguredMode: configured,
		EffectiveMode:  effective,
		ModelProvider:  model.Provider,
		ModelID:        model.ID,
		ContextWindow:  model.ContextWindow,
	}
	if s.db == nil {
		return state, nil
	}
	now := s.localModeNow().UTC()
	_, err = s.db.Exec(`INSERT INTO session_modes
		(session_id, configured_mode, effective_mode, model_provider, model_id, context_window, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(session_id) DO UPDATE SET
			configured_mode = excluded.configured_mode,
			effective_mode = excluded.effective_mode,
			model_provider = excluded.model_provider,
			model_id = excluded.model_id,
			context_window = excluded.context_window,
			updated_at = excluded.updated_at`,
		sessionID, configured, effective, model.Provider, model.ID, model.ContextWindow, now)
	if err != nil {
		return sessionModeState{}, err
	}
	return state, nil
}

func (s *Server) sessionMode(sessionID, provider, modelID string) sessionModeState {
	state := sessionModeState{ConfiguredMode: sessionModeAuto, EffectiveMode: sessionModeCloud}
	if s.db == nil {
		state.EffectiveMode = effectiveSessionMode(sessionModeAuto, modelModeMetadata{Provider: provider, ID: modelID})
		state.ModelProvider = provider
		state.ModelID = modelID
		return state
	}
	var lastActive *time.Time
	err := s.db.QueryRow(`SELECT configured_mode, effective_mode, model_provider, model_id,
		context_window, last_active_at FROM session_modes WHERE session_id = ?`, sessionID).
		Scan(&state.ConfiguredMode, &state.EffectiveMode, &state.ModelProvider, &state.ModelID,
			&state.ContextWindow, &lastActive)
	if err == nil {
		if lastActive != nil {
			state.LastActiveAt = lastActive.UTC().Format(time.RFC3339Nano)
		}
		return state
	}
	model := modelModeMetadata{Provider: provider, ID: modelID}
	state.EffectiveMode = effectiveSessionMode(sessionModeAuto, model)
	state.ModelProvider = provider
	state.ModelID = modelID
	return state
}

func (s *Server) updateAutoSessionMode(ctx context.Context, sessionID, provider, modelID string) sessionModeState {
	state := s.sessionMode(sessionID, provider, modelID)
	if state.ConfiguredMode != sessionModeAuto {
		return state
	}
	model := s.resolveModelMetadata(ctx, provider, modelID)
	updated, err := s.saveSessionMode(sessionID, sessionModeAuto, model)
	if err != nil {
		return state
	}
	return updated
}

func (s *Server) refreshSessionModeModel(ctx context.Context, session sessions.Session) sessionModeState {
	state := s.sessionMode(session.ID, session.ModelProvider, session.Model)
	if session.Model == "" || (state.ModelID == session.Model && state.ModelProvider == session.ModelProvider && state.ContextWindow > 0) {
		return state
	}
	model := s.resolveModelMetadata(ctx, session.ModelProvider, session.Model)
	updated, err := s.saveSessionMode(session.ID, state.ConfiguredMode, model)
	if err != nil {
		return state
	}
	return updated
}

func shouldCompactLocalSession(entries []map[string]any, contextWindow int, request chat.Request) bool {
	if contextWindow <= 0 {
		return false
	}
	latestCompaction := -1
	for i, entry := range entries {
		if entry["type"] == "compaction" {
			latestCompaction = i
		}
	}
	contextTokens := 0
	lastUsage := -1
	for i := len(entries) - 1; i > latestCompaction; i-- {
		entry := entries[i]
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		if message["role"] != "assistant" || message["stopReason"] == "error" || message["stopReason"] == "aborted" {
			continue
		}
		usage, _ := message["usage"].(map[string]any)
		contextTokens = usageTokenCount(usage)
		if contextTokens > 0 {
			lastUsage = i
			break
		}
	}
	if lastUsage >= 0 {
		for _, entry := range entries[lastUsage+1:] {
			encoded, _ := json.Marshal(entry)
			contextTokens += len(encoded) / 4
		}
	}
	contextTokens += len(request.Message) / 4
	contextTokens += len(request.Images) * 1200
	return contextTokens*100 >= contextWindow*65
}

func usageTokenCount(usage map[string]any) int {
	if usage == nil {
		return 0
	}
	if total := numberAsInt(usage["totalTokens"]); total > 0 {
		return total
	}
	return numberAsInt(usage["input"]) + numberAsInt(usage["output"]) +
		numberAsInt(usage["cacheRead"]) + numberAsInt(usage["cacheWrite"])
}

func numberAsInt(value any) int {
	switch n := value.(type) {
	case float64:
		return int(n)
	case int:
		return n
	case json.Number:
		v, _ := n.Int64()
		return int(v)
	default:
		return 0
	}
}

func (s *Server) markLocalSessionActive(sessionID string) {
	if sessionID == "" || s.db == nil {
		return
	}
	now := s.localModeNow().UTC()
	_, _ = s.db.Exec(`UPDATE session_modes SET last_active_at = ?, updated_at = ?
		WHERE session_id = ? AND effective_mode = ?`, now, now, sessionID, sessionModeLocal)
}

func (s *Server) resetLocalRecoveryBudget(sessionID string) {
	if sessionID == "" || s.db == nil {
		return
	}
	now := s.localModeNow().UTC()
	_, _ = s.db.Exec(`INSERT INTO local_recovery_state
		(session_id, automatic_attempts_since_user, updated_at) VALUES (?, 0, ?)
		ON CONFLICT(session_id) DO UPDATE SET automatic_attempts_since_user = 0, updated_at = excluded.updated_at`,
		sessionID, now)
}

func contextOverflowIncident(entries []map[string]any, contextWindow int) (string, bool) {
	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		if message["role"] == "user" {
			return "", false
		}
		if message["role"] != "assistant" {
			continue
		}
		stopReason, _ := message["stopReason"].(string)
		errorMessage, _ := message["errorMessage"].(string)
		if stopReason != "error" && stopReason != "length" {
			return "", false
		}
		// A plain length stop can also mean the model exhausted its output
		// allowance. It is not context-overflow evidence on its own.
		if !isContextFailure(errorMessage) {
			return "", false
		}
		if latestContextPercent(entries[:i+1], contextWindow) < 99 {
			return "", false
		}
		id, _ := entry["id"].(string)
		if id == "" {
			id = fmt.Sprintf("%v:%s", entry["timestamp"], errorMessage)
		}
		return id, true
	}
	return "", false
}

func transportInterruptionIncident(entries []map[string]any) (string, bool) {
	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		if message["role"] == "user" {
			return "", false
		}
		if message["role"] != "assistant" {
			continue
		}
		if message["stopReason"] != "error" {
			return "", false
		}
		errorMessage, _ := message["errorMessage"].(string)
		switch strings.ToLower(strings.TrimSpace(errorMessage)) {
		case "this operation was aborted", "the operation was aborted", "the operation was aborted.", "request aborted", "request was aborted":
			id, _ := entry["id"].(string)
			if id == "" {
				id = fmt.Sprintf("%v:%s", entry["timestamp"], errorMessage)
			}
			return id, true
		default:
			return "", false
		}
	}
	return "", false
}

func (s *Server) stopLocalRecovery(sessionID string) {
	s.localRecoveryMu.Lock()
	defer s.localRecoveryMu.Unlock()
	if s.localRecoverySession == sessionID && s.localRecoveryCancel != nil {
		s.localRecoveryCancel()
	}
	if s.db != nil {
		_, _ = s.db.Exec(`INSERT INTO local_recovery_state
			(session_id, automatic_attempts_since_user, updated_at) VALUES (?, ?, ?)
			ON CONFLICT(session_id) DO UPDATE SET automatic_attempts_since_user = excluded.automatic_attempts_since_user, updated_at = excluded.updated_at`,
			sessionID, localRecoveryUserStopped, s.localModeNow().UTC())
	}
}

func thinkingOnlyStopIncident(entries []map[string]any) (string, bool) {
	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		if message["role"] == "user" {
			return "", false
		}
		if message["role"] != "assistant" {
			continue
		}
		if message["stopReason"] != "stop" {
			return "", false
		}
		content, _ := message["content"].([]any)
		hasThinking := false
		for _, rawBlock := range content {
			block, _ := rawBlock.(map[string]any)
			switch block["type"] {
			case "thinking":
				thinking, _ := block["thinking"].(string)
				if strings.TrimSpace(thinking) != "" {
					hasThinking = true
				}
			case "text":
				text, _ := block["text"].(string)
				if strings.TrimSpace(text) != "" {
					return "", false
				}
			default:
				return "", false
			}
		}
		if !hasThinking {
			return "", false
		}
		id, _ := entry["id"].(string)
		if id == "" {
			id = fmt.Sprintf("%v:thinking-only-stop", entry["timestamp"])
		}
		return id, true
	}
	return "", false
}

func localRecoveryMadeHealthyProgress(entries []map[string]any, previousIncidentID, currentIncidentID string) bool {
	previousIndex := -1
	currentIndex := len(entries)
	for i, entry := range entries {
		switch {
		case entryMatchesRecoveryIncident(entry, previousIncidentID):
			previousIndex = i
		case entryMatchesRecoveryIncident(entry, currentIncidentID):
			currentIndex = i
		}
	}
	if previousIndex < 0 || currentIndex >= len(entries) || currentIndex <= previousIndex {
		return false
	}

	toolResults := 0
	for _, entry := range entries[previousIndex+1 : currentIndex] {
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		switch message["role"] {
		case "toolResult":
			if message["isError"] != true {
				toolResults++
			}
		case "assistant":
			if message["stopReason"] != "stop" {
				continue
			}
			content, _ := message["content"].([]any)
			for _, rawBlock := range content {
				block, _ := rawBlock.(map[string]any)
				if block["type"] == "text" {
					text, _ := block["text"].(string)
					if strings.TrimSpace(text) != "" {
						return true
					}
				}
			}
		}
	}
	if toolResults >= localRecoveryHealthyToolResults {
		return true
	}
	if toolResults == 0 {
		return false
	}
	previousTime, previousOK := entryTimestamp(entries[previousIndex])
	currentTime, currentOK := entryTimestamp(entries[currentIndex])
	return previousOK && currentOK && currentTime.Sub(previousTime) >= localRecoveryHealthyDuration
}

func entryMatchesRecoveryIncident(entry map[string]any, incidentID string) bool {
	if id, _ := entry["id"].(string); id != "" {
		return id == incidentID
	}
	message, _ := entry["message"].(map[string]any)
	errorMessage, _ := message["errorMessage"].(string)
	if errorMessage != "" && fmt.Sprintf("%v:%s", entry["timestamp"], errorMessage) == incidentID {
		return true
	}
	return fmt.Sprintf("%v:thinking-only-stop", entry["timestamp"]) == incidentID
}

func entryTimestamp(entry map[string]any) (time.Time, bool) {
	raw, _ := entry["timestamp"].(string)
	parsed, err := time.Parse(time.RFC3339Nano, raw)
	return parsed, err == nil
}

func latestContextPercent(entries []map[string]any, contextWindow int) int {
	if contextWindow <= 0 {
		return 0
	}
	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		if entry["type"] != "message" {
			continue
		}
		message, _ := entry["message"].(map[string]any)
		if message["role"] != "assistant" {
			continue
		}
		usage, _ := message["usage"].(map[string]any)
		if tokens := usageTokenCount(usage); tokens > 0 {
			return tokens * 100 / contextWindow
		}
	}
	return 0
}

func isContextFailure(message string) bool {
	message = strings.ToLower(message)
	for _, pattern := range []string{
		"context overflow",
		"context length",
		"context window",
		"context_length_exceeded",
		"maximum context",
		"max context",
		"prompt too long",
		"prompt is too long",
		"too many tokens",
		"token limit",
	} {
		if strings.Contains(message, pattern) {
			return true
		}
	}
	return false
}

func (s *Server) claimLocalRecovery(sessionID, incidentID string, entries []map[string]any) bool {
	tx, err := s.db.Begin()
	if err != nil {
		return false
	}
	defer tx.Rollback()
	now := s.localModeNow().UTC()
	var automaticAttemptsWithoutProgress int
	err = tx.QueryRow(`SELECT automatic_attempts_since_user FROM local_recovery_state
		WHERE session_id = ?`, sessionID).Scan(&automaticAttemptsWithoutProgress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false
	}
	if automaticAttemptsWithoutProgress == localRecoveryUserStopped {
		return false
	}
	if automaticAttemptsWithoutProgress > 0 {
		var previousIncidentID string
		if err := tx.QueryRow(`SELECT incident_id FROM local_recovery_incidents
			WHERE session_id = ? ORDER BY attempted_at DESC LIMIT 1`, sessionID).Scan(&previousIncidentID); err != nil {
			return false
		}
		if !localRecoveryMadeHealthyProgress(entries, previousIncidentID, incidentID) {
			return false
		}
	}
	result, err := tx.Exec(`INSERT OR IGNORE INTO local_recovery_incidents
		(session_id, incident_id, status, attempted_at) VALUES (?, ?, 'started', ?)`, sessionID, incidentID, now)
	if err != nil {
		return false
	}
	inserted, _ := result.RowsAffected()
	if inserted != 1 {
		return false
	}
	_, err = tx.Exec(`INSERT OR IGNORE INTO local_recovery_state
		(session_id, automatic_attempts_since_user, updated_at) VALUES (?, 0, ?)`, sessionID, now)
	if err != nil {
		return false
	}
	result, err = tx.Exec(`UPDATE local_recovery_state SET automatic_attempts_since_user = 1, updated_at = ?
		WHERE session_id = ?`, now, sessionID)
	if err != nil {
		return false
	}
	claimed, _ := result.RowsAffected()
	if claimed != 1 {
		return false
	}
	return tx.Commit() == nil
}

func (s *Server) finishLocalRecovery(sessionID, incidentID, status string) {
	_, _ = s.db.Exec(`UPDATE local_recovery_incidents SET status = ?
		WHERE session_id = ? AND incident_id = ?`, status, sessionID, incidentID)
}

func (s *Server) maybeStartLocalRecovery(sessionID string) bool {
	if s.disableBackgroundJobs || s.chatSender == nil || sessionID == "" {
		return false
	}
	if s.chatSender.Status(sessionID).State == workers.WorkerStateRunning {
		return false
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, sessionID)
	if err != nil {
		return false
	}
	mode := s.refreshSessionModeModel(context.Background(), resolved.Session)
	if mode.EffectiveMode != sessionModeLocal {
		return false
	}
	incidentID, contextOverflow := contextOverflowIncident(resolved.Session.Entries, mode.ContextWindow)
	if !contextOverflow {
		incidentID, _ = thinkingOnlyStopIncident(resolved.Session.Entries)
	}
	transportInterruption := false
	if incidentID == "" {
		incidentID, transportInterruption = transportInterruptionIncident(resolved.Session.Entries)
	}
	if incidentID == "" {
		return false
	}
	s.localRecoveryMu.Lock()
	if s.localRecoveryRunning {
		s.localRecoveryMu.Unlock()
		return false
	}
	if !s.claimLocalRecovery(sessionID, incidentID, resolved.Session.Entries) {
		s.localRecoveryMu.Unlock()
		return false
	}
	s.localRecoveryRunning = true
	recoveryCtx, cancelRecovery := context.WithCancel(context.Background())
	s.localRecoverySession = sessionID
	s.localRecoveryCancel = cancelRecovery
	s.localRecoveryMu.Unlock()
	if !s.startTask(func(ctx context.Context) {
		stopCancel := context.AfterFunc(ctx, cancelRecovery)
		defer stopCancel()
		ctx = recoveryCtx
		settled := false
		defer func() {
			checkNextIncident := settled && ctx.Err() == nil
			cancelRecovery()
			s.localRecoveryMu.Lock()
			s.localRecoveryRunning = false
			s.localRecoverySession = ""
			s.localRecoveryCancel = nil
			s.localRecoveryMu.Unlock()
			if checkNextIncident {
				s.maybeStartLocalRecovery(sessionID)
			}
		}()
		release, err := s.acquireSessionOperation(ctx, sessionID)
		if err != nil {
			s.finishLocalRecovery(sessionID, incidentID, "operation_cancelled")
			return
		}
		defer release()
		resolved, err = sessions.ResolveByID(s.sessionsDir, sessionID)
		if err != nil {
			s.finishLocalRecovery(sessionID, incidentID, "session_unavailable")
			return
		}
		if transportInterruption {
			currentIncident, _ := transportInterruptionIncident(resolved.Session.Entries)
			if currentIncident != incidentID || s.chatSender.Status(sessionID).State == workers.WorkerStateRunning {
				s.finishLocalRecovery(sessionID, incidentID, "superseded")
				return
			}
		}
		if err := s.prepareWorkerMode(sessionID, mode); err != nil {
			s.finishLocalRecovery(sessionID, incidentID, "prepare_failed")
			return
		}
		if contextOverflow || (transportInterruption && shouldCompactLocalSession(resolved.Session.Entries, mode.ContextWindow, chat.Request{})) {
			if err := s.forceCompactSessionLocked(ctx, sessionID, nil); err != nil {
				s.finishLocalRecovery(sessionID, incidentID, "compaction_failed")
				return
			}
			s.broadcast(sessionID, "reload")
		} else if !transportInterruption {
			if msg, err := formatSSEJSONEvent("local-recovery", map[string]any{
				"reason": "thinking-only-stop",
			}); err == nil {
				s.broadcast(sessionID, msg)
			}
		}
		if ctx.Err() != nil {
			s.finishLocalRecovery(sessionID, incidentID, "operation_cancelled")
			return
		}
		if err := s.chatSender.Send(ctx, sessionID, resolved.Path, chat.Request{Message: "continue if possible"}); err != nil {
			s.finishLocalRecovery(sessionID, incidentID, "continue_failed")
			return
		}
		waiter, ok := s.chatSender.(settledWaitingSender)
		if !ok {
			s.finishLocalRecovery(sessionID, incidentID, "settled_wait_unsupported")
			return
		}
		if err := waiter.WaitSettled(ctx, sessionID); err != nil {
			s.finishLocalRecovery(sessionID, incidentID, "continue_unsettled")
			return
		}
		s.finishLocalRecovery(sessionID, incidentID, "continued")
		settled = true
	}) {
		cancelRecovery()
		s.localRecoveryMu.Lock()
		s.localRecoveryRunning = false
		s.localRecoverySession = ""
		s.localRecoveryCancel = nil
		s.localRecoveryMu.Unlock()
		s.finishLocalRecovery(sessionID, incidentID, "server_stopping")
		return false
	}
	return true
}

func (s *Server) startLocalStartupRecovery() {
	if s.disableBackgroundJobs {
		return
	}
	var sessionID string
	err := s.db.QueryRow(`SELECT session_id FROM session_modes
		WHERE effective_mode = ? AND last_active_at IS NOT NULL
		ORDER BY last_active_at DESC LIMIT 1`, sessionModeLocal).Scan(&sessionID)
	if err == nil {
		s.maybeStartLocalRecovery(sessionID)
	}
}

func (s *Server) handleSessionMode(w http.ResponseWriter, r *http.Request) {
	resolved, err := sessions.ResolveByID(s.sessionsDir, r.URL.Query().Get("id"))
	if resolveOrWriteError(w, err) {
		return
	}
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 0, s.refreshSessionModeModel(r.Context(), resolved.Session))
	case http.MethodPost:
		var body struct {
			Mode string `json:"mode"`
		}
		if !decodeJSONBody(w, r, &body) {
			return
		}
		configured, err := normalizeSessionMode(body.Mode)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		current := s.sessionMode(resolved.Session.ID, resolved.Session.ModelProvider, resolved.Session.Model)
		if configured == current.ConfiguredMode {
			writeJSON(w, 0, current)
			return
		}
		if s.chatSender != nil && s.chatSender.Status(resolved.Session.ID).State == workers.WorkerStateRunning {
			writeJSONError(w, http.StatusConflict, "cannot change session mode while the session is running")
			return
		}
		model := s.resolveModelMetadata(r.Context(), resolved.Session.ModelProvider, resolved.Session.Model)
		state, err := s.saveSessionMode(resolved.Session.ID, configured, model)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, 0, state)
	default:
		w.Header().Set("Allow", "GET, POST")
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
