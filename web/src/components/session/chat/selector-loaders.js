export function chatSessionId({
  documentImpl = document,
  locationImpl = location,
  URLSearchParamsImpl = URLSearchParams,
} = {}) {
  return (
    new URLSearchParamsImpl(locationImpl.search).get('id') ||
    (documentImpl.getElementById('pi-chat-composer') || {}).dataset?.sessionId ||
    ''
  );
}

const noopKeydownSelector = { handleKeydown: () => false };

export function createChatSelectorLoaders({
  documentImpl = document,
  windowImpl = window,
  locationImpl = windowImpl.location,
  URLSearchParamsImpl = URLSearchParams,
  entries = [],
  chatApi,
  escapeHtml = String,
  modelSelector,
  thinkingSelector,
  slashSelector,
  mentionSelector,
  t = (key) => key,
  setModelLabel = () => {},
  setChatStatus = () => {},
  setThinkingLabel = () => {},
  setKnownModelLabel = () => {},
  getKnownModelLabel = () => '',
  setCurrentModelForThinking = () => {},
  setWorkerModelUpdate = () => {},
  getCurrentModelForThinking = () => null,
  getKnownThinkingLevel = () => '',
  setKnownThinkingLevel = () => {},
} = {}) {
  const getSessionId = () => chatSessionId({ documentImpl, locationImpl, URLSearchParamsImpl });

  function loadModelSelector() {
    return modelSelector.setupModelSelector({
      documentImpl,
      sessionId: getSessionId(),
      entries,
      chatApi,
      escapeHtml,
      t,
      setModelLabel,
      setChatStatus,
      setKnownModelLabel,
      getKnownModelLabel,
      setCurrentModelForThinking,
      setWorkerModelUpdate,
    });
  }

  function loadSlashSelector() {
    if (!slashSelector || typeof slashSelector.setupSlashCommands !== 'function') {
      return noopKeydownSelector;
    }
    return slashSelector.setupSlashCommands({
      documentImpl,
      sessionId: getSessionId(),
      chatApi,
      escapeHtml,
    });
  }

  function loadMentionSelector() {
    if (!mentionSelector || typeof mentionSelector.setupMentionAutocomplete !== 'function') {
      return noopKeydownSelector;
    }
    return mentionSelector.setupMentionAutocomplete({
      documentImpl,
      windowImpl,
      sessionId: getSessionId(),
      chatApi,
      escapeHtml,
    });
  }

  function loadThinkingSelector() {
    return thinkingSelector.setupThinkingLevelSelector({
      documentImpl,
      windowImpl,
      sessionId: getSessionId(),
      entries,
      getCurrentModel: getCurrentModelForThinking,
      getKnownThinkingLevel,
      setKnownThinkingLevel,
      setThinkingLabel,
      setChatStatus,
      chatApi,
    });
  }

  return {
    loadModelSelector,
    loadThinkingSelector,
    loadSlashSelector,
    loadMentionSelector,
  };
}
