<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · **မြန်မာ** · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

[F phone/tablet/laptop](https://pi.dev) (F) — F F network, Tailscale F.

PWA — F F. Claude Cowork — F. [F](../en/personal-assistant.md) F.

F: theme/font F, [my language](../en/personal-assistant.md) — pi-web F F. F — F F F.

</div>

## F F F?

pi-web F F F F. F F F, F F F LAN F F F, memory F, F F F.

| Area | Upstream pi-web | F |
|------|-----------------|--------------|
| Model/runtime policy | F pi-web F | **Auto / Local / Cloud** F, endpoint F F F, persistent manual override |
| Long-context handling | F pi compaction F | Local Mode **65%** F, long tool-call F F F |
| Compaction safety | F summaries | Bounded rolling checkpoints, F F F, F F |
| Interrupted runs | F worker error F | Context overflow, thinking-only F, transport F, persistent loop breakers |
| Manual rescue | F context F | **Force Compact** F F F, F F |
| Compatibility and releases | F F F | Local Mode F; Cloud Mode F F, F F F F |

F F F. F F F F. [F](../en/README.md) F [F](../../docs/dev/local-llm-development.md) F.

> [!TIP]
> F? **[F →](../en/README.md)** F. ([F →](../README.md))

## F

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## F F F

```
 pi (terminal)                 Browser (phone / tablet / laptop)
      │                                │
      │  writes JSONL                  │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           via MagicDNS)
```

- **pi** JSONL F `~/.pi/agent/sessions/`.
- **pi-web** Go F, browser F, SSE F.
- **pi --mode rpc** F F F — F, 10 F F.
- **fsnotify** F F, browser F F.
- **Tailscale Serve** localhost F HTTPS F tailnet.

## F

```bash
pi install npm:@timmygod/pi-web-local
```

F — binary F, F, `/web`, `/pi-web`, `/remote`, `/refresh` F.

`http://127.0.0.1:31415` F. pi F `/web` F. Tailscale F, pi-web HTTPS F tailnet — `/remote` F, QR F URL F.

> **macOS F:** Tailscale F, admin F, F. `/pi-web restart`, `/remote` F.

F, binary F, [F](../en/install.md) F.

## F F

`pi install npm:@timmygod/pi-web-local` F:

| Command | F |
|---------|--------------|
| `/web` | browser F (SSH: browser F URL F) |
| `/pi-web` | F, version, start/stop/restart, F |
| `/remote` | Tailscale F QR F URL F |
| `/refresh` | browser F F terminal F |

**F** F pi-web F, `/settings` F. **F** F. F:

- **F** — F F, F F.
- **F** — F F **F (no AI)**, F (F F) F.

pi-web F `~/.pi/agent/bin/pi-web` F, F F.

## F F

`pi install npm:@timmygod/pi-web-local` F:

| OS | F |
|----|-----------|
| macOS | launchd plist `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd F `~/.config/systemd/user/pi-web.service` |
| Windows | `HKCU` Run-key F `~/.config/pi-web/` |

F F, `~/.config/pi-web/env` F:

```
PI_WEB_TOKEN=your-token-here
```

F (F, F, F), [F](../en/install.md) F.

## F

```bash
make setup   # F F Go F
make check   # F test/build + Go test/vet
make build   # setup F, F, ./pi-web F
```

F F, F, F, [F](../../docs/dev/local-llm-development.md) F.
