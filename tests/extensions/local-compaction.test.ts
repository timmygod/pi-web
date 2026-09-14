import { describe, expect, it, vi } from 'vitest';

vi.mock('@earendil-works/pi-coding-agent', () => ({
  convertToLlm: (messages: unknown[]) => messages,
  serializeConversation: (messages: unknown[]) => JSON.stringify(messages),
}));
import install from '../../internal/rpc/local_compaction_guard.mjs';

const checkpoint =
  '## Goal\nShip project\n## Constraints\nKeep user files\n## State\nTests passed\n## Next\nDeploy\n## References\n/project\n<checkpoint-end>';
const result = (text = checkpoint, stopReason = 'stop') => ({
  content: [{ type: 'text', text }],
  stopReason,
  usage: { input: 50000, output: 100 },
});

function setup(responses = [result()]) {
  const handlers: Record<string, Function> = {};
  install({
    on: (name: string, fn: Function) => {
      handlers[name] = fn;
    },
  });
  const complete = vi.fn();
  for (const response of responses) complete.mockResolvedValueOnce(response);
  const controller = new AbortController();
  const ctx = {
    model: {
      api: 'openai-completions',
      maxTokens: 8192,
      contextWindow: 131072,
    },
    modelRegistry: { complete },
    ui: { notify: vi.fn() },
    abort: vi.fn(),
    sessionManager: {
      buildContextEntries: vi.fn(() => [
        { type: 'message', message: 'x'.repeat(80000) },
      ]),
    },
  };
  const event = {
    reason: 'threshold',
    signal: controller.signal,
    customInstructions: 'Preserve deployment constraints',
    preparation: {
      previousSummary: 'old '.repeat(10000),
      messagesToSummarize: [{ text: 'current goal' }],
      turnPrefixMessages: [{ text: 'unfinished tool turn' }],
      firstKeptEntryId: 'keep-me',
      tokensBefore: 86000,
    },
  };
  return { handlers, ctx, event, controller, complete };
}

describe('bounded Local Mode compaction', () => {
  it('replaces growing summaries while preserving the recent boundary and usage', async () => {
    const { handlers, ctx, event, complete } = setup();
    const answer = await handlers.session_before_compact(event, ctx);
    expect(answer.compaction).toMatchObject({
      summary: checkpoint,
      firstKeptEntryId: 'keep-me',
      tokensBefore: 86000,
      usage: { input: 50000, output: 100 },
      details: { piWebCheckpointVersion: 1 },
    });
    const [model, input, options] = complete.mock.calls[0];
    expect(model).toBe(ctx.model);
    expect(input.messages[0].content[0].text).toContain('unfinished tool turn');
    expect(input.messages[0].content[0].text).toContain(
      'Preserve deployment constraints',
    );
    expect(options.maxTokens).toBe(4096);
    expect(options.reasoningEffort).toBe('low');
    expect(ctx.abort).not.toHaveBeenCalled();
  });

  it.each([
    ['token cap', result(checkpoint, 'length')],
    ['empty', result('')],
    ['missing sections', result('summary <checkpoint-end>')],
    ['missing end', result(checkpoint.replace('<checkpoint-end>', ''))],
    ['character budget', result('x'.repeat(6000) + checkpoint)],
    ['Chinese byte budget', result('汉'.repeat(4100) + checkpoint)],
    ['provider error', result('', 'error')],
  ])(
    'retries %s once using the original evidence and a tighter target',
    async (_, invalid) => {
      const { handlers, ctx, event, complete } = setup([invalid, result()]);
      const answer = await handlers.session_before_compact(event, ctx);
      expect(answer.compaction.details.attempts).toBe(2);
      expect(complete).toHaveBeenCalledTimes(2);
      expect(complete.mock.calls[1][1].messages[0].content).toEqual(
        complete.mock.calls[0][1].messages[0].content,
      );
      expect(complete.mock.calls[1][1].systemPrompt).toContain('1440 tokens');
    },
  );

  it('stops after two failures without committing or falling back', async () => {
    const { handlers, ctx, event, complete } = setup([
      result('', 'length'),
      result('', 'length'),
    ]);
    expect(await handlers.session_before_compact(event, ctx)).toEqual({
      cancel: true,
    });
    expect(complete).toHaveBeenCalledTimes(2);
    expect(ctx.abort).toHaveBeenCalledOnce();
    expect(ctx.ui.notify).toHaveBeenCalledWith(
      expect.stringContaining('original context preserved'),
      'error',
    );
  });

  it('retries thrown provider failures', async () => {
    const { handlers, ctx, event, complete } = setup([]);
    complete
      .mockRejectedValueOnce(new Error('connection lost'))
      .mockResolvedValueOnce(result());
    expect(
      (await handlers.session_before_compact(event, ctx)).compaction.summary,
    ).toBe(checkpoint);
  });

  it('honors cancellation without retry or false success', async () => {
    const { handlers, ctx, event, controller, complete } = setup([]);
    complete.mockImplementationOnce(() => {
      controller.abort();
      return result();
    });
    expect(await handlers.session_before_compact(event, ctx)).toEqual({
      cancel: true,
    });
    expect(complete).toHaveBeenCalledOnce();
    expect(ctx.ui.notify).not.toHaveBeenCalled();
  });

  it('bounds manual compaction too, without aborting another run on failure', async () => {
    const { handlers, ctx, event } = setup([result(''), result('')]);
    event.reason = 'manual';
    expect(await handlers.session_before_compact(event, ctx)).toEqual({
      cancel: true,
    });
    expect(ctx.abort).not.toHaveBeenCalled();
  });

  it('scales the output budget to small context windows and model caps', async () => {
    const { handlers, ctx, event, complete } = setup();
    ctx.model.contextWindow = 8192;
    ctx.model.maxTokens = 1024;
    await handlers.session_before_compact(event, ctx);
    expect(complete.mock.calls[0][2].maxTokens).toBe(512);
  });

  it('permits repeated successful compactions but stops no-progress loops', async () => {
    const { handlers, ctx, event, complete } = setup([]);
    complete.mockResolvedValue(result());
    for (let round = 0; round < 10; round++) {
      ctx.sessionManager.buildContextEntries.mockReturnValue([
        { type: 'message', message: 'x'.repeat(80000) },
      ]);
      const answer = await handlers.session_before_compact(event, ctx);
      event.preparation.previousSummary = answer.compaction.summary;
      ctx.sessionManager.buildContextEntries.mockReturnValue([
        { type: 'message', message: checkpoint },
      ]);
      await handlers.session_compact(event, ctx);
    }
    expect(ctx.abort).not.toHaveBeenCalled();
    await handlers.session_before_compact(event, ctx);
    await handlers.session_compact(event, ctx);
    expect(ctx.abort).toHaveBeenCalledOnce();
  });
});
