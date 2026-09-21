# Sequence Flow: Schedules

This flow applies to the local-model edition; preserve the Local Mode boundary
when synchronizing shared runtime changes. See [Local-model edition development](../dev/local-llm-development.md).

Schedules run pi automatically on a cadence (or on demand). When a schedule
fires it creates a **fresh pi session**, sends the schedule's instructions as the
first message, and lets pi run autonomously. Each firing is recorded so the
created sessions can be tagged, filtered, and surfaced in a run log — and so a
schedule-specific push notification can be sent when the run finishes.

Scheduling state lives in SQLite (`pi-web.sqlite`), not in pi's session files.
The `internal/schedules` package owns the store and the cron math; the firing
loop and the session-creating runner live in `internal/server` because they need
the chat workers and SSE broadcast.

## Data model (SQLite)

`schedules` — one row per definition:

| column | notes |
|--------|-------|
| `id` | UUID |
| `name`, `instructions` | required |
| `model_provider`, `model_id` | legacy exact model selection; optional → pi defaults |
| `model_selector` | API model strategy: `local`, `free`, or a model name |
| `project_path` | optional → user home dir |
| `cron_expr` | empty = manual (Run-now only) |
| `run_at` | optional RFC3339 instant for a one-shot run; mutually exclusive with cron |
| `timezone` | IANA name; empty = server local |
| `callback_url` | optional HTTP(S) completion webhook |
| `enabled` | bool |
| `last_run_at` | last fire time |

`schedule_runs` — one row per firing; also the **session → schedule mapping**:

| column | notes |
|--------|-------|
| `schedule_id` | FK |
| `session_id` | created session UUID (filled after resolve) |
| `session_file` | created `.jsonl` filename |
| `fired_at`, `status`, `error`, `result` | `running` \| `succeeded` \| `failed` \| `cancelled` \| `skipped` |
| `completed_at`, `model_provider`, `model_id` | terminal time and actual selected model |
| `callback_status`, `callback_attempts`, `callback_error` | webhook delivery audit fields |

## Firing sequence

```
┌──────────┐   ┌──────────────┐   ┌──────────────┐   ┌────────────┐   ┌──────────┐
│ scheduler│   │ schedules    │   │   sessions   │   │  workers   │   │   push   │
│  (loop)  │   │  (store)     │   │ (create file)│   │ (manager)  │   │ (manager)│
└────┬─────┘   └──────┬───────┘   └──────┬───────┘   └─────┬──────┘   └────┬─────┘
     │                │                  │                 │               │
     │ every 30s: evaluateSchedules()    │                 │               │
     │── List() ─────▶│                  │                 │               │
     │◀── schedules ──│                  │                 │               │
     │                │                  │                 │               │
     │ for each enabled cron schedule:   │                 │               │
     │   next = NextFire(cron, tz, now)  │                 │               │
     │   (first sight only arms it —     │                 │               │
     │    missed past runs are skipped)  │                 │               │
     │                │                  │                 │               │
     │ when now >= next → fireSchedule() │                 │               │
     │── RecordRun(running) ────────────▶│                 │               │
     │── SetLastRun ────────────────────▶│                 │               │
     │── CreateSessionFileWithSettings ───────────────────▶│               │
     │   (project dir or home; implicit model/thinking     │               │
     │    entries are not restored on an empty session)    │               │
     │◀── filename ──────────────────────────────────────│               │
     │── ResolveByID ────────────────────────────────────▶│               │
     │◀── session UUID + path ───────────│                 │               │
     │── AttachSession(runID, uuid) ────▶│                 │               │
     │                │                  │                 │               │
     │── EnsureWorker(uuid, path) ───────────────────────▶│               │
     │── SetModel / SetThinkingLevel (when configured) ──▶│               │
     │── Send(uuid, path, {instructions}) ───────────────▶│─── pi runs ──▶│
     │                │                  │                 │               │
     │  (file watcher sees the new .jsonl → broadcasts `new-session`)     │
     │                │                  │                 │               │
     │  [run completes; worker → idle]   │                 │               │
     │  recomputeAndBroadcastStatus: running → idle        │               │
     │  scheduleNameForSession(uuid)? ──▶│                 │               │
     │◀── name, true ────────────────────│                 │               │
     │── NotifyScheduleDone(name, uuid) ─────────────────────────────────▶│
     │                │                  │                 │   web push ──▶ browser
```

## Manual / Run-now

A schedule with both `cron_expr` and `run_at` empty never fires on the timer.
Any schedule can be fired immediately via `POST /api/schedule/run?id=<id>`,
which calls the same `fireSchedule` path and returns the created `sessionId` so
the UI can navigate to it.

## One-shot schedules

Set `runAt` to an RFC3339 timestamp and leave `cronExpr` empty. When it becomes
due, pi-web atomically disables the schedule before dispatching the run. It
therefore executes at most once and remains available for run-history queries;
Jarvis does not need to delete it after receiving the callback. If pi-web was
offline at the requested instant, the enabled one-shot fires once when pi-web
next evaluates schedules.

## Missed runs

Schedules only fire while pi-web is running. On startup (and on first sight of
any recurring schedule) the loop computes the next fire time from *now*, so
cron occurrences that elapsed while the process was down are **skipped** rather
than replayed. An enabled overdue one-shot is different: it fires once after
startup, then disables itself.

## HTTP endpoints

| Method | Path | Purpose |
|--------|------|---------|
| GET | `/api/schedules` | list (with computed `nextRunAt`) |
| POST | `/api/schedules` | create |
| GET | `/api/schedule?id=` | read one |
| POST/PUT | `/api/schedule?id=` | update |
| DELETE | `/api/schedule?id=` | delete (and its runs) |
| POST | `/api/schedule/run?id=` | Run-now |
| GET | `/api/schedule/runs?id=` | run log |

The `/schedules` page itself is the SPA shell (served by the catch-all index
route); the Svelte router renders `SchedulesPage.svelte`. Create/update/delete
(and run-now) broadcast an SSE `schedules` event on `__all__` so an open
schedules page refetches. Agents can create schedules via `/skill:pi-web-schedule`
(`pi-web-ctl`); see [skills.md](./skills.md).

## Agent API contract

Create a schedule with `POST /api/schedules`. Authentication is the same as all
other pi-web API routes. The existing `modelProvider` + `modelId` pair remains
supported, but API integrations should normally use the single `model` field:

- `local`: prefer the most-used available local/LAN model, then other local
  models; available non-local models are fallback candidates.
- `free`: prefer starred models exposed by the model source, then recently and
  frequently used models.
- any other string: exact provider/id, id, or display-name matches first,
  followed by partial matches and recent known-working models.

Candidates are tried in order. A model that is no longer available is omitted;
if setting a candidate fails, the next candidate is tried. The actual model is
recorded on the run as `modelProvider` and `modelId`.

```json
{
  "name": "Jarvis reminder",
  "instructions": "Remind me to take a break.",
  "model": "local",
  "runAt": "2026-09-20T13:40:00Z",
  "callbackUrl": "http://127.0.0.1:8088/pi-web/callback",
  "enabled": true
}
```

Delete it with `DELETE /api/schedule?id=<schedule-id>`. `callbackUrl` must be an
absolute `http` or `https` URL. Treat it as sensitive configuration because
pi-web will make a server-side request to it.

### Callback

pi-web sends an HTTP `POST` with `Content-Type: application/json`,
`User-Agent: pi-web-scheduler/1`, and `X-Pi-Web-Event-ID`. A `2xx` response
acknowledges delivery. Other responses and network errors are retried up to
three times with short backoff. Consumers must deduplicate by `eventId` because
delivery is at-least-once across retries.

```json
{
  "version": "1",
  "event": "schedule.run.completed",
  "eventId": "schedule-run-42",
  "schedule": { "id": "…", "name": "Jarvis reminder" },
  "run": {
    "id": 42,
    "scheduleId": "…",
    "sessionId": "…",
    "firedAt": "2026-09-20T08:00:00Z",
    "completedAt": "2026-09-20T08:00:08Z",
    "status": "succeeded",
    "result": "Time to take a break.",
    "modelProvider": "ollama",
    "modelId": "qwen3",
    "callbackStatus": "",
    "callbackAttempts": 0
  }
}
```

`status` meanings:

| status | meaning |
|---|---|
| `succeeded` | an assistant response completed normally; `result` is its text |
| `failed` | setup/model/send failed or the assistant ended with an error; inspect `error` |
| `cancelled` | execution ended with an aborted response |
| `skipped` | reserved for a deliberately suppressed occurrence, such as a future overlap policy |

The callback payload describes delivery *before* the current attempt, so its
`callbackStatus` is normally empty. Query `GET /api/schedule/runs?id=<id>` for
the authoritative delivery state: `delivered` or `failed`, attempt count, and
the last delivery error.

### Relative reminders

For “remind me in one minute”, Jarvis computes an absolute RFC3339 instant from
its current clock and sends it as `runAt`; cron conversion and callback-time
cleanup are not needed:

```json
{
  "name": "Bathroom reminder",
  "instructions": "提醒我去拉屎",
  "model": "free",
  "runAt": "2026-09-20T13:40:00Z",
  "callbackUrl": "http://127.0.0.1:8088/pi-web/callback"
}
```

The create response returns the canonical UTC value in `schedule.runAt` and
`schedule.nextRunAt`. After firing, `enabled` becomes `false`; keep or delete
the schedule depending on whether Jarvis needs its history.

## Push notifications

Scheduled runs reuse the web-push subsystem ([share.md](./share.md) covers VAPID
setup). On the running→idle transition, `recomputeAndBroadcastStatus` checks
whether the session was schedule-created; if so it sends `NotifyScheduleDone`
(payload `type: "schedule-done"`) instead of the generic `session-done`. The
service worker (`internal/ui/embedded/assets/sw.js`) shows `schedule-done`
notifications **even when the app is foregrounded**, since a scheduled run is a
background event the user may not be watching.

## Frontend frequency presets

The editor offers presets (hourly, daily, weekdays, weekly) plus a raw custom
cron field and a manual option. Presets compile to a standard 5-field cron
expression client-side (`web/src/index/schedules.js` `buildCron`), and
`parseCron` recovers the preset + fields when editing an existing schedule.
