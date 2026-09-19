// @mention path autocomplete for the chat composer. Typing "@" opens a popup
// listing files and folders from the session's working directory, fetched from
// GET /api/files. Filtering, ranking, and bounds live on the server; this module
// anchors the trigger, debounces requests, and inserts the chosen path.

export function parseAtTrigger(text, caret) {
  if (typeof text !== 'string') return null;
  if (caret == null) caret = text.length;
  let at = -1;
  for (let i = caret - 1; i >= 0; i -= 1) {
    const ch = text[i];
    if (ch === '@') {
      at = i;
      break;
    }
    if (/\s/.test(ch)) return null;
  }
  if (at < 0) return null;
  if (at > 0 && !/\s/.test(text[at - 1])) return null;
  return { query: text.slice(at + 1, caret), start: at, end: caret };
}

export function renderFileList(files, { escapeHtml = String, loading = false } = {}) {
  if (loading) return '<div class="slash-empty">Searching files...</div>';
  const list = Array.isArray(files) ? files : [];
  if (list.length === 0) return '<div class="slash-empty">No files match</div>';
  let html = '';
  list.forEach((f) => {
    const path = f.path || '';
    const display = f.isDir ? path + '/' : path;
    html +=
      `<button type="button" class="slash-item" data-insert="${escapeHtml(path)}" data-isdir="${f.isDir ? '1' : ''}">` +
      `<span class="slash-item-name">${escapeHtml(display)}</span></button>`;
  });
  return html;
}

export function setupMentionAutocomplete({
  documentImpl = document,
  windowImpl = window,
  sessionId,
  chatApi,
  escapeHtml = String,
  debounceMs = 120,
  setTimeoutImpl = (windowImpl || globalThis).setTimeout.bind(windowImpl || globalThis),
  clearTimeoutImpl = (windowImpl || globalThis).clearTimeout.bind(windowImpl || globalThis),
  AbortControllerImpl = (windowImpl || globalThis).AbortController,
} = {}) {
  const textarea = documentImpl.getElementById('pi-chat-message');
  const popup = documentImpl.getElementById('pi-chat-mention-popup');
  const list = documentImpl.getElementById('pi-chat-mention-list');
  if (!textarea || !popup || !list) return { handleKeydown: () => false };

  let trigger = null;
  let debounceTimer = null;
  let inflight = null;
  let reqSeq = 0;

  function isOpen() {
    return popup.style.display !== 'none' && popup.style.display !== '';
  }

  function items() {
    return list.querySelectorAll('.slash-item');
  }

  function setActive(index) {
    const all = items();
    const clamped = Math.max(0, Math.min(index, all.length - 1));
    list.dataset.activeIndex = String(all.length ? clamped : -1);
    all.forEach((el, i) => el.classList.toggle('active', i === clamped));
    all[clamped]?.scrollIntoView?.({ block: 'nearest' });
  }

  function renderFiles(files, loading) {
    list.innerHTML = renderFileList(files, { escapeHtml, loading });
    setActive(0);
  }

  function positionPopup(anchor) {
    if (!anchor) return;
    const rect = anchor.getBoundingClientRect();
    const width = Math.min(rect.width, Math.max(240, (windowImpl?.innerWidth || 600) - 32));
    let left = rect.left;
    const maxLeft = Math.max(8, (windowImpl?.innerWidth || 600) - width - 8);
    if (left + width > maxLeft + 8) left = maxLeft;
    if (left < 8) left = 8;
    const bottom = Math.max(8, (windowImpl?.innerHeight || 800) - rect.top + 4);
    popup.style.setProperty('--slash-popup-left', `${left}px`);
    popup.style.setProperty('--slash-popup-width', `${width}px`);
    popup.style.setProperty('--slash-popup-bottom', `${bottom}px`);
  }

  function open() {
    positionPopup(textarea);
    popup.style.display = 'block';
  }

  function close() {
    popup.style.display = 'none';
    trigger = null;
    if (debounceTimer != null) {
      clearTimeoutImpl(debounceTimer);
      debounceTimer = null;
    }
    if (inflight) {
      inflight.abort();
      inflight = null;
    }
  }

  function fetchAndRender() {
    if (!trigger) return;
    const query = trigger.query;
    const seq = ++reqSeq;
    if (inflight) inflight.abort();
    inflight = AbortControllerImpl ? new AbortControllerImpl() : null;
    const signal = inflight ? inflight.signal : undefined;
    chatApi
      .getFiles(sessionId, query, { signal })
      .then((res) => (res.ok ? res.json() : Promise.reject(new Error('files error'))))
      .then((data) => {
        if (seq !== reqSeq || !isOpen()) return;
        renderFiles(data.files || [], false);
      })
      .catch((err) => {
        if (err && err.name === 'AbortError') return;
        if (seq !== reqSeq || !isOpen()) return;
        renderFiles([], false);
      });
  }

  function refresh() {
    const next = parseAtTrigger(textarea.value, textarea.selectionStart ?? textarea.value.length);
    if (!next) {
      if (isOpen()) close();
      return;
    }
    trigger = next;
    if (!isOpen()) {
      open();
      renderFiles([], true);
    }
    if (debounceTimer != null) clearTimeoutImpl(debounceTimer);
    debounceTimer = setTimeoutImpl(() => {
      debounceTimer = null;
      fetchAndRender();
    }, debounceMs);
  }

  function insert(path, isDir) {
    if (!trigger) return;
    const value = textarea.value;
    const replacement = isDir ? `@${path}/` : `${path} `;
    textarea.value = value.slice(0, trigger.start) + replacement + value.slice(trigger.end);
    const caret = trigger.start + replacement.length;
    textarea.selectionStart = textarea.selectionEnd = caret;
    if (!isDir) close();
    textarea.focus();
    textarea.dispatchEvent(new Event('input', { bubbles: true }));
  }

  // Called by the composer's keydown handler before its Enter-to-submit logic so
  // navigation/selection wins while the popup is open.
  function handleKeydown(event) {
    if (!isOpen()) return false;
    const all = items();
    let active = parseInt(list.dataset.activeIndex || '-1', 10);
    switch (event.key) {
      case 'ArrowDown':
        event.preventDefault();
        setActive(active + 1);
        return true;
      case 'ArrowUp':
        event.preventDefault();
        setActive(active - 1);
        return true;
      case 'Enter':
      case 'Tab':
        if (active >= 0 && all[active]) {
          event.preventDefault();
          all[active].click();
          return true;
        }
        return false;
      case 'Escape':
        event.preventDefault();
        close();
        return true;
      default:
        return false;
    }
  }

  textarea.addEventListener('input', refresh);

  list.addEventListener('click', (event) => {
    const item = event.target.closest('.slash-item');
    if (!item) return;
    insert(item.dataset.insert || '', item.dataset.isdir === '1');
  });

  documentImpl.addEventListener('click', (event) => {
    if (isOpen() && !popup.contains(event.target) && event.target !== textarea) close();
  });

  return { handleKeydown, open, close, isOpen, refresh };
}
