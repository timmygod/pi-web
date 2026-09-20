<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · **Filipino** · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Gamitin ang iyong [pi](https://pi.dev) coding agent mula sa iyong phone, tablet, o laptop — kahit saan sa iyong network, o nang remote sa pamamagitan ng Tailscale.

Ito ay isang buong PWA, kaya maaari mo itong i-install at gamitin na parang native app sa anumang device. Isipin mo itong ang iyong sariling personal na AI workspace — gaya ng Claude's Cowork, pero may iba't ibang models — mag-chat sa iba't ibang models, mag-code mula sa phone mo, o ginawan itong [personal assistant](../en/personal-assistant.md) na nakatira sa iyong machine.

Ginawan mo itong sarili mo: palitan ang themes at fonts, at gamitin sa iyong sariling wika — may dalang maraming wika ang pi-web at maaari mong magdagdag ng sarili mong wika. May mga feature pa na dumarating, pero hindi ito magiging bloated: anumang hindi mo kailangan ay maaaring patayin sa settings.

</div>

## Bakit itong local-model edition?

Ang orihinal na pi-web ang mananatiling upstream foundation para sa mga shared features at
fixes. Pinapanatili ng edition na ito ang karanasang iyon, tapos idinagdagan ng isang reliability layer para sa
mga model na tumatakbo sa iyong sariling machine o sa ibang lugar sa iyong LAN—kung saan ang generation ay
karaniwang mas mabagal, limitado ang memory, at maaaring huminto ang isang mahabang context na isang nangubos na
session.

| Area | Upstream pi-web | Ang edition na ito |
|------|-----------------|--------------|
| Model/runtime policy | Standard pi-web behavior | Per-session na **Auto / Local / Cloud** mode, may endpoint-aware na local detection at isang persistent manual override |
| Long-context handling | Karaniwang pi compaction behavior | Ang Local Mode ay proactively nagco-compact sa **65%** at inaalala muli sa loob ng mahabang tool-call loops bago ang susunod na model request |
| Compaction safety | Standard summaries | May limit na rolling checkpoints, isa pang mas malapit na rewrite para sa invalid/capped output, at no-progress detection imbes na walang hanggang re-compaction |
| Interrupted runs | Karaniwang worker at error handling | May limit na recovery para sa context overflow, thinking-only stops, at piniling transport interruptions, may persistent loop breakers |
| Manual rescue | Standard context details | Ang **Force Compact** ay nananatiling available bilang explicit na recovery path na hindi pinalilimot ang conversation |
| Compatibility at releases | Orihinal na proyekto at release line | Ang local-only na safeguards ay nasa likod ng Local Mode; pinapanatili ng Cloud Mode ang upstream behavior, at ang mga upstream changes ay sinusuri at inilalathala dito nang mag-isa |

Ito ay hindi isang rewrite o pambagi ng upstream. Ito ay isang sinadyang
pinapangalagaan na operating profile para sa mga gustong privacy at kontrol sa local model
na hindi tanggap ng mahina at mahahabang tumatakbo na sessions. Tingnan ang
[user guide](../en/README.md) para sa user-facing na workflow at
[local-model edition development](../../docs/dev/local-llm-development.md) para sa
implementation at synchronization policy.

> [!TIP]
> Bago ka pa rito? **[Basahin ang user guide →](../en/README.md)** para sa buong takdang-arak ng mga feature, mga hakbang ng install, at tips. ([Iba pang mga wika →](../README.md))

## Screenshots

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## Paano Ito Nakaigi

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

- Ang **pi** ay sumusulat ng conversation JSONL sa `~/.pi/agent/sessions/` habang nagtatrabaho.
- Ang **pi-web** ay isang Go server na bumabasa ng mga files na iyon, inir-render sa browser, at nag-stream ng live updates sa pamamagitan ng SSE.
- Ang mga **pi --mode rpc** workers ang nag-aalaga ng browser-initiated na chat — isa bawat session, tinatawag pagkatapos ng 10 min ng idle.
- Ang **fsnotify** ay nagbabantay sa sessions directory upang ma-reload ng browser sa loob ng ilang milliseconds ng bagong output.
- Ang **Tailscale Serve** ay inilalathala ang localhost server bilang isang HTTPS endpoint sa iyong tailnet.

## Install

```bash
pi install npm:@timmygod/pi-web-local
```

Yan na lang — binabawasan nito ang angkop na binary, inaayos ang auto-start, at tinatala ang `/web`, `/pi-web`, `/remote`, at `/refresh` na mga command.

Pagkatapos itong i-install, buksan ang `http://127.0.0.1:31415` sa iyong browser. Mula sa pi, gamitin ang `/web` upang buksan ang kasalukuyang session sa iyong browser agad. Kung tumatakbo ang Tailscale sa iyong machine, awtomatikong inilalathala ng pi-web ang isang HTTPS endpoint sa iyong tailnet — gamitin ang `/remote` mula sa pi upang makakuha ng QR code at URL para sa anumang device sa iyong tailnet.

> **Pag-access ng macOS mula sa malayo:** I-install at buksan ang Tailscale nang interactively, aprubahan ang administrator prompt, at mag-sign in. Tapos i-run ang `/pi-web restart`, na sinundan ng `/remote`.

Para sa manual na install, binary downloads, o pagbuo mula sa source, tingnan ang [user-docs/install.md](../en/install.md).

## Pi Integration

Pagkatapos ng `pi install npm:@timtygod/pi-web-local`, nakakakuha ka ng:

| Command | Ano ang ginagawa nito |
|---------|--------------|
| `/web` | Buksan ang kasalukuyang session sa iyong browser (SSH-aware: laktawan ang browser at ipakita ang URL lamang) |
| `/pi-web` | Ipakita ang status, version, i-start/stop/restart ang server, o i-update |
| `/remote` | Ipakita ang QR code at URL para sa remote access sa pamamagitan ng Tailscale |
| `/refresh` | Kunin ang mga bagong mensahe na isinulat mula sa mga remote browser at ibalik sa terminal session |

Ang **auto-titling** ng session ay nakabuti sa mismong pi-web at naka-configure sa `/settings` page. Ito ay **on by default** at awtomatikong nakapangalan ng mga session. Maaari mong pumili:

- **Kailan ititle** — isang beses bawat session, o sa bawat bagong mensahe (ang default).
- **Title model** — isang libre, agad na **built-in word heuristic (walang AI)** by default, o pumili ng model (hal. maliit/mabilis) para sa mas matalino, model-written na mga title.

Ang package ay nangungutang daragdag din ang pi-web binary sa `~/.pi/agent/bin/pi-web` at inaayos ang auto-start sa pag-login.

## Auto-Start sa Pag-login

Ang `pi install npm:@timmygod/pi-web-local` command ay awtomatikong inaayos ito:

| OS | Mechanism |
|----|-----------|
| macOS | launchd plist sa `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service sa `~/.config/systemd/user/pi-web.service` |
| Windows | `HKCU` Run-key entry na nagsisimula ng isang hidden starter sa `~/.config/pi-web/` |

Upang mag-set ng token para sa remote access, gumawa ng `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Para sa mas maraming detalye (manual setup, custom ports, non-loopback binds), tingnan ang [user-docs/install.md](../en/install.md).

## Development

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Para sa upstream synchronization, local-model testing, at ang parallel na release
workflow, tingnan ang [Local-model edition development](../../docs/dev/local-llm-development.md).
