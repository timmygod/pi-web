import { describe, expect, it, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import SessionInfoHeader from './SessionInfoHeader.svelte';
import { SessionDataModel } from '../../session/data/session-data.svelte.js';

function mount(overrides = {}, props = {}) {
  const model = new SessionDataModel({
    header: { id: 'sid-123', timestamp: '2026-01-01T00:00:00Z' },
    entries: [
      { id: 'u', type: 'message', message: { role: 'user' }, timestamp: '2026-01-01T00:00:00Z' },
      {
        id: 'a',
        parentId: 'u',
        type: 'message',
        timestamp: '2026-01-01T00:00:01Z',
        message: {
          role: 'assistant',
          model: 'm',
          usage: { input: 1200 },
          content: [{ type: 'toolCall' }],
        },
      },
    ],
    ...overrides,
  });
  return { model, ...render(SessionInfoHeader, { props: { model, ...props } }) };
}

describe('SessionInfoHeader', () => {
  it('renders session id, stats and the toggle/download buttons', () => {
    const { container } = mount();
    expect(screen.getByText('Session: sid-123')).toBeInTheDocument();
    expect(container.querySelector('[data-action="toggle-thinking"]')).toBeInTheDocument();
    expect(container.querySelector('[data-action="toggle-tools"]')).toBeInTheDocument();
    expect(container.querySelector('[data-action="toggle-tool-output"]')).toBeInTheDocument();
    expect(container.querySelector('.download-json-btn')).toBeInTheDocument();
    expect(container.textContent).toContain('Mode:');
    expect(container.textContent).toContain('Auto (Cloud)');
    expect(container.querySelector('.session-mode-select')).not.toBeInTheDocument();
    // messages summary reflects the entries
    expect(container.textContent).toContain('1 user, 1 assistant');
    expect(container.textContent).toContain('↑1.2k');
  });

  it('lets live sessions change mode and reflects the effective mode', async () => {
    const onModeChange = vi.fn().mockResolvedValue({
      configuredMode: 'auto',
      effectiveMode: 'local',
    });
    const { container, model } = mount(
      { configuredMode: 'cloud', effectiveMode: 'cloud' },
      { modeEditable: true, onModeChange },
    );
    const select = container.querySelector('.session-mode-select');

    expect(select.value).toBe('cloud');
    await userEvent.selectOptions(select, 'auto');

    expect(onModeChange).toHaveBeenCalledWith('auto');
    await waitFor(() => expect(select.value).toBe('auto'));
    expect(select.selectedOptions[0].textContent).toBe('Auto (Local)');
    expect(model.configuredMode).toBe('auto');
    expect(model.effectiveMode).toBe('local');
  });

  it('restores the previous mode when a live update fails', async () => {
    const onModeChange = vi.fn().mockRejectedValue(new Error('mode unavailable'));
    const { container } = mount(
      { configuredMode: 'auto', effectiveMode: 'local' },
      { modeEditable: true, onModeChange },
    );
    const select = container.querySelector('.session-mode-select');

    await userEvent.selectOptions(select, 'cloud');

    await waitFor(() => expect(select.value).toBe('auto'));
    expect(select).toHaveClass('error');
    expect(select).toHaveAttribute('title', 'mode unavailable');
  });

  it('renders an expandable system prompt and toggles on click', async () => {
    const { container } = mount({
      systemPrompt: Array.from({ length: 12 }, (_, i) => `line ${i}`).join('\n'),
    });
    const block = container.querySelector('.system-prompt.expandable');
    expect(block).toBeInTheDocument();
    expect(block).not.toHaveClass('expanded');
    expect(container.textContent).toContain('more lines, click to expand');
    await userEvent.click(block);
    expect(block).toHaveClass('expanded');
  });

  it('renders tools and expands params on click', async () => {
    const { container } = mount({
      tools: [
        {
          name: 'read',
          description: 'Read file',
          parameters: {
            required: ['path'],
            properties: { path: { type: 'string', description: 'the path' } },
          },
        },
      ],
    });
    const item = container.querySelector('.tool-item');
    expect(item).toBeInTheDocument();
    expect(container.querySelector('.tool-item-name').textContent).toBe('read');
    expect(container.querySelector('.tool-param-required')).toBeInTheDocument();
    await userEvent.click(item);
    expect(item).toHaveClass('params-expanded');
  });

  it('escapes session id and prompt text (no raw HTML injection)', () => {
    const { container } = mount({ header: { id: '<sid>' }, systemPrompt: '<b>x</b>' });
    expect(container.querySelector('h1').textContent).toBe('Session: <sid>');
    expect(container.querySelector('h1').innerHTML).toContain('&lt;sid&gt;');
  });
});
