package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/sessions"
	"pi-web/internal/workers"
)

type fakeSender struct {
	// Config fields below are set at construction (before the server starts any
	// goroutine) and only read afterwards, so they need no locking.
	state          workers.WorkerStatus
	status         workers.WorkerStatus
	getStateErr    error
	ensureWorkerCh chan struct{}
	sendCh         chan struct{}
	commands       []workers.SlashCommand
	commandsReady  bool
	commandsErr    error

	// mu guards every field below it. handleNewSession (and friends) initialize
	// workers on a background goroutine, so the sender methods run concurrently
	// with test assertions that read these fields.
	mu                      sync.Mutex
	sessionID               string
	sessionPath             string
	chat                    chat.Request
	getStateCalls           int
	ensureWorkerCalled      bool
	ensureWorkerSessionID   string
	ensureWorkerSessionPath string
	setModelSessionID       string
	setModelProvider        string
	setModelID              string
	setThinkingSessionID    string
	setThinkingLevel        string
	getCommandsCalls        int
}

type compactFakeSender struct {
	fakeSender
	events           chan string
	settled          chan struct{}
	waitStarted      chan struct{}
	modeConfig       workers.WorkerConfig
	compactStarted   chan struct{}
	compactRelease   <-chan struct{}
	compactCalls     int
	appendCompaction bool
}

func (f *compactFakeSender) Send(ctx context.Context, sessionID, sessionPath string, req chat.Request) error {
	if f.events != nil {
		f.events <- "send:" + req.Message
	}
	return f.fakeSender.Send(ctx, sessionID, sessionPath, req)
}

func (f *compactFakeSender) Abort(context.Context, string) error {
	if f.events != nil {
		f.events <- "abort"
	}
	return nil
}

func (f *compactFakeSender) Compact(ctx context.Context, _ string, sessionPath string) error {
	f.mu.Lock()
	f.compactCalls++
	call := f.compactCalls
	started := f.compactStarted
	release := f.compactRelease
	appendCompaction := f.appendCompaction
	f.mu.Unlock()
	if f.events != nil {
		f.events <- "compact"
	}
	if call == 1 && started != nil {
		close(started)
	}
	if release != nil {
		select {
		case <-release:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	if appendCompaction {
		file, err := os.OpenFile(sessionPath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		_, err = file.WriteString(`{"type":"compaction","id":"compact","parentId":"a","timestamp":"2026-09-08T00:00:00.000Z","summary":"summary","firstKeptEntryId":"a","tokensBefore":65000}` + "\n")
		closeErr := file.Close()
		if err != nil {
			return err
		}
		return closeErr
	}
	return nil
}

func (f *compactFakeSender) compactCallCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.compactCalls
}

func (f *compactFakeSender) PrepareMode(_ string, config workers.WorkerConfig) error {
	f.mu.Lock()
	f.modeConfig = config
	f.mu.Unlock()
	return nil
}

func (f *compactFakeSender) WaitSettled(ctx context.Context, _ string) error {
	if f.waitStarted != nil {
		close(f.waitStarted)
	}
	if f.settled == nil {
		return nil
	}
	select {
	case <-f.settled:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (f *fakeSender) Send(ctx context.Context, sessionID, sessionPath string, chat chat.Request) error {
	f.mu.Lock()
	f.sessionID = sessionID
	f.sessionPath = sessionPath
	f.chat = chat
	f.mu.Unlock()
	if f.sendCh != nil {
		f.sendCh <- struct{}{}
	}
	return nil
}

func (f *fakeSender) SetModel(ctx context.Context, sessionID, sessionPath, provider, modelID string) error {
	f.mu.Lock()
	f.setModelSessionID = sessionID
	f.setModelProvider = provider
	f.setModelID = modelID
	f.mu.Unlock()
	return nil
}

func (f *fakeSender) SetThinkingLevel(ctx context.Context, sessionID, sessionPath, level string) error {
	f.mu.Lock()
	f.setThinkingSessionID = sessionID
	f.setThinkingLevel = level
	f.mu.Unlock()
	return nil
}

func (f *fakeSender) Abort(ctx context.Context, sessionID string) error {
	return nil
}

func (f *fakeSender) GetState(ctx context.Context, sessionID string) (workers.WorkerStatus, error) {
	f.mu.Lock()
	f.getStateCalls++
	f.mu.Unlock()
	if f.getStateErr != nil {
		return workers.WorkerStatus{}, f.getStateErr
	}
	if f.state.State != "" || f.state.ThinkingLevel != "" || f.state.Model != "" || f.state.ModelProvider != "" {
		return f.state, nil
	}
	return workers.WorkerStatus{State: workers.WorkerStateIdle}, nil
}

func (f *fakeSender) GetCommands(ctx context.Context, sessionID string) ([]workers.SlashCommand, bool, error) {
	f.mu.Lock()
	f.getCommandsCalls++
	f.mu.Unlock()
	return f.commands, f.commandsReady, f.commandsErr
}

func (f *fakeSender) Status(sessionID string) workers.WorkerStatus {
	if f.status.State != "" || f.status.ThinkingLevel != "" || f.status.Model != "" || f.status.ModelProvider != "" {
		return f.status
	}
	return workers.WorkerStatus{State: workers.WorkerStateIdle}
}

func (f *fakeSender) EnsureWorker(ctx context.Context, sessionID, sessionPath string) error {
	f.mu.Lock()
	f.ensureWorkerCalled = true
	f.ensureWorkerSessionID = sessionID
	f.ensureWorkerSessionPath = sessionPath
	f.mu.Unlock()
	if f.ensureWorkerCh != nil {
		f.ensureWorkerCh <- struct{}{}
	}
	return nil
}

// Accessors for the mutex-guarded fields, so tests read them safely from a
// goroutine other than the one the sender method ran on.

func (f *fakeSender) sentInfo() (sessionID, sessionPath string, req chat.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.sessionID, f.sessionPath, f.chat
}

func (f *fakeSender) stateCalls() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.getStateCalls
}

func (f *fakeSender) ensureWorkerInfo() (called bool, sessionID, sessionPath string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.ensureWorkerCalled, f.ensureWorkerSessionID, f.ensureWorkerSessionPath
}

func (f *fakeSender) modelSessionID() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.setModelSessionID
}

func (f *fakeSender) modelSelection() (sessionID, provider, modelID string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.setModelSessionID, f.setModelProvider, f.setModelID
}

func (f *fakeSender) thinkingSessionID() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.setThinkingSessionID
}

func writeHighUsageLocalSession(t *testing.T, s *Server) sessions.ResolvedSession {
	t.Helper()
	project := filepath.Join(s.sessionsDir, "cwd")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(s.sessionsDir, "--project--")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "session.jsonl")
	content := `{"type":"session","version":3,"id":"sid","cwd":` + jsonString(project) + `}` + "\n" +
		`{"type":"message","id":"a","message":{"role":"assistant","model":"qwen","provider":"custom","stopReason":"stop","content":[{"type":"text","text":"ready"}],"usage":{"totalTokens":65000}}}` + "\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{
		Provider: "custom", ID: "qwen", ContextWindow: 100000,
	}); err != nil {
		t.Fatal(err)
	}
	resolved, err := sessions.ResolveByID(s.sessionsDir, "session.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	return resolved
}

func TestHandleChatQueuesResolvedSession(t *testing.T) {
	root := t.TempDir()
	wantPath := writeSessionFile(t, root, "--tmp--project--", "session.jsonl")
	fake := &fakeSender{sendCh: make(chan struct{}, 1)}
	s := &Server{sessionsDir: root, chatSender: fake}
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	mw.WriteField("message", "hello")
	mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/api/chat?id=session.jsonl", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	s.handleChat(w, req)
	if w.Code != http.StatusAccepted {
		t.Fatalf("status = %d body = %s", w.Code, w.Body.String())
	}
	var got map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["status"] != "queued" {
		t.Fatalf("status body = %#v, want queued", got)
	}
	select {
	case <-fake.sendCh:
	case <-time.After(time.Second):
		t.Fatal("Send was not called asynchronously")
	}
	sentID, sentPath, sentReq := fake.sentInfo()
	if sentID != "session.jsonl" || sentPath != wantPath || sentReq.Message != "hello" {
		t.Fatalf("sent id=%q path=%q msg=%q, want path %q", sentID, sentPath, sentReq.Message, wantPath)
	}
}

func TestHandleChatCompactsLocalSessionBeforeSendingAtSixtyFivePercent(t *testing.T) {
	s := newTestServer(t)
	writeHighUsageLocalSession(t, s)
	sender := &compactFakeSender{events: make(chan string, 2)}
	s.chatSender = sender

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	_ = mw.WriteField("message", "next")
	_ = mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/api/chat?id=session.jsonl", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	s.handleChat(w, req)
	if w.Code != http.StatusAccepted {
		t.Fatalf("status = %d: %s", w.Code, w.Body.String())
	}
	for i, want := range []string{"compact", "send:next"} {
		select {
		case got := <-sender.events:
			if got != want {
				t.Fatalf("event %d = %q, want %q", i, got, want)
			}
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for event %q", want)
		}
	}
}

func TestHandleChatDoesNotAbortRunningLocalCompaction(t *testing.T) {
	s := newTestServer(t)
	writeHighUsageLocalSession(t, s)
	sender := &compactFakeSender{
		fakeSender: fakeSender{status: workers.WorkerStatus{State: workers.WorkerStateRunning}},
		events:     make(chan string, 3),
	}
	s.chatSender = sender

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	_ = mw.WriteField("message", "steer while compacting")
	_ = mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/api/chat?id=session.jsonl", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	s.handleChat(w, req)
	if w.Code != http.StatusAccepted {
		t.Fatalf("status = %d: %s", w.Code, w.Body.String())
	}
	select {
	case got := <-sender.events:
		if got != "send:steer while compacting" {
			t.Fatalf("event = %q, want the prompt to be steered without aborting or compacting", got)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for steered prompt")
	}
	if calls := sender.compactCallCount(); calls != 0 {
		t.Fatalf("Compact calls = %d, want 0 while Pi is already running", calls)
	}
}

func TestConcurrentLocalChatRequestsCompactOnceFromFreshSessionState(t *testing.T) {
	s := newTestServer(t)
	resolved := writeHighUsageLocalSession(t, s)
	compactStarted := make(chan struct{})
	compactRelease := make(chan struct{})
	sender := &compactFakeSender{
		events:           make(chan string, 3),
		compactStarted:   compactStarted,
		compactRelease:   compactRelease,
		appendCompaction: true,
	}
	s.chatSender = sender

	errs := make(chan error, 2)
	go func() {
		errs <- s.sendSessionChat(context.Background(), resolved, chat.Request{Message: "first"})
	}()
	select {
	case <-compactStarted:
	case <-time.After(time.Second):
		t.Fatal("first request did not start compaction")
	}
	go func() {
		errs <- s.sendSessionChat(context.Background(), resolved, chat.Request{Message: "second"})
	}()

	select {
	case err := <-errs:
		t.Fatalf("a request completed before the active compaction was released: %v", err)
	case <-time.After(50 * time.Millisecond):
	}
	close(compactRelease)
	for range 2 {
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	}
	if calls := sender.compactCallCount(); calls != 1 {
		t.Fatalf("Compact calls = %d, want one shared compaction", calls)
	}
	for i, want := range []string{"compact", "send:first", "send:second"} {
		select {
		case got := <-sender.events:
			if got != want {
				t.Fatalf("event %d = %q, want %q", i, got, want)
			}
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for event %q", want)
		}
	}
}

func TestConcurrentForceCompactRequestsReuseCompletedCompaction(t *testing.T) {
	s := newTestServer(t)
	resolved := writeHighUsageLocalSession(t, s)
	compactStarted := make(chan struct{})
	compactRelease := make(chan struct{})
	sender := &compactFakeSender{
		compactStarted:   compactStarted,
		compactRelease:   compactRelease,
		appendCompaction: true,
	}
	s.chatSender = sender

	errs := make(chan error, 2)
	go func() {
		errs <- s.forceCompact(context.Background(), resolved.Session.ID, nil)
	}()
	select {
	case <-compactStarted:
	case <-time.After(time.Second):
		t.Fatal("first Force Compact request did not start")
	}
	secondStarted := make(chan struct{})
	go func() {
		close(secondStarted)
		errs <- s.forceCompact(context.Background(), resolved.Session.ID, nil)
	}()
	<-secondStarted
	select {
	case err := <-errs:
		t.Fatalf("a Force Compact request completed while the first was still active: %v", err)
	case <-time.After(50 * time.Millisecond):
	}
	close(compactRelease)
	for range 2 {
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	}
	if calls := sender.compactCallCount(); calls != 1 {
		t.Fatalf("Compact calls = %d, want one shared compaction", calls)
	}
}

func TestSessionOperationWaitHonorsCancellationAndReleasesItsEntry(t *testing.T) {
	s := &Server{}
	release, err := s.acquireSessionOperation(context.Background(), "session.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := s.acquireSessionOperation(ctx, "session.jsonl"); !errors.Is(err, context.Canceled) {
		t.Fatalf("waiting operation error = %v, want context.Canceled", err)
	}
	release()

	release, err = s.acquireSessionOperation(context.Background(), "session.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	release()
	s.sessionOperations.mu.Lock()
	remaining := len(s.sessionOperations.entries)
	s.sessionOperations.mu.Unlock()
	if remaining != 0 {
		t.Fatalf("session operation entries = %d, want 0 after all callers release", remaining)
	}
}

func TestForceCompactInterruptsRunningLocalSession(t *testing.T) {
	s := newTestServer(t)
	path := writeSessionFile(t, s.sessionsDir, "--project--", "session.jsonl")
	_ = path
	if _, err := s.saveSessionMode("session.jsonl", sessionModeLocal, modelModeMetadata{
		Provider: "custom", ID: "qwen", ContextWindow: 100000,
	}); err != nil {
		t.Fatal(err)
	}
	sender := &compactFakeSender{
		fakeSender: fakeSender{status: workers.WorkerStatus{State: workers.WorkerStateRunning}},
		events:     make(chan string, 2),
	}
	s.chatSender = sender
	req := httptest.NewRequest(http.MethodPost, "/api/force-compact?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleForceCompact(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d: %s", w.Code, w.Body.String())
	}
	if first, second := <-sender.events, <-sender.events; first != "abort" || second != "compact" {
		t.Fatalf("events = %q, %q; want abort, compact", first, second)
	}
}

func TestHandleChatRejectsUnknownSession(t *testing.T) {
	s := &Server{sessionsDir: t.TempDir(), chatSender: &fakeSender{}}
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	mw.WriteField("message", "hello")
	mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/api/chat?id=missing.jsonl", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	s.handleChat(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", w.Code)
	}
}

func TestHandleChatRejectsBrokenSession(t *testing.T) {
	root := t.TempDir()
	dir := root + "/--tmp-project--"
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dir+"/session.jsonl", []byte("{\"type\":\"session\",\"version\":3,\"id\":\"sid\",\"timestamp\":\"2026-05-06T00:00:00.000Z\",\"cwd\":\"/definitely/missing/path\"}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	s := &Server{sessionsDir: root, chatSender: &fakeSender{}}
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	mw.WriteField("message", "hello")
	mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/api/chat?id=session.jsonl", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	s.handleChat(w, req)
	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, want 409; body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "working directory no longer exists") {
		t.Fatalf("body = %q", w.Body.String())
	}
}

func TestHandleWorkerStatusDefaultsIdle(t *testing.T) {
	s := &Server{sessionsDir: t.TempDir(), chatSender: &fakeSender{}, now: time.Now}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"idle\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusUsesRecentSessionFileActivity(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	now := time.Date(2026, 5, 7, 21, 0, 0, 0, time.UTC)
	s := &Server{
		sessionsDir: root,
		chatSender:  &fakeSender{},
		fileMod:     map[string]time.Time{"session.jsonl": now.Add(-400 * time.Millisecond)},
		now:         func() time.Time { return now },
	}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()

	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"running\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusIgnoresStaleSessionFileActivity(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	now := time.Date(2026, 5, 7, 21, 0, 0, 0, time.UTC)
	s := &Server{
		sessionsDir: root,
		chatSender:  &fakeSender{},
		fileMod:     map[string]time.Time{"session.jsonl": now.Add(-10 * time.Second)},
		now:         func() time.Time { return now },
	}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()

	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"idle\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusSkipsGetStateWhenLocalStatusRunning(t *testing.T) {
	sender := &fakeSender{status: workers.WorkerStatus{State: workers.WorkerStateRunning}}
	s := &Server{sessionsDir: t.TempDir(), chatSender: sender}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()

	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if calls := sender.stateCalls(); calls != 0 {
		t.Fatalf("GetState calls = %d, want 0", calls)
	}
	if got := w.Body.String(); got != "{\"state\":\"running\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

// When an in-process chat worker has been resolved for the session
// (Model populated) and reports idle, the activity-window fallback must
// not override it — otherwise the Cancel button lingers after the
// assistant finishes because the JSONL write keeps the file mtime fresh.
func TestHandleWorkerStatusTrustsIdleWorkerOverRecentFileWrite(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	now := time.Date(2026, 5, 7, 21, 0, 0, 0, time.UTC)
	sender := &fakeSender{status: workers.WorkerStatus{State: workers.WorkerStateIdle, Model: "gpt-5.5"}}
	s := &Server{
		sessionsDir: root,
		chatSender:  sender,
		fileMod:     map[string]time.Time{"session.jsonl": now.Add(-100 * time.Millisecond)},
		now:         func() time.Time { return now },
	}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()

	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, `"state":"idle"`) {
		t.Fatalf("body = %q, want state=idle", body)
	}
}

func TestHandleCommandsRejectsNonGET(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	s := &Server{sessionsDir: root, chatSender: &fakeSender{}}
	req := httptest.NewRequest(http.MethodPost, "/api/commands?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleCommands(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", w.Code)
	}
}

func TestHandleCommandsReturnsWorkerCommands(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	sender := &fakeSender{
		commands:      []workers.SlashCommand{{Name: "skill:memory", Description: "mem", Source: "skill"}},
		commandsReady: true,
	}
	s := &Server{sessionsDir: root, chatSender: sender}
	req := httptest.NewRequest(http.MethodGet, "/api/commands?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleCommands(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	var body struct {
		Commands []workers.SlashCommand `json:"commands"`
		Ready    bool                   `json:"workerReady"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("body decode: %v", err)
	}
	if !body.Ready {
		t.Fatalf("workerReady = false, want true")
	}
	if len(body.Commands) != 1 || body.Commands[0].Name != "skill:memory" {
		t.Fatalf("commands = %#v", body.Commands)
	}
	if called, _, _ := sender.ensureWorkerInfo(); called {
		t.Fatalf("EnsureWorker called without ?load=1")
	}
}

func TestHandleCommandsReportsNotReadyWithoutWorker(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	s := &Server{sessionsDir: root, chatSender: &fakeSender{commandsReady: false}}
	req := httptest.NewRequest(http.MethodGet, "/api/commands?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleCommands(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); !strings.Contains(got, `"commands":[]`) || !strings.Contains(got, `"workerReady":false`) {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleCommandsLoadEnsuresWorker(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	sender := &fakeSender{commandsReady: true}
	s := &Server{sessionsDir: root, chatSender: sender}
	req := httptest.NewRequest(http.MethodGet, "/api/commands?id=session.jsonl&load=1", nil)
	w := httptest.NewRecorder()
	s.handleCommands(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if called, _, _ := sender.ensureWorkerInfo(); !called {
		t.Fatalf("EnsureWorker not called for ?load=1")
	}
}

func TestHandleCommandsDegradesOnQueryError(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	sender := &fakeSender{commandsReady: true, commandsErr: context.DeadlineExceeded}
	s := &Server{sessionsDir: root, chatSender: sender}
	req := httptest.NewRequest(http.MethodGet, "/api/commands?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleCommands(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (degraded)", w.Code)
	}
	if got := w.Body.String(); !strings.Contains(got, `"commands":[]`) {
		t.Fatalf("body = %q, want empty commands on error", got)
	}
}

func TestHandleSetThinkingLevelRequiresLevel(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	s := &Server{sessionsDir: root, chatSender: &fakeSender{}}
	req := httptest.NewRequest(http.MethodPost, "/api/set-thinking-level?id=session.jsonl", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleSetThinkingLevel(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
	if !strings.Contains(w.Body.String(), "level required") {
		t.Fatalf("body = %q", w.Body.String())
	}
}

func TestHandleSetThinkingLevelRejectsMissingSession(t *testing.T) {
	s := &Server{sessionsDir: t.TempDir(), chatSender: &fakeSender{}}
	req := httptest.NewRequest(http.MethodPost, "/api/set-thinking-level?id=missing.jsonl", strings.NewReader(`{"level":"high"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleSetThinkingLevel(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", w.Code)
	}
}

func TestHandleWorkerStatusUsesSessionStatusFile(t *testing.T) {
	root := t.TempDir()
	sessionsDir := filepath.Join(root, "sessions")
	statusDir := filepath.Join(root, "session-status")
	if err := os.MkdirAll(statusDir, 0755); err != nil {
		t.Fatal(err)
	}
	sessionID := "test-session.jsonl"
	status := map[string]any{
		"sessionId": sessionID,
		"state":     "running",
		"updatedAt": time.Now().UTC().Format(time.RFC3339),
	}
	data, _ := json.Marshal(status)
	if err := os.WriteFile(filepath.Join(statusDir, sessionID), data, 0644); err != nil {
		t.Fatal(err)
	}

	s := &Server{agentDir: root, sessionsDir: sessionsDir, chatSender: &fakeSender{}}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id="+sessionID, nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"running\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusIgnoresStaleSessionStatusFile(t *testing.T) {
	root := t.TempDir()
	sessionsDir := filepath.Join(root, "sessions")
	statusDir := filepath.Join(root, "session-status")
	if err := os.MkdirAll(statusDir, 0755); err != nil {
		t.Fatal(err)
	}
	sessionID := "test-session.jsonl"
	status := map[string]any{
		"sessionId": sessionID,
		"state":     "running",
		"updatedAt": time.Now().Add(-30 * time.Second).UTC().Format(time.RFC3339),
	}
	data, _ := json.Marshal(status)
	if err := os.WriteFile(filepath.Join(statusDir, sessionID), data, 0644); err != nil {
		t.Fatal(err)
	}

	s := &Server{sessionsDir: sessionsDir, chatSender: &fakeSender{}, now: time.Now}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id="+sessionID, nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"idle\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusFallsThroughForIdleStatusFile(t *testing.T) {
	root := t.TempDir()
	sessionsDir := filepath.Join(root, "sessions")
	statusDir := filepath.Join(root, "session-status")
	if err := os.MkdirAll(statusDir, 0755); err != nil {
		t.Fatal(err)
	}
	sessionID := "test-session.jsonl"
	status := map[string]any{
		"sessionId": sessionID,
		"state":     "idle",
		"updatedAt": time.Now().UTC().Format(time.RFC3339),
	}
	data, _ := json.Marshal(status)
	if err := os.WriteFile(filepath.Join(statusDir, sessionID), data, 0644); err != nil {
		t.Fatal(err)
	}

	s := &Server{agentDir: root, sessionsDir: sessionsDir, chatSender: &fakeSender{}, now: time.Now}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id="+sessionID, nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Body.String(); got != "{\"state\":\"idle\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestHandleWorkerStatusReturnsModelAndThinkingLevel(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	sender := &fakeSender{
		state: workers.WorkerStatus{
			State:         workers.WorkerStateIdle,
			Model:         "kimi-k2.6",
			ModelName:     "Kimi K2.6",
			ModelProvider: "opengo-work",
			ThinkingLevel: "medium",
		},
	}
	s := &Server{sessionsDir: root, chatSender: sender, now: time.Now}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	var got map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["model"] != "kimi-k2.6" {
		t.Fatalf("model = %q, want kimi-k2.6", got["model"])
	}
	if got["modelName"] != "Kimi K2.6" {
		t.Fatalf("modelName = %q, want Kimi K2.6", got["modelName"])
	}
	if got["modelProvider"] != "opengo-work" {
		t.Fatalf("modelProvider = %q, want opengo-work", got["modelProvider"])
	}
	if got["thinkingLevel"] != "medium" {
		t.Fatalf("thinkingLevel = %q, want medium", got["thinkingLevel"])
	}
}

func TestHandleWorkerStatusDoesNotSpawnWorkerWhenModelUnknown(t *testing.T) {
	root := t.TempDir()
	writeSessionFile(t, root, "test-project", "session.jsonl")
	sender := &fakeSender{
		ensureWorkerCh: make(chan struct{}, 1),
		state: workers.WorkerStatus{
			State:         workers.WorkerStateIdle,
			Model:         "kimi-k2.6",
			ModelProvider: "opengo-work",
			ThinkingLevel: "medium",
		},
	}
	s := &Server{sessionsDir: root, chatSender: sender, now: time.Now}
	req := httptest.NewRequest(http.MethodGet, "/api/worker-status?id=session.jsonl", nil)
	w := httptest.NewRecorder()
	s.handleWorkerStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if called, _, _ := sender.ensureWorkerInfo(); called {
		t.Fatal("worker-status should not prewarm/create workers")
	}
}

func TestHandleNewSessionPreinitializesWorker(t *testing.T) {
	root := t.TempDir()
	fake := &fakeSender{ensureWorkerCh: make(chan struct{}, 1)}
	s := &Server{
		sessionsDir: root,
		chatSender:  fake,
	}
	// A real absolute path: "/tmp/..." is not absolute on Windows.
	projectPath := filepath.Join(root, "test-project")
	req := httptest.NewRequest(http.MethodPost, "/api/new-session", strings.NewReader(`{"path":`+jsonString(projectPath)+`}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", w.Code, w.Body.String())
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["ok"] != true {
		t.Fatalf("ok = %v, want true", body["ok"])
	}
	id, _ := body["id"].(string)
	if id == "" {
		t.Fatal("missing id in response")
	}

	// Verify EnsureWorker was called
	select {
	case <-fake.ensureWorkerCh:
		called, sessionID, _ := fake.ensureWorkerInfo()
		if !called {
			t.Fatal("EnsureWorker not marked as called")
		}
		if sessionID == "" {
			t.Fatal("EnsureWorker called with empty sessionID")
		}
	case <-time.After(time.Second):
		t.Fatal("EnsureWorker was not called within 1s")
	}

	// Verify file was created
	projectDir := filepath.Join(root, sessions.EncodeProjectName(projectPath))
	entries, err := os.ReadDir(projectDir)
	if err != nil {
		t.Fatalf("expected project dir to exist: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("expected session file to be created")
	}
}

func TestHandleNewSessionCopiesSourceModelAndThinking(t *testing.T) {
	root := t.TempDir()
	_ = writeSessionFile(t, root, "--tmp--source--", "source.jsonl")
	fake := &fakeSender{state: workers.WorkerStatus{State: workers.WorkerStateIdle, ModelProvider: "openai", Model: "gpt-5", ThinkingLevel: "high"}}
	s := &Server{sessionsDir: root, chatSender: fake}

	projectPath := filepath.Join(root, "test-project")
	req := httptest.NewRequest(http.MethodPost, "/api/new-session", strings.NewReader(`{"path":`+jsonString(projectPath)+`,"sourceSessionId":"source.jsonl"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", w.Code, w.Body.String())
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	id, _ := body["id"].(string)
	if id == "" {
		t.Fatal("missing id in response")
	}

	waitForCondition(t, time.Second, func() bool {
		_, sessionID, _ := fake.ensureWorkerInfo()
		return sessionID == id
	})
	waitForCondition(t, time.Second, func() bool {
		modelSessionID, provider, modelID := fake.modelSelection()
		return modelSessionID == id && provider == "openai" && modelID == "gpt-5" && fake.thinkingSessionID() == id
	})
	if modelSessionID, provider, modelID := fake.modelSelection(); modelSessionID != id || provider != "openai" || modelID != "gpt-5" {
		t.Fatalf("worker model = (%q, %q, %q), want (%q, openai, gpt-5)", modelSessionID, provider, modelID, id)
	}
	projectDir := filepath.Join(root, sessions.EncodeProjectName(projectPath))
	data, err := os.ReadFile(filepath.Join(projectDir, id))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, `"type":"model_change"`) || !strings.Contains(content, `"implicit":true`) {
		t.Fatalf("new session file missing implicit model setting: %s", content)
	}
	if !strings.Contains(content, `"type":"thinking_level_change"`) || !strings.Contains(content, `"thinkingLevel":"high"`) {
		t.Fatalf("new session file missing implicit thinking setting: %s", content)
	}
}

func TestHandleNewSessionAppliesExplicitModelToWorker(t *testing.T) {
	root := t.TempDir()
	fake := &fakeSender{}
	s := &Server{sessionsDir: root, chatSender: fake}

	projectPath := filepath.Join(root, "test-project")
	body := `{"path":` + jsonString(projectPath) + `,"modelProvider":"anthropic","modelId":"claude-sonnet-4"}`
	req := httptest.NewRequest(http.MethodPost, "/api/new-session", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", w.Code, w.Body.String())
	}
	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	id, _ := response["id"].(string)
	waitForCondition(t, time.Second, func() bool {
		modelSessionID, provider, modelID := fake.modelSelection()
		return modelSessionID == id && provider == "anthropic" && modelID == "claude-sonnet-4"
	})
}

func TestHandleNewSessionWithoutChatSender(t *testing.T) {
	root := t.TempDir()
	s := &Server{sessionsDir: root}
	req := httptest.NewRequest(http.MethodPost, "/api/new-session", strings.NewReader(`{"path":`+jsonString(filepath.Join(root, "no-sender"))+`}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", w.Code, w.Body.String())
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["ok"] != true {
		t.Fatalf("ok = %v, want true", body["ok"])
	}
	id, _ := body["id"].(string)
	if id == "" {
		t.Fatal("missing id in response")
	}
}

func TestHandleNewSessionRejectsMissingPath(t *testing.T) {
	root := t.TempDir()
	s := &Server{sessionsDir: root}
	req := httptest.NewRequest(http.MethodPost, "/api/new-session", strings.NewReader(`{"path":""}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestHandleNewSessionRejectsGetMethod(t *testing.T) {
	root := t.TempDir()
	s := &Server{sessionsDir: root}
	req := httptest.NewRequest(http.MethodGet, "/api/new-session", nil)
	w := httptest.NewRecorder()
	s.handleNewSession(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusMethodNotAllowed)
	}
}

func waitForCondition(t *testing.T, timeout time.Duration, fn func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if fn() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	if !fn() {
		t.Fatal("condition not met before timeout")
	}
}
