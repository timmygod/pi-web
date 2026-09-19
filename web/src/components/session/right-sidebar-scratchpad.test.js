import { afterEach, describe, expect, it, vi } from 'vitest';
import { createScratchpadController } from './right-sidebar-scratchpad.js';

function renderScratchpad(value = '') {
  document.body.innerHTML = `
    <textarea id="scratchpad-textarea"></textarea>
    <span id="scratchpad-status"></span>
  `;
  const textarea = document.getElementById('scratchpad-textarea');
  const statusEl = document.getElementById('scratchpad-status');
  textarea.value = value;
  return { textarea, statusEl };
}

afterEach(() => {
  document.body.innerHTML = '';
  vi.restoreAllMocks();
});

describe('createScratchpadController', () => {
  it('loads scratchpad content and treats it as saved', async () => {
    const { textarea, statusEl } = renderScratchpad('initial');
    const fetchImpl = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ content: 'server notes' }),
    });
    const scratchpad = createScratchpadController({
      projectPath: '/proj a',
      textarea,
      statusEl,
      fetchImpl,
    });

    await scratchpad.load();
    await scratchpad.save();

    expect(fetchImpl).toHaveBeenCalledTimes(1);
    expect(fetchImpl).toHaveBeenCalledWith('/api/scratchpad?project=%2Fproj%20a');
    expect(textarea.value).toBe('server notes');
    expect(statusEl.textContent).toBe('Saved');
    expect(statusEl.className).toBe('scratchpad-status saved');
  });

  it('posts changed content and updates the saved baseline', async () => {
    const { textarea, statusEl } = renderScratchpad('initial');
    const fetchImpl = vi.fn().mockResolvedValue({ ok: true, json: async () => ({}) });
    const scratchpad = createScratchpadController({
      projectPath: '/proj',
      textarea,
      statusEl,
      fetchImpl,
    });

    scratchpad.adoptCurrentValue();
    textarea.value = 'changed';
    await scratchpad.save();
    await scratchpad.save();

    expect(fetchImpl).toHaveBeenCalledTimes(1);
    expect(fetchImpl).toHaveBeenCalledWith('/api/scratchpad', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ project: '/proj', content: 'changed' }),
    });
    expect(statusEl.textContent).toBe('Saved');
    expect(statusEl.className).toBe('scratchpad-status saved');
  });

  it('debounces input saves and cleanup removes the listener', async () => {
    const { textarea, statusEl } = renderScratchpad('initial');
    const fetchImpl = vi.fn().mockResolvedValue({ ok: true, json: async () => ({}) });
    const clearTimeoutImpl = vi.fn();
    let pendingSave = null;
    const setTimeoutImpl = vi.fn((callback) => {
      pendingSave = callback;
      return 12;
    });
    const scratchpad = createScratchpadController({
      projectPath: '/proj',
      textarea,
      statusEl,
      fetchImpl,
      setTimeoutImpl,
      clearTimeoutImpl,
      saveDelayMs: 250,
    });

    scratchpad.adoptCurrentValue();
    const cleanup = scratchpad.bind();
    textarea.value = 'queued';
    textarea.dispatchEvent(new Event('input'));

    expect(statusEl.textContent).toBe('Saving…');
    expect(statusEl.className).toBe('scratchpad-status saving');
    expect(setTimeoutImpl).toHaveBeenCalledWith(expect.any(Function), 250);
    expect(fetchImpl).not.toHaveBeenCalled();

    await pendingSave();
    expect(fetchImpl).toHaveBeenCalledTimes(1);

    cleanup();
    textarea.value = 'ignored';
    textarea.dispatchEvent(new Event('input'));
    expect(setTimeoutImpl).toHaveBeenCalledTimes(1);
    expect(clearTimeoutImpl).toHaveBeenCalledWith(12);
  });

  it('reports load and save failures', async () => {
    const { textarea, statusEl } = renderScratchpad('initial');
    const fetchImpl = vi
      .fn()
      .mockResolvedValueOnce({ ok: false, status: 500 })
      .mockResolvedValueOnce({ ok: false, status: 500 });
    const scratchpad = createScratchpadController({
      projectPath: '/proj',
      textarea,
      statusEl,
      fetchImpl,
    });

    await scratchpad.load();
    expect(statusEl.textContent).toBe('Load failed');
    expect(statusEl.className).toBe('scratchpad-status');

    textarea.value = 'changed';
    await scratchpad.save();
    expect(statusEl.textContent).toBe('Save failed');
    expect(statusEl.className).toBe('scratchpad-status');
  });

  it('applies remote content when clean and skips when dirty', () => {
    const { textarea, statusEl } = renderScratchpad('saved');
    const scratchpad = createScratchpadController({
      projectPath: '/proj',
      textarea,
      statusEl,
    });
    scratchpad.adoptCurrentValue();

    expect(scratchpad.applyRemote('from skill')).toBe(true);
    expect(textarea.value).toBe('from skill');

    textarea.value = 'local edit';
    expect(scratchpad.isDirty()).toBe(true);
    expect(scratchpad.applyRemote('ignored')).toBe(false);
    expect(textarea.value).toBe('local edit');
  });
});
