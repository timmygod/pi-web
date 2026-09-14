import { afterEach, describe, expect, it, vi } from 'vitest';
import { cleanup, fireEvent, render, waitFor } from '@testing-library/svelte';
import SessionHeader from './SessionHeader.svelte';

afterEach(() => {
  cleanup();
  vi.unstubAllGlobals();
  document.body.classList.remove('modal-sheet-open');
});

describe('SessionHeader', () => {
  it('opens the mode-aware new-session dialog and carries the source session', async () => {
    const fetchMock = vi.fn((url, options = {}) => {
      if (url === '/api/models') {
        return Promise.resolve(
          new Response(
            JSON.stringify({
              models: [
                {
                  provider: 'custom',
                  id: 'qwen',
                  name: 'Qwen',
                  baseUrl: 'http://127.0.0.1:8000/v1',
                },
              ],
            }),
          ),
        );
      }
      if (url === '/api/new-session' && options.method === 'POST') {
        return Promise.resolve(new Response(JSON.stringify({ ok: false, error: 'test stop' })));
      }
      return Promise.resolve(new Response('{}', { status: 404 }));
    });
    vi.stubGlobal('fetch', fetchMock);

    const { container } = render(SessionHeader, {
      props: {
        cwd: '/project',
        sessionId: 'source.jsonl',
        modelProvider: 'custom',
        modelId: 'qwen',
      },
    });

    await fireEvent.click(container.querySelector('#new-btn'));
    await waitFor(() => expect(container.querySelector('.modal-overlay.visible')).toBeTruthy());
    const fields = container.querySelectorAll('.new-session-field');
    await waitFor(() => expect(fields[0].querySelector('select').value).toBe('custom\u0000qwen'));
    expect(fields[1].querySelector('option[value="auto"]').textContent).toContain('Local');

    await fireEvent.click(container.querySelector('#createBtn'));
    await waitFor(() =>
      expect(fetchMock).toHaveBeenCalledWith('/api/new-session', expect.anything()),
    );
    const createCall = fetchMock.mock.calls.find(([url]) => url === '/api/new-session');
    expect(JSON.parse(createCall[1].body)).toMatchObject({
      path: '/project',
      mode: 'auto',
      modelProvider: 'custom',
      modelId: 'qwen',
      sourceSessionId: 'source.jsonl',
    });
  });
});
