<script>
  import { t } from '../../shared/i18n.js';
  import {
    effectiveModeForModel,
    modelKey as keyForModel, // NUL-joined canonical key (storage + DOM attribute boundary)
  } from '../../index/sessions.js';
  import { renderModelList } from '../session/chat/model-selector.js';
  import {
    loadRecentModels,
    loadStarredModels,
    saveStarredModels,
  } from '../../session/chat/chat-model-prefs.js';
  import { escapeHtml } from '../../shared/escape.js';
  import { onMount, tick } from 'svelte';

  let {
    open = false,
    recent = [],
    path = $bindable(''),
    creating = false,
    error = '',
    models = [],
    modelKey = $bindable(''),
    mode = $bindable('auto'),
    onClose = () => {},
    onCreate = () => {},
  } = $props();

  const selectedModel = $derived(models.find((model) => keyForModel(model) === modelKey) || null);
  const effectiveMode = $derived(effectiveModeForModel(mode, selectedModel));
  let modelFilter = $state('');

  let listEl = $state(null);
  // Renders the shared model-picker list into the host node. The html is built
  // in renderModelList: model names/providers are escapeHtml-escaped, star svls
  // are trusted lucide markup, section labels are i18n strings — the same trust
  // model the chat composer popups use via innerHTML. Done imperatively (not
  // {@html}) so that icon markup is not sanitized away.
  // The model picker is a static host node — Svelte never renders into it, so
  // it is safe to fill imperatively (and necessary: {@html} would sanitize the
  // trusted lucide star svls). Same trust model as the chat composer popups:
  // model names/providers are escapeHtml-escaped in renderModelList, stars are
  // trusted lucide markup, section labels are i18n strings.
  const setListHtml = (html) => {
    if (listEl) listEl.innerHTML = html; // eslint-disable-line svelte/no-dom-manipulating
  };
  // Star toggles write localStorage, which the $derived above doesn't track;
  // the shared star handler bumps this via window.__piWebRerenderModelPicker
  // right after saving, forcing a re-read.
  let starTick = $state(0);
  const modelPickerHtml = $derived.by(() => {
    void starTick;
    return renderModelList(models, {
      filter: modelFilter,
      selectedModel: modelKey ? selectedModel : null,
      starred: loadStarredModels(),
      recent: loadRecentModels(),
      escapeHtml,
      t,
    });
  });
  $effect(() => {
    if (listEl) setListHtml(modelPickerHtml);
  });
  const renderModelPicker = () => {
    starTick = starTick + 1;
    setListHtml(modelPickerHtml);
  };
  $effect(() => {
    if (open) void tick().then(renderModelPicker);
  });
  onMount(() => {
    // The model list is server-provided data rendered as escaped HTML (same
    // pattern as the chat composer popups); Svelte would sanitize {@html},
    // so it is written imperatively here.
    renderModelPicker();
    const prev = window.__piWebRerenderModelPicker;
    const mine = () => {
      renderModelPicker();
      prev?.();
    };
    window.__piWebRerenderModelPicker = mine;
    return () => {
      if (window.__piWebRerenderModelPicker === mine) {
        window.__piWebRerenderModelPicker = prev;
      }
    };
  });

  function chooseModel(event) {
    if (event.target.closest?.('.model-star')) return;
    const item = event.target.closest?.('.model-item');
    if (!item) return;
    if (!item.dataset.provider || !item.dataset.modelId) return;
    const model = models.find(
      (candidate) =>
        keyForModel(candidate) === `${item.dataset.provider}\u0000${item.dataset.modelId}`,
    );
    if (model) modelKey = keyForModel(model);
  }

  function handleModelPickerClick(event) {
    const star = event.target.closest?.('.model-star');
    if (!star) {
      chooseModel(event);
      return;
    }
    event.preventDefault();
    event.stopPropagation();
    const item = star.closest('.model-item');
    const provider = item?.dataset.provider;
    const modelId = item?.dataset.modelId;
    if (!provider || !modelId) return;
    const key = `${provider}\u0000${modelId}`;
    const starred = loadStarredModels();
    saveStarredModels(
      starred.includes(key) ? starred.filter((candidate) => candidate !== key) : [...starred, key],
    );
    renderModelPicker();
  }

  function chooseRecent(loc) {
    path = loc;
    requestAnimationFrame(() => document.getElementById('sessionPath')?.focus());
  }

  function handleKeydown(e) {
    if (e.key === 'Enter') {
      e.preventDefault();
      onCreate();
    }
  }
</script>

<div
  class="modal-overlay"
  id="modalOverlay"
  class:visible={open}
  class:open
  role="presentation"
  onclick={(e) => {
    if (e.currentTarget === e.target) onClose();
  }}
>
  <div class="modal">
    <div class="modal-sheet-header">
      <button
        class="modal-sheet-back"
        id="modalBackBtn"
        type="button"
        aria-label={t('index.closeNewSession')}
        onclick={onClose}
      >
        <span aria-hidden="true">←</span>
        <span>{t('index.startNewSession')}</span>
      </button>
    </div>
    <h2>{t('index.startNewSession')}</h2>
    <div class="recent-locations" id="recentLocations">
      {#each recent as loc (loc)}
        <button type="button" class="recent-chip" onclick={() => chooseRecent(loc)}>{loc}</button>
      {/each}
    </div>
    <input
      type="text"
      id="sessionPath"
      placeholder={t('index.sessionPathPlaceholder')}
      bind:value={path}
      onkeydown={handleKeydown}
    />
    <div class="new-session-field">
      <span>{t('index.model')}</span>
      <div class="model-picker" role="listbox" aria-label={t('index.model')}>
        <input
          type="text"
          id="newSessionModelSearch"
          class="model-picker-search"
          placeholder={t('composer.searchModels')}
          autocomplete="off"
          bind:value={modelFilter}
        />
        {#if !selectedModel}
          <button type="button" class="model-item" data-default="">{t('index.modelDefault')}</button
          >
        {/if}
        <div
          class="model-picker-list"
          id="newSessionModelList"
          role="group"
          bind:this={listEl}
          onclick={handleModelPickerClick}
          onkeydown={(event) => {
            if (
              (event.key === 'Enter' || event.key === ' ') &&
              event.target.classList.contains('model-item')
            ) {
              event.preventDefault();
              chooseModel(event);
            }
          }}
        ></div>
      </div>
    </div>
    <label class="new-session-field">
      <span>{t('index.mode')}</span>
      <select bind:value={mode}>
        <option value="auto"
          >{t('index.modeAuto')}{selectedModel
            ? ` (${effectiveMode === 'local' ? t('index.modeLocal') : t('index.modeCloud')})`
            : ''}</option
        >
        <option value="local">{t('index.modeLocal')}</option>
        <option value="cloud">{t('index.modeCloud')}</option>
      </select>
    </label>
    <div class="modal-actions">
      <button class="btn-secondary" id="cancelBtn" type="button" onclick={onClose}
        >{t('common.cancel')}</button
      >
      <button
        class="btn-primary"
        id="createBtn"
        type="button"
        disabled={creating}
        onclick={onCreate}>{t('common.create')}</button
      >
    </div>
    <div class="modal-error" id="modalError">{error}</div>
  </div>
</div>
