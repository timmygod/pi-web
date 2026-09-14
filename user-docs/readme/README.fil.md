<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dw/@timmygod/pi-web-local?label=downloads/wk&color=2ea043&cacheSeconds=86400)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb&cacheSeconds=86400)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · **Filipino** · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

I-drive ang iyong [pi](https://pi.dev) coding agent mula sa iyong telepono, tablet, o laptop — kahit saan sa iyong network, o malayuan sa pamamagitan ng Tailscale.

Isa itong ganap na PWA, kaya maaari mo itong i-install at gamitin tulad ng isang native app sa anumang device. Isipin mo ito bilang iyong sariling personal na AI workspace — tulad ng Cowork ng Claude, ngunit may iba't ibang modelo — mag-chat sa iba't ibang modelo, mag-code mula sa iyong telepono, o gawin itong isang [personal assistant](../en/personal-assistant.md) na naninirahan sa iyong makina.

Gawin itong sa iyo: magpalit ng mga tema at font, at gamitin ito sa iyong sariling wika — ang pi-web ay may kasamang maraming wika at maaari kang magdagdag ng sarili mo. Marami pang mga feature ang paparating, ngunit hindi ito magiging bloated: anumang hindi mo kailangan ay maaaring i-off sa settings.

</div>

## Bakit ang edisyong ito ng local-model?

Ang orihinal na pi-web ay nananatiling pundasyon ng upstream para sa mga katangian at pag-aayos na ibinahagi. Pinapanatili ng edisyong ito ang karanasang iyon, pagkatapos ay nagdadagdag ng layer ng pagiging maaasahan para sa mga modelong tumatakbo sa iyong sariling makina o sa ibang lugar sa iyong LAN—kung saan ang pagbuo ay madalas na mas mabagal, limitado ang memorya, at ang mahabang konteksto ay maaaring huminto sa isang malusog na session.

| Larangan | Upstream pi-web | Ang edisyong ito |
|------|-----------------|--------------|
| Patakaran ng model/runtime | Karaniwang pag-uugali ng pi-web | **Auto / Local / Cloud** na mode bawat session, na may endpoint-aware na pagtukoy ng local at isang persistent na manual override |
| Paghawak ng mahabang konteksto | Normal na pag-uugali ng pi compaction | Ang Local Mode ay aktibong nagko-compact sa **65%** at muling tinitingnan sa loob ng mahahabang loop ng tool-call bago ang isa pang kahilingan sa provider |
| Kaligtasan ng compaction | Karaniwang mga buod | Mga bounded rolling checkpoints, isang mas mahigpit na pagsulat muli para sa hindi wastong/capped na output, at pagtukoy ng walang pag-unlad sa halip na walang katapusang re-compaction |
| Mga pinutol na pagtakbo | Karaniwang paghawak ng worker at error | Bounded recovery para sa context overflow, thinking-only stops, at mga napiling interrupt ng transport, na may persistent loop breakers |
| Manual na rescue | Karaniwang detalye ng konteksto | Ang **Force Compact** ay nananatiling available bilang isang malinaw na landas ng pagbangon nang hindi binubura ang usapan |
| Pagkakatugma at mga release | Orihinal na proyekto at linya ng release | Ang mga safeguard na local-only ay nananatili sa likod ng Local Mode; Pinapanatili ng Cloud Mode ang pag-uugali ng upstream, at ang mga pagbabago ng upstream ay sinusuri at inilalathala dito nang hiwalay |

Hindi ito isang pagsulat muli o kapalit para sa upstream. Ito ay isang sinasadyang pinapanatiling operating profile para sa mga taong nais ng privacy at kontrol ng local-model nang hindi tinatanggap ang mga fragile na mahahabang tumatakbong session. Tingnan ang [gabay sa user](../en/README.md) para sa daloy ng trabaho na nakatuon sa user at [pag-unlad ng edisyong local-model](../../docs/dev/local-llm-development.md) para sa implementasyon at patakaran sa synchronization.

> [!WARNING]
> Ang pi-web ay kasalukuyang nasa **beta**. Magbabago at masisira ang mga bagay!

> [!TIP]
> Bago ka ba dito? **[Basahin ang user guide →](../en/README.md)** para sa isang buong tour ng mga feature, mga hakbang sa pag-install, at mga tip. ([Ibang mga wika →](../README.md))

## Mga Screenshot

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile PWA" width="90%" /><br />
  <em>Mobile PWA</em>
</div>

## Paano Ito Nagkakasya

```
 pi (terminal)                 Browser (telepono / tablet / laptop)
      │                                │
      │  sumusulat ng JSONL           │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           sa pamamagitan
                                                     ng MagicDNS)
```

- **pi** ay sumusulat ng conversation JSONL sa `~/.pi/agent/sessions/` habang ito ay gumagana.
- **pi-web** ay isang Go server na nagbabasa ng mga file na iyon, nire-render ang mga ito sa browser, at nag-stream ng mga live update sa pamamagitan ng SSE.
- Ang mga **pi --mode rpc** worker ay humahawak ng chat na pinasimulan ng browser — isa bawat session, tinatanggal pagkatapos ng 10 minutong idle.
- **fsnotify** ay binabantayan ang sessions directory upang ang browser ay mag-reload sa loob ng millisecond ng bagong output.
- **Tailscale Serve** ay nagpa-publish ng localhost server bilang isang HTTPS endpoint sa iyong tailnet.

## Pag-install

```bash
pi install npm:@timmygod/pi-web-local@beta
```

Iyon na — dina-download nito ang katugmang binary, nagse-set up ng auto‑start, at nirerehistro ang `/web`, `/pi-web`, `/remote`, at `/refresh` na mga command.

Kapag na-install na, buksan ang `http://127.0.0.1:31415` sa iyong browser. Mula sa pi, gamitin ang `/web` upang buksan agad ang kasalukuyang session sa iyong browser. Kung tumatakbo ang Tailscale sa iyong makina, awtomatikong nagpa-publish ang pi-web ng isang HTTPS endpoint sa iyong tailnet — gamitin ang `/remote` mula sa pi upang makakuha ng QR code at URL para sa anumang device sa iyong tailnet.

> **Malayuang access sa macOS:** I-install at buksan ang Tailscale nang interactive, aprubahan ang administrator prompt, at mag-sign in. Pagkatapos ay patakbuhin ang `/pi-web restart`, na susundan ng `/remote`.

Para sa mga manu-manong pag-install, pag-download ng binary, o pagbuo mula sa source, tingnan ang [user-docs/install.md](../en/install.md).

## Integrasyon ng Pi

Pagkatapos ng `pi install npm:@timmygod/pi-web-local@beta`, makukuha mo ang:

| Command | Kung ano ang ginagawa nito |
|---------|----------------------------|
| `/web` | Binubuksan ang kasalukuyang session sa iyong browser (may kamalayan sa SSH: nilalaktawan ang browser at ipinapakita lamang ang URL) |
| `/pi-web` | Ipinapakita ang katayuan, bersyon, pagsisimula/paghinto/pag-restart ng server, o pag-update |
| `/remote` | Ipinapakita ang isang QR code at URL para sa malayuang pag-access sa pamamagitan ng Tailscale |
| `/refresh` | Kinukuha ang mga bagong mensaheng isinulat mula sa mga remote browser pabalik sa terminal session |

Ang **auto-titling** ng session ay nakapaloob mismo sa pi-web at naka-configure sa pahina ng `/settings`. Ito ay **naka-on bilang default** at awtomatikong pinapangalanan ang mga session. Maaari kang pumili:

- **Kailan mag-title** — isang beses bawat session, o sa bawat bagong mensahe (ang default).
- **Modelo ng titulo** — isang libre, agarang **built-in word heuristic (walang AI)** bilang default, o pumili ng modelo (hal. isang maliit/mabilis) para sa mas matalinong, mga titulong isinulat ng modelo.

Ang package ay nag-i-install din ng pi-web binary sa `~/.pi/agent/bin/pi-web` at nagse-set up ng auto-start sa pag-login.

## Auto-Start sa Pag-login

Ang command na `pi install npm:@timmygod/pi-web-local@beta` ay awtomatikong nagse-set up nito:

| OS | Mekanismo |
|----|-----------|
| macOS | launchd plist sa `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service sa `~/.config/systemd/user/pi-web.service` |

Upang magtakda ng token para sa malayuang pag-access, lumikha ng `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=ang-iyong-token-dito
```

Para sa higit pang mga detalye (manu-manong pag-setup, mga custom port, mga non-loopback bind), tingnan ang [user-docs/install.md](../en/install.md).

## Development

```bash
make setup   # i-install ang frontend deps at i-download ang Go modules
make check   # frontend test/build + Go test/vet
make build   # setup kung kinakailangan, i-build ang frontend, pagkatapos ay i-build ang ./pi-web
```
