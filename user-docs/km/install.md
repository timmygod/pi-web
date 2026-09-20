# ការដំឡើង និងការប្រើប្រាស់

## មុខងារ

### ការគ្រប់គ្រងចម្ងាយ

- បន្តលំហូរ (session) ណាមួយពី browser ដោយប្រើអត្ថបទ ឬ temp file រូបភាព
- ចាប់ផ្ដើម session ថ្មីបំផុតទៅកាន់ project path ណាមួយ ដោយផ្ទាល់ពី web UI
- ការប្តូរ model និង thinking-level selector ក្នុង browser សម្រាប់ session ម្ដងមួយ
- ស្ថានភាព worker ក្នុង session ម្ដងមួយ (idle / running / error) ជាមួយ auto-recovery ពេល crash
- Session ជាច្រើនដំណើរការស្របគ្នា — ចាប់ផ្ដើមការងារនៅក្នុងមួយ ហើយមើលមួយទៀតកំពុង stream
- `PI_WEB_TOKEN` សម្រាប់ការបង្ហាញ LAN ដែលមានសុវត្ថិភាព — ត្រូវការដោយគេហលំនាំ ចំពោះ explicit non-loopback bind គ្រប់ទាំងអស់

### ការអាន sessions

- រុករក sessions រវាង projects ជាមួយ filters, search និងការរុករក branch ពេញលេញ
- ការធ្វើបច្ចុប្បន្នភាព incremental ភ្លាមៗ ខណៈដែល pi កំពុងដំណើរការនៅទៀត (តាមរយៈ fsnotify; latency ~ms)
- Follow mode សម្រាប់ការ tail session កំពុងធ្វើចលនា
- Deep links ទៅកាន់ messages ម្ដងមួយ
- ទាញយក session ជារាបស្តុក (format) JSONL
- ចែករំលែក snapshots ថេរជា secret GitHub Gists
- ការពង្រីក (extensions) `/web`, `/remote`, `/refresh`, `/pi-web token` និង `/pi-web set-token` របស់ pi សម្រាប់ការបើក sessions, remote QR, session sync និងការគ្រប់គ្រង token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) ដើម្បីឱ្យ session អាចគ្រប់គ្រង schedules, project scratchpad និង settings ដោយប្រើភាសាធម្មជាតិ

## ជ្រើសរើស session mode

ប្រហាប់នេះប្រើ providers និង models ដែលបានកំណត់រួចរាល់ក្នុង pi ហើយ; Local Mode
គឺជា runtime policy មិនមែនជា separate model installer ឬ second API-key screen នោះទេ។
ជ្រើសរើស mode ពេលបង្កើត session ឬផ្លាស់ប្ដូរវា បន្ទាប់ពីការដំណើរការបច្ចុប្បន្នបានស្រេចស្រួល៖

| Mode | ប្រើវាពេល | ឥរិយាបថ |
|------|-------------|----------|
| **Auto** | អ្នកចង់ឱ្យ pi-web សម្រេច | កំណត់ local/LAN endpoints ពី provider metadata ពេលអាច; បើមិនអាចទេ រក្សាផ្លូវធម្មតា |
| **Local** | Model កំពុងដំណើរការនៅលើម៉ាស៊ីននេះ ឬ LAN របស់អ្នក | បើក compaction boundary 65%, bounded checkpoints, Force Compact និង guarded automatic recovery |
| **Cloud** | Model ដែលបានជ្រើសរើសត្រូវបាន hosted និងគួរតាមឥរិយាបថ upstream | រក្សា compaction និង recovery policy ដែលជា local-only ឲ្យនៅក្រៅ session |

ការជ្រើសរើស Local ឬ Cloud ដោយដៃឈ្នះជាងការសម្គាល់ដោយស្វ័យប្រវត្តិ និងរក្សាទុកជារៀងរាល់
ពេល reload និង restart។ Session កំពុងដំណើរការបដិសេធនការផ្លាស់ប្ដូរ mode រហូតដល់ worker
របស់វាបានស្រេចស្រួល ដូច្នេះ mode ដែលបង្ហាញក្នុង UI ជានិច្ចត្រូវស្របនឹង policy ដែលកំពុងប្រើពិតប្រាកដ។

## តម្រូវការ

- [Go](https://go.dev) 1.25+ (ប្រើសម្រាប់ការក្រាង (build) ពី source ប៉ុណ្ណោះ)
- `pi` នៅលើ `PATH` របស់អ្នក សម្រាប់ browser chat/model switching
- ជម្រើស៖ `gh` សម្រាប់ការចែករំលែក
- នៅលើ Windows៖ pi ត្រូវការ bash shell សម្រាប់ shell tool របស់វា — [Git for Windows](https://git-scm.com/download/win) គ្រប់គ្រាន់ (មើល Windows docs របស់ pi)

## ការដំឡើង

### Pi package (ណែនាំ)

```bash
pi install npm:@timmygod/pi-web-local
```

បញ្ជាតែមួយនេះ៖
- ដំឡើង npm pi package ក្រោម package directory របស់ pi
- ដំណើរការ package `postinstall` script (`install.sh`, ឬ `install.ps1` នៅលើ Windows)
- ទាញយក pi-web binary ដែលត្រូវគ្នាសម្រាប់ package version និង platform របស់អ្នក ពី GitHub Releases
- ដំឡើងវាទៅ `~/.pi/agent/bin/pi-web` (`pi-web.exe` នៅលើ Windows)
- កំណត់ auto-start ពេល login (launchd នៅលើ macOS, systemd នៅលើ Linux, Run-key launcher នៅលើ Windows)
- ចុះឈ្មោះ pi commands `/web`, `/remote`, `/refresh`, `/pi-web token` និង `/pi-web set-token`

ការដាក់ឈ្មោះ session ដោយស្វ័យប្រវត្តិ (auto-titling) ត្រូវបានដាក់ចូលក្នុង pi-web (មិនមែន extension) ហើយកំណត់នៅលើ `/settings` page។ វាបើកដោយគេហលំនាំ៖ pi-web ដាក់ឈ្មោះ sessions ដោយស្វ័យប្រវត្តិដោយប្រើ word heuristic ដែលមានក្នុងខ្លួន ដោយមិនគិតថ្លៃ (គ្មាន AI) ហើយដាក់ឈ្មោះឡើងវិញនៅពេលមាន message ថ្មីនីមួយៗ។ អ្នកអាចប្តូរទៅ titling តែម្ដងក្នុង session មួយ និង/ឬជ្រើសរើស model ដើម្បីសរសេរ titles ដែលឆ្លើតជាង heuristic។

នៅលើ Linux auto-start ត្រូវបានកំណត់ជា user systemd service នៅ `~/.config/systemd/user/pi-web.service`។ Installer ប៉ុនប៉ុនប៉ិះ (rewrite) `ExecStart` របស់វាទៅ binary path ដែលបានដំឡើងពិតប្រាកដ។ បើ Tailscale មាននៅពេល runtime pi-web បង្ហាញ localhost server ដោយប្រើ Tailscale Serve HTTPS។ បើ user systemd មិនមានក៏ដោយ ដំណើរការវាដោយដៃជាមួយ `~/.pi/agent/bin/pi-web -o`។

ដើម្បីដំឡើងសម្រាប់ project ជាក់លាក់ណាមួយប៉ុណ្ណោះ (ចែករំលែកជាមួយក្រុមរបស់អ្នកតាមរយៈ `.pi/settings.json`)៖

```bash
pi install -l npm:@timmygod/pi-web-local
```

បន្ទាប់មក restart pi (ឬដំណើរការ `/reload`) ហើយប្រើ `/web`, `/pi-web`, `/remote`, `/refresh`។ គ្រប់គ្រង access token របស់អ្នកដោយប្រើ `/pi-web token` និង `/pi-web set-token`។

បើ npm បង្ខូត (abort) ដោយ `ENOTEMPTY` ខណៈកំពុងប្ដូរឈ្មោះ `@timmygod/pi-web-local` យក directories កំណត់ទុក (stale hidden backup) របស់ npm ចេញ និងដំឡើង package ឡើងវិញ៖

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### ការដំឡើងរហ័ស (មិនត្រូវការ build tools)

macOS / Linux៖

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell)៖

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

វាទាញយក pi-web binary ថ្មីបំផុត ដំឡើងវាទៅ `/usr/local/bin` (`~/.pi/agent/bin` នៅលើ Windows) និងកំណត់ auto-start ពេល login។ មិនត្រូវការ Go, Node ឬ pi។

### ការទាញយក binary

Pre-built binaries ត្រូវបានភ្ជាប់ជាមួយ [GitHub Release](https://github.com/timmygod/pi-web/releases) នីមួយៗ។

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

បន្ទាប់មក ផ្ទាច់វាទៅ PATH របស់អ្នក៖

```bash
cp pi-web ~/.pi/agent/bin/
# ឬ system-wide៖
sudo cp pi-web /usr/local/bin/
```

### ការក្រាង (Build) ពី source

Checkout នេះគឺជា local-model edition របស់ pi-web។ Build ធម្មតាផលិត
web application និង backend ជាមួយគ្នា; local-model safeguards ត្រូវបានបើក
នៅពេល runtime ដោយ session's effective Local Mode មិនមែនជា separate binary ទេ។

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # ក្រាង Vite bundle បន្ទាប់មក embed វាចូលក្នុង Go binary

# optional: ដាក់វានៅលើ PATH
cp pi-web ~/.pi/agent/bin/
```

Frontend bundle ត្រូវបាន embed ដោយ `web/assets_embed.go` ដូច្នេះ `go build` ត្រូវការ
`web/dist` ដែលមានជាមុនសិន។ `make build` ធ្វើជំហានទាំងពីរតាមលំដាប់; បើអ្នកក្រាង
ដោយដៃ ដំណើរការ `npm --prefix web install && npm --prefix web run build` មុន
`go build ./cmd/pi-web`។

សម្រាប់ maintained fork workflow, upstream synchronization និង Local Mode
verification checklist សូមមើល [ការណែនាំផ្ដាច់ (notes) អំពីការអភិវឌ្ឍ local-model](../../docs/dev/local-llm-development.md)។

### ការអភិវឌ្ឍ (Develop) ជាមួយ installed instance

បន្ទាប់ពី instance ដែលបានដំឡើង ឱ្យវាដំណើរការនៅលើ port `31415` បន្ទាប់មកចាប់ផ្ដើម source
checkout ក្នុង development mode៖

```bash
make dev
```

បើក `http://127.0.0.1:31416`។ `make dev` កំណត់ development environment 内部 `PI_WEB_DEV=1`
ដូច្នេះ source checkout ចែករំលែក sessions, settings និង SQLite data ជាមួយ installed instance ខណៈដែលរក្សា development runtime lock និង state file ដាច់ពីគ្នា។ Installed instances ដែលបើកដោយដៃធម្មតា មិនផ្លាស់ប្ដូរ ហើយរក្សា single-instance behavior ដើមនៅដដែល។

ដើម្បីការពារ autonomous work ដែលស៊ីគ្នា (duplicate) development mode មិនដំណើរការ schedule loop, chat-queue drainer, auto-titling ឬ push notifications ទេ។ ការស្នើ (Direct requests) ដែលធ្វើតាមរយៈ development UI នៅតែដំណើរការ។ កុំបញ្ជា chat session ដូចគ្នាពី instance ទាំងពីរនៅពេលតែមួយ; ដំណើរការ (process) ម្ដងមួយមាន RPC worker manager ផ្ទាល់ខ្លួន។

`make dev` ត្រូវការ [Air](https://github.com/air-verse/air) សម្រាប់ Go hot reload៖

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` គឺជា development harness plumbing មិនមែនជា supported production multi-instance mode ទេ។

## ការដោះកំណត់ (Uninstall)

```bash
pi remove npm:@timmygod/pi-web-local
```

វាដំណើរការ package `preuninstall` script (`uninstall.sh`, ឬ `uninstall.ps1`
នៅលើ Windows) ដែលឈប់ instance កំពុងដំណើរការ និងយកចេញ៖

- pi-web binary (`~/.pi/agent/bin/pi-web`, ឬ `/usr/local/bin/pi-web` សម្រាប់ standalone installs)
- version file (`~/.pi/agent/pi-web-version`)
- runtime state file (`~/.pi/agent/pi-web/pi-web-state.json`)
- auto-start config (launchd plist នៅលើ macOS, systemd user service នៅលើ Linux, Run-key entry + launcher scripts នៅលើ Windows)

ទិន្នន័យរបស់អ្នកត្រូវបានរក្សាទុក ដើម្បីឱ្យការដំឡើងឡើងវិញ ក្រោយនេះអាចចាប់ផ្ដើមពីកន្លែងដែលអ្នកបានទុកទុក៖
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, session files របស់អ្នកក្រោម `~/.pi/agent/sessions/` និង `~/.config/pi-web/env` (រួមទាំង
`PI_WEB_TOKEN`)។ យកវាចេញដោយដៃ បើអ្នកចង់បាន clean slate។

## ការប្រើប្រាស់

```bash
# ចាប់ផ្ដើមនៅលើ port ដោយគេហលំនាំ (31415)
pi-web

# ចាប់ផ្ដើម និងបើក browser
pi-web -o

# Port ផ្ទាល់ខ្លួន
pi-web -p 8080

# កំណត់ bind host ឡើងវិញ (loopback មិនត្រូវការ authentication ដោយគេហលំនាំ)
pi-web --host 127.0.0.1

# Non-loopback bind ត្រូវការ token — pi-web បដិសេធនការចាប់ផ្ដើមបើមិនមាន
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

ដោយគេហលំនាំ pi-web bind ទៅ `127.0.0.1`។ បើ Tailscale កំពុងដំណើរការជាមួយ MagicDNS **ហើយ `PI_WEB_TOKEN` ត្រូវបានកំណត់** pi-web ក៏ដំណើរការ `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` ហើយបោះពុម្ព HTTPS tailnet URL។ បើគ្មាន token pi-web នៅ loopback-only និងព្រម Tailscale Serve ដូច្នេះ tailnet peers មិនអាចទាញយក (reach) agent ដោយគ្មាន authentication ទេ។ Explicit non-loopback bind ណាមួយ ក៏ត្រូវការ `PI_WEB_TOKEN` ដែលត្រូវបានកំណត់ដែរ; បញ្ជូន `--insecure` ដើម្បី override សម្រាប់ការសាកល្បងក្នុងបណ្ដាញក្នុងផ្ទះ។

## ការចូលប្រើចម្ងាយ (Remote Access)

បន្ទាប់ពី pi-web កំពុងស្តាប់ (listen) ក្នុងផ្ទះ (local) បន្ទាប់មកប្រើ Tailscale HTTPS URL ដែលបានបោះពុម្ព ពីទូរស័ព្ទ ឬ laptop របស់អ្នកនៅលើ tailnet។

នៅលើ macOS ដំឡើង និងបើក Tailscale ដោយ interactive យល់ព្រម administrator prompt និងចូល (sign in)។ បន្ទាប់មកដំណើរការ `/pi-web restart` បន្ទាប់មក `/remote`។

នៅលើ Linux អនុញ្ញាតឱ្យ user របស់អ្នកគ្រប់គ្រង Tailscale មុនពេលដំឡើង/ដំណើរការ pi-web បើមិនដូច្នេះ `tailscale serve` អាចត្រូវការ sudo ហើយ auto-start អាចបរាជ័យ៖

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. ចាប់ផ្ដើម pi-web ជាមួយ token ដើម្បីបង្ហាញ Tailscale HTTPS endpoint
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. ពី device ផ្សេងទៀតដែលបានចត (connect) ជាមួយ Tailscale បើក
#    "Tailscale HTTPS" URL ដែលបានបោះពុម្ព និងបញ្ចូល token ម្ដង។
```

> ដោយគេហលំនាំ pi-web បដិសេធនការ bind ទៅ non-loopback address លើកលែងតែ `PI_WEB_TOKEN` ត្រូវបានកំណត់ — អ្នកណាដែលអាចទាញយក (reach) address ដែលបាន bind ក៏អាចមើល sessions និងបញ្ជូនព្រឹត្តិបត្រ (instructions) ទៅ pi ដែរ។ ដើម្បី override ការការពារនេះសម្រាប់ការសាកល្បងបណ្ដាញក្នុងផ្ទះ បញ្ជូន `--insecure`។ **កុំប្រើ `--insecure` នៅលើ Tailscale ឬ address ណាដែលអាចទាញយក (reachable) ពីខាងក្រៅម៉ាស៊ីនរបស់អ្នក។**
>
> Clients អាចបញ្ជូន token តាមរយៈ `Authorization: Bearer <token>` header, `X-Pi-Token` header ឬម្ដងតែ `?token=<token>` (ដែលកំណត់ `pi_token` cookie សម្រាប់ subsequent requests)។ Tokens ដែលបានបញ្ជូនតាមរយៈ `?token=` ចាប់ផ្ដើម (end up) នៅក្នុង browser history, server access logs និង `Referer` headers ពី links ណា ក្នុង page — ម្នាញ (prefer) header form សម្រាប់អ្វីដែលក្រៅពី initial bookmark។

## Browser Chat

បើក session page ហើយប្រើ composer នៅផ្នែកក្រោម ដើម្បីបន្ត session ជាក់លាក់នោះ។

- `Enter` ផ្ញើ, `Shift+Enter` បញ្ចូល newline
- អូស-ដាក់ (Drag-and-drop) ឬ paste រូបភាពដោយផ្ទាល់ចូលក្នុង composer
- Model picker និង thinking-level selector ស្ថិតនៅក្នុង header — ការផ្លាស់ប្ដូរមានអំពើ (apply) ទៅកាន់ pi worker ក្រោមបន្ទាប់ភ្លាមៗ
- ម្ដងមួយ session កំពុងធ្វើចលនា ទទួលបាន `pi --mode rpc` worker ដែលជាក់លាក់ផ្នែករបស់វា ដូច្នេះ sessions ផ្សេងៗ មិនបិទ (block) គ្នាទេ

## ការចែករំលែក Sessions

ចុច **Share** នៅលើ session page ដើម្បីបង្កើត secret GitHub Gist។

តម្រូវការ៖
- `gh` ត្រូវបានដំឡើង
- `gh auth login` បានបញ្ចប់

ការចែករំលែកផ្ដល់ត្រឡប់៖
- secret gist URL
- preview URL នៅ `https://pi.dev/session/#<gistId>`

Shared gists គឺជា snapshots និងមិនធ្វើបច្ចុប្បន្នភាព (live-update) ទេ។

## Auto-Start ពេល Login

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# ដំឡើង systemd user service
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optional: កំណត់ PI_WEB_TOKEN របស់អ្នកសម្រាប់ non-loopback binds
# (ឬប្រើ /pi-web set-token <token> ពីក្នុង pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# បើក និងចាប់ផ្ដើម
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# ពិនិត្យមើល status
systemctl --user status pi-web.service

# មើល logs
journalctl --user -u pi-web.service -f
```

> ដើម្បីឱ្យ service ចាប់ផ្ដើមពេល boot (មុន login) ប្រើ system service ជំនួសវិញ៖
> ចម្លង `init/pi-web.service` ទៅ `/etc/systemd/system/` និងប្រើ `sudo systemctl`។

### Windows

Installer កំណត់វាដោយស្វ័យប្រវត្តិ ដោយមិនត្រូវការ admin rights៖
`pi-web` entry ក្រោម `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
ចាប់ផ្ដើម `~/.config/pi-web/pi-web-start.vbs` ពេល login ដែលចាប់ផ្ដើម binary
ដោយ concealer (no console window) បន្ទាប់ពី load `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...)។

ដើម្បីគ្រប់គ្រងវាដោយដៃ៖

```powershell
# ចាប់ផ្ដើម / ឈប់
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# យក auto-start ចេញ
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

គ្មាន service supervision នៅលើ Windows៖ បើ pi-web crash វានៅ down
រហូតដល់ login លើកក្រោយ (launchd/systemd ចាប់ផ្ដើមវាឡើងវិញដោយស្វ័យប្រវត្តិ នៅលើ
platforms ផ្សេង)។
