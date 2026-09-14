package server

import (
	"context"
	"time"

	"pi-web/internal/sessions"
)

func (s *Server) initialSettingsFromSource(ctx context.Context, sourceSessionID string) sessions.InitialSettings {
	if s.chatSender == nil || sourceSessionID == "" {
		return sessions.InitialSettings{}
	}
	if _, err := sessions.ResolveByID(s.sessionsDir, sourceSessionID); err != nil {
		return sessions.InitialSettings{}
	}
	state, err := s.chatSender.GetState(ctx, sourceSessionID)
	if err != nil {
		return sessions.InitialSettings{}
	}
	return sessions.InitialSettings{
		ModelProvider: state.ModelProvider,
		ModelID:       state.Model,
		ThinkingLevel: state.ThinkingLevel,
	}
}

func (s *Server) initializeNewSessionWorker(ctx context.Context, sessionID, sessionPath string, settings sessions.InitialSettings) {
	if s.chatSender == nil {
		return
	}
	// The settings have already been written into the new session file as
	// implicit entries. Creating/switching the worker should pick them up from
	// the session history. Do not call SetModel/SetThinkingLevel here: those RPC
	// calls append visible "Switched to model" entries and duplicate the implicit
	// initial settings.
	workerCtx, cancel := context.WithTimeout(ctx, 35*time.Second)
	defer cancel()
	mode := s.sessionMode(sessionID, settings.ModelProvider, settings.ModelID)
	if mode.ContextWindow <= 0 {
		if err := s.chatSender.EnsureWorker(workerCtx, sessionID, sessionPath); err != nil {
			return
		}
		state, err := s.chatSender.GetState(workerCtx, sessionID)
		if err != nil || state.Model == "" {
			return
		}
		metadata := s.resolveModelMetadata(workerCtx, state.ModelProvider, state.Model)
		if updated, err := s.saveSessionMode(sessionID, mode.ConfiguredMode, metadata); err == nil {
			mode = updated
		}
	}
	if err := s.prepareWorkerMode(sessionID, mode); err != nil {
		return
	}
	_ = s.chatSender.EnsureWorker(workerCtx, sessionID, sessionPath)
}
