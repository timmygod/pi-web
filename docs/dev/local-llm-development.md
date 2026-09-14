# Local-model edition development

This checkout is a maintained pi-web edition for locally deployed and
LAN-hosted language models. It follows the upstream pi-web project for shared
functionality, while keeping a separate release line for local-model behavior,
testing, and operational safeguards.

The two lines are intentionally related but independent:

- upstream remains the source for broadly useful pi-web features and fixes;
- this edition periodically imports upstream changes;
- local-model changes are reviewed and released on this line without rewriting
  upstream history or tags;
- Cloud Mode should retain upstream behavior unless a change is a clearly safe,
  compatibility-preserving bug fix.

## Repository setup

Keep the original project as an `upstream` remote and your maintained fork as
`origin`:

```bash
git remote -v
git remote add upstream https://github.com/ygncode/pi-web.git  # once, if absent
git fetch --prune upstream
```

Use a dedicated branch for local-model work. Do not develop directly on the
temporary branch used to import upstream changes.

## Periodic upstream synchronization

Synchronize deliberately and review the result before merging it into the
local release branch:

```bash
git fetch --prune upstream
git switch main
git merge --no-ff upstream/main

# Review conflicts and local-mode boundaries before completing the merge.
make check
```

If the upstream default branch has a different name, substitute that ref. Pay
special attention to changes touching:

- `internal/rpc/` and `internal/workers/` — worker lifecycle and RPC behavior;
- `internal/server/chat.go` and `internal/server/local_mode.go` — compaction,
  recovery, and Cloud/Local isolation;
- `web/src/components/session/` — context usage, mode controls, and recovery UI;
- `docs/` — architecture and acceptance criteria that describe the contract.

When a conflict changes behavior, resolve the code and update the matching
English documentation in the same change. Never silently drop a local-model
guard just to make an upstream merge clean.

## Parallel release workflow

Keep upstream and local releases distinguishable. Import upstream first, then
apply and test local changes on the local release branch:

```bash
git switch main
git merge --no-ff upstream/main
git switch local-llm
git merge --no-ff main

make check
make build
git diff --check
```

Publish local tags/releases using the repository's local-edition naming
convention; never move or overwrite an upstream tag. Each release should state
the upstream revision it contains and call out local-model changes, migration
notes, and any known limitations.

## Local development

Install the normal build prerequisites, then build and test the source tree:

```bash
make setup
make build
make test
make check
```

For an interactive source checkout alongside an installed instance, use the
development harness:

```bash
make dev
```

The development instance listens on port `31416`, shares session data with the
installed instance, and disables autonomous background jobs. Do not drive the
same session from both instances at once; each process has its own RPC worker.
The user runs the server; documentation and test work should not start,
restart, or deploy the production instance.

## Local-model verification

Configure the local provider through pi, then verify the session in all three
mode states:

1. **Auto** detects a local or LAN endpoint when provider metadata makes that
   reliable.
2. **Local** applies the isolated worker settings and the model's context
   window.
3. **Cloud** keeps the upstream path and behavior.

Before releasing local-model changes, exercise the Local Mode acceptance path:

- proactive compaction at 65% before the next provider request;
- compaction checks during long tool-use loops;
- bounded rolling checkpoints and one tighter rewrite on invalid/capped output;
- exhausted retries or no-progress compaction stopping the current run;
- context-overflow recovery with one bounded continuation;
- Force Compact and refresh of post-compaction usage;
- one watchdog recovery at a time, held through `agent_settled`;
- no automatic revival of a backlog of historical sessions;
- mode changes rejected while a worker is running;
- Cloud Mode regression tests and the normal upstream UI behavior.

Use the focused tests while iterating, then run the complete Go, frontend, and
build checks before publishing:

```bash
go test ./internal/rpc ./internal/workers ./internal/server
go test ./...
cd web && npm run lint && npm run format:check && npm test
cd .. && make build
```

## Documentation and translations

English is the source of truth. Update `README.md`, `user-docs/en/`, and
`docs/` first. Do not hand-edit generated translations. Regenerate them with:

```bash
python3 scripts/build_readmes.py
python3 scripts/build_userdocs.py
```

The translation scripts run through pi's local OpenAI-compatible provider and
default to `llama-cpp/AtomicChat/Qwen3.8-Flash-Next-AD-4.27bpw-Q4_K_M-M64`.
Start the local inference server at `http://127.0.0.1:8080/v1` and register
that provider/model in `~/.pi/agent/models.json` before running them. The
scripts disable tools, extensions, context files, and reasoning for these
single-turn translation requests, so no cloud API key is needed. To use another
configured local model, set `PI_TRANSLATION_MODEL` to its full `provider/model`
reference, for example:

For Qwen models, include `"thinkingFormat": "qwen"` in the provider `compat`
settings so pi maps `--thinking off` to the model's `enable_thinking: false`
option. This keeps translation output focused on the translated text instead
of emitting a long reasoning trace.

```bash
PI_TRANSLATION_MODEL='llama-cpp/your-model-id' python3 scripts/build_readmes.py zh
```

Review generated Markdown links and code blocks before committing. Keep the
edition statement, upstream synchronization policy, and local-model acceptance
criteria consistent across the root README, user guide, architecture docs, and
this page.
