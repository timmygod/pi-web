<script>
  import { t } from '../../shared/i18n.js';
  import { effectiveModeForModel, modelKey as keyForModel } from '../../index/sessions.js';

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
    <label class="new-session-field">
      <span>{t('index.model')}</span>
      <select bind:value={modelKey}>
        <option value="">{t('index.modelDefault')}</option>
        {#each models as model (keyForModel(model))}
          <option value={keyForModel(model)}>
            {model.name || model.id || model.modelId} @ {model.provider}
          </option>
        {/each}
      </select>
    </label>
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
