# တပ်ဆင်ခြင်းနှင့် အသုံးပြုနည်း

## လုပ်ဆောင်ချက်များ

### ဓာတ်ပိုးထိန်းချုပ်ခြင်း (Remote control)

- browser ကနေ text သို့မဟုတ် image attachment ချိတ်ဆက်ထားပြီး မည်သည့် session မဆို ဆက်လက်လုပ်ဆောင်နိုင်သည်
- web UI တွင်နေ၍ project path မှာနေရာအနှံ့သို့ brand-new session စတင်ရန်
- session တစ်ခုချင်းစီအတွက် in-browser model switching နှင့် thinking-level selector
- session တစ်ခုချင်းစီအတွက် worker status (idle / running / error) နှင့် crash ဖြစ်လျှင် auto-recovery
- parallel ဖြင့် sessions များကို တစ်ပြိုင်နက်လိုက်လည်လည် အလုပ်လုပ်နိုင်သည် — တစ်ခုတွင် အလုပ်စတင်ပြီး တစ်ခုခြင်း၏ stream ကို ကြည့်ရှုနိုင်သည်
- `PI_WEB_TOKEN` သည် secure LAN exposure အတွက် — explicit non-loopback bind များအားလုံးအတွက် default ဖြင့်လိုအပ်သည်

### Sessions ကို အဖတ်ရှုခြင်း

- filters, search နှင့် full branch navigation ဖြင့် project များအနှံ့ sessions ကို ကြည့်ရှုနိုင်သည်
- pi က ဆက်လက်လည်ပတ်နေချိန်တွင် live incremental updates (fsnotify ကြona; ~ms latency)
- active session တွင်း tailing အတွက် follow mode
- individual message တစ်ခုချင်းစီအတွက် deep links
- session တစ်ခုကို JSONL အဖြစ် download
- static snapshots ကို secret GitHub Gists အနေဖြင့် share
- sessions မှာ ပြန်ဖွင့်ရန်, remote QR, session sync နှင့် token management အတွက် `/web`, `/remote`, `/refresh`, `/pi-web token` နှင့် `/pi-web set-token` pi extensions များ
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) သည် session က schedules, project scratchpad နှင့် settings များကို natural language ဖြင့် စီမံပေးနိုင်ရန်

## Session Mode ကို ရွေးချယ်ခြင်း

ဤ edition သည် pi တွင် ပြင်ဆင်ပြီးသား providers နှင့် models များကို အသုံးပြုသည်; Local Mode သည် runtime policy ဖြစ်ပြီး၊ တစ်ခြား model installer သို့မဟုတ် စတုတ္ထပိုင်း API-key screen မဟုတ်ပါ။ session သို့ ဖန်တီးစဉ် mode ကို ရွေးချယ်နိုင်သည် သို့မဟုတ် current run က stable ဖြစ်သွားပြီးနောက် ပြောင်းလဲနိုင်သည်:

| Mode | သုံးရန်အချိန် | Behavior |
|------|-------------|----------|
| **Auto** | pi-web က အာမခံပေးရန် | provider metadata မှ local/LAN endpoints များကို မဖြစ်နိုင်လျှင် normal path ကို ဆက်လက်ရန် |
| **Local** | model သည် ဤ machine တွင် သို့မဟုတ် LAN တွင် လည်ပတ်နေပါက | 65% compaction boundary, bounded checkpoints, Force Compact နှင့် guarded automatic recovery ကို ဖွင့်ရန် |
| **Cloud** | ရွေးချယ်ထားသော model သည် hosted ဖြစ်ပြီး upstream behavior ကို လိုက်နာရမည် | local-only compaction နှင့် recovery policy ကို session မှ ထုတ်ပစ်ရန် |

Manual Local သို့မဟုတ် Cloud selection သည် automatic detection ထက် အဆုံးအဖြတ်တွင် ရှိပြီး reload နှင့် restart များအတွင်း ပြောင်းလဲရန်။ running session သည် mode change များကို worker က stable ဖြစ်သွားသည်အထိ reject လုပ်သည်။ ထို့ကြောင့် UI တွင် ပြသထားသော mode သည် actually in use ဖြစ်သော policy နှင့် အမြဲကိုက်ညီသည်။

## Requirements

- [Go](https://go.dev) 1.25+ (source မှ တည်ဆောက်ရန်သာ)
- `pi` သည် `PATH` တွင် browser chat/model switching အတွက်
- Optional: `gh` သည် share အတွက်
- Windows တွင်: pi သည် shell tool အတွက် bash shell လိုအပ်သည် — [Git for Windows](https://git-scm.com/download/win) သည် လုံလောက်ပါသည် (pi ၏ Windows docs ကို ကြည့်ပါ)

## Install

### Pi package (အကြံပြုထားသည်)

```bash
pi install npm:@timmygod/pi-web-local
```

ဒီ command တစ်ခုတည်းသည်:
- pi ၏ package directory အောက်တွင် npm pi package ကို တပ်ဆင်သည်
- package `postinstall` script (`install.sh`, Windows တွင် `install.ps1`) ကို လည်ပတ်သည်
- package version နှင့် platform တွင် ကိုက်ညီသော pi-web binary ကို GitHub Releases မှ download
- `~/.pi/agent/bin/pi-web` (Windows တွင် `pi-web.exe`) သို့ တပ်ဆင်သည်
- login တွင် auto-start ကို စီစဉ်သည် (macOS တွင် launchd, Linux တွင် systemd, Windows တွင် Run-key launcher)
- `/web`, `/remote`, `/refresh`, `/pi-web token`, နှင့် `/pi-web set-token` pi commands များကို register

Session auto-titling သည် pi-web တွင် (extension မဟုတ်) ပါဝင်ပြီး `/settings` page တွင် configure လုပ်ထားသည်။ Default ဖြင့် on ဖြစ်သည်: pi-web သည် free built-in word heuristic (AI မဟုတ်) ကို အသုံးပြု၍ sessions များကို အလိုအလျောက် နာမည်ပေးပြီး၊ message အသစ်တစ်ခုစီအတွက် ပြန်လည်တပ်ဆင်သည်။ session တစ်ခုအတွက် titling တစ်ကြိမ်သာ ပြောင်းနိုင်သည်၊ နှင့်/သို့မဟုတ် heuristic ထက် ပိုမို smart ဖြစ်သော titles များ ရေးသားရန် model တစ်ခုကို ရွေးချယ်နိုင်သည်။

Linux တွင် auto-start သည် `~/.config/systemd/user/pi-web.service` တွင် user systemd service အဖြစ် configure လုပ်ထားသည်။ Installer သည် `ExecStart` ကို actually installed binary path ပြောင်းလဲသည်။ runtime တွင် Tailscale ရရှိနိုင်ပါက pi-web သည် localhost server ကို Tailscale Serve HTTPS ဖြင့် publish လုပ်သည်။ user systemd မရရှိပါက `~/.pi/agent/bin/pi-web -o` ဖြင့် manual ဖြင့် လည်ပတ်ပါ။

project တစ်ခုတည်းအတွက်သာ တပ်ဆင်ရန် (`.pi/settings.json` ကြောင့် team နှင့် share လုပ်ထားသည်):

```bash
pi install -l npm:@timmygod/pi-web-local
```

ထို့နောက် pi ကို restart လုပ်ပါ (သို့မဟုတ် `/reload` လည်ပတ်ပါ)၊ `/web`, `/pi-web`, `/remote`, `/refresh` များကို အသုံးပြုပါ။ `/pi-web token` နှင့် `/pi-web set-token` ဖြင့် access token ကို စီမံပါ။

npm သည် `ENOTEMPTY` ဖြင့် abort လျှင် `@timmygod/pi-web-local` ကို rename လုပ်ရင်းတွင် npm ၏ stale hidden backup directories များကိုဖယ်ရှားပြီး package ကို ပြန်လည်တပ်ဆင်ပါ:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Quick install (build tools မလိုအပ်)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

ဒါသည် latest pi-web binary ကို download လုပ်ပြီး `/usr/local/bin` (Windows တွင် `~/.pi/agent/bin`) သို့ တပ်ဆင်ပြီး login တွင် auto-start ကို စီစဉ်သည်။ Go, Node, သို့မဟုတ် pi မလိုအပ်ပါ။

### Download binary

Pre-built binaries များသည် foreach [GitHub Release](https://github.com/timmygod/pi-web/releases) တွင် attached ဖြစ်သည်။

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

ထို့နောက် PATH တွင် သွင်းပါ:

```bash
cp pi-web ~/.pi/agent/bin/
# သို့မဟုတ် system-wide:
sudo cp pi-web /usr/local/bin/
```

### Source မှ Build

ဒီ checkout သည် pi-web ၏ local-model edition ဖြစ်သည်။ Normal build သည် web application နှင့် backend ကို ခေတ်ထုတ်အောင် တည်ဆောက်သည်; local-model safeguards များသည် session ၏ effective Local Mode ဖြင့် runtime တွင် ဖွင့်ထားသည်။ တစ်ခြား binary ဖြင့် မဟုတ်ပါ။

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # Vite bundle ကို build လုပ်ပြီး Go binary သို့ embed လုပ်သည်

# optional: PATH တွင် ထားရန်
cp pi-web ~/.pi/agent/bin/
```

Frontend bundle သည် `web/assets_embed.go` ဖြင့် embed လုပ်ထားပြီး၊ `go build` သည် `web/dist` ကို အရင်ရှိရမည်။ `make build` သည် ခေတ်လက်ရှိ two steps ကို ဆက်တိုက်လည်ပတ်သည်; hand ဖြင့် build လုပ်ပါက `go build ./cmd/pi-web` ကို အရင် `npm --prefix web install && npm --prefix web run build` ကို လည်ပတ်ပါ။

maintained fork workflow, upstream synchronization နှင့် Local Mode verification checklist အတွက် [local-model development notes](../../docs/dev/local-llm-development.md) ကို ကြည့်ပါ။

### Installed instance နှင့်အတူ Develop

Installed instance ကို port `31415` တွင် လည်ပတ်နေရန် ခံထားရပြီး၊ source checkout ကို development mode တွင် စတင်ပါ:

```bash
make dev
```

`http://127.0.0.1:31416` ကို ဖွင့်ပါ။ `make dev` သည် internal `PI_WEB_DEV=1` development environment ကို set လုပ်ထားပြီး၊ source checkout သည် installed instance နှင့် sessions, settings, နှင့် SQLite data များကို share လုပ်သည်။ သို့သော် separate development runtime lock နှင့် state file ကို ကိုင်တွယ်ထားသည်။ Regular installed နှင့် manually launched instances များသည် မပြောင်းလဲဘဲ original single-instance behavior ကို ဆက်လက်ရိပ်ချက်ထားသည်။

duplicate autonomous work များကို တားဆီးရန် development mode သည် schedule loop, chat-queue drainer, auto-titling, သို့မဟုတ် push notifications များကို လည်ပတ်မည်မဟုတ်ပါ။ development UI မှ တင်သွင်းသော direct requests များသည် ဆက်လက်လည်ပတ်နိုင်သည်။ ခေတ်တို ခြောက်ဖွဲ့ chat session တစ်ခုကို instances နှစ်ခုမှ တစ်ပြိုင်နက် စီမံမေးသေးပါ။ foreach process သည် ကိုယ်ပိုင် RPC worker manager ကို ဖြစ်သည်။

`make dev` သည် Go hot reload အတွက် [Air](https://github.com/air-verse/air) လိုအပ်သည်:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` သည် development harness plumbing ဖြစ်ပြီး၊ production multi-instance mode မဟုတ်ပါ။

## Uninstall

```bash
pi remove npm:@timmygod/pi-web-local
```

ဒါသည် package `preuninstall` script (`uninstall.sh`, Windows တွင် `uninstall.ps1`) ကို လည်ပတ်ပြီး running instance ကို stop လုပ်ထားသည်။ အောက်ပါတို့ကို ဖယ်ရှားသည်:

- pi-web binary (`~/.pi/agent/bin/pi-web`, သို့မဟုတ် standalone installs အတွက် `/usr/local/bin/pi-web`)
- version file (`~/.pi/agent/pi-web-version`)
- runtime state file (`~/.pi/agent/pi-web/pi-web-state.json`)
- auto-start config (macOS တွင် launchd plist, Linux တွင် systemd user service, Windows တွင် Run-key entry + launcher scripts)

သင့် data များသည် ခေတ်တို ပြန်လည်တပ်ဆင်မှုအတွက် ခေတ်တို ပိတ်ဆို့ထားသောနေရာမှ ဆက်လက်အသုံးပြုနိုင်ရန် preserve ဖြစ်သည်: `~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, `~/.pi/agent/sessions/` အောက်တွင် session files များနှင့် `~/.config/pi-web/env` (including `PI_WEB_TOKEN`)။ clean slate ဖြစ်စေရန် manually ဖယ်ရှားပါ။

## Usage

```bash
# default port (31415) တွင် စတင်ရန်
pi-web

# browser ဖွင့်ပြီး စတင်ရန်
pi-web -o

# Custom port
pi-web -p 8080

# bind host ကို override လုပ်ပါ (loopback သည် default ဖြင့် unauthenticated ဖြစ်သည်)
pi-web --host 127.0.0.1

# non-loopback bind သည် token လိုအပ်သည် — pi-web သည် အခြားအခြင်းတွင် start ဖြင့် reject လုပ်သည်
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Default ဖြင့် pi-web သည် `127.0.0.1` သို့ bind လုပ်ထားသည်။ Tailscale သည် MagicDNS ကြona **and `PI_WEB_TOKEN` သည် set ဖြစ်နေပါက** pi-web သည် `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` ကို ဆက်လက်လည်ပတ်ပြီး HTTPS tailnet URL ကို ပုံဖော်သည်။ token မရှိပါက pi-web သည် loopback-only ဖြစ်နေပြီး Tailscale Serve ကို ခေတ်ဖုံးပြီး၊ tailnet peers များက unauthenticated ဖြင့် agent ကို ခေတ်ထိစပ်မရအောင်။ explicit non-loopback bind များအားလုံးသည် `PI_WEB_TOKEN` သည် set ဖြစ်နေရန် လိုအပ်သည်။ local testing အတွက် override လုပ်ရန် `--insecure` ကို pass လုပ်ပါ။

## Remote Access

pi-web ကို local တွင် ကြည့်နေရင်း ခေတ်တို tailnet တွင် phone သို့မဟုတ် laptop ကနေ printed Tailscale HTTPS URL ကို အသုံးပြုပါ။

macOS တွင် Tailscale ကို အသုံးတပ်ပြီး interactively ဖွင့်ပါ၊ administrator prompt ကို ခေတ်ကိုင်သည့်သော သောဟိုသော sign in လုပ်ပါ။ ထို့နောက် `/pi-web restart` ကို ခေတ်လည်ပတ်ပြီး `/remote` ကို ခေတ်လည်ပတ်ပါ။

Linux တွင် pi-web ကို install/lówပတ်ရန်အရင် သင့် user က Tailscale ကို စီမံနိုင်ရန် လုပ်ပါ။ အခြားအခြင်းတွင် `tailscale serve` သည် sudo လိုအပ်ပြီး auto-start သည် fail ဖြစ်နိုင်သည်:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Tailscale HTTPS endpoint ကို publish လုပ်ရန် token ဖြင့် pi-web ကို စတင်ပါ
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Tailscale connected device တစ်ခုမှ printed
#    "Tailscale HTTPS" URL ကို ဖွင့်ပြီး token ကို တစ်ကြိမ် ခေတ်ဝှန်ပါ။
```

> Default ဖြင့် pi-web သည် `PI_WEB_TOKEN` သည် set ဖြစ်နေပါက non-loopback address သို့ bind လုပ်ရန် reject လုပ်သည်။ ခေတ်တို bind address ကို ခေတ်ထိစပ်နိုင်သောသူတိုင်း sessions များကို ခေတ်ကြည့်ရှုပြီး pi သို့ instructions များကို ခေတ်တင်နိုင်သည်။ local-network testing အတွက် ဤ guard ကို override လုပ်ရန် `--insecure` ကို pass လုပ်ပါ။ **Tailscale တွင် သို့မဟုတ် machine ကျော်အပြင် ထိစပ်နိုင်သော address များတွင် `--insecure` ကို အသုံးမပြုပါနှင့်။**
>
> Clients များသည် `Authorization: Bearer <token>` header, `X-Pi-Token` header, သို့မဟုတ် တစ်ကြိမ်တည်း `?token=<token>` (ခေတနာ `pi_token` cookie ကို subsequent requests များအတွက် set လုပ်သည်) ကြona token ကို pass လုပ်နိုင်သည်။ `?token=` ကြona pass လုပ်ထားသော tokens များသည် browser history, server access logs, နှင့် page တွင် links များမှ `Referer` headers တွင် ရှိနေသည်။ initial bookmark များအပြင် header form ကို ခေတ်ထောက်ပေးပါ။

## Browser Chat

Session page တစ်ခုကို ဖွင့်ပြီး၊ ခေတ် exact session ကို ဆက်လက်လုပ်ဆောင်ရန် အောက်ဘက် composer ကို အသုံးပြုပါ။

- `Enter` သည် send လုပ်သည်၊ `Shift+Enter` သည် newline ကို insert လုပ်သည်
- composer တွင် images များကို drag-and-drop သို့မဟုတ် paste လုပ်ပါ
- model picker နှင့် thinking-level selector သည် header တွင်ရှိပြီး — changes များသည် underlying pi worker တွင် ခေတ်တို apply လုပ်သည်
- active session တစ်ခုချင်းစီသည် ကိုယ်ပိုင် dedicated `pi --mode rpc` worker ကို ရရှိပြီး၊ sessions များတစ်ခုနှင့်တစ်ခုကို ခေတ်တို block မလုပ်နိုင်အောင်

## Sharing Sessions

Session page တွင် **Share** ကို ခေတ်နှိပ်၍ secret GitHub Gist တစ်ခုကို ဖန်တီးပါ။

Requirements:
- `gh` ကို တပ်ဆင်ထားပါ
- `gh auth login` ကို ပြီးမြောက်ပါ

Sharing သည် အောက်ပါတို့ကို ခေတ်ပြန်ရန်:
- secret gist URL
- `https://pi.dev/session/#<gistId>` တွင် preview URL

Shared gists များသည် snapshots များဖြစ်ပြီး live update မလုပ်ပါ။

## Login တွင် Auto-Start

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# systemd user service ကို တပ်ဆင်ပါ
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optional: non-loopback binds အတွက် PI_WEB_TOKEN ကို set လုပ်ပါ
# (သို့မဟုတ် pi အတွင်းမှ /pi-web set-token <token> ကို အသုံးပြုပါ)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# enable လုပ်ပြီး start လုပ်ပါ
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# status ကို ကြည့်ပါ
systemctl --user status pi-web.service

# logs များကို ခေတ်ကြည့်ပါ
journalctl --user -u pi-web.service -f
```

> boot တွင် (login မတိုင်မီ) service ကို start လုပ်ရန် system service ကို အသုံးပြုပါ:
> `init/pi-web.service` ကို `/etc/systemd/system/` သို့ copy လုပ်ပြီး `sudo systemctl` ကို အသုံးပြုပါ။

### Windows

Installer သည် admin rights မလိုအပ်ဘဲ automatically configure လုပ်ထားသည်: `HKCU\Software\Microsoft\Windows\CurrentVersion\Run` အောက်တွင် `pi-web` entry တစ်ခုသည် login တွင် `~/.config/pi-web/pi-web-start.vbs` ကို launch လုပ်ထားပြီး၊ `~/.config/pi-web/env` (`PI_WEB_TOKEN`, `PATH`, ...) ကို load လုပ်ပြီး binary ကို hidden (console window မရှိ) ဖြင့် start လုပ်ထားသည်။

hand ဖြင့် စီမံရန်:

```powershell
# start / stop လုပ်ရန်
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# auto-start ကို ဖယ်ရှားရန်
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Windows တွင် service supervision မရှိပါ: pi-web သည် crash ဖြစ်ပါက next login အထိ down ဖြစ်နေသည် (launchd/systemd သည် အခြား platforms တွင် automatically restart လုပ်ထားသည်)။
