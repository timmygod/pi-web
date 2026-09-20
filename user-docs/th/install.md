# การติดตั้งและการใช้งาน

## ฟีเจอร์

### ควบคุมระยะไกล

- ดำเนินการต่อเซสชันใดๆ จากเบราว์เซอร์พร้อมแนบข้อความหรือรูปภาพ
- เริ่มต้นเซสชันใหม่สำหรับ路径โปรเจกต์ใดๆ ได้เลยจาก UI บนเว็บ
- สลับโมเดลและเลือก ระดับการคิด (thinking-level) ในเบราว์เซอร์แยกตามแต่ละเซสชัน
- สถานะ worker ของแต่ละเซสชัน (ว่าง / กำลังทำงาน / เกิดข้อผิดพลาด) พร้อมการกู้คืนโดยอัตโนมัติเมื่อระบบล้ม
- รันหลายเซสชันพร้อมกันได้ — เริ่มงานในเซสชันหนึ่ง และดูอีกเซสชันที่สตรีมอยู่
- `PI_WEB_TOKEN` เพื่อการเปิดเข้าถึงใน LAN อย่างปลอดภัย — ถูกบังคับใช้เป็นค่าเริ่มต้นสำหรับ bind ที่ไม่ใช่ loopback ทุกกรณี

### การอ่านเซสชัน

- ค้นดูเซสชันข้ามโปรเจกต์พร้อมตัวกรอง การค้นหา และการนำทาง branch แบบเต็ม
- อัปเดตแบบเพิ่มตามช่วงเวลาแบบเรียลไทม์ขณะที่ pi ยังทำงานอยู่ (ผ่าน fsnotify; ความล่าช้า ~ms)
- โหมดติดตาม (Follow mode) สำหรับไล่ดูเซสชันที่กำลังทำงานอยู่
- ลิงก์ตรง (Deep links) ไปยังข้อความแต่ละตัว
- ดาวน์โหลดเซสชันเป็นไฟล์ JSONL
- แชร์ภาพนิ่ง (static snapshots) เป็น secret GitHub Gists
- พิชension ของ pi ได้แก่ `/web`, `/remote`, `/refresh`, `/pi-web token` และ `/pi-web set-token` สำหรับเปิดเซสชัน QR ระยะไกล ซิงก์เซสชัน และการจัดการโทเคน
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) เพื่อให้เซสชันสามารถจัดการตารางเวลา สเก็ตช์แพดของโปรเจกต์ และการตั้งค่าได้ด้วยภาษาธรรมชาติ

## การเลือกโหมดของเซสชัน

รุ่นนี้ใช้ provider และโมเดลที่ถูกกำหนดค่าไว้ใน pi แล้ว; Local Mode เป็นนโยบายของ runtime ไม่ใช่ตัวติดตั้งโมเดลแยกหรือหน้ากรอก API-key อีกหน้าหนึ่ง
เลือกโหมดเมื่อสร้างเซสชัน หรือเปลี่ยนภายหลังเมื่อการรันปัจจุบันเสร็จสิ้น:

| โหมด | ใช้เมื่อ | การทำงาน |
|------|-------------|----------|
| **Auto** | คุณต้องการให้ pi-web ตัดสินใจ | จะ resolve เอ็นด์พอยต์ local/LAN จาก metadata ของ provider เมื่อทำได้ หากไม่ได้จะคงเส้นทางปกติไว้ |
| **Local** | โมเดลกำลังรันอยู่บนเครื่องนี้หรือใน LAN ของคุณ | เปิดใช้งานขอบเขตการ compact ที่ 65%, checkpoints แบบจำกัด, Force Compact และมาตรการกู้คืนโดยอัตโนมัติแบบมีป้องกัน |
| **Cloud** | โมเดลที่เลือกถูกโฮสต์และควรยึดตามพฤติกรรม upstream | เก็บนโยบายการ compact และการกู้คืนเฉพาะ local ออกนอกเซสชัน |

การเลือก Local หรือ Cloud แบบ manual จะชนะการตรวจจับอัตโนมัติและจะคงอยู่ข้ามการ reload และ restart ด้วย เซสชันที่กำลังทำงานอยู่จะปฏิเสธการเปลี่ยนโหมดจนกว่า worker ของมันจะสิ้นสุดลง ดังนั้นโหมดที่แสดงใน UI จะตรงตามนโยบายที่ถูกใช้งานจริงเสมอ

## ข้อกำหนด

- [Go](https://go.dev) 1.25+ (เฉพาะกรณี build จากซอร์สโค้ดเท่านั้น)
- `pi` ที่อยู่ใน `PATH` ของคุณสำหรับการแชท/สลับโมเดลในเบราว์เซอร์
- opsional: `gh` สำหรับการแชร์
- ใน Windows: pi ต้องการ bash shell เพื่อเครื่องมือ shell — [Git for Windows](https://git-scm.com/download/win) เพียงพอก็ได้ (ดูเอกสาร Windows ของ pi)

## การติดตั้ง

### แพ็กเกจของ Pi (แนะนำ)

```bash
pi install npm:@timmygod/pi-web-local
```

คำสั่งเดียวนี้จะ:
- ติดตั้ง npm pi package ภายใต้ directori ของแพ็กเกจของ pi
- รันสคริปต์ `postinstall` ของแพ็กเกจ (`install.sh` หรือ `install.ps1` ใน Windows)
- โหลด pi-web binary ที่ตรงกับเวอร์ชันแพ็กเกจและ platform ของคุณจาก GitHub Releases
- ติดตั้งไปที่ `~/.pi/agent/bin/pi-web` (`pi-web.exe` ใน Windows)
- จัดการ auto-start ตอน log in (launchd บน macOS, systemd บน Linux, launcher คีย์ Run บน Windows)
- ลงทะเบียนคำสั่ง pi ได้แก่ `/web`, `/remote`, `/refresh`, `/pi-web token` และ `/pi-web set-token`

การตั้งชื่อเซสชันอัตโนมัติถูกสร้างอยู่ใน pi-web (ไม่ใช่ใน extension) และกำหนดค่าในหน้า `/settings` เปิดไว้เป็นค่าเริ่มต้น: pi-web จะตั้งชื่อเซสชันอัตโนมัติโดยใช้ heuristic คำภายในที่ไม่เสียค่าใช้จ่าย (ไม่มี AI) และตั้งชื่อใหม่เมื่อมีข้อความใหม่ทุกครั้ง คุณสลับไปเป็นตั้งชื่อครั้งเดียวต่อเซสชันได้ และ/หรือเลือกโมเดลเพื่อสร้างชื่อที่ฉลาดขึ้นแทน heuristic

ใน Linux auto-start จะถูกกำหนดเป็น user systemd service อยู่ที่ `~/.config/systemd/user/pi-web.service` ตัวติดตั้งจะเขียน `ExecStart` ของมันใหม่เป็นเส้นทาง binary ที่ติดตั้งจริง ถ้า Tailscale ใช้งานได้ในขณะ runtime pi-web จะเผยแพร่ local host server ด้วย Tailscale Serve HTTPS ถ้า user systemd ใช้งานไม่ได้ ให้รันด้วยตัวเองโดยใช้ `~/.pi/agent/bin/pi-web -o`

เพื่อติดตั้งเฉพาะโปรเจกต์ที่ทำงานร่วมกัน (แชร์กับทีมผ่าน `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

จากนั้น restart pi (หรือรัน `/reload`) และใช้ `/web`, `/pi-web`, `/remote`, `/refresh` จัดการ access token ของคุณด้วย `/pi-web token` และ `/pi-web set-token`

ถ้า npm พังลงด้วย `ENOTEMPTY` ระหว่างการ rename `@timmygod/pi-web-local` ให้ลบ directori backup แบบซ่อนของ npm ที่ล้าสมัยและติดตั้งแพ็กเกจใหม่:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### ติดตั้งแบบเร็ว (ไม่ต้องใช้เครื่องมือ build)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

คำสั่งนี้จะโหลด pi-web binary ล่าสุด ติดตั้งไปที่ `/usr/local/bin` (`~/.pi/agent/bin` ใน Windows) และจัดการ auto-start ตอน log in ไม่ต้องใช้ Go, Node หรือ pi

### ดาวน์โหลด binary

binary ที่ build ไว้ล่วงหน้าจะถูกแนบกับ [GitHub Release](https://github.com/timmygod/pi-web/releases) แต่ละเวอร์ชัน

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

จากนั้นย้ายไปใส่ใน PATH ของคุณ:

```bash
cp pi-web ~/.pi/agent/bin/
# หรือทั้งระบบ:
sudo cp pi-web /usr/local/bin/
```

### Build จากซอร์สโค้ด

การ checkout นี้คือรุ่น local-model ของ pi-web การ build ปกติจะสร้าง web application และ backend ร่วมกัน; มาตรการป้องกันสำหรับ local-model จะถูกเปิดใช้งานในขณะ runtime ด้วย Local Mode ที่มีผลของเซสชัน ไม่ใช่ด้วย binary แยก

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # builds the Vite bundle, then embeds it into the Go binary

# optional: put it on PATH
cp pi-web ~/.pi/agent/bin/
```

frontend bundle จะถูกฝังโดย `web/assets_embed.go` ดังนั้น `go build` ต้องมี `web/dist` ให้มีอยู่ก่อน `make build` จะทำทั้งสองขั้นตามลำดับ; ถ้า build ด้วยตัวเอง ให้รัน `npm --prefix web install && npm --prefix web run build` ก่อน `go build ./cmd/pi-web`

สำหรับ workflow ของ fork ที่ดูแล การซิงก์ upstream และรายการตรวจสอบการตรวจสอบ Local Mode ดู [บันทึกการพัฒนา local-model](../../docs/dev/local-llm-development.md)

### พัฒนาควบคู่กับ instance ที่ติดตั้งแล้ว

ปล่อย instance ที่ติดตั้งแล้วทำงานต่อในพอร์ท `31415` จากนั้นเริ่มการ checkout ของซอร์สในโหมดพัฒนา:

```bash
make dev
```

เปิด `http://127.0.0.1:31416` `make dev` จะตั้งค่าสภาพแวดล้อมการพัฒนาภายใน `PI_WEB_DEV=1` ทำให้การ checkout ของซอร์สใช้เซสชัน การตั้งค่า และข้อมูล SQLite ร่วมกับ instance ที่ติดตั้งแล้ว ขณะเดียวกันก็เก็บ runtime lock และไฟล์สถานะของพัฒนาแยกต่างหาก Instance ที่ติดตั้งปกติและ instance ที่รันด้วยมือจะไม่เปลี่ยนแปลงและคงพฤติกรรม single-instance เดิมไว้

เพื่อป้องกันงานอัตโนมัติที่ซ้ำซ้อน โหมดพัฒนาจะไม่รัน schedule loop, chat-queue drainer, auto-titling หรือ push notifications คำสั่งตรงที่ดำเนินการผ่าน UI ของพัฒนา ยังคงทำงานได้ ไม่ควรขับเคลื่อนเซสชันแชทเดียวกันจากทั้งสอง instance พร้อมกัน; แต่ละ process มี RPC worker manager ของมันเอง

`make dev` ต้องใช้ [Air](https://github.com/air-verse/air) เพื่อ Go hot reload:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` เป็นกลไกของชุดทดสอบการพัฒนา ไม่ใช่โหมด multi-instance ที่รองรับใน production

## การถอนการติดตั้ง

```bash
pi remove npm:@timmygod/pi-web-local
```

คำสั่งนี้จะรันสคริปต์ `preuninstall` ของแพ็กเกจ (`uninstall.sh` หรือ `uninstall.ps1`
ใน Windows) ซึ่งจะหยุด instance ที่กำลังทำงานและลบ:

- pi-web binary (`~/.pi/agent/bin/pi-web` หรือ `/usr/local/bin/pi-web` สำหรับ standalone installs)
- ไฟล์เวอร์ชัน (`~/.pi/agent/pi-web-version`)
- ไฟล์สถานะ runtime (`~/.pi/agent/pi-web/pi-web-state.json`)
- การตั้งค่า auto-start (launchd plist บน macOS, systemd user service บน Linux, คีย์ Run + สคริปต์ launcher บน Windows)

ข้อมูลของคุณจะถูกเก็บรักษาไว้เพื่อให้การติดตั้งใหม่ในภายหลังต่อได้ตรงจุดที่คุณหยุดไว้:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, ไฟล์
เซสชันภายใต้ `~/.pi/agent/sessions/` และ `~/.config/pi-web/env` (รวมถึง
`PI_WEB_TOKEN`) หากต้องการเริ่มใหม่แบบสะอาด ให้ลบพวกนี้ด้วยตนเอง

## การใช้งาน

```bash
# Start on the default port (31415)
pi-web

# Start and open a browser
pi-web -o

# Custom port
pi-web -p 8080

# Override bind host (loopback is unauthenticated by default)
pi-web --host 127.0.0.1

# Non-loopback bind requires a token — pi-web refuses to start otherwise
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

เป็นค่าเริ่มต้น pi-web จะ bind ที่ `127.0.0.1` ถ้า Tailscale กำลังทำงานพร้อม MagicDNS **และตั้งค่า `PI_WEB_TOKEN` แล้ว** pi-web จะรัน `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` ด้วย และพิมพ์ URL HTTPS ของ tailnet ออกมา หากไม่มีโทเคน pi-web จะคง loopback-only และข้าม Tailscale Serve จึงทำให้ peer ใน tailnet เข้าถึง agent แบบไม่มีการยืนยันตัวตนไม่ได้ การ bind ที่ไม่ใช่ loopback อย่างชัดเจนก็ต้องการ `PI_WEB_TOKEN` ที่ถูกตั้งค่าเช่นกัน; ใส่ `--insecure` เพื่อ override สำหรับการทดสอบภายใน

## การเข้าถึงระยะไกล

ปล่อยให้ pi-web รับฟังภายใน แล้วใช้ URL Tailscale HTTPS ที่พิมพ์ออกมาจากโทรศัพท์หรือแล็ปท็อปที่อยู่ใน tailnet

บน macOS ให้ติดตั้งและเปิด Tailscale แบบ interactively อนุมัติ prompted ของผู้ดูแลระบบ และ log in จากนั้นรัน `/pi-web restart` ต่อด้วย `/remote`

บน Linux อนุญาตให้ user ของคุณจัดการ Tailscale ก่อนติดตั้ง/รัน pi-web มิฉะนั้น `tailscale serve` อาจต้องการ sudo และ auto-start อาจล้มเหลว:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Start pi-web with a token so it publishes the Tailscale HTTPS endpoint
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. From any other Tailscale-connected device, open the printed
#    "Tailscale HTTPS" URL and enter the token once.
```

> โดยค่าเริ่มต้น pi-web จะปฏิเสธการ bind ที่ไม่ใช่ loopback เว้นแต่จะตั้งค่า `PI_WEB_TOKEN` — มิฉะนั้นใครก็ตามที่สามารถเข้าถึง address ที่ bind ได้จะดูเซสชันและส่งคำสั่งไปยัง pi ได้ หากต้องการ override มาตรการนี้สำหรับการทดสอบใน local-network ให้ใส่ `--insecure` **อย่าใช้ `--insecure` บน Tailscale หรือ address ใดๆ ที่เข้าถึงได้จากนอกเครื่องของคุณ**
>
> Client สามารถส่งโทเคนผ่าน header `Authorization: Bearer <token>`, header `X-Pi-Token` หรือครั้งเดียวผ่าน `?token=<token>` (ซึ่งจะตั้ง cookie `pi_token` สำหรับคำขอครั้งต่อๆ ไป) โทเคนที่ถูกส่งผ่าน `?token=` จะถูกบันทึกไว้ใน history ของเบราว์เซอร์ access logs ของเซิร์ฟเวอร์ และ header `Referer` จากลิงก์ใดๆ บนหน้าเว็บ — ควรใช้รูปแบบ header สำหรับทุกอย่างที่ไม่ใช่ bookmark ครั้งแรก

## แชทในเบราว์เซอร์

เปิดหน้าของเซสชัน และใช้ composer ด้านล่างเพื่อดำเนินการต่อเซสชันนั้นๆ

- `Enter` ส่งข้อความ, `Shift+Enter` chèn换行
- ลากและวางหรือ paste รูปภาพโดยตรงลงใน composer
- Model picker และ thinking-level selector อยู่ในส่วน header — การเปลี่ยนแปลงจะใช้ผลกับ pi worker ล่างทันที
- เซสชันที่กำลังทำงานแต่ละเซสชันจะมี worker `pi --mode rpc` ของตัวเอง ดังนั้นเซสชันที่แตกต่างกันจึงไม่กีดกันซึ่งกันและกัน

## การแชร์เซสชัน

คลิก **Share** บนหน้าเซสชันเพื่อสร้าง secret GitHub Gist

ข้อกำหนด:
- ต้องติดตั้ง `gh`
- ต้องเสร็จสิ้น `gh auth login`

การแชร์จะคืน:
- URL ของ secret gist
- preview URL ที่ `https://pi.dev/session/#<gistId>`

gist ที่แชร์เป็นภาพนิ่ง (snapshot) และไม่อัปเดตแบบเรียลไทม์

## Auto-Start ตอน Log in

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Install the systemd user service
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optional: set your PI_WEB_TOKEN for non-loopback binds
# (or use /pi-web set-token <token> from inside pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Enable and start
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Check status
systemctl --user status pi-web.service

# View logs
journalctl --user -u pi-web.service -f
```

> เพื่อให้ service เริ่มตอนบูต (ก่อน log in) ให้ใช้ system service แทน:
> คัดลอก `init/pi-web.service` ไปยัง `/etc/systemd/system/` และใช้ `sudo systemctl`

### Windows

ตัวติดตั้งจะจัดการให้โดยอัตโนมัติโดยไม่ต้องใช้สิทธิ์ admin:
คีย์ `pi-web` ภายใต้ `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
จะเปิด `~/.config/pi-web/pi-web-start.vbs` ตอน log in ซึ่งจะเริ่มต้น binary
แบบซ่อน (ไม่มี console window) หลังจากโหลด `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...)

หากต้องการจัดการด้วยมือ:

```powershell
# Start / stop
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Remove auto-start
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

ใน Windows ไม่มี service supervision: ถ้า pi-web crash จะค้างอยู่จนกระทั่ง log in
ครั้งถัดไป (launchd/systemd จะ restart ให้โดยอัตโนมัติบน platform อื่น)
