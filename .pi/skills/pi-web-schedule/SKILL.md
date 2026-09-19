---
name: pi-web-schedule
description: Create, list, update, disable, delete, or run-now pi-web schedules. Use when the user wants a recurring or timed job, cron, reminder, or "every day at 2am" automation in pi-web. Slash command /skill:pi-web-schedule.
---

# pi-web schedules

Manage pi-web schedules through `pi-web-ctl`. Do not write SQLite yourself and do not call curl with `?token=` — the CLI talks to the local HTTP API.

```bash
pi-web-ctl schedule list
pi-web-ctl schedule create --name NAME --instructions TEXT --daily --hour 2 --timezone sg
pi-web-ctl schedule get NAME_OR_ID
pi-web-ctl schedule update NAME_OR_ID --paused
pi-web-ctl schedule enable|disable|delete|run NAME_OR_ID
pi-web-ctl schedule runs NAME_OR_ID
```

If `pi-web-ctl` is not on `PATH`, run `python3 ~/.pi/agent/bin/pi-web-ctl` (Unix) or `py -3 %USERPROFILE%\.pi\agent\bin\pi-web-ctl.py` (Windows).

Cadence flags: `--daily`, `--hourly`, `--weekdays`, `--weekly --weekday mon`, `--every-hours N`, `--manual`, or `--cron "0 2 * * *"`. Prefer flags over raw cron. Pass `--timezone` as an IANA name or alias (`sg`, `jst`, `utc`, `pt`, `et`, `ct`, `london`). The CLI resolves aliases.

`--project` defaults to the current working directory. `--model provider/id` and `--thinking` are optional (pi defaults). `--name` defaults to the first line of `--instructions`.

## Rules

- A fire starts a **new empty session**. `--instructions` must be a standalone prompt. Never "continue what we were doing."
- "2am sg daily" (or similar, with a clear time + timezone + cadence) → create immediately.
- Ambiguous time ("tomorrow morning", missing timezone when the user implied a place) → ask before creating.
- After create, echo name, cadence, timezone, next run, project, and a one-line instruction summary.
- Match existing schedules by name; if several share a name, use the id from `list`.
- Recurring jobs only fire while pi-web is running.
