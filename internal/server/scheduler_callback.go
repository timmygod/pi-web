package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"pi-web/internal/schedules"
	"pi-web/internal/sessions"
)

type scheduleCallbackPayload struct {
	Version  string `json:"version"`
	Event    string `json:"event"`
	EventID  string `json:"eventId"`
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	Run schedules.Run `json:"run"`
}

func (s *Server) completeScheduleRun(sessionID string) {
	run, sc, err := s.schedules.RunForSession(sessionID)
	if err != nil || run.Status != schedules.RunStatusRunning {
		return
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, run.SessionFile)
	if err != nil {
		_ = s.schedules.CompleteRun(run.ID, schedules.RunStatusFailed, "", err.Error(), run.ModelProvider, run.ModelID)
		s.deliverScheduleCallback(context.Background(), sc, run.ID)
		return
	}
	status, result, errMessage := scheduleResult(resolved.Session.Entries)
	_ = s.schedules.CompleteRun(run.ID, status, result, errMessage, run.ModelProvider, run.ModelID)
	s.deliverScheduleCallback(context.Background(), sc, run.ID)
}

func scheduleResult(entries []map[string]any) (string, string, string) {
	for i := len(entries) - 1; i >= 0; i-- {
		if entries[i]["type"] != "message" {
			continue
		}
		message, _ := entries[i]["message"].(map[string]any)
		if message["role"] != "assistant" {
			continue
		}
		stopReason, _ := message["stopReason"].(string)
		errMessage, _ := message["errorMessage"].(string)
		result := callbackMessageText(message["content"])
		switch stopReason {
		case "aborted":
			return schedules.RunStatusCancelled, result, errMessage
		case "error":
			return schedules.RunStatusFailed, result, errMessage
		default:
			return schedules.RunStatusSucceeded, result, ""
		}
	}
	return schedules.RunStatusFailed, "", "scheduled session ended without an assistant response"
}

func callbackMessageText(content any) string {
	if text, ok := content.(string); ok {
		return text
	}
	parts, _ := content.([]any)
	var out []string
	for _, part := range parts {
		item, _ := part.(map[string]any)
		if item["type"] != "text" {
			continue
		}
		if text, _ := item["text"].(string); text != "" {
			out = append(out, text)
		}
	}
	return strings.Join(out, "\n")
}

func (s *Server) deliverScheduleCallback(ctx context.Context, sc schedules.Schedule, runID int64) {
	if sc.CallbackURL == "" {
		return
	}
	run, err := s.schedules.GetRun(runID)
	if err != nil {
		return
	}
	payload := scheduleCallbackPayload{Version: "1", Event: "schedule.run.completed", EventID: fmt.Sprintf("schedule-run-%d", run.ID), Run: run}
	payload.Schedule.ID, payload.Schedule.Name = sc.ID, sc.Name
	body, _ := json.Marshal(payload)
	client := &http.Client{Timeout: 10 * time.Second}
	if s.callbackHTTPClient != nil {
		client = s.callbackHTTPClient
	}
	var lastErr error
	for attempt := 1; attempt <= 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, sc.CallbackURL, bytes.NewReader(body))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "pi-web-scheduler/1")
			req.Header.Set("X-Pi-Web-Event-ID", payload.EventID)
			resp, doErr := client.Do(req)
			if doErr == nil {
				_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 64<<10))
				resp.Body.Close()
				if resp.StatusCode >= 200 && resp.StatusCode < 300 {
					_ = s.schedules.SetCallbackDelivery(runID, "delivered", "", attempt)
					return
				}
				lastErr = fmt.Errorf("callback returned HTTP %d", resp.StatusCode)
			} else {
				lastErr = doErr
			}
		} else {
			lastErr = err
		}
		if attempt < 3 {
			select {
			case <-ctx.Done():
				lastErr = ctx.Err()
				attempt = 3
			case <-time.After(time.Duration(attempt) * time.Second):
			}
		}
	}
	_ = s.schedules.SetCallbackDelivery(runID, "failed", lastErr.Error(), 3)
}
