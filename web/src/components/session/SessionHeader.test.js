import { afterEach, describe, expect, it, vi } from 'vitest';
import { cleanup, fireEvent, render, waitFor } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import SessionHeader from './SessionHeader.svelte';
import { sessionModals, resetSessionModals } from '../../session/session-modals.svelte.js';

afterEach(() => {
  cleanup();
  vi.unstubAllGlobals();
  vi.restoreAllMocks();
  resetSessionModals();
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
    await waitFor(() =>
      expect(fields[0].querySelector('.model-item[data-provider="custom"]')).toBeTruthy(),
    );
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

  it('navigates to /schedules from the header control', async () => {
    const user = userEvent.setup();
    const pushState = vi.spyOn(window.history, 'pushState');
    render(SessionHeader, {
      props: { title: 'How', cwd: '/tmp', sessionId: 's.jsonl', sessionUUID: 'uuid' },
    });

    const link = document.querySelector('[data-schedules-btn]');
    expect(link).toBeTruthy();
    expect(link.getAttribute('href')).toBe('/schedules');
    expect(link.textContent).toContain('Schedules');

    const right = document.querySelector('.session-header-right');
    expect(right?.children[0]).toBe(link);
    expect(right?.children[1]?.id).toBe('new-session-header-btn');

    await user.click(link);
    expect(pushState).toHaveBeenCalledWith({ back: '/' }, '', '/schedules');
  });

  it('opens the manage-projects sheet from the header control', async () => {
    const user = userEvent.setup();
    render(SessionHeader, {
      props: { title: 'How', cwd: '/tmp', sessionId: 's.jsonl', sessionUUID: 'uuid' },
    });

    const button = document.querySelector('[data-manage-projects-btn]');
    expect(button).toBeTruthy();
    expect(button.textContent).toContain('Manage Projects');
    expect(button.previousElementSibling?.id).toBe('tree-toggle');

    expect(sessionModals.projects).toBe(false);
    await user.click(button);
    expect(sessionModals.projects).toBe(true);
  });
});
