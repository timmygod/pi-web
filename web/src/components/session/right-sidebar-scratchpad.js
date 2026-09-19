export function createScratchpadController({
  projectPath = '',
  textarea,
  statusEl,
  fetchImpl = fetch,
  setTimeoutImpl = setTimeout,
  clearTimeoutImpl = clearTimeout,
  saveDelayMs = 1000,
} = {}) {
  let saveTimer = null;
  let lastSaved = textarea ? textarea.value : '';

  function setStatus(text, cls) {
    if (!statusEl) return;
    statusEl.textContent = text;
    statusEl.className = `scratchpad-status ${cls || ''}`.trim();
  }

  async function load() {
    if (!projectPath || !textarea) return;
    try {
      const res = await fetchImpl(`/api/scratchpad?project=${encodeURIComponent(projectPath)}`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      const content = data.content ?? '';
      textarea.value = content;
      lastSaved = content;
      setStatus('Saved', 'saved');
    } catch {
      setStatus('Load failed', '');
    }
  }

  async function save() {
    if (!projectPath || !textarea) return;
    const content = textarea.value;
    if (content === lastSaved) return;
    setStatus('Saving…', 'saving');
    try {
      const res = await fetchImpl('/api/scratchpad', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ project: projectPath, content }),
      });
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      lastSaved = content;
      setStatus('Saved', 'saved');
    } catch {
      setStatus('Save failed', '');
    }
  }

  function onInput() {
    setStatus('Saving…', 'saving');
    clearTimeoutImpl(saveTimer);
    saveTimer = setTimeoutImpl(save, saveDelayMs);
  }

  function adoptCurrentValue() {
    if (textarea) lastSaved = textarea.value;
  }

  function isDirty() {
    return !!(textarea && textarea.value !== lastSaved);
  }

  function applyRemote(content) {
    if (!textarea) return false;
    if (isDirty()) return false;
    const next = content ?? '';
    // Our own debounced save echoes back over SSE; reassigning value would
    // move the caret for no reason.
    if (textarea.value === next) return true;
    textarea.value = next;
    lastSaved = next;
    setStatus('Saved', 'saved');
    return true;
  }

  function bind() {
    textarea?.addEventListener('input', onInput);
    return () => {
      textarea?.removeEventListener('input', onInput);
      clearTimeoutImpl(saveTimer);
    };
  }

  return {
    load,
    save,
    setStatus,
    adoptCurrentValue,
    isDirty,
    applyRemote,
    bind,
  };
}
