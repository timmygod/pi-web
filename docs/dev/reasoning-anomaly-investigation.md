# Tool-turn reasoning anomaly investigation

Investigation date: 2026-09-10. Status: message transport and template checks
completed; controlled generation experiments remain pending while the single
inference slot is serving an active task.

## Symptom and current conclusion

During a continuing coding task, some assistant thinking blocks claim that no
substantive user task or previous analysis exists. The same assistant messages
still contain tool calls related to the ongoing task.

The evidence strongly localizes the symptom to the model/inference side of the
pipeline. It does not support a pi-web display error or context truncation as the
explanation for the reported occurrence. The evidence does **not** yet distinguish
model behavior from an inference-runtime/cache issue, and does not establish a
sampling or quantization cause. Calling this a proven training-data artifact
would be premature.

No application code, model settings, running services, or existing session files
were changed during this investigation. Diagnostic scripts and private payload
copies were written under `/tmp`; this report excludes conversation transcripts.

## Evidence snapshot

The local session scan found 89 distinct assistant entries matching a deliberately
narrow family of phrases about absent prior analysis or an imagined initial
software-engineer instruction. These occurred in 10 Jarvis sessions, across 6,836
thinking-bearing assistant entries recorded with the AtomicChat model identifier.
The approximately 1.3% figure is a keyword-match rate, not a measured failure rate;
the matcher can miss other wording and does not establish task correctness.

- 88 matches immediately followed a tool-result entry in the session file.
  The remaining match followed a compaction entry.
- All 89 matching assistant entries contained tool calls. This demonstrates
  continued tool use, not that every individual action was correct.
- Four affected sessions had a matching entry before their first compaction.
- The recorded total-token counts ranged from 26,736 to 121,464. Compaction and
  approaching the context limit are therefore not necessary conditions.
- Reconstructed contexts contained no matching phrase in their non-assistant
  conversation messages. In 31 cases, no matching earlier assistant thinking was
  present either. Repetition of earlier anomalous thinking is a plausible
  amplifier, but cannot explain every first occurrence.
- 48 records could be uniquely correlated with the available llama.cpp log by
  uncached input tokens, output tokens, and final token count. All 48 had
  `truncated = 0`. This excludes reported context shifting for those matching
  records, not every possible internal inference fault.

The specific user-reported record is session
`a41fbdba-1820-4eb2-aa5c-29a39c27318e`, entry `a7968c70`, at
`2026-09-10T12:03:04.593Z`:

| Field | Recorded value |
| --- | --- |
| Provider/API | `llama-cpp` / `openai-completions` |
| Model | `AtomicChat/Qwen3.8-Flash-Next-AD-4.27bpw-Q4_K_M-M64` |
| Thinking signature | `reasoning_content` |
| Stop reason | `toolUse` |
| Uncached input / cached input / output | 479 / 57,331 / 183 tokens |
| Total tokens | 57,993 |
| Matching server task | `2445242` |
| Server final tokens / truncation | 57,992 / `0` |

The server's final token count is one less than the API usage total. Both the
input/output counts and the neighboring response's counts also match the log.

## Checks performed

### Stored content and rendering

The reported text is already in the session JSONL's assistant `thinking` block.
`web/src/components/session/SessionEntry.svelte` renders `block.thinking` as text.
It does not construct the anomalous prose. The phrase was not found as a canned
message in the project or the installed Pi bundle.

### Pi message reconstruction and parser

The currently installed Pi version is `0.85.1`. Its bundled
`buildSessionContext`, `convertToLlm`, and OpenAI-completions `stream` functions
were used to reconstruct all 89 pre-response conversation bodies. A diagnostic
`onPayload` hook captured each body and stopped before network access; a separate
fetch guard rejected any accidental network call.

All 89 reconstructions preserved their non-empty assistant thinking blocks. The
reported example reconstructed 107 conversation messages: one user-role
compaction summary, 53 assistant messages, and 53 tool messages. All 53 thinking
blocks and the final tool result survived conversion; three earlier thinking
blocks already matched the anomaly pattern.

A separate synthetic SSE test against the installed provider adapter verified
that split `reasoning_content` deltas are preserved, tool arguments are parsed,
and a second request does not inherit the first request's reasoning buffer.

These are current-code reconstruction tests, **not historical HTTP captures**.
They do not recover historical system prompts, active tool definitions,
extension mutations, or the exact Pi version used by every older worker.

### Running server template

The running server reports build `b1-c9ca51c`, one inference slot, and a
131,072-token context. The source checkout's HEAD differs from the running build;
runtime endpoint results take precedence over the current checkout.

The `/apply-template` endpoint was used without inference. This endpoint uses the
same chat-parameter parser as chat completions. Synthetic user/assistant/tool
messages produced proper role boundaries, a tool-response wrapper, preserved
reasoning, real newlines, and the expected assistant thinking prefix. A suspected
double-escaped-newline problem was disproved by the rendered result.

Two reconstructed historical message bodies were also rendered:

| Entry | Messages | Prior thinking blocks | Result |
| --- | --- | --- | --- |
| `d7591eaa` | 41 | 18 | All thinking and tool-result text preserved |
| `a7968c70` | 107 | 53 | All thinking and tool-result text preserved |

The second probe is not an exact historical full prompt: historical system and
tool-definition text were not available. It tests preservation and structure of
the recovered conversation body.

Preservation probes confirmed:

| Template setting | Thinking before latest user query | Thinking within current tool loop |
| --- | --- | --- |
| Default | Preserved | Preserved |
| `preserve_thinking: true` | Preserved | Preserved |
| `preserve_thinking: false` | Removed | Preserved |

Thus, simply disabling `preserve_thinking` does not remove the anomalous thinking
between tool calls in the same user turn. It is not a demonstrated remedy.

### Sampling and upstream guidance

The server reports thinking-mode defaults of temperature 1.0, top-p about 0.95,
top-k 20, min-p about 0.05, and repetition penalty 1.0. The
[Qwen model card](https://huggingface.co/Qwen/Qwen3.8-Flash-Next#api-usage)
recommends min-p 0.0 with otherwise matching values. It also supports preserving
thinking and the `medium` reasoning effort used here.

The min-p difference is a testable configuration discrepancy, **not an established
cause**. There is no controlled evidence here that changing it will eliminate
this symptom. The
[official template](https://huggingface.co/Qwen/Qwen3.8-Flash-Next/blob/main/chat_template.jinja)
also preserves thinking in the current tool loop when historical preservation is
disabled.

## Remaining causal tests

The active task continuously occupies the only model slot. Inserting unrelated
generation can evict its prompt cache and add substantial re-prefill latency.
The user has been asked whether to perform that test immediately or wait for the
task to finish. No generation experiment has been run under that uncertainty.

The next experiment should capture raw HTTP responses without executing any
returned tool calls, using bounded output, fixed paired seeds, and a recovered
context with no earlier anomaly. Compare:

1. Current settings as the baseline, across multiple seeds.
2. The same requests with only min-p changed to 0.0.
3. If reproducible, a raw completion on the rendered prompt to separate generated
   text from chat-response parsing.
4. If needed, a cold-cache versus warm-cache comparison and an independently
   verified runtime/model build. These require additional compute and must not
   be silently substituted for the user's working deployment.

Failure to reproduce in a small sample would not clear the model or runtime:
the observed keyword-match rate is low and the complete historical request and
random state are unavailable. A definitive attribution requires a reproducing
request or a captured new occurrence plus controlled comparisons.

UI suppression would only hide the symptom. No string blacklist, prompt patch,
global thinking disablement, or unverified sampling change has been applied.
