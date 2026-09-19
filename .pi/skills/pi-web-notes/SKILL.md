---
name: pi-web-notes
description: Read or write the pi-web per-project scratchpad (right-sidebar notes). Use when the user says notes, scratchpad, jot this down, write this here. Not for long-term memory ("remember that") and not for transcript annotations. Slash command /skill:pi-web-notes.
---

# pi-web notes (scratchpad)

The right sidebar is **one markdown blob per project**. Read, append, or replace it with `pi-web-ctl`. Do not write SQLite yourself.

```bash
pi-web-ctl notes read [--project PATH]
pi-web-ctl notes append --text TEXT [--project PATH]
pi-web-ctl notes replace --text TEXT [--project PATH]
```

`--project` defaults to the current working directory.

If `pi-web-ctl` is not on `PATH`, run `python3 ~/.pi/agent/bin/pi-web-ctl`.

## Rules

- Default verb is **append**. Replace only when the user clearly wants a rewrite ("replace the notes", "clear the scratchpad").
- Do not confuse with `/skill:memory` ("remember that") or session annotations (highlights on a transcript).
- After write, say that it landed in the project scratchpad (right sidebar).
