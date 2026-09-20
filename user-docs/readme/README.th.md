<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · **ไทย** · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

ใช้งาน [pi](https://pi.dev) coding agent ของคุณได้จากมือถือ แท็บเล็ต หรือแล็ปท็อป — ที่ไหนก็ได้บนเครือข่ายของคุณ หรือแบบระยะไกลผ่าน Tailscale

เป็น PWA เต็มรูปแบบ ทำให้คุณติดตั้งและใช้งานเหมือนแอปเนทีฟได้บนทุกอุปกรณ์ คิดว่านี่คือพื้นที่ทำงาน AI ส่วนตัวของคุณ — คล้าย Claude's Cowork แต่ใช้โมเดลอื่นได้ — แชทข้ามโมเดล โคดจากมือถือ หรือเปลี่ยนให้เป็น [ผู้ช่วยส่วนตัว](../en/personal-assistant.md) ที่อยู่บนเครื่องของคุณ

ปรับให้เป็นของคุณ: สลับธีมและฟอนต์ และใช้งานได้ในภาษาของคุณเอง — pi-web มาพร้อมหลายภาษาและคุณสามารถเพิ่มภาษาของตัวเองได้ คุณลักษณะเพิ่มเติมกำลังจะมา แต่จะไม่เต็มไปด้วยของที่ไม่จำเป็น: สิ่งที่เลือกใช้ได้ทั้งหมดสามารถปิดได้ใน settings

</div>

## ทำไมต้อง edeทื่ลโมเดลลอคคานี้?

pi-web ต้นฉบับยังคงเป็นฐานพื้นฐาน upstream สำหรับฟีเจอร์ที่ share ร่วมกันและ
fixes ทิศนี้รักษาก_experience นั้น แล้วเพิ่มชั้น reliability
สำหรับโมเดลที่รันอยู่บนเครื่องของคุณเองหรือที่อื่นบน LAN ของคุณ—ซึ่งการ generate
มักช้ากว่า ความจำมีจำกัด และ context ยาวอาจทำให้ session ที่สุขภาพดีปกติหยุดชะงัก

| Area | Upstream pi-web | ทิศนี้ |
|------|-----------------|--------------|
| Model/runtime policy | พฤติกรรมมาตรฐานของ pi-web | โหมด **Auto / Local / Cloud** ต่อ session พร้อม local detection ที่คำนึงถึง endpoint และ manual override แบบถาวร |
| Long-context handling | พฤติกรรม compaction ปกติของ pi | Local Mode ทำ compaction อย่าง proactively ที่ **65%** และตรวจสอบอีกครั้งในลูป tool-call ยาวก่อนคำขอโมเดลครั้งถัดไป |
| Compaction safety | summary มาตรฐาน | rolling checkpoints แบบ bounded, rewrite อีกครั้งสำหรับ output ที่ไม่ถูกต้อง/ถูกจำกัด และตรวจจับ no-progress แทนการ re-compactionEndless |
| Interrupted runs | การจัดการ worker และ error ปกติ | recovery แบบ bounded สำหรับ context overflow, thinking-only stops และ transport interruptions ที่เลือก พร้อม loop breakers แบบถาวร |
| Manual rescue | context details มาตรฐาน | **Force Compact** ยังพร้อมใช้ในฐานะ explicit recovery path โดยไม่ลบการสนทนา |
| Compatibility and releases | โปรเจกต์และ release line ต้นฉบับ | safeguards สำหรับ local-only ยังอยู่หลัง Local Mode; Cloud Mode รักษาวีติกรรม upstream, และ upstream changes ถูก review และ release ที่นี่อย่างอิสระ |

นี่ไม่ใช่การ rewrite หรือตัวแทนของ upstream แต่เป็น operating profile
ที่บำรุงรักษาอย่างตั้งใจสำหรับคนที่ต้องการ privacy และ control ของ local-model
โดยไม่ยอมรับ session ที่รันยาวและเปราะบาง ดู
[user guide](../en/README.md) สำหรับ workflow ที่ผู้ใช้เห็น และ
[local-model edition development](../../docs/dev/local-llm-development.md) สำหรับ
implementation และ synchronization policy

> [!TIP]
> ใหม่ที่นี่? **[อ่าน user guide →](../en/README.md)** เพื่อทัวร์เต็มรูปแบบของฟีเจอร์ ขั้นตอนติดตั้ง และเคล็ดลับ. ([ภาษาอื่น →](../README.md))

## Screenshots

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## How It Fits Together

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

- **pi** เขียน conversation JSONL ไปที่ `~/.pi/agent/sessions/` ขณะทำงาน
- **pi-web** เป็น Go server ที่อ่านไฟล์เหล่านั้น แสดงใน browser และส่ง live updates แบบ streaming ผ่าน SSE
- **pi --mode rpc** workers จัดการ chat ที่ browser เริ่ม — sessionละหนึ่ง แล้วถูกเก็บ (reaped) หลัง idle 10 นาที
- **fsnotify** ตรวจจับ directory sessions เพื่อที่ browser จะ reload ภายใน milliseconds หลังมี output ใหม่
- **Tailscale Serve** เผยแพร่ localhost server เป็น endpoint HTTPS บน tailnet ของคุณ

## Install

```bash
pi install npm:@timmygod/pi-web-local
```

แค่นี้เอง — มันดาวน์โหลด binary ที่ตรงกัน ตั้งค่า auto-start และ register คำสั่ง `/web`, `/pi-web`, `/remote`, และ `/refresh`

หลังติดตั้ง เปิด `http://127.0.0.1:31415` ใน browser จาก pi ใช้ `/web` เพื่อเปิด session ปัจจุบันใน browser ทันที ถ้า Tailscale รันอยู่บนเครื่องของคุณ pi-web จะเผยแพร่ endpoint HTTPS บน tailnet ของคุณอัตโนมัติ — ใช้ `/remote` จาก pi เพื่อดู QR code และ URL สำหรับทุกอุปกรณ์บน tailnet ของคุณ

> **การเข้าถึงระยะไกลบน macOS:** ติดตั้งและเปิด Tailscale แบบ interactive ยืนยัน administrator prompt และ sign in จากนั้นรัน `/pi-web restart` ตามด้วย `/remote`

สำหรับการติดตั้ง manual binary downloads หรือการ build จาก source ดู [user-docs/install.md](../en/install.md)

## Pi Integration

หลัง `pi install npm:@timmygod/pi-web-local` คุณจะได้:

| Command | What it does |
|---------|--------------|
| `/web` | เปิด session ปัจจุบันใน browser (รองรับ SSH: ข้าม browser และแสดง URL เท่านั้น) |
| `/pi-web` | แสดงสถานะ version เริ่มต้น/หยุด/รีสตาร์ท server หรือ update |
| `/remote` | แสดง QR code และ URL สำหรับการเข้าถึงระยะไกลผ่าน Tailscale |
| `/refresh` | ดึง messages ใหม่ที่เขียนจาก remote browsers กลับเข้า terminal session |

**auto-titling** ของ session ถูก built-in ในตัว pi-web และกำหนดค่าในหน้า `/settings` มัน **เปิดโดย default** และตั้งชื่อ session อัตโนมัติ คุณสามารถเลือก:

- **When to title** — ครั้งเดียวต่อ session หรือทุก message ใหม่ (default)
- **Title model** — **built-in word heuristic (ไม่ใช้ AI)** ฟรีและ immediate โดย default หรือเลือกโมเดล (เช่น โมเดลเล็ก/เร็ว) สำหรับชื่อที่ฉลาดขึ้นเขียนโดยโมเดล

package นี้ยังติดตั้ง binary pi-web ที่ `~/.pi/agent/bin/pi-web` และตั้งค่า auto-start ตอน login

## Auto-Start on Login

คำสั่ง `pi install npm:@timmygod/pi-web-local` ตั้งค่านี้ให้อัตโนมัติ:

| OS | Mechanism |
|----|-----------|
| macOS | launchd plist ที่ `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service ที่ `~/.config/systemd/user/pi-web.service` |
| Windows | `HKCU` Run-key entry ที่ запуска hidden starter ใน `~/.config/pi-web/` |

เพื่อตั้ง token สำหรับการเข้าถึงระยะไกล สร้าง `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

สำหรับรายละเอียดเพิ่มเติม (manual setup, custom ports, non-loopback binds) ดู [user-docs/install.md](../en/install.md)

## Development

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

สำหรับ upstream synchronization local-model testing และ parallel release
workflow ดู [Local-model edition development](../../docs/dev/local-llm-development.md)
