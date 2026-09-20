# Pagganap at Paggamit

## Mga Tampok

### Remote control

- Ituloy ang anumang session mula sa browser na may text o image attachments
- Simulan ang bagong session para sa anumang project path, direkta mula sa web UI
- Paggustong model sa loob ng browser at selector ng antas ng pag-iisip, bawat session
- Status ng worker bawat session (idle / running / error) na may auto-recovery sa crash
- Maraming sessions ang tumatakbo nang sabay-sabay — simulan ang trabaho sa isa, manood ng ibang stream
- `PI_WEB_TOKEN` para sa ligtas na LAN exposure — kinakailangan nang default para sa anumang eksplisitong non-loopback bind

### Pagbasa ng mga session

- Mag-browse ng mga session sa iba't ibang project na may mga filter, search, at buong branch navigation
- Live incremental updates habang tumatakbo pa rin ang pi (sa pamamagitan ng fsnotify; ~ms latency)
- Follow mode para sa pag-follow ng mga aktibong session
- Deep links sa mga indibidwal na mensahe
- I-download ang isang session bilang JSONL
- Magbahagi ng mga static snapshot bilang secret GitHub Gists
- `/web`, `/remote`, `/refresh`, `/pi-web token`, at `/pi-web set-token` pi extensions para sa pagbubukas ng mga session, remote QR, pag-sync ng session, at pamamahala ng token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) upang maupo ng isang session ang mga schedule, ang project scratchpad, at mga setting sa natural language

## Piliin ang session mode

Ginagamit ng edisyong ito ang mga provider at model na nakakonfigure na sa pi; ang Local Mode ay isang runtime policy, hindi isang hiwalay na model installer o ikaapat na API-key screen. Pumili ng mode sa paggawa ng session, o baguhin ito pagkatapos tumigil ng kasalukuyang run:

| Mode | Gamitin kapag | Pag-uugali |
|------|-------------|----------|
| **Auto** | Kung nais ng pi-web ang pagpapasya | Resolves ang local/LAN endpoints mula sa provider metadata kapag posibleng; kung hindi, pananatili ang normal na path |
| **Local** | Kapag ang model ay nakatakbo sa makina o LAN mo | Paganahin ang 65% compaction boundary, bounded checkpoints, Force Compact, at guarded automatic recovery |
| **Cloud** | Kapag ang napiling model ay naka-host at dapat sundin ang upstream na ugali | Ipinapanatili na ang local-only compaction at recovery policy ay hindi kasama sa session |

Ang manual na Local o Cloud na pagpili ay nagwawagi sa automatic detection at mananatili sa pamamagitan ng mga reload at restart. Ang tumatakbo na session ay tumatanggi sa mga pagbabago ng mode hangga't hindi na tumigil ang kanyang worker, kaya ang mode na ipinapakita sa UI ay laging sumasunod sa aktwal na ginagamit na policy.

## Mga Pangangailangan

- [Go](https://go.dev) 1.25+ (sa pag-build mula sa source lamang)
- `pi` sa iyong `PATH` para sa browser chat/model switching
-opsyonal: `gh` para sa pagbabahagi
- Sa Windows: kailangan ng pi ng isang bash shell para sa shell tool nito — ang [Git for Windows](https://git-scm.com/download/win) sapat na (tingnan ang pi's Windows docs)

## Pag-install

### Pi package (inirerekomenda)

```bash
pi install npm:@timmygod/pi-web-local
```

Ang isang utos na ito:
- I-install ang npm pi package sa ilalim ng package directory ng pi
- Tinatakbo ang `postinstall` script ng package (`install.sh`, o `install.ps1` sa Windows)
- Inidownload ang naaangkop na pi-web binary para sa iyong package version at platform mula sa GitHub Releases
- I-install ito sa `~/.pi/agent/bin/pi-web` (`pi-web.exe` sa Windows)
- Itinatayo ang auto-start sa login (launchd sa macOS, systemd sa Linux, isang Run-key launcher sa Windows)
- Inarehistro ang `/web`, `/remote`, `/refresh`, `/pi-web token`, at `/pi-web set-token` pi commands

Ang session auto-titling ay nakatago sa loob ng pi-web (hindi sa extension) at itinatag sa `/settings` page. Buksan ito nang default: automatic naming ng mga session ng pi-web gamit ang free built-in word heuristic (walang AI), na nagre-retitle sa bawat bagong mensahe. Maaari mong baguhin sa titling na isa lamang beses bawat session, at/o pumili ng model upang isulat ang mas matalinong mga title sa halip na heuristic.

Sa Linux, ang auto-start ay iniset bilang isang user systemd service sa `~/.config/systemd/user/pi-web.service`. Uulitin ng installer ang `ExecStart` nito patungo sa actual na installed binary path. Kung ang Tailscale ay available sa runtime, ipapalabas ng pi-web ang localhost server gamit ang Tailscale Serve HTTPS. Kung hindi available ang user systemd, ilagay itong manual gamit ang `~/.pi/agent/bin/pi-web -o`.

Upang i-install para sa partikular na project lamang (na hinihibay sa iyong team sa pamamagitan ng `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Tapos ay i-restart ang pi (o i-run ang `/reload`), at gamitin ang `/web`, `/pi-web`, `/remote`, `/refresh`. Pamahalaan ang iyong access token gamit ang `/pi-web token` at `/pi-web set-token`.

Kung tumigil ang npm sa `ENOTEMPTY` habang pinapalitan ng pangalan ang `@timmygod/pi-web-local`, alisin ang mga lumang hidden backup directory ng npm at i-reinstall ang package:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Mabilisang pag-install (walang build tools ang kinakailangan)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Inidownload nito ang pinakabagong pi-web binary, ininisital nito ito sa `/usr/local/bin` (`~/.pi/agent/bin` sa Windows), at itinatayo ang auto-start sa login. Walang kinakailangang Go, Node, o pi.

### I-download ang binary

Ang mga pre-built na binary ay nakadikit sa bawat [GitHub Release](https://github.com/timmygod/pi-web/releases).

```bash
# macOS (Apple Silicon)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-arm64
chmod +x pi-web

# macOS (Intel)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-amd64
chmod +x pi-web

# Linux (amd64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-amd64
chmod +x pi-web

# Linux (arm64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-arm64
chmod +x pi-web
```

```powershell
# Windows (x64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-amd64.exe

# Windows (ARM64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-arm64.exe
```

Tapos iilipat ito sa iyong PATH:

```bash
cp pi-web ~/.pi/agent/bin/
# o system-wide:
sudo cp pi-web /usr/local/bin/
```

### Pag-build mula sa source

Ang checkout na ito ay ang local-model edisyon ng pi-web. Ang normal na pag-build ay gumagawa ng web application at backend nang sabay; ang local-model safeguards ay pinagkakaganap sa runtime ng effective Local Mode ng session, hindi ng hiwalay na binary.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # nagbabuild ng Vite bundle, at itong inilalagay sa Go binary

# opsyonal: ilagay sa PATH
cp pi-web ~/.pi/agent/bin/
```

Ang frontend bundle ay naka-embed ng `web/assets_embed.go`, kaya kinakailangan ng `go build` na umiral muna ang `web/dist`. Ang `make build` ay gumagawa ng dalawang hakbang nang isang-bang; kung gagawan mo ng kamay, i-run ang `npm --prefix web install && npm --prefix web run build` bago ang `go build ./cmd/pi-web`.

Para sa maintained fork workflow, upstream synchronization, at Local Mode verification checklist, tingnan ang [local-model development notes](../../docs/dev/local-llm-development.md).

### Pag-develop kasama ang installed instance

Iwanang tumakbo ang installed instance sa port `31415`, tapos ay simulan ang source checkout sa development mode:

```bash
make dev
```

Buksan ang `http://127.0.0.1:31416`. Inilalagay ng `make dev` ang internal `PI_WEB_DEV=1` development environment, kaya ang source checkout ay hinihibay ang mga session, settings, at SQLite data sa installed instance habang pinapanatili ang hiwalay na development runtime lock at state file. Ang regular na installed at manual na launch na instances ay walang pagbabago at nananatili ang orihinal na single-instance behavior.

Upang mabalanse ang duplicate autonomous work, ang development mode ay hindi tumatakbo ng schedule loop, chat-queue drainer, auto-titling, o push notifications. Ang mga direkta na kahilingan na ginawa sa pamamagitan ng development UI ay gumagahanap pa rin. Huwag gamitin ang parehong chat session mula sa magkaibang instances sa sabay; bawat proseso ay may sariling RPC worker manager.

Kailangan ng `make dev` ang [Air](https://github.com/air-verse/air) para sa Go hot reload:

```bash
go install github.com/air-verse/air@latest
```

Ang `PI_WEB_DEV` ay isang development harness plumbing, hindi isang supported production multi-instance mode.

## Pag tanggal

```bash
pi remove npm:@timmygod/pi-web-local
```

Tinitingnan nito ang package `preuninstall` script (`uninstall.sh`, o `uninstall.ps1` sa Windows), na tumitigil ng tumatakbo na instance at nanonormal:

- ang pi-web binary (`~/.pi/agent/bin/pi-web`, o `/usr/local/bin/pi-web` sa standalone installs)
- ang version file (`~/.pi/agent/pi-web-version`)
- ang runtime state file (`~/.pi/agent/pi-web/pi-web-state.json`)
- ang auto-start config (launchd plist sa macOS, systemd user service sa Linux, Run-key entry + launcher scripts sa Windows)

Ang iyong data ay napanatili para sa pagkatapos na pag-install ay tumama kung saan mo nagsimula: `~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, ang iyong session files sa ilalim ng `~/.pi/agent/sessions/`, at `~/.config/pi-web/env` (kasama ang `PI_WEB_TOKEN`). I-alis itong manual kung nais mong ma-manurin ang clean slate.

## Paggamit

```bash
# Simulan sa default port (31415)
pi-web

# Simulan at buksan ang browser
pi-web -o

# Custom port
pi-web -p 8080

# I-override ang bind host (loopback ay unauthenticated nang default)
pi-web --host 127.0.0.1

# Ang non-loopback bind ay kinakailangan ng token — ang pi-web ay tumatanggi sa pagsimula kung wala
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Nang default, binabind ng pi-web ang `127.0.0.1`. Kung tumatakbo ang Tailscale na may MagicDNS **at ang `PI_WEB_TOKEN` ay naka-set**, pinanunusunin din ng pi-web ang `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` at ipinapakita ang HTTPS tailnet URL. Kapag walang token, nananatili ang pi-web na loopback-only at hinahanap ang Tailscale Serve, kaya hindi mababarang ng tailnet peers ang agent na unauthenticated. Ang anumang eksplisitong non-loopback bind ay nangangailangan rin ng `PI_WEB_TOKEN` na naka-set; i-pass ang `--insecure` upang i-override para sa local testing.

## Remote Access

Iwanang nakikinig ang pi-web sa lokal, tapos ay gamitin ang naka-print na Tailscale HTTPS URL mula sa iyong telepono o laptop sa tailnet.

Sa macOS, i-install at buksan ang Tailscale nang interactive, aprubahan ang administrator prompt, at mag-sign in. Tapos ay i-run ang `/pi-web restart`, na sinundan ng `/remote`.

Sa Linux, payagan ang iyong user na pamahalaan ang Tailscale bago mag-install/run ng pi-web, kung hindi maaaring ang `tailscale serve` ay nangangailangan ng sudo at mabigo ang auto-start:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Simulan ang pi-web na may token upang ipalabas ang Tailscale HTTPS endpoint
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Mula sa ibang Tailscale-connected device, buksan ang naka-print na
#    "Tailscale HTTPS" URL at ilagay ang token isang beses.
```

> Nang default, tumatanggi ang pi-web sa pag-bind ng non-loopback address maliban kapag ang `PI_WEB_TOKEN` ay naka-set — ang sinuman na maaabot ang bound na address ay maaaring tingnan ang mga session at magpadala ng mga utos sa pi. Upang i-override ang guard na ito para sa local-network testing, i-pass ang `--insecure`. **Huwag gamitin ang `--insecure` sa Tailscale o anumang address na maaabot mula sa labas ng iyong makina.**
>
> Maaaring ilagay ng mga client ang token sa pamamagitan ng `Authorization: Bearer <token>` header, `X-Pi-Token` header, o isang beses sa pamamagitan ng `?token=<token>` (na nagseset ng `pi_token` cookie para sa mga sumunod na kahilingan). Ang mga token na inilagay sa pamamagitan ng `?token=` ay umaabot sa browser history, server access logs, at `Referer` headers mula sa mga link sa page — inirerekomenda ang header form para sa anumang bagay na higit sa unang bookmark.

## Browser Chat

Buksan ang session page at gamitin ang composer sa ibaba upang ituloy ang eksaktong session na iyon.

- Ang `Enter` ay nagpapadala, ang `Shift+Enter` ay naglalagay ng newline
- I-drag-and-drop o i-paste ang mga image direkta sa composer
- Ang model picker at thinking-level selector ay nasa header — ang mga pagbabago ay agad na nasa ilalim ng pi worker
- Ang bawat aktibong session ay may sariling `pi --mode rpc` worker, kaya ang magkaibang session ay hindi naghihiwalay

## Pagbabahagi ng mga Session

Pindutin ang **Share** sa isang session page upang gumawa ng secret GitHub Gist.

Mga pangangailangan:
- Naka-install ang `gh`
- Nacocomplete ang `gh auth login`

Ang pagbabahagi ay nagbabalik:
- ng secret gist URL
- ng preview URL sa `https://pi.dev/session/#<gistId>`

Ang mga bahagiang gist ay mga snapshot at hindi live-update.

## Auto-Start sa Login

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# I-install ang systemd user service
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Opsyonal: ilagay ang iyong PI_WEB_TOKEN para sa non-loopback binds
# (o gamitin ang /pi-web set-token <token> mula sa loob ng pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# I-enable at simulan
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Tignan ang status
systemctl --user status pi-web.service

# Tignan ang mga log
journalctl --user -u pi-web.service -f
```

> Para sa serbisyo upang simulan sa boot (bago mag-login), gamitin ang isang system service sa halip: i-copy ang `init/pi-web.service` sa `/etc/systemd/system/` at gamitin ang `sudo systemctl`.

### Windows

Inaautomate nito ng installer nito nang automatic, walang kinakailangang admin rights: ang `pi-web` entry sa ilalim ng `HKCU\Software\Microsoft\Windows\CurrentVersion\Run` ay nagsisimula sa `~/.config/pi-web/pi-web-start.vbs` sa login, na nagsisimula sa binary nang nakatago (walang console window) pagkatapos ng pag-load ng `~/.config/pi-web/env` (`PI_WEB_TOKEN`, `PATH`, ...).

Para pamahalaan nito ng kamay:

```powershell
# Simulan / itigil
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Alisin ang auto-start
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Walang service supervision sa Windows: kung sumikip ang pi-web, mananatili itong naka-down hanggang sa susunod na login (auto-restart ng launchd/systemd sa iba pang mga platform).
