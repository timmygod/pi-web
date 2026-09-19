import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { cleanup, fireEvent, render } from '@testing-library/svelte';
import NewSessionModal from './NewSessionModal.svelte';
import { loadStarredModels } from '../../session/chat/chat-model-prefs.js';

afterEach(cleanup);
beforeEach(() => localStorage.clear());

const MODELS = [
  {
    provider: 'custom',
    id: 'qwen',
    name: 'Qwen',
    baseUrl: 'http://192.168.1.5:8000/v1',
  },
  {
    provider: 'openai',
    id: 'gpt',
    name: 'GPT',
    baseUrl: 'https://api.openai.com/v1',
  },
];

describe('NewSessionModal', () => {
  it('renders a searchable model list in place of the native select', async () => {
    const { container } = render(NewSessionModal, {
      props: { open: true, models: MODELS },
    });
    const modelField = container.querySelectorAll('.new-session-field')[0];
    expect(modelField.querySelector('select')).toBeNull();
    const search = container.querySelector('#newSessionModelSearch');
    expect(search).toBeTruthy();
    const items = Array.from(container.querySelectorAll('#newSessionModelList .model-item'));
    expect(items.map((item) => item.dataset.provider)).toEqual(['custom', 'openai']);
    await fireEvent.input(search, { target: { value: 'gpt' } });
    expect(Array.from(container.querySelectorAll('#newSessionModelList .model-item'))).toHaveLength(
      1,
    );
    expect(container.querySelector('#newSessionModelList .model-item').dataset.provider).toBe(
      'openai',
    );
  });

  it('selecting a model updates Mode Auto to the endpoint-derived mode', async () => {
    const { container } = render(NewSessionModal, {
      props: { open: true, models: MODELS },
    });
    const fields = container.querySelectorAll('.new-session-field');
    expect(fields[0].textContent).toContain('Model');
    expect(fields[1].textContent).toContain('Mode');
    await fireEvent.click(
      container.querySelector('#newSessionModelList .model-item[data-provider="custom"]'),
    );
    expect(fields[1].querySelector('option[value="auto"]').textContent).toContain('Local');
    expect(Array.from(fields[1].querySelectorAll('option')).map((option) => option.value)).toEqual([
      'auto',
      'local',
      'cloud',
    ]);
  });

  it('shows the default (no model) option when none is selected', () => {
    const { container } = render(NewSessionModal, {
      props: { open: true, models: MODELS },
    });
    expect(container.querySelector('.model-item[data-default]')?.textContent).toContain(
      'Default (pi chooses)',
    );
  });

  it('toggles a star without selecting the model', async () => {
    const { container } = render(NewSessionModal, {
      props: { open: true, models: MODELS },
    });
    const model = container.querySelector(
      '#newSessionModelList .model-item[data-provider="openai"]',
    );

    await fireEvent.click(model.querySelector('.model-star'));

    expect(loadStarredModels()).toEqual(['openai\u0000gpt']);
    expect(container.querySelector('.model-star.starred')).toBeTruthy();
    expect(container.querySelector('option[value="auto"]').textContent).not.toContain('Cloud');

    await fireEvent.click(container.querySelector('.model-star.starred'));
    expect(loadStarredModels()).toEqual([]);
  });
});
