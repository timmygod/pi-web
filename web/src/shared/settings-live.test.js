import { describe, expect, it, vi } from 'vitest';
import { applyRemoteSettings } from './settings-live.js';

function fakeStorage(seed = {}) {
  const map = new Map(Object.entries(seed));
  return {
    getItem: (k) => (map.has(k) ? map.get(k) : null),
    setItem: (k, v) => map.set(k, String(v)),
  };
}

function fakeDocument() {
  const props = {};
  return {
    documentElement: {
      dataset: {},
      style: { setProperty: vi.fn((k, v) => (props[k] = v)), _props: props },
    },
    cookie: '',
  };
}

describe('applyRemoteSettings', () => {
  it('applies theme and fonts without reloading', () => {
    const storage = fakeStorage({ 'pi-web:v1:locale': 'en' });
    const documentImpl = fakeDocument();
    const reload = vi.fn();
    applyRemoteSettings(
      {
        settings: {
          'pi-web-theme': 'nord',
          'pi-web:v1:locale': 'en',
          'pi-web:v1:font-ui': 'sans',
        },
      },
      { storage, documentImpl, windowImpl: {}, reload },
    );
    expect(documentImpl.documentElement.dataset.theme).toBe('nord');
    expect(documentImpl.documentElement.style._props['--font-sans']).toContain('Inter');
    expect(reload).not.toHaveBeenCalled();
  });

  it('reloads when locale changes, matching the Settings language picker', () => {
    const storage = fakeStorage({ 'pi-web:v1:locale': 'en' });
    const reload = vi.fn();
    applyRemoteSettings(
      { settings: { 'pi-web:v1:locale': 'ja' } },
      { storage, documentImpl: fakeDocument(), windowImpl: {}, reload },
    );
    expect(storage.getItem('pi-web:v1:locale')).toBe('ja');
    expect(reload).toHaveBeenCalledTimes(1);
  });

  it('does not reload when locale is unchanged', () => {
    const storage = fakeStorage({ 'pi-web:v1:locale': 'ja' });
    const reload = vi.fn();
    applyRemoteSettings(
      { settings: { 'pi-web:v1:locale': 'ja', 'pi-web-theme': 'dark' } },
      { storage, documentImpl: fakeDocument(), windowImpl: {}, reload },
    );
    expect(reload).not.toHaveBeenCalled();
  });
});
