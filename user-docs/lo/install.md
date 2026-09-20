# การติดตั้งและการใช้งาน

## ฟีเจอร์

### ควบคุมระยะไกล

- ต่อสายเซสชันใดก็ได้จากเบราว์เซอร์พร้อมไฟล์แนบข้อความหรือรูปภาพ
- เริ่มเซสชันใหม่ตั้งต้นที่ส่งตรงไปยังพาธโครงการใดๆ ได้เลย จาก UI บนเว็บ
- สลับโมเดลภายในเบราว์เซอร์และตัวเลือกระดับการคิด (thinking-level) รายเซสชัน
- สถานะ worker ของแต่ละเซสชัน (idle / running / error) พร้อมฟื้นตัวอัตโนมัติเมื่อ crash
- รันเซสชันหลายตัวขนานกัน — สั่งงานตัวหนึ่ง แล้วดูอีกตัวที่กำลังสตรีมอยู่
- `PI_WEB_TOKEN` เพื่อความปลอดภัยเมื่อเปิดเผยบน LAN — บังคับใช้โดยค่าเริ่มต้นสำหรับการ bind ชัดเจนที่ไม่ใช่ loopback

### การอ่านเซสชัน

- ไล่ดูเซสชันข้ามโครงการพร้อมตัวกรอง การค้นหา และการนำทาง branch เต็มรูปแบบ
- อัปเดตแบบ_incremental โดยสด ในขณะที่ pi ยังกำลังทำงาน (ผ่าน fsnotify; ความล่าช้าประมาณ ~ms)
- โหมด Follow สำหรับ追尾 (tail) เซสชันที่กำลังทำงาน
- ลิงก์直达 (deep link) ไปยังข้อความเฉพาะ
- ดาวน์โหลดเซสชันเป็น JSONL
- แชร์ snapshot แบบ static เป็น secret GitHub Gists
- การขยาย (pi extensions) ของ pi ได้แก่ `/web`, `/remote`, `/refresh`, `/pi-web token` และ `/pi-web set-token` สำหรับเปิดเซสชัน, QR ระยะไกล, ซิงค์เซสชัน และการจัดการ token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) เพื่อให้เซสชันจัดการตารางเวลา, กระดานสเก็ตช์ (scratchpad) ของโครงการ และการตั้งค่าด้วยภาษาธรรมชาติ

## เลือกโหมดเซสชัน

รุ่นนี้ใช้ providers และโมเดลที่ตั้งค่าไว้แล้วใน pi; โหมด Local เป็นนโยบาย runtime ไม่ใช่ตัวติดตั้งโมเดลแยกต่างหาก หรือหน้า API-key ตัวที่สอง
เลือกโหมดเมื่อสร้างเซสชัน หรือเปลี่ยนภายหลังเมื่อการรันปัจจุบันเสร็จนิ่ง:

| โหมด | ใช้เมื่อ | พฤติกรรม |
|------|-------------|----------|
| **Auto** | คุณต้องการให้ pi-web ตัดสินใจ | ปรับปลายทาง local/LAN จาก metadata ของ provider เมื่อเป็นไปได้; หรือไม่งั้นก็ยังคงเส้นทางปกติ |
| **Local** | โมเดลกำลังรันบนเครื่องนี้หรือ LAN ของคุณ | เปิดขอบเขต compaction 65%, checkpoints ที่จำกัดขนาด, Force Compact และการฟื้นตัวอัตโนมัติภายใต้การป้องกัน |
| **Cloud** | โมเดลที่เลือกเป็นแบบ hosted และควรใช้พฤติกรรมแบบ upstream | ยีงเว้นนโยบาย compaction และ recovery ที่เฉพาะ local ไว้ไม่ให้อยู่ในเซสชัน |

การเลือก Local หรือ Cloud แบบ manual ชนะการตรวจจับอัตโนมัติ และคงอยู่ข้ามการโหลดใหม่และการรีสตาร์ท เซสชันที่กำลังรันจะปฏิเสธการเปลี่ยนโหมดจนกว่า worker ของมันจะนิ่ง ดังนั้นโหมดที่แสดงใน UI จึงตรงกับนโยบายที่ถูกใช้จริงเสมอ

## ความต้องการ

- [Go](https://go.dev) 1.25+ (ใช้เฉพาะในการ build จาก source)
- `pi` บน `PATH` ของคุณ สำหรับ chat/การสลับโมเดลในเบราว์เซอร์
- เพิ่มเติม (optional): `gh` สำหรับแชร์
- บน Windows: pi ต้องการ bash shell สำหรับ shell tool ของมัน — [Git for Windows](https://git-scm.com/download/win) เพียงพอ (ดูเอกสาร Windows ของ pi)

## ติดตั้ง

### ปิ แพ็กเกจ (แนะนำ)

```bash
pi install npm:@timmygod/pi-web-local
```

คำสั่งเดียวนี้:
- ติดตั้ง npm pi package ภายใต้ไดเรกทอรี package ของ pi
- รันสคริปต์ `postinstall` ของ package (`install.sh` หรือ `install.ps1` บน Windows)
- ดาวน์โหลด binary pi-web ที่ตรงกับเวอร์ชัน package และแพลตฟอร์มของคุณ จาก GitHub Releases
- ติดตั้งไปที่ `~/.pi/agent/bin/pi-web` (`pi-web.exe` บน Windows)
- ตั้งค่าเริ่มทำงานอัตโนมัติเมื่อเข้าสู่ระบบ (launchd บน macOS, systemd บน Linux, Run-key launcher บน Windows)
- จดทะเบียนคำสั่ง pi ได้แก่ `/web`, `/remote`, `/refresh`, `/pi-web token` และ `/pi-web set-token`

การตั้งชื่อเซสชันอัตโนมัติถูกรวมอยู่ใน pi-web (ไม่ใช่ใน extension) และตั้งค่าได้ที่หน้า `/settings` เปิดไว้โดยค่าเริ่มต้น: pi-web ตั้งชื่อเซสชันอัตโนมัติโดยใช้ heuristic คำศัพท์ builtin ที่ฟรี (ไม่ใช้ AI) และตั้งชื่อใหม่ทุกข้อความใหม่ คุณสามารถสลับไปตั้งชื่อหนึ่งครั้งต่อเซสชัน และ/หรือเลือกโมเดลเพื่อเขียนชื่อที่ชาญฉลาดขึ้นแทน heuristic

บน Linux การเริ่มทำงานอัตโนมัติถูกตั้งค่าเป็น user systemd service ที่ `~/.config/systemd/user/pi-web.service` ตัวติดตั้งจะเขียน `ExecStart` ใหม่เป็นพาธ binary ที่ติดตั้งจริง หาก Tailscale มีอยู่ขณะ runtime pi-web จะเผยแพร่เซิร์ฟเวอร์ localhost ด้วย Tailscale Serve HTTPS หาก user systemd ใช้ไม่ได้ ให้รันเองด้วย `~/.pi/agent/bin/pi-web -o`

เพื่อติดตั้งเฉพาะสำหรับโครงการเฉพาะ (แชร์กับทีมของคุณผ่าน `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

จากนั้นรีสตาร์ท pi (หรือรัน `/reload`) และใช้ `/web`, `/pi-web`, `/remote`, `/refresh` จัดการ access token ของคุณด้วย `/pi-web token` และ `/pi-web set-token`

หาก npm หยุดด้วย `ENOTEMPTY` ขณะ rename `@timmygod/pi-web-local` ให้ลบไดเรกทอรี backup ซ่อนค้างของ npm แล้วติดตั้ง package ใหม่:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### ติดตั้งด่วน (ไม่ต้องใช้เครื่องมือ build)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

สิ่งนี้ดาวน์โหลด binary pi-web ล่าสุด ติดตั้งไปที่ `/usr/local/bin` (`~/.pi/agent/bin` บน Windows) และตั้งค่าเริ่มทำงานอัตโนมัติเมื่อเข้าสู่ระบบ ไม่ต้องใช้ Go, Node หรือ pi

### ดาวน์โหลด binary

binary ที่ build ไว้ล่วงหน้าถูกแนบมากับแต่ละ [GitHub Release](https://github.com/timmygod/pi-web/releases)

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

จากนั้นย้ายมันไปที่ PATH ของคุณ:

```bash
cp pi-web ~/.pi/agent/bin/
# หรือทั้งระบบ:
sudo cp pi-web /usr/local/bin/
```

### Build จาก source

checkout นี้คือรุ่น local-model ของ pi-web การ build ปกติสร้าง web application และ backend ด้วยกัน; ความปลอดภัยสำหรับ local-model ถูกเปิดใน runtime ด้วย Local Mode ที่มีผลของเซสชัน ไม่ใช่ด้วย binary แยกต่างหาก

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # build Vite bundle จากนั้นฝังมันเข้าไปใน binary ของ Go

# optional: วางมันบน PATH
cp pi-web ~/.pi/agent/bin/
```

frontend bundle ถูกฝังโดย `web/assets_embed.go` ดังนั้น `go build` ต้องมี
`web/dist` ก่อน `make build` ทำทั้งสองขั้นตอนตามลำดับ; หากคุณ build
ด้วยมือ ให้รัน `npm --prefix web install && npm --prefix web run build` ก่อน
`go build ./cmd/pi-web`

สำหรับ workflow fork ที่บำรุงรักษาไว้ การซิงค์กับ upstream และรายการตรวจสอบการยืนยัน Local Mode ดู [ заметки การพัฒนา local-model](../../docs/dev/local-llm-development.md)

### พัฒนาไปพร้อมกันกับ instance ที่ติดตั้งไว้

ปล่อย instance ที่ติดตั้งไว้ทำงานต่อที่พอร์ต `31415` จากนั้นเปิด
checkout ของ source ในโหมด development:

```bash
make dev
```

เปิด `http://127.0.0.1:31416` `make dev` ตั้งค่าสภาพแวดล้อม `PI_WEB_DEV=1`
ภายในสำหรับ development ดังนั้น checkout ของ source จึงแชร์เซสชัน การตั้งค่า และข้อมูล SQLite กับ instance ที่ติดตั้งไว้ โดยยังคง runtime lock และไฟล์ state ของ development ไว้แยกต่างหาก instance ที่ติดตั้งไว้ปกติและที่เริ่มต้นด้วยมือไม่เปลี่ยนและคงพฤติกรรม single-instance แบบเดิม

เพื่อป้องกันการทำงาน autonomous ซ้ำซ้อน โหมด development ไม่ได้รัน
schedule loop, chat-queue drainer, auto-titling หรือ push notifications คำขอโดยตรงที่ทำผ่าน UI ของ development ยังคงทำงานได้ อย่าขับเซสชัน chat เดียวกันจากทั้งสอง instance พร้อมกัน; แต่ละกระบวนการมีตัวจัดการ RPC worker ของตัวเอง

`make dev` ต้องการ [Air](https://github.com/air-verse/air) สำหรับ Go hot reload:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` เป็นเครื่องมือของ harness สำหรับ development ไม่ใช่โหมด multi-instance อย่างเป็นทางการของ production

## เลิกติดตั้ง

```bash
pi remove npm:@timmygod/pi-web-local
```

สิ่งนี้รันสคริปต์ `preuninstall` ของ package (`uninstall.sh` หรือ `uninstall.ps1`
บน Windows) ซึ่งหยุด instance ที่กำลังทำงานและลบ:

- binary ของ pi-web (`~/.pi/agent/bin/pi-web` หรือ `/usr/local/bin/pi-web` สำหรับ installation แบบ standalone)
- ไฟล์เวอร์ชัน (`~/.pi/agent/pi-web-version`)
- ไฟล์ state ของ runtime (`~/.pi/agent/pi-web/pi-web-state.json`)
- การตั้งค่าเริ่มทำงานอัตโนมัติ (launchd plist บน macOS, user service ของ systemd บน Linux, Run-key entry + สคริปต์ launcher บน Windows)

ข้อมูลของคุณยังคงอยู่ ดังนั้นการติดตั้งใหม่ภายหลังจะต่อตรงที่คุณหยุดไว้:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, ไฟล์เซสชัน
ของคุณใต้ `~/.pi/agent/sessions/` และ `~/.config/pi-web/env` (รวมถึง
`PI_WEB_TOKEN`) หากต้องการเริ่มใหม่โดยสะอาด ให้ลบสิ่งเหล่านี้ด้วยมือ

## การใช้งาน

```bash
# เริ่มที่พอร์ตค่าเริ่มต้น (31415)
pi-web

# เริ่มและเปิดเบราว์เซอร์
pi-web -o

# พอร์ตกำหนดเอง
pi-web -p 8080

# override bind host (loopback ไม่มี authentication โดยค่าเริ่มต้น)
pi-web --host 127.0.0.1

# bind ที่ไม่ใช่ loopback ต้องการ token — pi-web ปฏิเสธที่จะเริ่มในทางกลับกัน
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

โดยค่าเริ่มต้น pi-web bind ไปที่ `127.0.0.1` หาก Tailscale กำลังรันด้วย MagicDNS **และ `PI_WEB_TOKEN` ถูกตั้งไว้** pi-web จะรัน `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` ด้วย และพิมพ์ URL ของ HTTPS tailnet หากไม่มี token pi-web จะคง loopback-only และข้าม Tailscale Serve ดังนั้น peer ใน tailnet จึงไม่สามารถเข้าถึง agent โดยไม่มี authentication การ bind ที่ไม่ใช่ loopback ที่ระบุชัดก็ต้องการ `PI_WEB_TOKEN` ให้ตั้งเช่นกัน; ส่ง `--insecure` เพื่อ override สำหรับทดสอบภายใน

## การเข้าถึงระยะไกล

ปล่อย pi-web รับฟังในระดับ local แล้วใช้ URL ของ Tailscale HTTPS ที่พิมพ์มาจากโทรศัพท์หรือแล็ปท็อปบน tailnet

บน macOS ติดตั้งและเปิด Tailscale แบบ interactive ยืนยันคำสั่งของ administrator และเข้าสู่ระบบ จากนั้นรัน `/pi-web restart` ตามด้วย `/remote`

บน Linux อนุญาตให้ผู้ใช้ของคุณจัดการ Tailscale ก่อนติดตั้ง/รัน pi-web มิฉะนั้น `tailscale serve` อาจต้องการ sudo และการเริ่มทำงานอัตโนมัติอาจล้มเหลว:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. เริ่ม pi-web ด้วย token เพื่อให้มันเผยแพร่ endpoint ของ Tailscale HTTPS
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. จากอุปกรณ์อื่นที่เชื่อม Tailscale แล้ว ให้เปิด
#    URL "Tailscale HTTPS" ที่พิมพ์ไว้ และป้อน token หนึ่งครั้ง
```

> โดยค่าเริ่มต้น pi-web ปฏิเสธที่จะ bind ไปที่ address ที่ไม่ใช่ loopback เว้นแต่ `PI_WEB_TOKEN` จะถูกตั้งไว้ — โดยไม่เช่นนั้นผู้ใดก็ตามที่สามารถเข้าถึง address ที่ bind ได้สามารถดูเซสชันและส่งคำสั่งไปที่ pi ได้ เพื่อ override การป้องกันนี้สำหรับการทดสอบบนเครือข่ายภายใน ให้ส่ง `--insecure` **อย่าใช้ `--insecure` บน Tailscale หรือ address ใดที่เข้าถึงได้จากภายนอกเครื่องของคุณ**
>
> Client สามารถส่ง token ผ่าน header `Authorization: Bearer <token>`, header `X-Pi-Token` หรือหนึ่งครั้งผ่าน `?token=<token>` (ซึ่งตั้ง cookie `pi_token` สำหรับการขอขอกับกำลังจะมาถึง) token ที่ส่งผ่าน `?token=` จะเข้าไปอยู่ใน history ของเบราว์เซอร์ server access logs และ header `Referer` จากลิงก์ทุกตัวบนหน้า — แนะนำให้ใช้รูปแบบ header สำหรับทุกสิ่งเกินจาก bookmark เริ่มต้น

## Chat ในเบราว์เซอร์

เปิดหน้าของเซสชันแล้วใช้ composer ที่ด้านล่างเพื่อต่อเซสชันนั้นพอดี

- `Enter` ส่งข้อความ `Shift+Enter` แทรกบรรทัดใหม่
- ลาก-วางหรือวางรูปโดยตรงเข้าไปใน composer
- ตัวเลือกโมเดลและตัวเลือกระดับการคิดอยู่ที่ header — การเปลี่ยนแปลงมีผลกับ pi worker ฐานทันที
- แต่ละเซสชันที่กำลังทำงานได้ worker `pi --mode rpc` ส่วนตัวของมันเอง ดังนั้นเซสชันต่างๆ จึงไม่บล็อกกัน

## แชร์เซสชัน

คลิก **แชร์** บนหน้าของเซสชันเพื่อสร้าง secret GitHub Gist

ความต้องการ:
- ติดตั้ง `gh` แล้ว
- `gh auth login` เสร็จแล้ว

การแชร์คืน:
- URL ของ secret gist
- URL preview ที่ `https://pi.dev/session/#<gistId>`

gist ที่แชร์เป็น snapshot และไม่ได้ live-update

## เริ่มทำงานอัตโนมัติเมื่อเข้าสู่ระบบ

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# ติดตั้ง user service ของ systemd
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# optional: ตั้ง PI_WEB_TOKEN ของคุณสำหรับ bind ที่ไม่ใช่ loopback
# (หรือใช้ /pi-web set-token <token> จากภายใน pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# เปิดใช้งานและเริ่ม
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# ดูสถานะ
systemctl --user status pi-web.service

# ดู log
journalctl --user -u pi-web.service -f
```

> เพื่อให้ service เริ่มเมื่อเครื่องเปิด (ก่อนเข้าสู่ระบบ) ให้ใช้ระบบ service แทน:
> คัดลอก `init/pi-web.service` ไปที่ `/etc/systemd/system/` และใช้ `sudo systemctl`

### Windows

ตัวติดตั้งตั้งค่าสิ่งนี้อย่างอัตโนมัติ โดยไม่ต้องใช้สิทธิ์ admin:
entry `pi-web` ภายใต้ `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
จะเปิด `~/.config/pi-web/pi-web-start.vbs` เมื่อเข้าสู่ระบบ ซึ่งเริ่ม binary
แบบซ่อน (ไม่มี console window) หลังจากโหลด `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...)

เพื่อจัดการด้วยมือ:

```powershell
# เริ่ม / หยุด
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# ลบการเริ่มทำงานอัตโนมัติ
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

ไม่มี service supervision บน Windows: หาก pi-web crash มันจะหยุดทำงาน
จนถึงการเข้าสู่ระบบครั้งถัดไป (launchd/systemd restart มันอัตโนมัติบนแพลตฟอร์ม
อื่น)
