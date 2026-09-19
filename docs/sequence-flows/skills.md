# Sequence Flow: pi-web skills

pi-web ships pi skills under `.pi/skills/`. They do not write `pi-web.sqlite`
themselves. A shared CLI (`pi-web-ctl`) discovers the running server and calls
the existing HTTP APIs.

`pi install` copies `.pi/skills/common/pi_web.py` to `~/.pi/agent/bin/pi-web-ctl`
so the agent can invoke it from any project directory.

## Discover → auth → API → SSE

```
Agent (pi session)
     │
     │  pi-web-ctl schedule create …
     ▼
Read ~/.pi/agent/pi-web/pi-web-state.json   (regular file, not *-dev.json)
     │
     ├── port → http://127.0.0.1:{port}     (always loopback)
     └── PI_WEB_TOKEN / ~/.config/pi-web/env
              │
              ▼
     X-Pi-Token header   (never ?token= — that 302s past the handler)
              │
              ▼
     POST /api/schedules
              │
              ▼
     SSE event: schedules on __all__
              │
              ▼
     /schedules page silent-refetches /api/schedules
```

v1 skills:

| Skill | CLI | API |
|---|---|---|
| `pi-web-schedule` | `pi-web-ctl schedule …` | `/api/schedules`, `/api/schedule` |
| `pi-web-notes` | `pi-web-ctl notes …` | `/api/scratchpad` |
| `pi-web-settings` | `pi-web-ctl settings …` | `/api/settings` |

If the server is down the CLI exits with “pi-web is not running. Start it with
`/pi-web start`.” Schedules cannot fire unless the server is up anyway.
