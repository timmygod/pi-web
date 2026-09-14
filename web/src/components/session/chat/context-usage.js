export function buildContextWindows(models = []) {
  const windows = {};
  (models || []).forEach((model) => {
    const provider = model.provider || '';
    const id = model.id || model.modelId || '';
    if (!id) return;
    windows[id.toLowerCase()] = model.contextWindow || 0;
    if (provider) {
      windows[`${provider}/${id}`.toLowerCase()] = model.contextWindow || 0;
    }
  });
  return windows;
}

export function getModelContextLimit(modelId, provider = '', contextWindows = {}) {
  if (!modelId) return 128000;
  const id = modelId.toLowerCase();
  const prov = provider.toLowerCase();

  if (prov && contextWindows[`${prov}/${id}`]) {
    return contextWindows[`${prov}/${id}`];
  }
  if (contextWindows[id]) {
    return contextWindows[id];
  }

  if (id.includes('deepseek')) return 1000000;
  if (
    id.includes('gemini-1.5-pro') ||
    id.includes('gemini-2.0-pro') ||
    id.includes('gemini-2.5-pro') ||
    id.includes('gemini-3.1-pro') ||
    id.includes('agy-gemini-pro')
  ) {
    return 1000000;
  }
  if (id.includes('gemini-')) return 1000000;
  if (id.includes('claude-') || id.includes('sonnet') || id.includes('opus')) return 200000;
  if (id.includes('gpt-5')) return 272000;
  if (
    id.includes('gpt-4') ||
    id.includes('gpt4') ||
    id.includes('gpt-3.5') ||
    id.includes('o1') ||
    id.includes('o3')
  )
    return 128000;
  if (
    id.includes('llama-3') ||
    id.includes('llama3') ||
    id.includes('qwen') ||
    id.includes('glm') ||
    id.includes('mimo')
  )
    return 128000;
  if (id.includes('llama-2') || id.includes('llama2')) return 4096;
  return 128000;
}

export function collectContextUsage(entries = []) {
  let inputTokens = 0;
  let outputTokens = 0;
  let cacheReadTokens = 0;
  let cacheWriteTokens = 0;

  entries.forEach((entry) => {
    if (entry?.type !== 'message' || !entry.message) return;
    const message = entry.message;
    if (message.role === 'assistant' && message.usage) {
      inputTokens += message.usage.input || 0;
      outputTokens += message.usage.output || 0;
      cacheReadTokens += message.usage.cacheRead || 0;
      cacheWriteTokens += message.usage.cacheWrite || 0;
    }
  });

  let latestCompaction = -1;
  for (let i = entries.length - 1; i >= 0; i -= 1) {
    if (entries[i]?.type === 'compaction') {
      latestCompaction = i;
      break;
    }
  }

  let contextTokens = 0;
  let contextKnown = false;
  for (let i = entries.length - 1; i > latestCompaction; i -= 1) {
    const entry = entries[i];
    if (entry?.type !== 'message' || !entry.message) continue;
    const msg = entry.message;
    if (msg.role === 'assistant' && msg.usage) {
      contextTokens =
        msg.usage.totalTokens ||
        (msg.usage.input || 0) +
          (msg.usage.output || 0) +
          (msg.usage.cacheRead || 0) +
          (msg.usage.cacheWrite || 0);
      contextKnown = contextTokens > 0;
      break;
    }
  }

  return {
    inputTokens,
    outputTokens,
    cacheReadTokens,
    cacheWriteTokens,
    totalIOTokens: inputTokens + outputTokens + cacheReadTokens + cacheWriteTokens,
    contextTokens,
    contextKnown,
    hasCompaction: latestCompaction >= 0,
  };
}

function entryKey(entry) {
  if (!entry) return '';
  return String(entry.id || `${entry.type || ''}:${entry.timestamp || ''}`);
}

export function invalidateContextUsage(documentImpl = document) {
  const el = documentImpl.getElementById('pi-chat-context-usage');
  if (!el) return;
  el.dataset.contextInvalidated = el.dataset.lastEntryKey || 'pending';
  renderUnknownContextUsage(documentImpl, el);
}

function renderUnknownContextUsage(documentImpl, el) {
  const fillPath = el.querySelector('.pi-context-fill');
  const textSpan = el.querySelector('.pi-context-text');
  if (fillPath) fillPath.setAttribute('stroke-dasharray', '0, 100');
  if (textSpan) textSpan.textContent = '—';
  el.setAttribute('title', 'Context usage will refresh after the next model response');
  el.classList.remove('warning', 'danger');

  const popoverBox = documentImpl.getElementById('pi-chat-context-popover');
  const usedSpan = popoverBox?.querySelector('.pi-popover-used');
  const popoverBar = popoverBox?.querySelector('.pi-popover-progress-bar');
  if (usedSpan) usedSpan.textContent = '—';
  if (popoverBar) popoverBar.style.width = '0%';
  popoverBox?.classList.remove('warning', 'danger');
  el.style.display = 'inline-flex';
}

function splitModelLabel(label = '') {
  const modelName = label ? label.split(' @ ')[0].trim() : '';
  const providerName = label && label.includes(' @ ') ? label.split(' @ ')[1].trim() : '';
  return { modelName, providerName };
}

function formatTokensDetail(n) {
  if (n >= 1000000) return (n / 1000000).toFixed(1) + 'M';
  if (n >= 1000) return (n / 1000).toFixed(1) + 'k';
  return n.toLocaleString();
}

function formatLimit(n) {
  if (n >= 1000000) return (n / 1000000).toFixed(1) + 'M';
  if (n >= 1000) return (n / 1000).toFixed(0) + 'k';
  return n.toLocaleString();
}

export function updateContextUsage({
  documentImpl = document,
  entries = [],
  knownModelLabel = '',
  contextWindows = {},
  positionPopover = () => {},
} = {}) {
  const el = documentImpl.getElementById('pi-chat-context-usage');
  if (!el) return;

  const usage = collectContextUsage(entries);
  if (!usage.hasCompaction && usage.contextTokens <= 0 && usage.totalIOTokens <= 0) {
    el.style.display = 'none';
    return;
  }

  const lastEntryKey = entryKey(entries[entries.length - 1]);
  const invalidatedAt = el.dataset.contextInvalidated || '';
  el.dataset.lastEntryKey = lastEntryKey;
  if (!usage.contextKnown || (invalidatedAt && invalidatedAt === lastEntryKey)) {
    renderUnknownContextUsage(documentImpl, el);
    return;
  }
  delete el.dataset.contextInvalidated;

  const { modelName, providerName } = splitModelLabel(knownModelLabel);
  const limit = getModelContextLimit(modelName, providerName, contextWindows);
  const percent = Math.min(100, Math.max(0, Math.round((usage.contextTokens / limit) * 100)));

  const fillPath = el.querySelector('.pi-context-fill');
  const textSpan = el.querySelector('.pi-context-text');

  if (fillPath) fillPath.setAttribute('stroke-dasharray', `${percent}, 100`);
  if (textSpan) textSpan.textContent = `${percent}%`;

  const formatNumber = (num) => num.toLocaleString();
  el.setAttribute(
    'title',
    `Click for details (${formatNumber(usage.contextTokens)} / ${formatNumber(limit)} tokens used in context)`,
  );

  el.classList.remove('warning', 'danger');
  if (percent >= 90) el.classList.add('danger');
  else if (percent >= 70) el.classList.add('warning');

  const popoverBox = documentImpl.getElementById('pi-chat-context-popover');
  const valInput = popoverBox ? popoverBox.querySelector('#pi-popover-val-input') : null;
  const valCacheRead = popoverBox ? popoverBox.querySelector('#pi-popover-val-cache-read') : null;
  const valCacheWrite = popoverBox ? popoverBox.querySelector('#pi-popover-val-cache-write') : null;
  const valOutput = popoverBox ? popoverBox.querySelector('#pi-popover-val-output') : null;
  const valTotal = popoverBox ? popoverBox.querySelector('#pi-popover-val-total') : null;

  const usedSpan = popoverBox ? popoverBox.querySelector('.pi-popover-used') : null;
  const limitSpan = popoverBox ? popoverBox.querySelector('.pi-popover-limit') : null;
  const popoverBar = popoverBox ? popoverBox.querySelector('.pi-popover-progress-bar') : null;

  if (valInput) valInput.textContent = formatTokensDetail(usage.inputTokens);
  if (valCacheRead) valCacheRead.textContent = formatTokensDetail(usage.cacheReadTokens);
  if (valCacheWrite) valCacheWrite.textContent = formatTokensDetail(usage.cacheWriteTokens);
  if (valOutput) valOutput.textContent = formatTokensDetail(usage.outputTokens);
  if (valTotal) valTotal.textContent = formatTokensDetail(usage.totalIOTokens);

  if (usedSpan) usedSpan.textContent = formatTokensDetail(usage.contextTokens);
  if (limitSpan) limitSpan.textContent = formatLimit(limit);
  if (popoverBar) popoverBar.style.width = `${percent}%`;

  if (popoverBox) {
    popoverBox.classList.remove('warning', 'danger');
    if (percent >= 90) popoverBox.classList.add('danger');
    else if (percent >= 70) popoverBox.classList.add('warning');
  }

  if (popoverBox && popoverBox.style.display !== 'none') {
    positionPopover();
  }

  el.style.display = 'inline-flex';
}

export function createContextUsageController({
  documentImpl = document,
  entries = [],
  chatApi,
  getKnownModelLabel = () => '',
  positionPopover = () => {},
  isCurrentSession = () => true,
} = {}) {
  let contextWindows = {};

  const update = () => {
    if (!isCurrentSession()) return;
    updateContextUsage({
      documentImpl,
      entries,
      knownModelLabel: getKnownModelLabel(),
      contextWindows,
      positionPopover,
    });
  };

  if (chatApi && typeof chatApi.listModels === 'function') {
    chatApi
      .listModels()
      .then((res) => {
        if (res.ok) return res.json();
        throw new Error();
      })
      .then((data) => {
        if (!isCurrentSession()) return;
        contextWindows = buildContextWindows(data.models || []);
        update();
      })
      .catch(() => {});
  }

  return {
    update,
    getContextWindows: () => contextWindows,
  };
}
