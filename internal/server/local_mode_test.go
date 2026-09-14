package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/workers"
)

func TestNormalizeSessionMode(t *testing.T) {
	tests := []struct {
		input string
		want  string
		ok    bool
	}{
		{"", sessionModeAuto, true},
		{"AUTO", sessionModeAuto, true},
		{" local ", sessionModeLocal, true},
		{"cloud", sessionModeCloud, true},
		{"hybrid", "", false},
	}
	for _, tt := range tests {
		got, err := normalizeSessionMode(tt.input)
		if (err == nil) != tt.ok || got != tt.want {
			t.Fatalf("normalizeSessionMode(%q) = %q, %v; want %q, ok=%v", tt.input, got, err, tt.want, tt.ok)
		}
	}
}

func TestShouldCompactLocalSessionAtSixtyFivePercent(t *testing.T) {
	entries := []map[string]any{
		{
			"type": "message",
			"message": map[string]any{
				"role":  "assistant",
				"usage": map[string]any{"totalTokens": float64(65000)},
			},
		},
	}
	if !shouldCompactLocalSession(entries, 100000, chat.Request{}) {
		t.Fatal("expected compaction at 65%")
	}
	entries[0]["message"].(map[string]any)["usage"] = map[string]any{"totalTokens": float64(64900)}
	if shouldCompactLocalSession(entries, 100000, chat.Request{}) {
		t.Fatal("did not expect compaction below 65%")
	}
}

func TestShouldCompactLocalSessionIgnoresPreCompactionUsage(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "message": map[string]any{
			"role": "assistant", "usage": map[string]any{"totalTokens": float64(99000)},
		}},
		{"type": "compaction"},
	}
	if shouldCompactLocalSession(entries, 100000, chat.Request{}) {
		t.Fatal("pre-compaction usage must not trigger another compaction")
	}
}

func TestContextOverflowIncidentRequiresRecentFailureAndFullContext(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "message": map[string]any{
			"role": "assistant", "usage": map[string]any{"totalTokens": float64(99500)},
		}},
		{"type": "message", "id": "failure-1", "message": map[string]any{
			"role": "assistant", "stopReason": "error", "errorMessage": "maximum context length exceeded",
		}},
	}
	id, ok := contextOverflowIncident(entries, 100000)
	if !ok || id != "failure-1" {
		t.Fatalf("incident = %q, %v", id, ok)
	}
	entries[0]["message"].(map[string]any)["usage"] = map[string]any{"totalTokens": float64(98000)}
	if _, ok := contextOverflowIncident(entries, 100000); ok {
		t.Fatal("98% context should not satisfy watchdog threshold")
	}
}

func TestContextOverflowIncidentDoesNotTreatPlainLengthStopAsOverflow(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "message": map[string]any{
			"role": "assistant", "usage": map[string]any{"totalTokens": float64(100000)},
		}},
		{"type": "message", "id": "length-1", "message": map[string]any{
			"role": "assistant", "stopReason": "length",
		}},
	}
	if _, ok := contextOverflowIncident(entries, 100000); ok {
		t.Fatal("plain output-length stop must not be treated as context-overflow evidence")
	}
}

func TestContextOverflowIncidentDoesNotCrossANewerUserMessage(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "message": map[string]any{
			"role": "assistant", "usage": map[string]any{"totalTokens": float64(100000)},
		}},
		{"type": "message", "id": "old-overflow", "message": map[string]any{
			"role": "assistant", "stopReason": "error", "errorMessage": "context_length_exceeded",
		}},
		{"type": "message", "message": map[string]any{"role": "user", "content": "new attempt"}},
	}
	if _, ok := contextOverflowIncident(entries, 100000); ok {
		t.Fatal("an overflow before the latest user message is not the immediately preceding run")
	}
}

func TestThinkingOnlyStopIncident(t *testing.T) {
	tests := []struct {
		name    string
		message map[string]any
		want    bool
	}{
		{
			name: "thinking only stop",
			message: map[string]any{
				"role": "assistant", "stopReason": "stop",
				"content": []any{map[string]any{"type": "thinking", "thinking": "still working"}},
			},
			want: true,
		},
		{
			name: "empty trailing text is harmless",
			message: map[string]any{
				"role": "assistant", "stopReason": "stop",
				"content": []any{
					map[string]any{"type": "thinking", "thinking": "still working"},
					map[string]any{"type": "text", "text": "  "},
				},
			},
			want: true,
		},
		{
			name: "completed text",
			message: map[string]any{
				"role": "assistant", "stopReason": "stop",
				"content": []any{
					map[string]any{"type": "thinking", "thinking": "done"},
					map[string]any{"type": "text", "text": "answer"},
				},
			},
		},
		{
			name: "tool call",
			message: map[string]any{
				"role": "assistant", "stopReason": "stop",
				"content": []any{
					map[string]any{"type": "thinking", "thinking": "done"},
					map[string]any{"type": "toolCall", "name": "read"},
				},
			},
		},
		{
			name: "output length",
			message: map[string]any{
				"role": "assistant", "stopReason": "length",
				"content": []any{map[string]any{"type": "thinking", "thinking": "still working"}},
			},
		},
		{
			name: "empty thinking",
			message: map[string]any{
				"role": "assistant", "stopReason": "stop",
				"content": []any{map[string]any{"type": "thinking", "thinking": ""}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries := []map[string]any{{
				"type": "message", "id": "candidate", "message": tt.message,
			}}
			id, got := thinkingOnlyStopIncident(entries)
			if got != tt.want {
				t.Fatalf("matched = %v, want %v", got, tt.want)
			}
			if got && id != "candidate" {
				t.Fatalf("incident id = %q, want candidate", id)
			}
		})
	}
}

func TestThinkingOnlyStopIncidentDoesNotCrossANewerUserMessage(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "id": "old", "message": map[string]any{
			"role": "assistant", "stopReason": "stop",
			"content": []any{map[string]any{"type": "thinking", "thinking": "unfinished"}},
		}},
		{"type": "message", "message": map[string]any{"role": "user", "content": "new attempt"}},
	}
	if _, ok := thinkingOnlyStopIncident(entries); ok {
		t.Fatal("a thinking-only stop before the latest user message must not be recovered")
	}
}

func TestLocalRecoveryClaimRequiresHealthyProgressAfterAutomaticAttempt(t *testing.T) {
	s := newTestServer(t)
	entries := []map[string]any{
		{"type": "message", "id": "incident-1", "timestamp": "2026-09-08T00:00:00Z"},
		{"type": "message", "id": "incident-2", "timestamp": "2026-09-08T00:00:01Z"},
	}
	if !s.claimLocalRecovery("session", "incident-1", entries) {
		t.Fatal("first incident should be claimed")
	}
	if s.claimLocalRecovery("session", "incident-2", entries) {
		t.Fatal("immediate second incident should be blocked without healthy progress")
	}
	s.resetLocalRecoveryBudget("session")
	if !s.claimLocalRecovery("session", "incident-2", entries) {
		t.Fatal("new user action should reset the bounded recovery budget")
	}
	if s.claimLocalRecovery("session", "incident-2", entries) {
		t.Fatal("same incident must never be claimed twice")
	}
}

func TestLocalRecoveryClaimAllowsNewIncidentAfterHealthyProgress(t *testing.T) {
	s := newTestServer(t)
	entries := []map[string]any{
		{"type": "message", "id": "incident-1", "timestamp": "2026-09-08T00:00:00Z"},
		{"type": "message", "id": "tool-1", "message": map[string]any{"role": "toolResult"}},
		{"type": "message", "id": "tool-2", "message": map[string]any{"role": "toolResult"}},
		{"type": "message", "id": "tool-3", "message": map[string]any{"role": "toolResult"}},
		{"type": "message", "id": "incident-2", "timestamp": "2026-09-08T00:01:00Z"},
	}
	if !s.claimLocalRecovery("session", "incident-1", entries) {
		t.Fatal("first incident should be claimed")
	}
	if !s.claimLocalRecovery("session", "incident-2", entries) {
		t.Fatal("a new incident after three completed tool results should be claimed")
	}

	entries = append(entries,
		map[string]any{"type": "message", "id": "answer", "message": map[string]any{
			"role": "assistant", "stopReason": "stop",
			"content": []any{map[string]any{"type": "text", "text": "checkpoint complete"}},
		}},
		map[string]any{"type": "message", "id": "incident-3", "timestamp": "2026-09-08T00:02:00Z"},
	)
	if !s.claimLocalRecovery("session", "incident-3", entries) {
		t.Fatal("a new incident after a completed assistant answer should be claimed")
	}
}

func TestLocalRecoveryHealthyProgressAllowsOneLongToolCycle(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "id": "incident-1", "timestamp": "2026-09-08T00:00:00Z"},
		{"type": "message", "id": "tool", "message": map[string]any{"role": "toolResult"}},
		{"type": "message", "id": "incident-2", "timestamp": "2026-09-08T00:05:00Z"},
	}
	if !localRecoveryMadeHealthyProgress(entries, "incident-1", "incident-2") {
		t.Fatal("a completed tool cycle sustained for five minutes should count as healthy progress")
	}
}

func TestLocalRecoveryContinuationPromptIsNotHealthyProgress(t *testing.T) {
	entries := []map[string]any{
		{"type": "message", "id": "incident-1"},
		{"type": "message", "id": "continue", "message": map[string]any{
			"role": "user", "content": "continue if possible",
		}},
		{"type": "message", "id": "incident-2"},
	}
	if localRecoveryMadeHealthyProgress(entries, "incident-1", "incident-2") {
		t.Fatal("the watchdog's own continuation prompt must not re-arm recovery")
	}
}

func TestEndpointIsLocal(t *testing.T) {
	for _, endpoint := range []string{
		"http://localhost:11434/v1",
		"http://127.0.0.1:8080",
		"http://[::1]:8080",
		"http://192.168.1.20:8000/v1",
		"http://10.0.0.8/v1",
		"http://modelbox.local/v1",
	} {
		if !endpointIsLocal(endpoint) {
			t.Errorf("endpointIsLocal(%q) = false", endpoint)
		}
	}
	for _, endpoint := range []string{"", "not a url", "https://api.openai.com/v1", "https://example.com/v1"} {
		if endpointIsLocal(endpoint) {
			t.Errorf("endpointIsLocal(%q) = true", endpoint)
		}
	}
}

func TestEffectiveSessionMode(t *testing.T) {
	lan := modelModeMetadata{Provider: "custom", BaseURL: "http://192.168.0.4:8000/v1"}
	cloud := modelModeMetadata{Provider: "openai", BaseURL: "https://api.openai.com/v1"}
	if got := effectiveSessionMode(sessionModeAuto, lan); got != sessionModeLocal {
		t.Fatalf("auto LAN = %q", got)
	}
	if got := effectiveSessionMode(sessionModeAuto, cloud); got != sessionModeCloud {
		t.Fatalf("auto cloud = %q", got)
	}
	if got := effectiveSessionMode(sessionModeCloud, lan); got != sessionModeCloud {
		t.Fatalf("manual cloud override = %q", got)
	}
	if got := effectiveSessionMode(sessionModeLocal, cloud); got != sessionModeLocal {
		t.Fatalf("manual local override = %q", got)
	}
}

func TestModelMetadataFromPayload(t *testing.T) {
	payload := json.RawMessage(`{"models":[{"provider":"custom","id":"qwen","baseUrl":"http://localhost:8080/v1","contextWindow":65536}]}`)
	got := modelMetadataFromPayload(payload, "custom", "qwen")
	if got.BaseURL != "http://localhost:8080/v1" || got.ContextWindow != 65536 {
		t.Fatalf("metadata = %#v", got)
	}
}

func TestSessionModeManualOverridePersistsAndWinsOverModelDetection(t *testing.T) {
	s := newTestServer(t)
	s.models = func(context.Context) (json.RawMessage, error) {
		return json.RawMessage(`{"models":[{"provider":"custom","id":"qwen","baseUrl":"http://localhost:8080/v1","contextWindow":65536}]}`), nil
	}
	project := filepath.Join(s.sessionsDir, "project")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatal(err)
	}
	body, _ := json.Marshal(map[string]any{
		"path": project, "mode": "cloud", "modelProvider": "custom", "modelId": "qwen",
	})
	w := httptest.NewRecorder()
	s.handleNewSession(w, httptest.NewRequest(http.MethodPost, "/api/new-session", bytes.NewReader(body)))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d: %s", w.Code, w.Body.String())
	}
	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	id, _ := response["id"].(string)
	state := s.sessionMode(id, "custom", "qwen")
	if state.ConfiguredMode != sessionModeCloud || state.EffectiveMode != sessionModeCloud {
		t.Fatalf("mode = %#v", state)
	}
	updated := s.updateAutoSessionMode(context.Background(), id, "custom", "qwen")
	if updated.ConfiguredMode != sessionModeCloud || updated.EffectiveMode != sessionModeCloud {
		t.Fatalf("manual override changed after model detection: %#v", updated)
	}
	modeBody := bytes.NewBufferString(`{"mode":"auto"}`)
	w = httptest.NewRecorder()
	s.handleSessionMode(w, httptest.NewRequest(http.MethodPost, "/api/session-mode?id="+id, modeBody))
	if w.Code != http.StatusOK {
		t.Fatalf("switch to auto status = %d: %s", w.Code, w.Body.String())
	}
	state = s.sessionMode(id, "custom", "qwen")
	if state.ConfiguredMode != sessionModeAuto || state.EffectiveMode != sessionModeLocal {
		t.Fatalf("auto mode was not restored/persisted: %#v", state)
	}
}

func TestSessionModeChangeIsRejectedWhileWorkerIsRunning(t *testing.T) {
	s := newTestServer(t)
	s.chatSender = &fakeSender{status: workers.WorkerStatus{State: workers.WorkerStateRunning}}
	dir := filepath.Join(s.sessionsDir, "--project--")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "session.jsonl")
	if err := os.WriteFile(path, []byte(`{"type":"session","version":3,"id":"sid","cwd":"/tmp"}`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := s.saveSessionMode("session.jsonl", sessionModeCloud, modelModeMetadata{}); err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	s.handleSessionMode(w, httptest.NewRequest(http.MethodPost, "/api/session-mode?id=session.jsonl", bytes.NewBufferString(`{"mode":"local"}`)))
	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, want 409: %s", w.Code, w.Body.String())
	}
	if got := s.sessionMode("session.jsonl", "", "").ConfiguredMode; got != sessionModeCloud {
		t.Fatalf("persisted mode = %q, want unchanged cloud", got)
	}
}

func TestWatchdogCompactsAndContinuesOnceForContextIncident(t *testing.T) {
	s := newTestServer(t)
	project := filepath.Join(s.sessionsDir, "project")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(s.sessionsDir, "--project--")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "session.jsonl")
	content := `{"type":"session","version":3,"id":"sid","cwd":` + jsonString(project) + `}` + "\n" +
		`{"type":"message","id":"usage","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"stop","usage":{"totalTokens":99500},"content":[{"type":"text","text":"working"}]}}` + "\n" +
		`{"type":"message","id":"overflow","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"error","errorMessage":"maximum context length exceeded","content":[]}}` + "\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{
		Provider: "custom", ID: "qwen", ContextWindow: 100000,
	}); err != nil {
		t.Fatal(err)
	}
	sender := &compactFakeSender{
		events: make(chan string, 2), settled: make(chan struct{}), waitStarted: make(chan struct{}),
	}
	s.chatSender = sender
	if !s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("watchdog did not claim qualifying incident")
	}
	for i, want := range []string{"compact", "send:continue if possible"} {
		select {
		case got := <-sender.events:
			if got != want {
				t.Fatalf("event %d = %q, want %q", i, got, want)
			}
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for %q", want)
		}
	}
	select {
	case <-sender.waitStarted:
	case <-time.After(time.Second):
		t.Fatal("watchdog did not wait for agent_settled")
	}
	s.localRecoveryMu.Lock()
	runningBeforeSettled := s.localRecoveryRunning
	s.localRecoveryMu.Unlock()
	if !runningBeforeSettled {
		t.Fatal("watchdog released the global recovery slot before agent_settled")
	}
	sender.mu.Lock()
	preparedWindow := sender.modeConfig.LocalContextWindow
	sender.mu.Unlock()
	if preparedWindow != 100000 {
		t.Fatalf("recovery worker context window = %d, want 100000", preparedWindow)
	}
	close(sender.settled)
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		s.localRecoveryMu.Lock()
		running := s.localRecoveryRunning
		s.localRecoveryMu.Unlock()
		if !running {
			break
		}
		time.Sleep(time.Millisecond)
	}
	if s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("same persisted incident was recovered twice")
	}
}

func TestWatchdogContinuesOnceForThinkingOnlyStopWithoutCompaction(t *testing.T) {
	s := newTestServer(t)
	project := filepath.Join(s.sessionsDir, "project")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(s.sessionsDir, "--project--")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "session.jsonl")
	content := `{"type":"session","version":3,"id":"sid","cwd":` + jsonString(project) + `}` + "\n" +
		`{"type":"message","id":"thinking-stop","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"stop","content":[{"type":"thinking","thinking":"unfinished reasoning"}]}}` + "\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{
		Provider: "custom", ID: "qwen", ContextWindow: 100000,
	}); err != nil {
		t.Fatal(err)
	}
	sender := &compactFakeSender{events: make(chan string, 2)}
	s.chatSender = sender
	client := s.addClient("session.jsonl")
	defer s.removeClient(client)

	if !s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("watchdog did not claim thinking-only stop")
	}
	select {
	case got := <-client.ch:
		want := "event: local-recovery\ndata: {\"reason\":\"thinking-only-stop\"}"
		if got != want {
			t.Fatalf("recovery event = %q, want %q", got, want)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for recovery event")
	}
	select {
	case got := <-sender.events:
		if got != "send:continue if possible" {
			t.Fatalf("event = %q, want continuation without compaction", got)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for continuation")
	}

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		s.localRecoveryMu.Lock()
		running := s.localRecoveryRunning
		s.localRecoveryMu.Unlock()
		if !running {
			break
		}
		time.Sleep(time.Millisecond)
	}
	if s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("same thinking-only incident was recovered twice")
	}
}

func TestWatchdogTransportInterruption(t *testing.T) {
	for _, tc := range []struct {
		name                string
		stopReason          string
		usage               int
		compacted           bool
		running             bool
		stopped             bool
		cancelDuringCompact bool
		want                []string
	}{
		{name: "high context", stopReason: "error", usage: 87000, want: []string{"compact", "send:continue if possible"}},
		{name: "low context", stopReason: "error", usage: 1000, want: []string{"send:continue if possible"}},
		{name: "already compacted", stopReason: "error", usage: 87000, compacted: true, want: []string{"send:continue if possible"}},
		{name: "user aborted", stopReason: "aborted", usage: 87000},
		{name: "worker still recovering", stopReason: "error", usage: 87000, running: true},
		{name: "user stopped", stopReason: "error", usage: 87000, stopped: true},
		{name: "stop during compact", stopReason: "error", usage: 87000, cancelDuringCompact: true, want: []string{"compact", "abort"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestServer(t)
			dir := filepath.Join(s.sessionsDir, "--project--")
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatal(err)
			}
			content := `{"type":"session","version":3,"id":"sid","cwd":` + jsonString(dir) + "}\n" +
				fmt.Sprintf(`{"type":"message","id":"usage","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"stop","usage":{"totalTokens":%d},"content":[{"type":"text","text":"working"}]}}`, tc.usage) + "\n" +
				fmt.Sprintf(`{"type":"message","id":"interrupted","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":%q,"errorMessage":"This operation was aborted","content":[]}}`, tc.stopReason) + "\n"
			if tc.compacted {
				content += `{"type":"compaction","id":"checkpoint","summary":"checkpoint","firstKeptEntryId":"interrupted","tokensBefore":87000}` + "\n"
			}
			if err := os.WriteFile(filepath.Join(dir, "session.jsonl"), []byte(content), 0644); err != nil {
				t.Fatal(err)
			}
			if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{Provider: "custom", ID: "qwen", ContextWindow: 100000}); err != nil {
				t.Fatal(err)
			}
			sender := &compactFakeSender{events: make(chan string, 4)}
			if tc.running {
				sender.status = workers.WorkerStatus{State: workers.WorkerStateRunning}
			}
			if tc.cancelDuringCompact {
				sender.compactRelease = make(chan struct{})
			}
			s.chatSender = sender
			if tc.stopped {
				s.stopLocalRecovery("session.jsonl")
			}
			if got := s.maybeStartLocalRecovery("session.jsonl"); got != (len(tc.want) > 0) {
				t.Fatalf("started = %v", got)
			}
			for _, want := range tc.want {
				select {
				case got := <-sender.events:
					if got != want {
						t.Fatalf("event = %q, want %q", got, want)
					}
					if tc.cancelDuringCompact && got == "compact" {
						w := httptest.NewRecorder()
						s.handleCancelChat(w, httptest.NewRequest(http.MethodPost, "/api/chat/cancel?id=session.jsonl", nil))
						if w.Code != http.StatusOK {
							t.Fatalf("cancel: %d %s", w.Code, w.Body.String())
						}
					}
				case <-time.After(time.Second):
					t.Fatalf("missing %q", want)
				}
			}
			deadline := time.Now().Add(time.Second)
			for {
				s.localRecoveryMu.Lock()
				running := s.localRecoveryRunning
				s.localRecoveryMu.Unlock()
				if !running {
					break
				}
				if time.Now().After(deadline) {
					t.Fatal("recovery did not settle")
				}
				time.Sleep(time.Millisecond)
			}
			if s.maybeStartLocalRecovery("session.jsonl") {
				t.Fatal("recovered the same incident again")
			}
			select {
			case e := <-sender.events:
				t.Fatalf("unexpected operation %s", e)
			default:
			}
		})
	}
}

func TestTransportInterruptionDoesNotRecoverOlderErrors(t *testing.T) {
	failure := map[string]any{"type": "message", "id": "failure", "message": map[string]any{"role": "assistant", "stopReason": "error", "errorMessage": "This operation was aborted"}}
	for _, latest := range []map[string]any{
		{"role": "user"},
		{"role": "assistant", "stopReason": "stop"},
		{"role": "assistant", "stopReason": "error", "errorMessage": "authentication failed"},
		{"role": "assistant", "stopReason": "aborted", "errorMessage": "This operation was aborted"},
	} {
		if _, ok := transportInterruptionIncident([]map[string]any{failure, {"type": "message", "message": latest}}); ok {
			t.Fatalf("recovered superseded error: %v", latest)
		}
	}
}

type progressRecoverySender struct {
	compactFakeSender
	calls int
}

func (f *progressRecoverySender) Send(ctx context.Context, sessionID, sessionPath string, request chat.Request) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls++
	var output strings.Builder
	if f.calls < 4 {
		for i := 0; i < 3; i++ {
			fmt.Fprintf(&output, `{"type":"message","id":"tool-%d-%d","message":{"role":"toolResult","isError":false}}`+"\n", f.calls, i)
		}
	}
	fmt.Fprintf(&output, `{"type":"message","id":"failure-%d","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"error","errorMessage":"This operation was aborted","content":[]}}`+"\n", f.calls)
	file, err := os.OpenFile(sessionPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	_, err = file.WriteString(output.String())
	closeErr := file.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func TestWatchdogKeepsRecoveringWithProgressAndStopsWithoutProgress(t *testing.T) {
	s := newTestServer(t)
	dir := filepath.Join(s.sessionsDir, "--project--")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	content := `{"type":"session","version":3,"id":"sid","cwd":` + jsonString(dir) + "}\n" +
		`{"type":"message","id":"failure-0","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"error","errorMessage":"This operation was aborted","content":[]}}` + "\n"
	if err := os.WriteFile(filepath.Join(dir, "session.jsonl"), []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{Provider: "custom", ID: "qwen", ContextWindow: 100000}); err != nil {
		t.Fatal(err)
	}
	sender := &progressRecoverySender{}
	s.chatSender = sender
	if !s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("initial recovery not started")
	}
	deadline := time.Now().Add(3 * time.Second)
	for {
		sender.mu.Lock()
		calls := sender.calls
		sender.mu.Unlock()
		s.localRecoveryMu.Lock()
		running := s.localRecoveryRunning
		s.localRecoveryMu.Unlock()
		if calls == 4 && !running {
			break
		}
		if calls > 4 || time.Now().After(deadline) {
			t.Fatalf("calls=%d running=%v; want four recoveries then circuit breaker", calls, running)
		}
		time.Sleep(time.Millisecond)
	}
	if s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("no-progress incident bypassed circuit breaker")
	}
	s.stopLocalRecovery("session.jsonl")
	s.resetLocalRecoveryBudget("session.jsonl")
	if !s.maybeStartLocalRecovery("session.jsonl") {
		t.Fatal("a new user action should permit recovery again")
	}
}
