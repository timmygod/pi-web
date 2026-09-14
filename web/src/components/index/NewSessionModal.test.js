import { afterEach, describe, expect, it } from 'vitest';
import { cleanup, fireEvent, render } from '@testing-library/svelte';
import NewSessionModal from './NewSessionModal.svelte';

afterEach(cleanup);

describe('NewSessionModal', () => {
  it('renders Model then Mode and resolves Auto from the selected endpoint', async () => {
    const { container } = render(NewSessionModal, {
      props: {
        open: true,
        models: [
          {
            provider: 'custom',
            id: 'qwen',
            name: 'Qwen',
            baseUrl: 'http://192.168.1.5:8000/v1',
          },
        ],
      },
    });
    const fields = container.querySelectorAll('.new-session-field');
    expect(fields[0].textContent).toContain('Model');
    expect(fields[1].textContent).toContain('Mode');
    await fireEvent.change(fields[0].querySelector('select'), {
      target: { value: 'custom\u0000qwen' },
    });
    expect(fields[1].querySelector('option[value="auto"]').textContent).toContain('Local');
    expect(Array.from(fields[1].querySelectorAll('option')).map((option) => option.value)).toEqual([
      'auto',
      'local',
      'cloud',
    ]);
  });
});
