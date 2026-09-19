import { describe, expect, it, vi, beforeEach } from 'vitest';
import { createAppEvents } from './app-events.js';

class FakeEventSource {
  constructor(url) {
    this.url = url;
    this.listeners = {};
    this.close = vi.fn();
    FakeEventSource.instances.push(this);
  }
  addEventListener(name, fn) {
    (this.listeners[name] ||= []).push(fn);
  }
  emit(name, data) {
    for (const fn of this.listeners[name] || []) fn({ data });
  }
}
FakeEventSource.instances = [];

function fakeWindow() {
  const handlers = {};
  return {
    handlers,
    addEventListener: (name, fn) => (handlers[name] ||= []).push(fn),
    removeEventListener: (name, fn) => {
      handlers[name] = (handlers[name] || []).filter((h) => h !== fn);
    },
    fire: (name) => {
      for (const fn of handlers[name] || []) fn();
    },
  };
}

function subscribe(event, onEvent, windowImpl = fakeWindow()) {
  const sub = createAppEvents({ event, onEvent, EventSourceImpl: FakeEventSource, windowImpl });
  sub.connect();
  return sub;
}

describe('createAppEvents', () => {
  beforeEach(() => {
    FakeEventSource.instances = [];
  });

  it('subscribes to __all__ and forwards parsed payloads', () => {
    const onEvent = vi.fn();
    const sub = subscribe('schedules', onEvent);

    const es = FakeEventSource.instances[0];
    expect(es.url).toBe('/events?id=__all__');

    es.emit('schedules', JSON.stringify({ action: 'created', id: 'abc' }));
    expect(onEvent).toHaveBeenCalledWith({ action: 'created', id: 'abc' });

    es.emit('schedules', 'not-json');
    expect(onEvent).toHaveBeenLastCalledWith(null);

    sub.cleanup();
    expect(es.close).toHaveBeenCalled();
  });

  it('shares one stream across events and closes it with the last listener', () => {
    const onSettings = vi.fn();
    const onScratchpad = vi.fn();
    const windowImpl = fakeWindow();
    const settings = subscribe('settings', onSettings, windowImpl);
    const scratchpad = subscribe('scratchpad', onScratchpad, windowImpl);

    expect(FakeEventSource.instances).toHaveLength(1);
    const es = FakeEventSource.instances[0];
    es.emit('settings', JSON.stringify({ settings: { 'pi-web-theme': 'nord' } }));
    es.emit('scratchpad', JSON.stringify({ project: '/p', content: 'hi' }));
    expect(onSettings).toHaveBeenCalledWith({ settings: { 'pi-web-theme': 'nord' } });
    expect(onScratchpad).toHaveBeenCalledWith({ project: '/p', content: 'hi' });

    settings.cleanup();
    expect(es.close).not.toHaveBeenCalled();
    es.emit('scratchpad', JSON.stringify({ project: '/p', content: 'still live' }));
    expect(onScratchpad).toHaveBeenLastCalledWith({ project: '/p', content: 'still live' });

    scratchpad.cleanup();
    expect(es.close).toHaveBeenCalled();
  });

  it('drops the stream on pagehide and reconnects on pageshow', () => {
    const onEvent = vi.fn();
    const windowImpl = fakeWindow();
    const sub = subscribe('settings', onEvent, windowImpl);

    windowImpl.fire('pagehide');
    expect(FakeEventSource.instances[0].close).toHaveBeenCalled();

    windowImpl.fire('pageshow');
    expect(FakeEventSource.instances).toHaveLength(2);
    FakeEventSource.instances[1].emit('settings', JSON.stringify({ settings: {} }));
    expect(onEvent).toHaveBeenCalledWith({ settings: {} });

    sub.cleanup();
  });
});
