import { describe, expect, it, vi } from 'vitest';
import { isLocalModel, renderModelList, setupModelSelector } from './chat/model-selector.js';
import {
  loadRecentModels,
  loadStarredModels,
  saveStarredModels,
  touchRecentModel,
} from '../../session/chat/chat-model-prefs.js';
import { t } from '../../shared/i18n.js';

function createDom() {
  const div = document.createElement('div');
  div.className = 'content-container';
  div.innerHTML = `
    <button id="pi-chat-model-label">Model</button>
    <div id="pi-chat-model-popup">
      <div class="pi-chat-model-header">
        <input id="pi-chat-model-search" />
        <button id="pi-chat-model-close" type="button">×</button>
      </div>
      <div id="pi-chat-model-list"></div>
    </div>
    <textarea id="pi-chat-message"></textarea>
  `;
  document.body.appendChild(div);
  return div;
}

function cleanupDom(el) {
  el.remove();
}

function makeModel(provider, id, extra = {}) {
  return { provider, id, name: id, ...extra };
}

describe('isLocalModel', () => {
  it('detects local providers', () => {
    expect(isLocalModel(makeModel('ollama', 'x'))).toBe(true);
    expect(isLocalModel(makeModel('LLama', 'x'))).toBe(true);
    expect(isLocalModel(makeModel('anthropic', 'x'))).toBe(false);
  });

  it('detects local baseUrl hosts', () => {
    expect(isLocalModel(makeModel('custom', 'x', { baseUrl: 'http://localhost:8080/v1' }))).toBe(
      true,
    );
    expect(isLocalModel(makeModel('custom', 'x', { baseUrl: 'http://192.168.1.10:11434' }))).toBe(
      true,
    );
    expect(isLocalModel(makeModel('custom', 'x', { baseUrl: 'http://home.local/v1' }))).toBe(true);
    expect(isLocalModel(makeModel('custom', 'x', { baseUrl: 'https://api.anthropic.com' }))).toBe(
      false,
    );
    expect(isLocalModel(makeModel('custom', 'x', { baseUrl: 'http://api.example.com/v1' }))).toBe(
      false,
    );
  });
});

describe('renderModelList', () => {
  const baseOpts = { t, starred: [], recent: [] };

  it('orders local, starred, recent, then cloud sections', () => {
    const models = [
      makeModel('openai', 'gpt-4o'),
      makeModel('anthropic', 'claude-sonnet'),
      makeModel('ollama', 'qwen3'),
      makeModel('anthropic', 'claude-3'),
    ];
    const html = renderModelList(models, {
      ...baseOpts,
      starred: ['anthropic\u0000claude-3'],
      recent: ['openai\u0000gpt-4o'],
    });
    const labels = [...html.matchAll(/model-section-label">([^<]+)</g)].map((m) => m[1]);
    expect(labels).toEqual(['Local models', 'Starred', 'Recently used', 'All models']);
    // star button present and marked starred (lucide svg) for the starred row;
    // the key is the NUL-joined canonical key (sessions.js)
    expect(html).toContain('data-star-key="anthropic\u0000claude-3"');
    expect(html).toMatch(/model-star starred[\s\S]*data-star-key="anthropic/);
    expect(html).toMatch(/class="star-icon"/);
  });

  it('omits empty sections', () => {
    const models = [makeModel('openai', 'gpt-4o')];
    const html = renderModelList(models, baseOpts);
    expect(html).toContain('All models');
    expect(html).not.toContain('Local models');
    expect(html).not.toContain('Starred');
    expect(html).not.toContain('Recently used');
  });

  it('filters by name or provider', () => {
    const models = [makeModel('openai', 'gpt-4o'), makeModel('anthropic', 'claude-sonnet')];
    const html = renderModelList(models, { ...baseOpts, filter: 'claude' });
    expect(html).toContain('claude-sonnet');
    expect(html).not.toContain('gpt-4o');
  });

  it('marks the selected model', () => {
    const models = [makeModel('openai', 'gpt-4o'), makeModel('anthropic', 'claude-3')];
    const html = renderModelList(models, {
      ...baseOpts,
      selectedModel: makeModel('anthropic', 'claude-3'),
    });
    expect(html).toContain('model-item selected');
    expect(html).toMatch(/data-provider="anthropic" data-model-id="claude-3"/);
  });
});

describe('chat-model-prefs', () => {
  it('round-trips starred and recents', () => {
    const storage = {
      _m: new Map(),
      getItem(k) {
        return this._m.has(k) ? this._m.get(k) : null;
      },
      setItem(k, v) {
        this._m.set(k, String(v));
      },
    };
    saveStarredModels(['a/b', 'c/d'], storage);
    expect(loadStarredModels(storage)).toEqual(['a/b', 'c/d']);

    touchRecentModel('x/y', storage);
    touchRecentModel('x/y', storage); // idempotent
    touchRecentModel('p/q', storage);
    const recents = loadRecentModels(storage);
    expect(recents).toEqual(['p/q', 'x/y']);
  });
});

describe('setupModelSelector', () => {
  function setup({ models = [], starStorage } = {}) {
    const el = createDom();
    const chatApi = {
      listModels: () => Promise.resolve({ ok: true, json: () => Promise.resolve({ models }) }),
      setModel: vi.fn(() => Promise.resolve({ ok: true, json: () => Promise.resolve({}) })),
    };
    const storage = starStorage ?? {
      _m: new Map(),
      getItem(k) {
        return this._m.has(k) ? this._m.get(k) : null;
      },
      setItem(k, v) {
        this._m.set(k, String(v));
      },
    };
    const api = setupModelSelector({
      documentImpl: document,
      windowImpl: { matchMedia: () => ({ matches: false }), localStorage: storage },
      sessionId: 's',
      chatApi,
      t,
      storage,
    });
    return { el, chatApi, api, storage };
  }

  it('returns { open, close } API', () => {
    const { el, api } = setup();
    expect(api).toHaveProperty('open');
    expect(api).toHaveProperty('close');
    cleanupDom(el);
  });

  it('open shows the overlay and focuses the search on fine-pointer devices', () => {
    const { el, api } = setup();
    const popup = document.getElementById('pi-chat-model-popup');
    const search = document.getElementById('pi-chat-model-search');
    search.focus = vi.fn();
    api.open();
    expect(popup.classList.contains('open')).toBe(true);
    expect(el.classList.contains('model-picker-open')).toBe(true);
    expect(search.focus).toHaveBeenCalled();
    cleanupDom(el);
  });

  it('close hides the overlay', () => {
    const { el, api } = setup();
    const popup = document.getElementById('pi-chat-model-popup');
    api.open();
    expect(popup.classList.contains('open')).toBe(true);
    api.close();
    expect(popup.classList.contains('open')).toBe(false);
    expect(el.classList.contains('model-picker-open')).toBe(false);
    cleanupDom(el);
  });

  it('close button closes the popup and re-focuses the textarea', () => {
    const { el, api } = setup();
    const popup = document.getElementById('pi-chat-model-popup');
    const closeBtn = document.getElementById('pi-chat-model-close');
    const textarea = document.getElementById('pi-chat-message');
    textarea.focus = vi.fn();
    api.open();
    expect(popup.classList.contains('open')).toBe(true);
    closeBtn.click();
    expect(popup.classList.contains('open')).toBe(false);
    expect(textarea.focus).toHaveBeenCalled();
    cleanupDom(el);
  });

  it('re-focuses the chat textarea when close(true) is called', () => {
    const { el, api } = setup();
    const textarea = document.getElementById('pi-chat-message');
    textarea.focus = vi.fn();
    api.open();
    api.close(true);
    expect(textarea.focus).toHaveBeenCalled();
    cleanupDom(el);
  });

  it('clicking a model row calls setModel and closes the popup', async () => {
    const { el, api, chatApi } = setup({
      models: [makeModel('openai', 'gpt-4o')],
    });
    // wait for async listModels
    await new Promise((r) => setTimeout(r, 0));
    api.open();
    const item = document.querySelector('.model-item[data-model-id="gpt-4o"]');
    expect(item).toBeTruthy();
    item.click();
    expect(chatApi.setModel).toHaveBeenCalledWith('s', {
      provider: 'openai',
      modelId: 'gpt-4o',
    });
    expect(document.getElementById('pi-chat-model-popup').classList.contains('open')).toBe(false);
    cleanupDom(el);
  });

  it('clicking the star toggles the star without selecting the model', async () => {
    const { el, api, chatApi, storage } = setup({
      models: [makeModel('openai', 'gpt-4o'), makeModel('anthropic', 'claude-3')],
    });
    await new Promise((r) => setTimeout(r, 0));
    api.open();
    const star = document.querySelector('.model-item[data-model-id="gpt-4o"] .model-star');
    expect(star).toBeTruthy();
    star.click();
    // star toggled on
    const updated = document.querySelector('.model-item[data-model-id="gpt-4o"] .model-star');
    expect(updated.classList.contains('starred')).toBe(true);
    expect(updated.querySelector('svg.star-icon')).toBeTruthy();
    // starred persisted
    expect(loadStarredModels(storage)).toEqual(['openai\u0000gpt-4o']);
    // model was NOT selected
    expect(chatApi.setModel).not.toHaveBeenCalled();
    // row itself is still in the list (now under Starred section)
    expect(document.querySelector('.model-item[data-model-id="gpt-4o"]')).toBeTruthy();
    cleanupDom(el);
  });
});
