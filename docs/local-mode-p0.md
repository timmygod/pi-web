# Local Mode P0: Context Stability and Local-Model UX

## Purpose

Pi Web was originally designed around cloud models. This fork also needs to support local models, which are generally slower, more sensitive to long context, more likely to stall or crash, and often constrained to one practical active session because of memory pressure.

The P0 goal is **not** to replace or break the original cloud workflow. Instead, add a session-level **Local Mode** and put local-model-specific behavior behind that mode. Cloud Mode should preserve upstream behavior unless a change is a clearly safe bug fix that does not alter cloud semantics.

The core reliability principle is:

> Local model sessions must not die merely because context grew too large. Prevent the problem early, recover automatically if prevention fails, and always leave a manual recovery path.

---

## 1. Session Mode: Auto / Local / Cloud

### New Session UI

Add a `Mode` field directly below `Model` on the New Session UI:

- `Auto` (default)
- `Local`
- `Cloud`

### Auto detection

When `Mode = Auto`, choosing/changing the model should resolve the effective mode automatically.

Requirements:

- Prefer provider / endpoint information over model-name heuristics.
- A local or LAN/OpenAI-compatible endpoint should resolve to Local when that can be determined reliably.
- Do not hard-code specific model names as the primary detection mechanism.
- If the user manually chooses `Local` or `Cloud`, that choice has highest priority and must not be silently overridden by later auto-detection.
- Switching back to `Auto` re-enables automatic resolution.

### Persistence

Mode is session-level state and must survive:

- page reloads,
- reopening an old session,
- Pi Web restart,
- remote access/reconnect.

Persist enough information to distinguish the configured mode (`Auto` / `Local` / `Cloud`) and the effective resolved mode when needed.

All requirements below apply only to effective **Local Mode** unless explicitly noted otherwise.

---

## 2. P0 Context Stability: proactive compaction at 65%

Local models can become dramatically slower as active context grows. Therefore Local Mode must compact much earlier than the normal cloud-oriented behavior.

### Trigger rule

Use **context utilization percentage**, not a fixed token-number threshold:

```text
contextUsage = currentContextTokens / modelContextWindow
```

When Local Mode reaches **65%**, compaction must be triggered before another model/provider request is allowed to proceed.

### Required check timing

The 65% rule must be checked at every relevant boundary, including **inside a long autonomous tool-call loop**.

In particular, the check must run:

- before a normal new model request,
- after tool results are appended and before the next assistant/model response,
- during long uninterrupted chains of `toolUse`, not only after the whole agent run ends,
- after a user prompt when applicable.

Do not rely on `agent_end` / a new user prompt as the only opportunities to compact.

### Upstream compaction bug to account for

Pi previously had a bug where automatic compaction was not checked mid-loop during long autonomous tool-calling sessions. The relevant upstream report/fix is:

- issue: `earendil-works/pi#8884`
- fix: `earendil-works/pi#8782` (`fix(coding-agent): compact before post-tool model requests`)

That fix moved threshold compaction into the between-tool-turn path (`prepareNextTurnWithContext`) and was reported working in Pi 0.84.4.

Codex must inspect the currently supported Pi behavior/version and make sure Local Mode's 65% guarantee is true in practice. **Percentage-based thresholding is not sufficient if the check itself is skipped mid-loop.**

### Loop protection

Compaction must not enter a tight repeat loop.

After compacting:

- recompute utilization,
- continue normally if utilization was reduced,
- do not repeatedly compact the same unchanged state forever,
- bound automatic recovery attempts when compaction cannot make progress.

---

## 3. Context-overflow fallback: compact and retry

The 65% proactive rule is the first line of defense. If it fails for any reason, Local Mode needs an explicit fallback.

If a model/provider/tool execution path fails because of a context-related error, for example:

- context overflow,
- maximum context length exceeded,
- prompt too long,
- provider-equivalent context-limit errors,

Pi Web must not simply surface the error and leave the session stuck.

Required behavior:

```text
context-related failure
    -> stop the normal failed path
    -> force compaction
    -> if compaction succeeds, retry/resume the interrupted run
```

Rules:

- Treat context failure as a recovery condition, not an ordinary tool error to be fed back into an already-full context.
- Automatic compact+retry must be bounded. Do not create an infinite overflow -> compact -> retry loop.
- Record enough recovery state/reason to distinguish one incident from another and to support watchdog recovery below.

---

## 4. Manual emergency recovery: Force Compact

The existing context/details popover is useful and should remain.

In Local Mode, add a prominent action at the top:

**Force Compact**

Behavior:

- Ignore the current percentage and explicitly request compaction.
- Reuse the same underlying compaction mechanism as automatic compaction where possible.
- After completion, refresh the displayed context usage.
- This is an emergency recovery control and must not become useless merely because the session is currently marked running.
- If a currently running/stuck generation must be safely interrupted before compaction, implement the safest supported sequence and restore the session to a usable state afterward.

Do **not** call this action `Reset Context`; it must compact/summarize context, not wipe the conversation.

---

## 5. Context percentage UI must not flash stale session data

Current behavior can briefly show the previous session's context percentage after navigating to another session, then replace it with the new session's value.

In Local Mode this must not happen.

On session change:

- immediately invalidate/clear the old session's context-usage UI state,
- show an empty/loading state (`—` is acceptable) until the selected session's usage is known,
- only apply asynchronous usage results if they still belong to the currently selected session,
- never render `previous session % -> current session %` as a transient flash.

If the underlying stale-state bug can be fixed globally without changing cloud semantics, a global fix is acceptable; otherwise gate the behavior to Local Mode.

---

## 6. Local reasoning-block actions

For reasoning/thinking blocks in Local Mode, simplify the three existing actions.

Current actions:

- Fork
- Label
- Copy (currently copies a session permalink containing IDs such as `leafId` / `targetId`)

Local Mode behavior:

- Hide `Fork`.
- Hide `Label`.
- Keep only `Copy`.
- Redefine `Copy` to copy the **complete textual content of that reasoning block** to the system clipboard.
- Do not copy a session URL, permalink, `leafId`, `targetId`, internal message IDs, button text, or other UI metadata.
- Copy must still copy the full reasoning block when the block is visually collapsed.
- Provide normal short success feedback such as `Copied` / checkmark.

Do not delete the underlying upstream Fork/Label functionality globally; Cloud Mode should keep the original behavior.

---

## 7. Two-layer watchdog / recovery model

### Existing outer watchdog

On macOS, Pi Web is already restarted by `launchd` because `init/com.pi-web.plist` uses `KeepAlive=true`.

Keep that responsibility simple:

- `launchd`: process-level supervisor; restart Pi Web if the Pi Web process dies.
- Pi Web Local Mode: session-level recovery; determine whether the active local session needs rescue.

Do not try to put session intelligence into `launchd` itself.

### Runtime session recovery

Pi Web already has a running-status transition path. When a Local Mode session transitions from `running` to `idle`, inspect **that session only** as a possible recovery candidate.

Automatic rescue should require all of the following:

1. Effective mode is Local.
2. The session was running and has just become idle / settled unexpectedly.
3. Context utilization is at or above **99%** (or effectively full if the provider reports an equivalent terminal value).
4. The immediately preceding run contains evidence of a context-overflow/context-limit failure.
5. This specific failure incident has not already been watchdog-recovered.

Recovery behavior:

```text
force compact
    -> if compaction succeeds
    -> send user message: "continue if possible"
```

Use exactly that simple continuation intent; do not invent a large recovery prompt that burns more context.

### Never revive all historical sessions

A local machine may have many old sessions but enough memory for only one practical active local model session.

Hard rule:

> The watchdog must never scan old sessions and revive every matching session.

During normal runtime, only consider the session that just transitioned from running to idle.

### Pi Web restart / startup recovery

A process crash may lose the in-memory running -> idle transition. On startup, recovery may inspect **at most one** Local Mode candidate: the most recently actually-active local session, not merely the newest-created session.

Preferred behavior:

- persist/track the identity of the last active Local Mode session or an equivalent single recovery candidate,
- on startup inspect only that candidate,
- require the same 99% + recent context-overflow evidence before recovery,
- if there is no trustworthy single candidate, do not broadly revive historical sessions.

Hard safety limit:

```text
max concurrent / startup watchdog recovery sessions = 1
```

### Recovery loop protection

Persist a per-incident recovery marker/state so this cannot happen forever:

```text
overflow -> compact -> continue -> overflow -> compact -> continue -> ...
```

One failed recovery incident must eventually stop and remain available for manual intervention instead of endlessly consuming local compute/memory.

---

## 8. Conditional investigation: reasoning language

There is a suspected bug where using Chinese reasoning appears correlated with compaction not triggering and the session dying. **Do not assume Chinese is the cause.**

The upstream mid-loop compaction issue above is a plausible alternative explanation.

### Investigation first

After implementing/verifying the percentage-based 65% mid-loop behavior, perform an A/B reproduction with the same model/task/context settings:

- A: English reasoning
- B: Chinese reasoning

Record/compare at minimum:

- context utilization over time,
- tool-call / `toolUse` sequence,
- compaction checks and triggers,
- compaction reason,
- overflow events,
- turn/agent settlement behavior.

### Decision rule

**Case A: the problem disappears after the compaction fix**

- Treat language as not proven causal.
- Do not add a reasoning-language system prompt.

**Case B: a language-specific failure is reproducible and the root cause is identified/fixed**

- Fix the root cause first.
- Stress-test Chinese reasoning with long tool-call loops and automatic compaction.
- Only after it is demonstrably stable, add a system instruction equivalent to:

```text
Reasoning/thinking should use {language}.
```

`{language}` must come from the current Pi Web UI language setting.

**Case C: no reliable root cause can be found**

- Do not implement reasoning-language forcing.
- Keep the current/default English reasoning behavior.

Reliability wins over localized reasoning output.

---

## 9. Cloud Mode compatibility / non-goals

P0 is intentionally scoped.

Do not:

- remove upstream Cloud Mode features globally,
- globally remove Fork/Label/permalink behavior,
- turn Pi Web into a single-session-only application,
- revive multiple old local sessions on startup,
- add a reasoning-language prompt without proving it is safe,
- wipe/reset conversation history as a substitute for compaction,
- create unbounded auto-retry/recovery loops.

Local Mode may optimize for the common case of one active local session, but the architecture must remain session-correct and must not corrupt behavior if multiple sessions exist.

---

## 10. Acceptance criteria

P0 is complete only when the following are demonstrated with tests where feasible and manual/stub integration coverage where required:

1. New Session exposes `Auto / Local / Cloud` under Model; manual override persists and wins over auto-detection.
2. Local Mode auto-compaction is percentage-based and triggers at **65%**.
3. The 65% check runs before subsequent provider calls inside long uninterrupted tool-call loops.
4. A context-limit failure triggers bounded forced compaction and a bounded retry/resume path rather than leaving the session dead.
5. Context details expose **Force Compact** and it remains a usable emergency action for a stuck/running local session.
6. Switching sessions does not flash another session's context percentage.
7. Local reasoning blocks show only `Copy`, and Copy puts the full reasoning text (not a permalink) on the clipboard.
8. A Local Mode session that unexpectedly becomes idle at >=99% after a context-overflow incident can be recovered by compact + `continue if possible`.
9. Runtime recovery only touches the session that just failed; startup recovery inspects at most one most-recently-active Local Mode candidate.
10. Recovery cannot loop forever or revive a backlog of historical sessions.
11. Cloud Mode retains existing upstream behavior.
12. Reasoning-language forcing is implemented **only** if the Chinese-specific failure is reproduced, root-caused, fixed, and stress-tested; otherwise it is intentionally omitted.

---

## 11. Implementation guidance for Codex

Before structural changes, follow `AGENTS.md` and read the relevant architecture docs under `docs/`, especially the session/RPC/status flow. Useful existing areas include:

- `internal/rpc/worker.go`
- `internal/workers/manager.go`
- `internal/server/status.go`
- `internal/server/status_sweeper.go`
- `internal/server/chat.go`
- `internal/server/new_session.go`
- `web/src/components/session/chat/ContextUsage.svelte`
- `web/src/components/session/chat/context-usage.js`
- `web/src/session/render/session-entry-actions.js`
- New Session components under `web/src/components/index/`

Do not blindly follow these filenames if the current code has moved; trace the current behavior first.

Prefer behavior-driven tests for the failure modes above. Run the repository-required validation from `AGENTS.md` (`make test`, `make check`, and the appropriate build/E2E coverage) before considering the task complete.
