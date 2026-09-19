// Shared `/events` subscriptions for named SSE events (schedules, scratchpad,
// settings). Consumers share one EventSource per topic: browsers cap parallel
// HTTP/1.1 connections per host at six, and a session tab already holds the
// per-session stream, so a stream per consumer starves fetches once a couple of
// tabs are open.
function parseJSON(data) {
  try {
    return JSON.parse(data);
  } catch {
    return null;
  }
}

const streamsByImpl = new Map();

function topicsFor(EventSourceImpl) {
  let topics = streamsByImpl.get(EventSourceImpl);
  if (!topics) {
    topics = new Map();
    streamsByImpl.set(EventSourceImpl, topics);
  }
  return topics;
}

function attachDispatcher(entry, event) {
  if (entry.attached.has(event)) return;
  entry.attached.add(event);
  entry.stream.addEventListener(event, (message) => {
    const payload = parseJSON(message.data);
    for (const handler of entry.listeners.get(event) ?? []) handler(payload);
  });
}

function openStream(EventSourceImpl, topic, entry) {
  if (!entry.stream) {
    entry.stream = new EventSourceImpl(`/events?id=${encodeURIComponent(topic)}`);
    entry.attached = new Set();
    bindPageLifecycle(EventSourceImpl, topic, entry);
  }
  for (const event of entry.listeners.keys()) attachDispatcher(entry, event);
}

function closeStream(entry) {
  if (!entry.stream) return;
  entry.stream.close();
  entry.stream = null;
  entry.attached = new Set();
}

// `beforeunload` would opt the page out of the bfcache; `pagehide`/`pageshow`
// drop and restore the stream without hurting back/forward navigation.
function bindPageLifecycle(EventSourceImpl, topic, entry) {
  const windowImpl = entry.windowImpl;
  if (!windowImpl?.addEventListener || entry.pagehideHandler) return;
  entry.pagehideHandler = () => closeStream(entry);
  entry.pageshowHandler = () => {
    if (!entry.stream) openStream(EventSourceImpl, topic, entry);
  };
  windowImpl.addEventListener('pagehide', entry.pagehideHandler);
  windowImpl.addEventListener('pageshow', entry.pageshowHandler);
}

function unbindPageLifecycle(entry) {
  const windowImpl = entry.windowImpl;
  if (!windowImpl?.removeEventListener) return;
  if (entry.pagehideHandler) windowImpl.removeEventListener('pagehide', entry.pagehideHandler);
  if (entry.pageshowHandler) windowImpl.removeEventListener('pageshow', entry.pageshowHandler);
  entry.pagehideHandler = null;
  entry.pageshowHandler = null;
}

/**
 * Subscribe to one named SSE event on the shared `/events` stream for `topic`.
 * Returns the same `{ connect, cleanup }` shape as createStatusEvents.
 */
export function createAppEvents({
  event,
  topic = '__all__',
  EventSourceImpl = globalThis.EventSource,
  windowImpl = globalThis.window,
  onEvent = () => {},
} = {}) {
  let subscribed = false;

  function connect() {
    if (subscribed || !EventSourceImpl || !event) return;
    subscribed = true;
    const topics = topicsFor(EventSourceImpl);
    let entry = topics.get(topic);
    if (!entry) {
      entry = {
        stream: null,
        attached: new Set(),
        listeners: new Map(),
        windowImpl,
        pagehideHandler: null,
        pageshowHandler: null,
      };
      topics.set(topic, entry);
    }
    let handlers = entry.listeners.get(event);
    if (!handlers) {
      handlers = new Set();
      entry.listeners.set(event, handlers);
    }
    handlers.add(onEvent);
    openStream(EventSourceImpl, topic, entry);
  }

  function cleanup() {
    if (!subscribed) return;
    subscribed = false;
    const topics = streamsByImpl.get(EventSourceImpl);
    const entry = topics?.get(topic);
    if (!entry) return;
    const handlers = entry.listeners.get(event);
    handlers?.delete(onEvent);
    if (handlers && handlers.size === 0) entry.listeners.delete(event);
    if (entry.listeners.size > 0) return;
    closeStream(entry);
    unbindPageLifecycle(entry);
    topics.delete(topic);
    if (topics.size === 0) streamsByImpl.delete(EventSourceImpl);
  }

  return { connect, cleanup };
}
