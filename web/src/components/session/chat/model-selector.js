import {
  detectCurrentModel,
  findModel,
  isScopedModel,
  modelDisplayLabel,
} from '../../../session/chat/chat-selectors.js';
import { icon } from '../../../shared/icons.js';
import { Star, StarOff } from 'lucide';
import {
  loadRecentModels,
  loadStarredModels,
  saveStarredModels,
  touchRecentModel,
} from '../../../session/chat/chat-model-prefs.js';

import { endpointIsLocal } from '../../../index/sessions.js';

// A "local" model is one served from the machine or LAN: the pi-known local
// provider names, or a baseUrl on a private/loopback endpoint (shared with
// effectiveModeForModel in index/sessions.js so the index and the picker
// agree).
const LOCAL_PROVIDERS = new Set(['llama', 'ollama', 'lmstudio', 'lm-studio', 'local']);

export function isLocalModel(model) {
  const provider = String(model?.provider || '')
    .trim()
    .toLowerCase();
  if (LOCAL_PROVIDERS.has(provider)) return true;
  return endpointIsLocal(model?.baseUrl);
}

// `provider\u0000id` key — byte-identical to sessions.js `modelKey` (the
// canonical storage/index key), so localStorage entries match models from the
// server. Must stay in sync with sessions.js; do not use a DOM separator in
// keys (model ids may contain '/', e.g. Llama/LFM2-8B-LFIAW-MTP).
export function modelKey(model) {
  return `${model?.provider || ''}\u0000${model?.id || model?.modelId || ''}`;
}

export function renderModelList(models, options = {}) {
  const {
    filter = '',
    selectedModel = null,
    starred = [],
    recent = [],
    escapeHtml = String,
    t: tFn = (key) => key,
  } = options;
  const starredSet = new Set(starred);
  const recentsSet = new Set(recent);
  // Keys everywhere are the NUL-joined canonical keys (sessions.js). In the DOM
  // they travel via escapeHtml-escaped attributes, which round-trip losslessly
  // through HTML parsing (no / < > & \" in a NUL-joined key).
  const q = filter.trim().toLowerCase();

  // Stable display order: priority (local > starred > recent > cloud) first,
  // then provider, then name — so a star click or a recent-use update
  // re-ranks without disturbing the alphabetical order inside each section.
  const sections = new Map(); // label -> models
  const push = (label, models) => {
    if (!sections.has(label)) sections.set(label, []);
    sections.get(label).push(...models);
  };

  const sorted = [...models].sort(
    (a, b) =>
      (a?.provider || '').localeCompare(b?.provider || '') ||
      String(a?.name || a?.id || '').localeCompare(String(b?.name || b?.id || '')),
  );
  const localModels = sorted.filter(
    (m) => isLocalModel(m) && !starredSet.has(modelKey(m)) && !recentsSet.has(modelKey(m)),
  );
  const starredModels = sorted.filter((m) => starredSet.has(modelKey(m)));
  const recentModels = sorted.filter(
    (m) => !isLocalModel(m) && !starredSet.has(modelKey(m)) && recentsSet.has(modelKey(m)),
  );
  const cloudModels = sorted.filter(
    (m) => !isLocalModel(m) && !starredSet.has(modelKey(m)) && !recentsSet.has(modelKey(m)),
  );
  push(tFn('composer.modelSectionLocal'), localModels);
  push(tFn('composer.modelSectionStarred'), starredModels);
  push(tFn('composer.modelSectionRecent'), recentModels);
  push(tFn('composer.modelSectionCloud'), cloudModels);

  const selectedKey = selectedModel ? modelKey(selectedModel) : '';
  const matches = (m) => {
    if (!q) return true;
    const name = String(m?.name || m?.id || m?.modelId || '').toLowerCase();
    const provider = String(m?.provider || '').toLowerCase();
    return name.includes(q) || provider.includes(q);
  };

  let html = '';
  for (const [label, modelsInSection] of sections) {
    const visible = modelsInSection.filter(matches);
    if (visible.length === 0) continue;
    html += `<div class="model-section-label">${escapeHtml(label)}</div>`;
    visible.forEach((model) => {
      const provider = model?.provider || 'unknown';
      const id = model?.id || model?.modelId || '';
      const name = model?.name || id;
      const key = modelKey(model);
      const starredNow = starredSet.has(key);
      const starTitle = starredNow ? tFn('composer.modelUnstar') : tFn('composer.modelStar');
      const scoped = isScopedModel(model) ? '<span class="model-scope-badge">scoped</span>' : '';
      const selected = selectedKey !== '' && modelKey(model) === selectedKey;
      const starIcon = starredNow
        ? icon(Star, { size: 14, class: 'star-icon' })
        : icon(StarOff, { size: 14, class: 'star-icon' });
      html += `<div class="model-item${selected ? ' selected' : ''}" role="button" tabindex="0" data-provider="${escapeHtml(provider)}" data-model-id="${escapeHtml(id)}"><span class="model-name-wrap">${escapeHtml(name)}${scoped}</span><button type="button" class="model-star${starredNow ? ' starred' : ''}" data-star-key="${escapeHtml(key)}" title="${escapeHtml(starTitle)}" aria-label="${escapeHtml(starTitle)}">${starIcon}</button></div>`;
    });
  }
  if (!html) return '<div class="model-empty">No models match</div>';
  return html;
}

export function setupModelSelector({
  documentImpl = document,
  windowImpl = documentImpl.defaultView,
  sessionId,
  entries = [],
  chatApi,
  escapeHtml = String,
  t = (key) => key,
  storage = globalThis.localStorage,
  setModelLabel = () => {},
  setChatStatus = () => {},
  setKnownModelLabel = () => {},
  getKnownModelLabel = () => '',
  setCurrentModelForThinking = () => {},
  setWorkerModelUpdate = () => {},
} = {}) {
  let allModels = [];
  let selectedModel = null;
  let starred = loadStarredModels(storage);
  let recents = loadRecentModels(storage);

  const popup = documentImpl.getElementById('pi-chat-model-popup');
  const popupSearch = documentImpl.getElementById('pi-chat-model-search');
  const popupList = documentImpl.getElementById('pi-chat-model-list');
  const popupClose = documentImpl.getElementById('pi-chat-model-close');
  const modelLabelBtn = documentImpl.getElementById('pi-chat-model-label');
  const popupStackingRoot = popup?.closest('.content-container');

  // Always show the label button so the user can open the model picker.
  // Server may have hidden it when no model was detected at page load.
  if (modelLabelBtn) modelLabelBtn.style.display = '';

  function renderPopupList(filter) {
    if (!popupList) return;
    popupList.innerHTML = renderModelList(allModels, {
      filter,
      selectedModel,
      starred,
      recent: recents,
      escapeHtml,
      t,
    });
    popupList.dataset.activeIndex = '-1';
  }

  function setSelected(model) {
    selectedModel = model;
    setCurrentModelForThinking(model || null);
  }

  function openPopup({ focusSearch = true } = {}) {
    if (!popup) return;
    popupStackingRoot?.classList.add('model-picker-open');
    popup.classList.add('open');
    renderPopupList('');
    if (popupSearch) {
      popupSearch.value = '';
      // Touch: focusing the search field pops up the on-screen keyboard and
      // scrolls the composer into view, which reads as the picker "re-"
      // selecting the model. Keyboard/desktop flows keep the focus so the
      // search field and Arrow/Enter nav work.
      if (focusSearch && !windowImpl?.matchMedia?.('(pointer: coarse)').matches) {
        popupSearch.focus();
      }
    }
  }

  function closePopup(focusTextarea = false) {
    if (popup) popup.classList.remove('open');
    popupStackingRoot?.classList.remove('model-picker-open');
    if (focusTextarea) {
      const textarea = documentImpl.getElementById('pi-chat-message');
      if (textarea) textarea.focus();
    }
  }

  const api = {
    open: () => openPopup({ focusSearch: true }),
    close: closePopup,
    attachStarToggle: (node) => node?.addEventListener?.('click', starToggleFromEvent),
  };

  // The button click handler stays responsive even before the model list
  // finishes loading. Touch taps use pointer events, so `click` still fires
  // for a tap — this keeps a single open/close code path.
  modelLabelBtn?.addEventListener('click', (e) => {
    e.stopPropagation();
    if (popup && popup.classList.contains('open')) closePopup();
    else openPopup({ focusSearch: false });
  });

  // Explicit close button (×) in the popup header.
  popupClose?.addEventListener('click', (e) => {
    e.stopPropagation();
    closePopup(true);
  });

  popupSearch?.addEventListener('input', () => renderPopupList(popupSearch.value));
  popupSearch?.addEventListener('keydown', (e) => {
    const items = popupList ? popupList.querySelectorAll('.model-item') : [];
    let popupActive = parseInt((popupList && popupList.dataset.activeIndex) || '-1', 10);
    if (e.key === 'ArrowDown') {
      e.preventDefault();
      popupActive = Math.min(popupActive + 1, items.length - 1);
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      popupActive = Math.max(popupActive - 1, 0);
    } else if (e.key === 'Enter') {
      e.preventDefault();
      if (popupActive >= 0 && items[popupActive]) items[popupActive].click();
      return;
    } else if (e.key === 'Escape') {
      e.preventDefault();
      e.stopPropagation();
      closePopup();
      modelLabelBtn?.focus();
      return;
    }
    if (popupList) popupList.dataset.activeIndex = popupActive;
    items.forEach((item, i) => item.classList.toggle('active', i === popupActive));
    items[popupActive]?.scrollIntoView?.({ block: 'nearest' });
  });

  // Refresh the chat popup list; the new-session modal re-renders via its
  // own hook (window.__piWebRerenderModelPicker) registered alongside.
  const rerenderAll = () => {
    if (popupList) renderPopupList(popupSearch ? popupSearch.value : '');
    if (documentImpl.defaultView?.__piWebRerenderModelPicker) {
      documentImpl.defaultView.__piWebRerenderModelPicker();
    }
  };

  // Star clicks: toggle the star without selecting the row. The key is
  // reconstructed from the row's data-provider/data-model-id attributes
  // (plain identifiers, no NUL), because a NUL-joined key does not
  // round-trip through HTML parsing (\u0000 → U+FFFD in attributes).
  function starToggleFromEvent(e) {
    const star = e.target.closest?.('.model-star');
    if (!star) return;
    e.preventDefault();
    e.stopPropagation();
    const row = star.closest?.('.model-item');
    const provider = row?.dataset?.provider;
    const modelId = row?.dataset?.modelId;
    if (!provider || !modelId) return;
    const key = modelKey({ provider, id: modelId });
    if (starred.includes(key)) starred = starred.filter((k) => k !== key);
    else starred = [...starred, key];
    saveStarredModels(starred, storage);
    rerenderAll();
  }

  popupList?.addEventListener('click', starToggleFromEvent);

  async function applyModel(provider, modelId) {
    recents = touchRecentModel(modelKey({ provider, id: modelId }), storage);
    closePopup(true);
    try {
      const setRes = await chatApi.setModel(sessionId, { provider, modelId });
      const setData = await setRes.json();
      if (!setRes.ok) throw new Error(setData.error || 'set model failed');
      const model = findModel(allModels, provider, modelId);
      const chosen = model || { provider, id: modelId, name: modelId };
      setSelected(chosen);
      const newLabel = modelDisplayLabel(chosen);
      setKnownModelLabel(newLabel);
      setModelLabel(newLabel);
    } catch (err) {
      setChatStatus(err.message || String(err), 'error');
    }
  }

  popupList?.addEventListener('click', async (e) => {
    const item = e.target.closest('.model-item');
    if (!item) return;
    if (e.target.closest('.model-star')) return; // handled above
    const provider = item.dataset.provider;
    const modelId = item.dataset.modelId;
    if (!provider || !modelId) return;
    await applyModel(provider, modelId);
  });

  // Keyboard activation: div.model-item is not a native button, so
  // Enter/Space must trigger the same click handler.
  popupList?.addEventListener('keydown', (e) => {
    if (e.key !== 'Enter' && e.key !== ' ') return;
    const item = e.target.closest?.('.model-item');
    if (!item || e.target !== item) return;
    e.preventDefault();
    item.click();
  });

  // Clicking the popup backdrop (not the search or list) closes the popup.
  popup?.addEventListener('click', (e) => {
    if (e.target === popup) closePopup(true);
  });

  // Esc key closes the popup from anywhere. Listened on the overlay itself
  // (bubble phase, not capture): keydown from the search input or a model row
  // bubbles up to the overlay, and any other handler on the same node runs
  // before this one because it was registered first — so this only runs when
  // nothing else claimed the key, which also blocks the document-level
  // keyboard-nav shortcut handler.
  popup?.addEventListener('keydown', (e) => {
    if (e.key !== 'Escape' || !popup.classList.contains('open')) return;
    e.preventDefault();
    e.stopPropagation();
    closePopup(true);
  });

  // Load the model list asynchronously; the button is already wired.
  // Fire-and-forget: the popup opens immediately (Ctrl+L) and renders
  // available models as they load.
  chatApi
    .listModels()
    .then((res) => {
      if (!res.ok) throw new Error('api error');
      return res.json();
    })
    .then((data) => {
      if (!data.models || data.models.length === 0) {
        allModels = [];
        if (popupList) {
          popupList.innerHTML =
            '<div class="model-empty">No models configured<br><small>Run <code>pi setup</code> to configure</small></div>';
        }
        return;
      }
      allModels = data.models;
      if (popup && popup.classList.contains('open')) {
        renderPopupList(popupSearch ? popupSearch.value : '');
      }
      function updateToggleFromStatus(provider, modelId) {
        if (!provider || !modelId) return;
        const model = findModel(allModels, provider, modelId);
        if (model) setSelected(model);
      }
      setWorkerModelUpdate(updateToggleFromStatus);
      const detected = detectCurrentModel(entries);
      if (detected.modelId) {
        const model = findModel(allModels, detected.provider, detected.modelId);
        if (model) {
          setSelected(model);
          const detectedLabel = modelDisplayLabel(model);
          if (detectedLabel && !getKnownModelLabel()) {
            setKnownModelLabel(detectedLabel);
            setModelLabel(detectedLabel);
          }
        }
      }
    })
    .catch(() => {
      // Model list fetch failed; button still works (shows empty list).
      if (popupList) {
        popupList.innerHTML =
          '<div class="model-empty">Failed to load models<br><small>Check that <code>pi</code> is on PATH</small></div>';
      }
    });

  return api;
}
