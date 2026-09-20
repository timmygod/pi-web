<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · **ລາວ**

</div>

<div align="center">

ใช้ [pi](https://pi.dev) coding agent ของคุณได้ตั้งแต่โทรศัพท์, แท็บเล็ต, หรือแล็ปท็อป — ได้ทุกที่ในเครือข่ายของคุณ หรือทางไกลผ่าน Tailscale

มันเป็น PWA เต็มรูปแบบ ดังนั้นคุณจึงติดตั้งและใช้งานได้เหมือนแอปเนทีฟบนอุปกรณ์ใดก็ได้ ลองนึกภาพว่าเป็น AI workspace ส่วนตัวของคุณ — คล้าย Cowork ของ Claude แต่ใช้โมเดลอื่นได้ — แชทข้ามโมเดล, เขียนโค้ดจากโทรศัพท์, หรือเปลี่ยนมันเป็น [ผู้ช่วยส่วนตัว](../en/personal-assistant.md) ที่ทำงานอยู่บนเครื่องของคุณ

ทำให้เป็นของคุณ: สลับธีมและฟอนต์ และใช้งานในภาษาของคุณเอง — pi-web รองรับหลายภาษา และคุณสามารถเพิ่มภาษาของคุณเองได้ มีฟีเจอร์อื่น ๆ อีกมากมายที่กำลังจะมาถึง แต่จะไม่臃肿 (ไม่เทอะทะ): สิ่งที่คุณไม่ต้องการสามารถปิดได้ในหน้าตั้งค่า

</div>

## ทำไมต้องเป็น local-model edition นี้?

pi-web ต้นฉบับยังคงเป็นฐานราก upstream สำหรับฟีเจอร์และการแก้ไขร่วมกัน
edition นี้คงประสบการณ์นั้นไว้ จากนั้นเพิ่มชั้นความน่าเชื่อถือสำหรับโมเดลที่
ทำงานบนเครื่องของคุณเองหรือที่อื่นใน LAN ของคุณ — ซึ่งการประมวลผลมักช้ากว่า
หน่วยความจำมีจำกัด และ context ยาว ๆ อาจทำให้เซสชันที่สุขภาพดีอยู่แล้วยังติดค้าง

| หัวข้อ | pi-web upstream | edition นี้ |
|------|-----------------|--------------|
| นโยบายโมเดล/runtime | พฤติกรรม pi-web มาตรฐาน | โหมด **Auto / Local / Cloud** ต่อเซสชัน พร้อมการตรวจตรา local ที่รู้ endpoint และการ override การทำงานด้วยมือแบบถาวร |
| การจัดการ long-context | พฤติกรรมการ compact ของ pi ปกติ | Local Mode จะ compact ล่วงหน้าเมื่อ **65%** และตรวจสอบอีกครั้งภายในลูป tool-call ยาว ๆ ก่อนคำขอ model ครั้งถัดไป |
| ความปลอดภัยของการ compact | สรุปแบบมาตรฐาน | rolling checkpoints แบบมีขอบเขต, การเขียนใหม่แบบเข้มงวดขึ้นหนึ่งครั้งสำหรับ output ที่ไม่ถูกต้อง/ถูกจำกัด และการตรวจจับว่าไม่มีความคืบหน้าแทนที่จะ re-compaction อย่างไม่สิ้นสุด |
| runs ที่ถูกขัดจังหวะ | การจัดการ worker และ error ปกติ | การกู้คืนแบบมีขอบเขตสำหรับ context overflow, การหยุด thinking-only และการรบกวนขนส่งที่เลือก พร้อมตัวทำลายลูปแบบถาวร |
| การช่วยด้วยมือ | รายละเอียด context มาตรฐาน | **Force Compact** ยังพร้อมใช้ในฐานะเส้นทางกู้คืนที่ชัดเจน โดยไม่ลบการสนทนา |
| ความเข้ากันได้และการเผยแพร่ | โปรเจกต์และสายการเผยแพร่ต้นฉบับ | มาตรการป้องกัน local-only ยังคงอยู่หลัง Local Mode; Cloud Mode คงพฤติกรรมการ upstream ไว้ และการเปลี่ยนแปลงของ upstream ได้รับการทบทวนและเผยแพร่ที่นี่แยกกัน |

นี่ไม่ใช่การเขียนใหม่หรือสิ่งแทนที่ upstream แต่เป็น operating profile ที่
ดูแลอย่างตั้งใจสำหรับผู้คนที่ต้องการความเป็นส่วนตัวและควบคุมของ local-model โดยไม่ต้องยอมรับเซสชันยาว ๆ ที่เปราะบาง ดู
[คู่มือผู้ใช้](../en/README.md) สำหรับ workflow ระดับผู้ใช้ และ
[การพัฒนา local-model edition](../../docs/dev/local-llm-development.md) สำหรับ
การนำไปใช้และนโยบายการซิงค์

> [!TIP]
> ใหม่ที่นี่? **[อ่านคู่มือผู้ใช้ →](../en/README.md)** สำหรับทัวร์เต็มรูปแบบของขั้นตอนการติดตั้ง และเคล็ดลับต่าง ๆ ([ภาษาอื่น →](../README.md))

## ภาพหน้าจอ

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>เดสก์ท็อป</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>มือถือ</em>
</div>

## ทุกอย่างประกอบกันอย่างไร

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

- **pi** เขียน JSONL ของบทสนทนาไปที่ `~/.pi/agent/sessions/` ระหว่างทำงาน
- **pi-web** เป็นเซิร์ฟเวอร์ Go ที่อ่านไฟล์เหล่านั้น แสดงผลในเบราว์เซอร์ และส่งอัปเดตแบบสดผ่าน SSE
- **pi --mode rpc** workers จัดการแชทที่เกิดจากเบราว์เซอร์ — เซสชันละหนึ่งตัว เก็บเมื่อเว้นว่างเกิน 10 นาที
- **fsnotify** เฝ้าดูโฟลเดอร์ sessions ทำให้เบราว์เซอร์โหลดใหม่ภายในไม่กี่มิลลิวินาทีเมื่อมี output ใหม่
- **Tailscale Serve** เผยแพร่เซิร์ฟเวอร์ localhost เป็น endpoint HTTPS บน tailnet ของคุณ

## ติดตั้ง

```bash
pi install npm:@timmygod/pi-web-local
```

เพียงเท่านี้ — จะดาวน์โหลด binary ที่ตรงกัน ตั้งค่า auto-start และลงทะเบียนคำสั่ง `/web`, `/pi-web`, `/remote` และ `/refresh`

เมื่อติดตั้งเสร็จ ให้เปิด `http://127.0.0.1:31415` ในเบราว์เซอร์ของคุณ จาก pi ให้ใช้ `/web` เพื่อบอกเซสชันปัจจุบันในเบราว์เซอร์ทันที หาก Tailscale กำลังทำงานบนเครื่องของคุณ pi-web จะเผยแพร่ endpoint HTTPS บน tailnet ของคุณอัตโนมัติ — ใช้ `/remote` จาก pi เพื่อดึง QR code และ URL สำหรับอุปกรณ์ใดก็ตามบน tailnet ของคุณ

> **การเข้าถึงระยะไกลบน macOS:** ติดตั้งและเปิด Tailscale แบบ interactive ยืนยัน prompt ของผู้ดูแลระบบ และเข้าสู่ระบบ จากนั้นรัน `/pi-web restart` ตามด้วย `/remote`

สำหรับการติดตั้งด้วยมือ, ดาวน์โหลด binary หรือสร้างจาก source ดู [user-docs/install.md](../en/install.md)

## Pi Integration

หลัง `pi install npm:@timmygod/pi-web-local` คุณจะได้:

| คำสั่ง | สิ่งที่ทำได้ |
|---------|--------------|
| `/web` | เปิดเซสชันปัจจุบันในเบราว์เซอร์ (รับรู้ SSH: ข้ามเบราว์เซอร์และแสดงเฉพาะ URL) |
| `/pi-web` | แสดงสถานะ, เวอร์ชัน, เริ่ม/หยุด/รีสตาร์ทเซิร์ฟเวอร์ หรืออัปเดต |
| `/remote` | แสดง QR code และ URL สำหรับการเข้าถึงระยะไกลผ่าน Tailscale |
| `/refresh` | ดึงข้อความใหม่ที่ถูกเขียนจากเบราว์เซอร์ระยะไกลกลับเข้าเซสชันในเทอร์มินัล |

**การตั้งชื่อเซสชันอัตโนมัติ** (auto-titling) ติดตั้งมากับตัว pi-web เอง และตั้งค่าได้ในหน้า `/settings` **เปิดโดยค่าเริ่มต้น** และตั้งชื่อเซสชันอัตโนมัติ คุณสามารถเลือกได้:

- **ตอนไหนตั้งชื่อ** — ต่อเซสชันครั้งเดียว หรือทุกข้อความใหม่ (ค่าเริ่มต้น)
- **โมเดลตั้งชื่อ** — **heuristic แบบสร้างในตัว (ไม่มี AI)** ฟรีและทันที โดยค่าเริ่มต้น หรือเลือกโมเดล (เช่น โมเดลเล็ก/เร็ว) เพื่อชื่อที่ชาญฉลาดขึ้นจากโมเดล

แพ็กเกจยังติดตั้ง binary ของ pi-web ไปที่ `~/.pi/agent/bin/pi-web` และตั้งค่า auto-start เมื่อเข้าสู่ระบบ

## Auto-Start เมื่อเข้าสู่ระบบ

คำสั่ง `pi install npm:@timmygod/pi-web-local` ตั้งค่านี้ให้โดยอัตโนมัติ:

| ระบบปฏิบัติการ | กลไก |
|----|-----------|
| macOS | launchd plist ที่ `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service ที่ `~/.config/systemd/user/pi-web.service` |
| Windows | `HKCU` Run-key entry เริ่มการเรียกสตาร์ทเตอร์แบบซ่อนใน `~/.config/pi-web/` |

เพื่อตั้งค่า token สำหรับการเข้าถึงระยะไกล ให้สร้าง `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

สำหรับรายละเอียดเพิ่มเติม (การตั้งค่าด้วยมือ, พอร์ตแบบ custom, การ bind แบบ non-loopback) ดู [user-docs/install.md](../en/install.md)

## การพัฒนา

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

สำหรับการซิงค์กับ upstream, การทดสอบ local-model และ workflow การเผยแพร่แบบคู่ขนาน ดู [Local-model edition development](../../docs/dev/local-llm-development.md)
