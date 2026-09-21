package server

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"pi-web/internal/schedules"

	_ "modernc.org/sqlite"
)

func newScheduleTestServer(t *testing.T) (*Server, *fakeSender) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	db.SetMaxOpenConns(1)
	for _, ddl := range []string{
		schedules.SchedulesTableDDL,
		schedules.RunsTableDDL,
		schedules.RunsScheduleIndexDDL,
		schedules.RunsSessionIndexDDL,
		sessionModesSchema,
	} {
		if _, err := db.Exec(ddl); err != nil {
			t.Fatalf("schema: %v", err)
		}
	}
	t.Cleanup(func() { db.Close() })

	sender := &fakeSender{}
	s := &Server{
		db:          db,
		schedules:   schedules.NewStore(db),
		sessionsDir: t.TempDir(),
		chatSender:  sender,
		now:         func() time.Time { return time.Date(2026, 6, 15, 8, 0, 0, 0, time.UTC) },
	}
	return s, sender
}

func TestFireScheduleCreatesSessionAndSends(t *testing.T) {
	s, sender := newScheduleTestServer(t)
	project := t.TempDir()

	sc, err := s.schedules.Create(schedules.Schedule{
		ID:           "sched1",
		Name:         "Nightly digest",
		Instructions: "Summarize the day",
		ProjectPath:  project,
		Enabled:      true,
	})
	if err != nil {
		t.Fatalf("create schedule: %v", err)
	}

	sessionID, err := s.fireSchedule(sc)
	if err != nil {
		t.Fatalf("fireSchedule: %v", err)
	}
	if sessionID == "" {
		t.Fatal("expected a session id")
	}

	_, _, req := sender.sentInfo()
	if req.Message != "Summarize the day" {
		t.Errorf("Send message = %q, want instructions", req.Message)
	}
	if !sender.ensureWorkerCalled {
		t.Error("EnsureWorker was not called")
	}

	// The run is recorded and mapped to the created session.
	name, ok := s.scheduleNameForSession(sessionID)
	if !ok || name != "Nightly digest" {
		t.Errorf("scheduleNameForSession = %q, %v", name, ok)
	}
	runs, err := s.schedules.ListRuns("sched1", 10)
	if err != nil || len(runs) != 1 {
		t.Fatalf("ListRuns = %v (err %v)", runs, err)
	}
	if runs[0].SessionID != sessionID {
		t.Errorf("run session = %q, want %q", runs[0].SessionID, sessionID)
	}
}

func TestFireScheduleAppliesConfiguredModel(t *testing.T) {
	s, sender := newScheduleTestServer(t)

	sc, err := s.schedules.Create(schedules.Schedule{
		ID:            "sched-model",
		Name:          "Model run",
		Instructions:  "go",
		ModelProvider: "opencode-go",
		ModelID:       "deepseek-v4.1-flash",
		ThinkingLevel: "high",
		ProjectPath:   t.TempDir(),
		Enabled:       true,
	})
	if err != nil {
		t.Fatalf("create schedule: %v", err)
	}

	sessionID, err := s.fireSchedule(sc)
	if err != nil {
		t.Fatalf("fireSchedule: %v", err)
	}

	sender.mu.Lock()
	provider, modelID := sender.setModelProvider, sender.setModelID
	modelSession := sender.setModelSessionID
	thinking, thinkingSession := sender.setThinkingLevel, sender.setThinkingSessionID
	sender.mu.Unlock()

	if provider != "opencode-go" || modelID != "deepseek-v4.1-flash" {
		t.Errorf("SetModel = %s/%s, want opencode-go/deepseek-v4.1-flash", provider, modelID)
	}
	if modelSession != sessionID {
		t.Errorf("SetModel session = %q, want %q", modelSession, sessionID)
	}
	if thinking != "high" {
		t.Errorf("SetThinkingLevel = %q, want high", thinking)
	}
	if thinkingSession != sessionID {
		t.Errorf("SetThinkingLevel session = %q, want %q", thinkingSession, sessionID)
	}
}

func TestFireScheduleKeepsDefaultsWhenUnset(t *testing.T) {
	s, sender := newScheduleTestServer(t)

	sc, err := s.schedules.Create(schedules.Schedule{
		ID:           "sched-default",
		Name:         "Default run",
		Instructions: "go",
		ProjectPath:  t.TempDir(),
		Enabled:      true,
	})
	if err != nil {
		t.Fatalf("create schedule: %v", err)
	}

	if _, err := s.fireSchedule(sc); err != nil {
		t.Fatalf("fireSchedule: %v", err)
	}

	sender.mu.Lock()
	provider, modelID := sender.setModelProvider, sender.setModelID
	thinking := sender.setThinkingLevel
	sender.mu.Unlock()

	if provider != "" || modelID != "" {
		t.Errorf("SetModel called with %s/%s, want no call", provider, modelID)
	}
	if thinking != "" {
		t.Errorf("SetThinkingLevel called with %q, want no call", thinking)
	}
}

func TestScheduleModelCandidates(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	s.models = func(context.Context) (json.RawMessage, error) {
		return json.RawMessage(`{"models":[
			{"provider":"openai","id":"gpt-5","name":"GPT 5"},
			{"provider":"ollama","id":"qwen-local","name":"Qwen Local","baseUrl":"http://127.0.0.1:11434/v1"},
			{"provider":"anthropic","id":"claude-sonnet","name":"Claude Sonnet","starred":true}
		]}`), nil
	}
	local := s.scheduleModelCandidates(context.Background(), schedules.Schedule{ModelSelector: "local"})
	if len(local) != 3 || local[0].ID != "qwen-local" {
		t.Fatalf("local candidates = %+v", local)
	}
	free := s.scheduleModelCandidates(context.Background(), schedules.Schedule{ModelSelector: "free"})
	if len(free) != 3 || free[0].ID != "claude-sonnet" {
		t.Fatalf("free candidates = %+v", free)
	}
	named := s.scheduleModelCandidates(context.Background(), schedules.Schedule{ModelSelector: "sonnet"})
	if len(named) != 3 || named[0].ID != "claude-sonnet" {
		t.Fatalf("named candidates = %+v", named)
	}
}

func TestScheduleCallbackDelivery(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	received := make(chan scheduleCallbackPayload, 1)
	s.callbackHTTPClient = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		var payload scheduleCallbackPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Error(err)
		}
		received <- payload
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(strings.NewReader("")), Header: make(http.Header)}, nil
	})}
	sc, err := s.schedules.Create(schedules.Schedule{ID: "callback", Name: "Callback", Instructions: "go", CallbackURL: "http://callback.test/result", Enabled: true})
	if err != nil {
		t.Fatal(err)
	}
	runID, err := s.schedules.RecordRun(schedules.Run{ScheduleID: sc.ID, FiredAt: time.Now().UTC().Format(time.RFC3339), Status: schedules.RunStatusRunning})
	if err != nil {
		t.Fatal(err)
	}
	if err := s.schedules.CompleteRun(runID, schedules.RunStatusSucceeded, "done", "", "openai", "gpt-5"); err != nil {
		t.Fatal(err)
	}
	s.deliverScheduleCallback(context.Background(), sc, runID)
	payload := <-received
	if payload.Run.Status != schedules.RunStatusSucceeded || payload.Run.Result != "done" || payload.EventID == "" {
		t.Fatalf("payload = %+v", payload)
	}
	run, _ := s.schedules.GetRun(runID)
	if run.CallbackStatus != "delivered" || run.CallbackAttempts != 1 {
		t.Fatalf("delivery = %+v", run)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestScheduleResultStatuses(t *testing.T) {
	entry := func(reason, text, errMessage string) []map[string]any {
		return []map[string]any{{
			"type": "message",
			"message": map[string]any{
				"role": "assistant", "stopReason": reason, "errorMessage": errMessage,
				"content": []any{map[string]any{"type": "text", "text": text}},
			},
		}}
	}
	if status, result, _ := scheduleResult(entry("stop", "ok", "")); status != schedules.RunStatusSucceeded || result != "ok" {
		t.Fatalf("success = %s %q", status, result)
	}
	if status, _, _ := scheduleResult(entry("error", "", "boom")); status != schedules.RunStatusFailed {
		t.Fatalf("failure = %s", status)
	}
	if status, _, _ := scheduleResult(entry("aborted", "", "stopped")); status != schedules.RunStatusCancelled {
		t.Fatalf("cancelled = %s", status)
	}
}

func TestEvaluateSchedulesSkipsMissedRuns(t *testing.T) {
	s, sender := newScheduleTestServer(t)
	// A daily 09:00 schedule; "now" is 08:00. First evaluation must only arm the
	// next fire, never fire immediately for a missed past occurrence.
	if _, err := s.schedules.Create(schedules.Schedule{
		ID:           "daily",
		Name:         "Daily",
		Instructions: "go",
		CronExpr:     "0 9 * * *",
		Timezone:     "UTC",
		Enabled:      true,
	}); err != nil {
		t.Fatal(err)
	}

	state := map[string]scheduleState{}
	s.evaluateSchedules(state)

	if _, _, req := sender.sentInfo(); req.Message != "" {
		t.Errorf("schedule fired on first evaluation; message=%q", req.Message)
	}
	st, ok := state["daily"]
	if !ok {
		t.Fatal("expected next-fire state to be armed")
	}
	want := time.Date(2026, 6, 15, 9, 0, 0, 0, time.UTC)
	if !st.next.Equal(want) {
		t.Errorf("next = %v, want %v", st.next, want)
	}
}

func TestEvaluateSchedulesIgnoresManualAndDisabled(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	if _, err := s.schedules.Create(schedules.Schedule{ID: "m", Name: "Manual", Instructions: "x", Enabled: true}); err != nil {
		t.Fatal(err)
	}
	if _, err := s.schedules.Create(schedules.Schedule{ID: "d", Name: "Disabled", Instructions: "x", CronExpr: "0 9 * * *", Enabled: false}); err != nil {
		t.Fatal(err)
	}
	state := map[string]scheduleState{}
	s.evaluateSchedules(state)
	if len(state) != 0 {
		t.Errorf("manual/disabled schedules should not be armed; state=%v", state)
	}
}

func TestEvaluateSchedulesRunsAndDisablesOneShot(t *testing.T) {
	s, sender := newScheduleTestServer(t)
	if _, err := s.schedules.Create(schedules.Schedule{
		ID: "once", Name: "Once", Instructions: "remind me", ProjectPath: t.TempDir(),
		RunAt: "2026-06-15T07:59:00Z", Enabled: true,
	}); err != nil {
		t.Fatal(err)
	}
	s.evaluateSchedules(map[string]scheduleState{})
	s.wg.Wait()
	sc, err := s.schedules.Get("once")
	if err != nil {
		t.Fatal(err)
	}
	if sc.Enabled {
		t.Fatal("one-shot schedule remained enabled")
	}
	if _, _, req := sender.sentInfo(); req.Message != "remind me" {
		t.Fatalf("message = %q", req.Message)
	}
	runs, err := s.schedules.ListRuns("once", 10)
	if err != nil || len(runs) != 1 {
		t.Fatalf("runs = %+v, err = %v", runs, err)
	}
	s.evaluateSchedules(map[string]scheduleState{})
	s.wg.Wait()
	runs, _ = s.schedules.ListRuns("once", 10)
	if len(runs) != 1 {
		t.Fatalf("one-shot fired more than once: %+v", runs)
	}
}

func TestSchedulesAPICreateListRun(t *testing.T) {
	s, sender := newScheduleTestServer(t)
	project := t.TempDir()

	body, _ := json.Marshal(map[string]any{
		"name":         "API sched",
		"instructions": "do it",
		"projectPath":  project,
		"cronExpr":     "0 9 * * *",
		"timezone":     "UTC",
		"enabled":      true,
	})
	w := httptest.NewRecorder()
	s.handleApiSchedules(w, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body)))
	if w.Code != http.StatusCreated {
		t.Fatalf("create status = %d, body %s", w.Code, w.Body.String())
	}
	var created struct {
		Schedule schedules.Schedule `json:"schedule"`
	}
	json.Unmarshal(w.Body.Bytes(), &created)
	if created.Schedule.ID == "" {
		t.Fatal("expected created schedule id")
	}
	if created.Schedule.NextRunAt == "" {
		t.Error("expected nextRunAt to be computed")
	}

	// List returns it.
	lw := httptest.NewRecorder()
	s.handleApiSchedules(lw, httptest.NewRequest(http.MethodGet, "/api/schedules", nil))
	var list struct {
		Schedules []schedules.Schedule `json:"schedules"`
	}
	json.Unmarshal(lw.Body.Bytes(), &list)
	if len(list.Schedules) != 1 {
		t.Fatalf("list len = %d", len(list.Schedules))
	}

	// Run-now fires regardless of cadence.
	rw := httptest.NewRecorder()
	s.handleApiScheduleRun(rw, httptest.NewRequest(http.MethodPost, "/api/schedule/run?id="+created.Schedule.ID, nil))
	if rw.Code != http.StatusAccepted {
		t.Fatalf("run status = %d, body %s", rw.Code, rw.Body.String())
	}
	if _, _, req := sender.sentInfo(); req.Message != "do it" {
		t.Errorf("run-now Send message = %q", req.Message)
	}
}

func TestSchedulesAPICreateBroadcastsSSE(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	client := s.addClient(globalSessID)
	defer s.removeClient(client)

	body, _ := json.Marshal(map[string]any{
		"name":         "SSE sched",
		"instructions": "do it",
		"cronExpr":     "0 9 * * *",
		"timezone":     "UTC",
	})
	w := httptest.NewRecorder()
	s.handleApiSchedules(w, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body)))
	if w.Code != http.StatusCreated {
		t.Fatalf("create status = %d, body %s", w.Code, w.Body.String())
	}

	select {
	case msg := <-client.ch:
		if !strings.Contains(msg, "event: schedules") {
			t.Fatalf("sse = %q, want schedules event", msg)
		}
		if !strings.Contains(msg, `"action":"created"`) {
			t.Fatalf("sse = %q, want action=created", msg)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for schedules SSE")
	}
}

func TestSchedulesAPIDeleteBroadcastsSSE(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	created, err := s.schedules.Create(schedules.Schedule{
		ID: "del-1", Name: "Gone", Instructions: "x", Enabled: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	client := s.addClient(globalSessID)
	defer s.removeClient(client)

	w := httptest.NewRecorder()
	s.handleApiSchedule(w, httptest.NewRequest(http.MethodDelete, "/api/schedule?id="+created.ID, nil))
	if w.Code != http.StatusOK {
		t.Fatalf("delete status = %d, body %s", w.Code, w.Body.String())
	}

	select {
	case msg := <-client.ch:
		if !strings.Contains(msg, `"action":"deleted"`) {
			t.Fatalf("sse = %q, want action=deleted", msg)
		}
		if !strings.Contains(msg, created.ID) {
			t.Fatalf("sse = %q, want id %s", msg, created.ID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for schedules SSE")
	}
}

func TestSchedulesAPIValidation(t *testing.T) {
	s, _ := newScheduleTestServer(t)
	// Missing name.
	body, _ := json.Marshal(map[string]any{"instructions": "x"})
	w := httptest.NewRecorder()
	s.handleApiSchedules(w, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body)))
	if w.Code != http.StatusBadRequest {
		t.Errorf("missing name status = %d, want 400", w.Code)
	}
	// Bad cron.
	body2, _ := json.Marshal(map[string]any{"name": "n", "instructions": "x", "cronExpr": "nope"})
	w2 := httptest.NewRecorder()
	s.handleApiSchedules(w2, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body2)))
	if w2.Code != http.StatusBadRequest {
		t.Errorf("bad cron status = %d, want 400", w2.Code)
	}
	// One-shot timestamps are RFC3339 and cannot be combined with cron.
	body3, _ := json.Marshal(map[string]any{"name": "n", "instructions": "x", "runAt": "tomorrow"})
	w3 := httptest.NewRecorder()
	s.handleApiSchedules(w3, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body3)))
	if w3.Code != http.StatusBadRequest {
		t.Errorf("bad runAt status = %d, want 400", w3.Code)
	}
	body4, _ := json.Marshal(map[string]any{"name": "n", "instructions": "x", "runAt": "2026-09-20T14:00:00Z", "cronExpr": "0 9 * * *"})
	w4 := httptest.NewRecorder()
	s.handleApiSchedules(w4, httptest.NewRequest(http.MethodPost, "/api/schedules", bytes.NewReader(body4)))
	if w4.Code != http.StatusBadRequest {
		t.Errorf("runAt+cron status = %d, want 400", w4.Code)
	}
}
