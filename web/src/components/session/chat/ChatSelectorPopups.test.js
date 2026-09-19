import { afterEach, describe, expect, it } from 'vitest';
import { cleanup, render } from '@testing-library/svelte';
import { t } from '../../../shared/i18n.js';
import ChatSelectorPopups from './ChatSelectorPopups.svelte';

afterEach(() => {
  cleanup();
});

describe('ChatSelectorPopups', () => {
  it('renders selector runtime anchors with stable ids and classes', () => {
    render(ChatSelectorPopups);

    expect(document.getElementById('pi-chat-model-popup').className).toBe('pi-chat-model-popup');
    expect(document.getElementById('pi-chat-model-close')).toBeTruthy();
    expect(document.getElementById('pi-chat-model-search').getAttribute('autocomplete')).toBe(
      'off',
    );
    expect(document.getElementById('pi-chat-model-search').getAttribute('placeholder')).toBe(
      t('composer.searchModels'),
    );
    expect(document.getElementById('pi-chat-model-list').className).toBe('pi-chat-model-list');

    expect(document.getElementById('pi-chat-thinking-popup').className).toBe(
      'pi-chat-thinking-popup',
    );
    expect(document.getElementById('pi-chat-thinking-list').className).toBe(
      'pi-chat-thinking-list',
    );
    expect(document.getElementById('pi-chat-slash-popup').className).toBe('pi-chat-slash-popup');
    expect(document.getElementById('pi-chat-slash-list').className).toBe('pi-chat-slash-list');
    expect(document.getElementById('pi-chat-mention-popup').className).toBe('pi-chat-slash-popup');
    expect(document.getElementById('pi-chat-mention-list').className).toBe('pi-chat-slash-list');
  });
});
