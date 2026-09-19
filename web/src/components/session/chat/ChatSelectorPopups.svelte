<script>
  import { t } from '../../../shared/i18n.js';
  import { icon } from '../../../shared/icons.js';
  import { X } from 'lucide';
</script>

<!-- eslint-disable svelte/no-at-html-tags -- trusted: Lucide icon SVG and rendered session markdown -->
<div
  id="pi-chat-model-popup"
  class="pi-chat-model-popup"
  role="dialog"
  aria-label={t('composer.searchModels')}
>
  <div class="pi-chat-model-header">
    <input
      type="text"
      id="pi-chat-model-search"
      class="pi-chat-model-search"
      placeholder={t('composer.searchModels')}
      autocomplete="off"
    />
    <button
      type="button"
      id="pi-chat-model-close"
      class="pi-chat-model-close"
      aria-label={t('common.close')}
      title={t('common.close')}>{@html icon(X, { size: 18, class: 'close-icon' })}</button
    >
  </div>
  <div id="pi-chat-model-list" class="pi-chat-model-list"></div>
</div>
<div id="pi-chat-thinking-popup" class="pi-chat-thinking-popup" style="display: none">
  <div id="pi-chat-thinking-list" class="pi-chat-thinking-list"></div>
</div>
<div id="pi-chat-slash-popup" class="pi-chat-slash-popup" style="display: none">
  <div id="pi-chat-slash-list" class="pi-chat-slash-list"></div>
</div>
<div id="pi-chat-mention-popup" class="pi-chat-slash-popup" style="display: none">
  <div id="pi-chat-mention-list" class="pi-chat-slash-list"></div>
</div>
