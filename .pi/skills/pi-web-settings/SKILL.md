---
name: pi-web-settings
description: Read or change pi-web settings (theme, language, fonts, auto-title, notifications, artifact visibility, session display defaults, cat/pomodoro). Use when the user wants to change how pi-web looks or behaves. Slash command /skill:pi-web-settings.
---

# pi-web settings

Change server-backed pi-web settings with `pi-web-ctl`. Do not write SQLite yourself.

```bash
pi-web-ctl settings get
pi-web-ctl settings get theme
pi-web-ctl settings set theme dark
pi-web-ctl settings set language ja
pi-web-ctl settings set auto-title off
```

If `pi-web-ctl` is not on `PATH`, run `python3 ~/.pi/agent/bin/pi-web-ctl`.

Aliases (CLI also accepts the raw storage key): theme, language, font-ui, font-content, font-code, auto-title, auto-title-mode, auto-title-model, notify-on-done, artifacts, thinking, tools, tool-outputs, cat, bedtime, wakeup, layout, spinner.

Booleans accept on/off. Theme and fonts apply live; language reloads open pi-web tabs (same as the Settings picker) so chrome re-renders. Report the new value. Do not rewrite custom-languages JSON unless the user is adding a language.
