import { afterEach, describe, expect, it } from 'vitest';
import { cleanup, fireEvent, render, waitFor } from '@testing-library/svelte';
import { vi } from 'vitest';
import ContextUsage from './ContextUsage.svelte';

afterEach(cleanup);

describe('ContextUsage', () => {
  it('renders the usage capsule IDs used by the composer runtime', () => {
    render(ContextUsage);

    expect(document.getElementById('pi-chat-context-usage')).toBeTruthy();
    expect(document.querySelector('#pi-chat-context-usage .pi-context-fill')).toBeTruthy();
    expect(document.querySelector('#pi-chat-context-usage .pi-context-text')?.textContent).toBe(
      '0%',
    );
  });

  it('renders the popover IDs used by the composer runtime', () => {
    render(ContextUsage, { props: { popover: true } });

    expect(document.getElementById('pi-chat-context-popover')).toBeTruthy();
    expect(document.getElementById('pi-popover-val-input')?.textContent).toBe('0');
    expect(document.getElementById('pi-popover-val-cache-read')?.textContent).toBe('0');
    expect(document.getElementById('pi-popover-val-cache-write')?.textContent).toBe('0');
    expect(document.getElementById('pi-popover-val-output')?.textContent).toBe('0');
    expect(document.getElementById('pi-popover-val-total')?.textContent).toBe('0');
  });

  it('shows Force Compact only in Local Mode and invokes the session endpoint', async () => {
    const fetchMock = vi.fn().mockResolvedValue({ ok: true, json: async () => ({ ok: true }) });
    vi.stubGlobal('fetch', fetchMock);
    render(ContextUsage);
    const { container } = render(ContextUsage, {
      props: { popover: true, localMode: true, sessionId: 'local-session' },
    });
    await fireEvent.click(container.querySelector('.pi-force-compact'));
    await waitFor(() => expect(fetchMock).toHaveBeenCalledTimes(1));
    expect(fetchMock.mock.calls[0][0]).toContain('/api/force-compact?id=local-session');
    expect(document.querySelector('.pi-context-text')?.textContent).toBe('—');
    vi.unstubAllGlobals();
  });
});
