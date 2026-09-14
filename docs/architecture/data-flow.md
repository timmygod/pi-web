# Data Flow & Session File Format

This data-flow reference is maintained on the local-model edition release line.
The append-only session format remains upstream-compatible, while Local Mode
worker isolation, context protection, and recovery boundaries are documented in
[Local-model edition development](../dev/local-llm-development.md).

## Session File Format

Sessions are stored as **JSONL** files (one JSON object per line):

```
~/.pi/agent/sessions/--project-name--/
└── 2026-01-15T10-30-00.000Z_a1b2c3d4.jsonl
```

### Example JSONL Content

```jsonl
{"type":"session","version":3,"id":"uuid","timestamp":"2026-01-15T10:30:00Z","cwd":"/Users/me/project","name":"My Session"}
{"type":"message","timestamp":"2026-01-15T10:30:01Z","message":{"role":"user","content":"Hello"}}
{"type":"message","timestamp":"2026-01-15T10:30:05Z","message":{"role":"assistant","content":"Hi!"},"usage":{"totalTokens":42,"cost":{"total":0.0001}}}
{"type":"session_info","timestamp":"2026-01-15T10:30:06Z","name":"Renamed Session"}
{"type":"tool_call","timestamp":"2026-01-15T10:30:06Z","tool":"bash","command":"ls -la"}
{"type":"tool_result","timestamp":"2026-01-15T10:30:07Z","tool":"bash","output":"..."}
{"type":"branch_summary","timestamp":"2026-01-15T10:35:00Z","branch":"main","summary":"..."}
{"type":"compaction","timestamp":"2026-01-15T10:40:00Z","before":"...","after":"..."}
```

### Entry Types

| `type`                  | Description                                                     |
| ----------------------- | --------------------------------------------------------------- |
| `session`               | Header metadata (cwd, name, version, id)                        |
| `message`               | User or assistant message with optional `usage` and `cost`      |
| `session_info`          | Session metadata update; latest `name` is used as display title |
| `tool_call`             | Agent invoked a tool                                            |
| `tool_result`           | Tool execution result                                           |
| `bash` / `bash_output`  | Shell command and its output                                    |
| `branch_summary`        | Summary of work on a git branch                                 |
| `compaction`            | Conversation history was compacted                              |
| `model_change`          | Model switched mid-session                                      |
| `thinking_level_change` | Thinking level changed mid-session                              |
| `diff`                  | Code diff output from edit/tool operations                      |

### Project Directory Encoding

Project names are filesystem-safe encoded:

```go
EncodeProjectName("/Users/me/project") → "--Users-me-project--"
DecodeProjectName("--Users-me-project--") → "/Users/me/project"
```

## Parse Flow

```
File on disk
     │
     ▼
sessions.ParseFile(path, dirName, fileName)
     │
     ├──▶ stream file line-by-line
     │
     ├──▶ json.Unmarshal each line
     │        ├──▶ type=="session" → sess.Header
     │        ├──▶ type=="message" → increment MessageCount, sum tokens/cost
     │        ├──▶ type=="session_info" → latest rename/title metadata
     │        └──▶ all types → append to Entries
     │
     ├──▶ Name = latest session_info.name, else session.name, else first user text, else filename
     ├──▶ Model = last message model or last model_change modelId
     ├──▶ ModelProvider = provider for last-known model
     ├──▶ LastActivity = latest timestamp (or file modtime fallback)
     │
     └──▶ ChatAvailable = cwd still exists?
```

## Cache Strategy

`sessions.Cache` avoids re-parsing unchanged files:

```
LoadAll(dir)
    │
    ├──▶ ReadDir all project subdirs
    │
    ├──▶ For each .jsonl file:
    │         ├──▶ Check modtime against cache
    │         ├──▶ MATCH → return cached SessionSummary
    │         └──▶ MISMATCH → ParseSummary + store in cache
    │
    ├──▶ Evict files no longer on disk
    │
    └──▶ SortByActivity (descending by timestamp)
```

## Data Flow: Viewing a Session

```
Browser GET /session?id=<id>
           │
           ▼
    server.handleSession → SPA shell
           │
           ▼
Browser GET /api/session?id=<id>
           │
           ▼
    server.handleApiSession
           │
           ├──▶ sessions.Cache.Resolve → find file path + parse/cache by modtime
           │
           ├──▶ sessions.ParseFile → Session struct when cache is stale
           │
           └──▶ Write JSON response for SessionPage.svelte
```

## Data Flow: Chat Message

```
Browser POST /api/chat?id=<id>
           │
           ▼
    server.handleChat
           │
           ├──▶ sessions.ResolveByID → Session + Path
           │
           ├──▶ chat.ParseRequest(r)
           │         ├──▶ ParseMultipartForm
           │         ├──▶ Extract text + image files
           │         └──▶ Validate (not empty, image size, mime type)
           │
           ├──▶ Resolve configured/effective mode from SQLite + model endpoint
           ├──▶ Acquire the per-session compact/send gate
           ├──▶ Reload the session JSONL after waiting for the gate
           ├──▶ Local idle worker + projected context >= 65%: force compact first
           ├──▶ Local running worker: steer/queue without aborting active compaction
           ├──▶ workers.Manager.Send(ctx, sessionID, sessionPath, chatReq)
           │         │
           │         ├──▶ Get or create ChatWorker for session
           │         │         └──▶ rpc.NewPiWorkerWithStream(sessionPath, streamSink)
           │         │               ├──▶ exec.Command("pi", "--mode", "rpc")
           │         │               ├──▶ Start subprocess
           │         │               ├──▶ switch_session RPC
           │         │               └──▶ Background goroutines: consume stdout, wait
           │         │
           │         └──▶ worker.Prompt(ctx, chatReq)
           │               ├──▶ BuildPromptCommand (JSONL to stdin)
           │               ├──▶ Await response on pending channel
           │               └──▶ Update status → running
           │
           └──▶ Return {"ok": true, "status": "queued"}
```

For effective Local sessions, the worker's isolated Pi settings enforce the same
65% boundary between tool results and the next model call. Pi 0.85.1 installs
this check through `prepareNextTurnWithContext`, so uninterrupted tool-use loops
do not wait for `agent_end`. A Local-only extension aborts the run if automatic
compaction fails or reduces the active message context by less than 5%/512
tokens, preventing Pi from sending the unchanged context to the provider.
Pi Web serializes server-initiated compact/send preflight per session and reloads
the session after acquiring that gate. This prevents concurrent chat, queue,
scheduler, watchdog, and Force Compact requests from running Pi's non-reentrant
manual compaction against one session. A chat arriving while Pi is already
running is submitted as steering input; it does not abort an automatic
compaction. The cancel endpoint remains outside the gate so an explicit cancel
can still interrupt a long operation.
Pi's native context-overflow path removes the failed
assistant response, compacts, and retries once. The pi-web watchdog is a second
layer: a running→idle transition with both >=99% utilization and a context-limit
error is compacted once per incident, then sent exactly `continue if possible`.
The same transition also catches a Local assistant `stop` containing reasoning
but no response text or tool call. That path skips compaction, emits a live
recovery notice, and sends the same bounded continuation. Both paths share
persisted per-incident deduplication and a progress-aware circuit breaker: an
immediate repeat failure is blocked, while a new failure after a completed
answer, multiple successful tool results, or sustained successful work can be
rescued again. A user action also resets the breaker. The global watchdog
recovery slot remains held until Pi emits `agent_settled`, not merely until it
acknowledges the continuation prompt.

## Data Flow: Rename Session

```
Browser POST /api/rename-session?id=<id>
           │
           ▼
    server.handleRenameSession
           │
           ├──▶ Decode JSON body → {"name":"New Name"}
           ├──▶ Resolve session ID → filesystem path
           ├──▶ sessions.RenameSession(path, name, now)
           │         └──▶ Append JSONL line: {"type":"session_info","timestamp":"...","name":"New Name"}
           ├──▶ record modtime + broadcast "reload" to session SSE clients
           │
           └──▶ Return {"ok": true, "name": "New Name"}
```

Rename is the only intentional pi-web write to an existing session JSONL file. It appends metadata history; it does not rewrite existing entries. Creating a new session is the other direct write path, but it only creates a fresh JSONL file.

## Data Flow: Live Reload

```
Editor saves session file
           │
           ▼
    fsnotify detects Write event
           │
           ▼
    debouncer.schedule(path)  (50ms debounce)
           │
           ▼
    Server.recordModTime(sessID, modTime)
           │
           ├──▶ Update fileMod map
           ├──▶ Broadcast "reload" to SSE clients for this sessID
           └──▶ Recompute running status → broadcast status-delta
           │
           ▼
    Browser EventSource receives "reload"
           └──▶ fetch /api/session
                └──▶ append/upsert canonical entries and clear preview
```

## Data Flow: Share to Gist

```
Browser POST /share?id=<id>
           │
           ▼
    server.handleShare
           │
           ├──▶ share.FindGh → locate `gh` CLI
           │
           ├──▶ gh auth status → verify login
           │
           ├──▶ deps.Resolve(id) → find matching session
           │
           ├──▶ renderExportSessionPage(session)  (no live chrome)
           │
           ├──▶ Write to temp file
           │
           ├──▶ gh gist create --public=false <tmpfile>
           │
           └──▶ Return {gistUrl, gistId, previewUrl}
```

## Data Flow: Create New Session

```
Browser POST /api/new-session
           │
           ▼
    server.handleNewSession
           │
           ├──▶ Decode path, optional sourceSessionId, model, and Auto/Local/Cloud mode
           │
           ├──▶ If sourceSessionId is present, read current worker model/thinking state
           │
           ├──▶ Validate path (absolute, exists or create)
           │
           ├──▶ Encode project name → create directory under sessionsDir
           │
           ├──▶ Generate UUID + timestamp → write fresh JSONL file
           │         ├──▶ session header entry
           │         └──▶ implicit model_change / thinking_level_change entries when copied
           │              from the source session. These entries include normal entry `id`
           │              and `parentId` fields so `pi --mode rpc switch_session` restores
           │              the same initial model/thinking state.
           │
           ├──▶ Persist configured/effective mode + resolved context window in SQLite
           ├──▶ Pre-initialize mode-configured chat worker (EnsureWorker)
           │         └──▶ So the session page can read default model/thinking level immediately
           │
           └──▶ Return {"ok": true, "id": <filename>}
```

## Data Flow: Fork Session

```
Browser POST /api/fork-session?id=<sourceId>
           │
           ▼
    server.handleApiForkSession
           │
           ├──▶ Decode JSON body → {"entryId": "..."}
           ├──▶ Resolve source session ID → filesystem path
           ├──▶ sessions.ForkSessionFile(sessionsDir, sourcePath, entryId, now)
           │         ├──▶ Parse source session into by-ID map
           │         ├──▶ Walk from entryId back to root (via parentId)
           │         ├──▶ Reverse to chronological order
           │         ├──▶ Create new session header with parentSession reference + forkedFrom
           │         └──▶ Write new JSONL file in same project directory
           ├──▶ Initialize worker for the new session (async)
           │
           └──▶ Return {"ok": true, "id": <newFilename>}
```

## Data Flow: Clone Session

```
Browser POST /api/clone-session?id=<sourceId>
           │
           ▼
    server.handleApiCloneSession
           │
           ├──▶ Decode JSON body → {"leafId": "..."}  (optional, defaults to last entry)
           ├──▶ Resolve source session ID → filesystem path
           ├──▶ sessions.CloneSessionFile(sessionsDir, sourcePath, leafId, now)
           │         ├──▶ Parse source session into by-ID map
           │         ├──▶ Walk from leafId back to root (via parentId)
           │         ├──▶ Reverse to chronological order
           │         ├──▶ Create new session header with parentSession reference
           │         └──▶ Write new JSONL file in same project directory
           ├──▶ Initialize worker for the new session (async)
           │
           └──▶ Return {"ok": true, "id": <newFilename>}
```

## Data Flow: Scratchpad (Notes)

```
Browser GET /api/scratchpad?project=<cwd>
           │
           ▼
    server.handleGetScratchpad
           │
           ├──▶ Query SQLite: SELECT content FROM scratchpads WHERE project_path = ?
           │
           └──▶ Return {"content": "..."}  (empty string if no notes exist)

Browser POST /api/scratchpad
           │
           ▼
    server.handleSaveScratchpad
           │
           ├──▶ Decode JSON body → {"project": "...", "content": "..."}
           ├──▶ UPSERT into SQLite scratchpads table (INSERT ... ON CONFLICT DO UPDATE)
           │
           └──▶ Return {"ok": true}
```
