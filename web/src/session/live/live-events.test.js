import { describe, expect, it, vi } from 'vitest';
import {
  createSessionEventSource,
  getSessionIdFromLocation,
  handleSessionReload,
  wireSessionEvents,
} from './live-events.js';

describe('live events', () => {
  it('gets session id and creates event source', () => {
    expect(getSessionIdFromLocation({ locationImpl: { search: '?id=a%20b&x=1' } })).toBe('a%20b');
    const EventSourceImpl = vi.fn();
    createSessionEventSource('a b', { EventSourceImpl });
    expect(EventSourceImpl).toHaveBeenCalledWith('/events?id=a%20b');
  });

  it('handles reload entries, title, and follow behavior', async () => {
    const entries = [{ id: 'a' }, { id: 'r', message: { role: 'toolResult' } }];
    const fetchImpl = vi.fn(() =>
      Promise.resolve(
        new Response(JSON.stringify({ name: 'New Title', entries }), { status: 200 }),
      ),
    );
    const entryState = { seen: new Set(), liveRendered: new Set() };
    const appendEntry = vi.fn((entry) => {
      entryState.seen.add(entry.id);
      return true;
    });
    const refresh = vi.fn();
    const updateStats = vi.fn();
    const updateTitle = vi.fn();
    const scrollAfterLayout = vi.fn();
    const onReloaded = vi.fn();

    const result = await handleSessionReload({
      sessionId: 's',
      fetchImpl,
      entryState,
      clearChatPreview: vi.fn(),
      appendEntry,
      upsertEntry: vi.fn(),
      refreshEntriesAffectedByToolResult: refresh,
      updateStats,
      updateTitle,
      isFollowing: () => true,
      scrollAfterLayout,
      onReloaded,
    });

    expect(fetchImpl).toHaveBeenCalledWith('/api/session?id=s');
    expect(result.newCount).toBe(2);
    expect(onReloaded).toHaveBeenCalledWith({ name: 'New Title', entries });
    expect(refresh).toHaveBeenCalledWith(entries[1], entries);
    expect(updateStats).toHaveBeenCalledWith(entries);
    expect(updateTitle).toHaveBeenCalledWith('New Title');
    expect(scrollAfterLayout).toHaveBeenCalledWith(true);
  });

  it('reconciles via the model in reactive mode (no DOM patching)', async () => {
    // No appendEntry/upsertEntry → the Svelte <SessionContent> owns #messages.
    // handleSessionReload only tracks new ids and flags them via onNewEntries.
    const entries = [{ id: 'a' }, { id: 'b' }];
    const fetchImpl = vi.fn(() =>
      Promise.resolve(new Response(JSON.stringify({ entries }), { status: 200 })),
    );
    const entryState = { seen: new Set(['a']), liveRendered: new Set() };
    const onReloaded = vi.fn();
    const onNewEntries = vi.fn();
    const scrollAfterLayout = vi.fn();
    const clearChatPreview = vi.fn();

    const result = await handleSessionReload({
      sessionId: 's',
      fetchImpl,
      entryState,
      clearChatPreview,
      isFollowing: () => true,
      scrollAfterLayout,
      onReloaded,
      onNewEntries,
    });

    expect(onReloaded).toHaveBeenCalledWith({ entries });
    expect(result.newCount).toBe(1);
    expect(entryState.seen.has('b')).toBe(true);
    expect(onNewEntries).toHaveBeenCalledWith(['b']);
    expect(clearChatPreview).toHaveBeenCalledWith(entries, ['b']);
    expect(scrollAfterLayout).toHaveBeenCalledWith(true);
  });

  it('scrolls instead of showing the button when at the bottom but follow flag is stale', async () => {
    // Regression: agent finishes while the viewport is pinned to the bottom, but
    // `following` was knocked false by a relayout clamp. The reload must use the
    // live scroll position and not pop the "scroll to bottom" button.
    const entries = [{ id: 'a' }, { id: 'b' }];
    const fetchImpl = vi.fn(() =>
      Promise.resolve(new Response(JSON.stringify({ entries }), { status: 200 })),
    );
    const entryState = { seen: new Set(['a']), liveRendered: new Set() };
    const scrollAfterLayout = vi.fn();
    const incrementPending = vi.fn();
    const showFollowButton = vi.fn();

    await handleSessionReload({
      sessionId: 's',
      fetchImpl,
      entryState,
      clearChatPreview: vi.fn(),
      isFollowing: () => false,
      isAtBottom: () => true,
      scrollAfterLayout,
      incrementPending,
      showFollowButton,
      onReloaded: vi.fn(),
    });

    expect(scrollAfterLayout).toHaveBeenCalledWith(true);
    expect(showFollowButton).not.toHaveBeenCalled();
    expect(incrementPending).not.toHaveBeenCalled();
  });

  it('shows the button when not following and not at the bottom', async () => {
    const entries = [{ id: 'a' }, { id: 'b' }];
    const fetchImpl = vi.fn(() =>
      Promise.resolve(new Response(JSON.stringify({ entries }), { status: 200 })),
    );
    const entryState = { seen: new Set(['a']), liveRendered: new Set() };
    const scrollAfterLayout = vi.fn();
    const incrementPending = vi.fn();
    const showFollowButton = vi.fn();

    await handleSessionReload({
      sessionId: 's',
      fetchImpl,
      entryState,
      clearChatPreview: vi.fn(),
      isFollowing: () => false,
      isAtBottom: () => false,
      scrollAfterLayout,
      incrementPending,
      showFollowButton,
      onReloaded: vi.fn(),
    });

    expect(showFollowButton).toHaveBeenCalledTimes(1);
    expect(incrementPending).toHaveBeenCalledWith(1);
    expect(scrollAfterLayout).not.toHaveBeenCalled();
  });

  it('wires event source messages', () => {
    const eventSource = { addEventListener: vi.fn() };
    const onReload = vi.fn();
    const onChatPreview = vi.fn();
    const onError = vi.fn();
    wireSessionEvents({ eventSource, onReload, onChatPreview, onError });
    eventSource.onmessage({ data: 'noop' });
    eventSource.onmessage({ data: 'reload' });
    expect(onReload).toHaveBeenCalledTimes(1);
    const previewHandler = eventSource.addEventListener.mock.calls[0][1];
    previewHandler({ data: JSON.stringify({ content: 'x' }) });
    expect(onChatPreview).toHaveBeenCalledWith({ content: 'x' });
    previewHandler({ data: '{bad' });
    expect(onError).toHaveBeenCalled();
  });

  it('reconciles on a chat-preview done (covers a dropped first-write reload)', () => {
    const eventSource = { addEventListener: vi.fn() };
    const onReload = vi.fn();
    const onChatPreview = vi.fn();
    wireSessionEvents({ eventSource, onReload, onChatPreview });
    const previewHandler = eventSource.addEventListener.mock.calls[0][1];

    previewHandler({ data: JSON.stringify({ content: 'streaming', done: false }) });
    expect(onReload).not.toHaveBeenCalled();

    previewHandler({ data: JSON.stringify({ content: 'final', done: true }) });
    expect(onChatPreview).toHaveBeenLastCalledWith({ content: 'final', done: true });
    expect(onReload).toHaveBeenCalledTimes(1);
  });

  it('dispatches pi-session-reload window event on reload', () => {
    const eventSource = { addEventListener: vi.fn() };
    const dispatched = [];
    const windowImpl = {
      dispatchEvent: (e) => {
        dispatched.push(e);
        return true;
      },
    };
    class FakeCustomEvent {
      constructor(type) {
        this.type = type;
      }
    }
    wireSessionEvents({
      eventSource,
      onReload: vi.fn(),
      onChatPreview: vi.fn(),
      windowImpl,
      CustomEventImpl: FakeCustomEvent,
    });
    eventSource.onmessage({ data: 'reload' });
    expect(dispatched.length).toBe(1);
    expect(dispatched[0].type).toBe('pi-session-reload');
    eventSource.onmessage({ data: 'noop' });
    expect(dispatched.length).toBe(1);
  });

  it('reports local watchdog recovery events', () => {
    const eventSource = { addEventListener: vi.fn() };
    const onLocalRecovery = vi.fn();
    const onError = vi.fn();
    wireSessionEvents({
      eventSource,
      onReload: vi.fn(),
      onChatPreview: vi.fn(),
      onLocalRecovery,
      onError,
    });
    const recoveryHandler = eventSource.addEventListener.mock.calls.find(
      ([name]) => name === 'local-recovery',
    )[1];

    recoveryHandler({ data: JSON.stringify({ reason: 'thinking-only-stop' }) });
    expect(onLocalRecovery).toHaveBeenCalledWith({ reason: 'thinking-only-stop' });
    recoveryHandler({ data: '{bad' });
    expect(onError).toHaveBeenCalledTimes(1);
  });
});
