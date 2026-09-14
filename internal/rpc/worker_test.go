package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"time"

	"pi-web/internal/chat"
	"pi-web/internal/workers"
)

type nopWriteCloser struct{ w io.Writer }

func (n nopWriteCloser) Write(p []byte) (int, error) { return n.w.Write(p) }
func (n nopWriteCloser) Close() error                { return nil }

type writeFunc func([]byte) (int, error)

func (f writeFunc) Write(p []byte) (int, error) { return f(p) }

func waitForPending(t *testing.T, w *piRPCWorker, id string) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		w.mu.Lock()
		_, ok := w.pending[id]
		w.mu.Unlock()
		if ok {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("pending request %q never registered", id)
}

func TestStatusReportsRunningDuringRecentStreamActivity(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"hello"}}`)

	if got := w.Status(); got.State != workers.WorkerStateRunning {
		t.Fatalf("status = %q, want running", got.State)
	}
}

func TestStatusStaysRunningAfterAgentEndUntilSettled(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateRunning},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"hello"}}`)
	w.handleRPCLine(`{"type":"agent_end"}`)

	if got := w.Status(); got.State != workers.WorkerStateRunning {
		t.Fatalf("status = %q, want running while post-agent work is pending", got.State)
	}

	w.handleRPCLine(`{"type":"agent_settled"}`)

	if got := w.Status(); got.State != workers.WorkerStateIdle {
		t.Fatalf("status = %q, want idle after agent settles", got.State)
	}
}

func TestWaitSettledDoesNotReturnAtAgentEnd(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateRunning},
		pending: make(map[string]chan response),
		settled: make(chan struct{}),
	}
	done := make(chan error, 1)
	go func() { done <- w.WaitSettled(context.Background()) }()
	w.handleRPCLine(`{"type":"agent_end"}`)
	select {
	case <-done:
		t.Fatal("WaitSettled returned at agent_end")
	case <-time.After(20 * time.Millisecond):
	}
	w.handleRPCLine(`{"type":"agent_settled"}`)
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("WaitSettled did not return after agent_settled")
	}
}

func TestWaitSettledReportsWorkerFailure(t *testing.T) {
	settled := make(chan struct{})
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateRunning},
		pending: make(map[string]chan response),
		settled: settled,
	}
	done := make(chan error, 1)
	go func() { done <- w.WaitSettled(context.Background()) }()
	w.mu.Lock()
	w.status = workers.WorkerStatus{State: workers.WorkerStateError, Error: "worker exited"}
	close(settled)
	w.settled = nil
	w.mu.Unlock()
	select {
	case err := <-done:
		if err == nil || err.Error() != "worker exited" {
			t.Fatalf("WaitSettled error = %v, want worker exited", err)
		}
	case <-time.After(time.Second):
		t.Fatal("WaitSettled did not return after worker failure")
	}
}

func TestInteractiveExtensionUIRequestIsCancelledAfterTimeout(t *testing.T) {
	originalTimeout := extensionUIRequestTimeout
	extensionUIRequestTimeout = 10 * time.Millisecond
	defer func() { extensionUIRequestTimeout = originalTimeout }()

	writes := make(chan []byte, 1)
	w := &piRPCWorker{
		stdin: nopWriteCloser{writeFunc(func(p []byte) (int, error) {
			writes <- append([]byte(nil), p...)
			return len(p), nil
		})},
		pending:  make(map[string]chan response),
		uiTimers: make(map[string]*time.Timer),
		status:   workers.WorkerStatus{State: workers.WorkerStateRunning},
	}
	w.handleRPCLine(`{"type":"extension_ui_request","id":"ui-1","method":"confirm","title":"Permission","message":"Allow?"}`)

	var data []byte
	select {
	case data = <-writes:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timed out waiting for timeout response")
	}
	var response map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(data), &response); err != nil {
		t.Fatalf("timeout response is invalid JSON: %v", err)
	}
	if response["type"] != "extension_ui_response" || response["id"] != "ui-1" || response["cancelled"] != true {
		t.Fatalf("timeout response = %#v", response)
	}
}

func TestFireAndForgetExtensionUIRequestDoesNotNeedResponse(t *testing.T) {
	var buf bytes.Buffer
	w := &piRPCWorker{
		stdin:    nopWriteCloser{&buf},
		pending:  make(map[string]chan response),
		uiTimers: make(map[string]*time.Timer),
	}
	w.handleRPCLine(`{"type":"extension_ui_request","id":"ui-2","method":"notify","message":"hello"}`)
	time.Sleep(20 * time.Millisecond)
	if buf.Len() != 0 {
		t.Fatalf("fire-and-forget request wrote response: %q", buf.String())
	}
}

func TestStatusDoesNotStayRunningAfterStreamActivityExpires(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"hello"}}`)
	time.Sleep(2200 * time.Millisecond)

	if got := w.Status(); got.State != workers.WorkerStateIdle {
		t.Fatalf("status = %q, want idle after stream activity expires", got.State)
	}
}

func TestHandleRPCLineTracksTurnEndAsStreamActivity(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{"type":"turn_end"}`)

	if got := w.Status(); got.State != workers.WorkerStateRunning {
		t.Fatalf("status = %q, want running", got.State)
	}
}

func TestHandleRPCLineEmitsStreamPreviewCallbacks(t *testing.T) {
	var previews []StreamPreview
	w := &piRPCWorker{
		status:        workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending:       make(map[string]chan response),
		streamSink:    func(preview StreamPreview) { previews = append(previews, preview) },
		streamPreview: &streamPreviewAccumulator{},
	}

	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"hel"}}`)
	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"lo"}}`)

	if len(previews) != 2 {
		t.Fatalf("previews = %d, want 2", len(previews))
	}
	if previews[0].Content != "hel" || previews[0].Done {
		t.Fatalf("first preview = %+v", previews[0])
	}
	if previews[1].Content != "hello" || previews[1].Done {
		t.Fatalf("second preview = %+v", previews[1])
	}
}

func TestHandleRPCLineEmitsDonePreviewOnAgentEnd(t *testing.T) {
	var previews []StreamPreview
	w := &piRPCWorker{
		status:        workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending:       make(map[string]chan response),
		streamSink:    func(preview StreamPreview) { previews = append(previews, preview) },
		streamPreview: &streamPreviewAccumulator{},
	}

	w.handleRPCLine(`{"type":"message_update","assistantMessageEvent":{"type":"text_delta","delta":"hello"}}`)
	w.handleRPCLine(`{"type":"agent_end"}`)

	if len(previews) != 2 {
		t.Fatalf("previews = %d, want 2", len(previews))
	}
	if previews[1].Content != "hello" || !previews[1].Done {
		t.Fatalf("done preview = %+v", previews[1])
	}
}

func TestHandleRPCLineTracksMessageEndAsStreamActivity(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{"type":"message_end"}`)

	if got := w.Status(); got.State != workers.WorkerStateRunning {
		t.Fatalf("status = %q, want running", got.State)
	}
}

func TestGetCommandsReturnsCachedWithoutRPC(t *testing.T) {
	w := &piRPCWorker{
		pending:        make(map[string]chan response),
		commands:       []workers.SlashCommand{{Name: "skill:memory", Source: "skill"}},
		commandsCached: true,
	}
	// stdin is nil: the cache path must not attempt any RPC write.
	got, err := w.GetCommands(context.Background())
	if err != nil {
		t.Fatalf("GetCommands error: %v", err)
	}
	if len(got) != 1 || got[0].Name != "skill:memory" {
		t.Fatalf("got = %#v", got)
	}
}

func TestGetCommandsParsesResponseAndCaches(t *testing.T) {
	var buf bytes.Buffer
	w := &piRPCWorker{
		stdin:   nopWriteCloser{&buf},
		pending: make(map[string]chan response),
	}

	type result struct {
		cmds []workers.SlashCommand
		err  error
	}
	resCh := make(chan result, 1)
	go func() {
		cmds, err := w.GetCommands(context.Background())
		resCh <- result{cmds, err}
	}()

	waitForPending(t, w, "req-1")
	w.handleRPCLine(`{"type":"response","id":"req-1","command":"get_commands","success":true,"data":{"commands":[{"name":"skill:memory","description":"mem","source":"skill"},{"name":"btw","description":"side chat","source":"extension"}]}}`)

	got := <-resCh
	if got.err != nil {
		t.Fatalf("GetCommands error: %v", got.err)
	}
	if len(got.cmds) != 2 {
		t.Fatalf("commands = %#v", got.cmds)
	}
	if got.cmds[0].Name != "skill:memory" || got.cmds[0].Source != "skill" || got.cmds[0].Description != "mem" {
		t.Fatalf("first command = %#v", got.cmds[0])
	}

	// Second call must hit the cache: no further RPC write to stdin.
	buf.Reset()
	cached, err := w.GetCommands(context.Background())
	if err != nil {
		t.Fatalf("cached GetCommands error: %v", err)
	}
	if len(cached) != 2 {
		t.Fatalf("cached commands = %#v", cached)
	}
	if buf.Len() != 0 {
		t.Fatalf("cache hit wrote to stdin: %q", buf.String())
	}
}

func TestHandleRPCLineIgnoresMalformedJSON(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	w.handleRPCLine(`{not-json}`)

	if got := w.Status(); got.State != workers.WorkerStateIdle {
		t.Fatalf("status = %q, want idle", got.State)
	}
}

func TestHandleRPCLineTracksThinkingAndTextStreamEvents(t *testing.T) {
	w := &piRPCWorker{
		status:  workers.WorkerStatus{State: workers.WorkerStateIdle},
		pending: make(map[string]chan response),
	}

	for _, line := range []string{
		`{"type":"message_update","assistantMessageEvent":{"type":"thinking_end"}}`,
		`{"type":"message_update","assistantMessageEvent":{"type":"text_start"}}`,
		`{"type":"message_update","assistantMessageEvent":{"type":"text_end","content":"done"}}`,
	} {
		w.handleRPCLine(line)
		if got := w.Status(); got.State != workers.WorkerStateRunning {
			t.Fatalf("line %s => status = %q, want running", strings.TrimSpace(line), got.State)
		}
	}
}

func TestCancelledRPCDoesNotWriteAfterWaitingForWriter(t *testing.T) {
	var buf bytes.Buffer
	w := &piRPCWorker{stdin: nopWriteCloser{&buf}, pending: make(map[string]chan response)}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w.writeMu.Lock()
	done := make(chan error, 1)
	go func() {
		done <- w.sendAndAwait(ctx, BuildPromptCommand("cancelled", chat.Request{Message: "continue if possible"}, false))
	}()
	waitForPending(t, w, "cancelled")
	cancel()
	w.writeMu.Unlock()
	select {
	case err := <-done:
		if err != context.Canceled {
			t.Fatalf("error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("cancelled RPC did not finish")
	}
	if buf.Len() != 0 {
		t.Fatalf("cancelled continuation was sent: %s", buf.String())
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if len(w.pending) != 0 {
		t.Fatal("cancelled RPC leaked pending response")
	}
}
