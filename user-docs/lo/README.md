# ยินดีต้อนรับสู่ pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · **ລາວ**

</div>

**กำลังคิดจะลองใช้ pi-web ไหม? ลองเลย — คุณจะต้องหลงรักมัน**

pi-web คือ web UI และ PWA ที่สวยงามสำหรับ [pi](https://pi.dev) — AI coding agent แบบโอเพนซอร์ส มันช่วยให้คุณเรียกดู อ่าน และดำเนินเซสชัน pi ของคุณต่อได้จากเบราว์เซอร์ใดก็ได้ บนอุปกรณ์ใดก็ได้ พร้อมฟีเจอร์ที่คิดมาอย่างดีในทุกจุด

## อะไรที่แตกต่างในเวอร์ชันนี้?

รีโพสิทอรีนี้ยังคงรักษาอินเทอร์เฟซ pi-web ที่มาจาก upstream และฟีเจอร์ที่ใช้ร่วมกันไว้ แต่
เปลี่ยนวิธีปกป้องเซสชันเมื่อโมเดลที่เลือกทำงานในระดับ local หรือบน LAN ของคุณ

- **เลือกนโยบาย runtime ต่อเซสชัน.** ตรวจจับ endpoint ที่ local/LAN โดยอัตโนมัติเมื่อ
  provider metadata ชัดเจน; Local และ Cloud เป็นการบังคับแบบ manual ที่เก็บถาวร
- **ป้องกันความล้มเหลวของ context ก่อนหน้า.** Local Mode จะ compact ที่ usage 65% และตรวจ
  เสริมอีกทีระหว่าง tool calls ก่อนส่ง request ไปยังโมเดลครั้งถัดไป
- **จำกัดขอบเขตของสรุป.** Rolling checkpoints กันไม่ให้สรุปเก่าโตขึ้นเรื่อย ๆ มาลองใหม่
  ด้วย budget ที่กระชับขึ้น และหยุดอย่างปลอดภัยเมื่อการ compact ไม่ได้ทำให้ก้าวหน้าอย่าง
  มีนัยสำคัญ
- **กู้คืนแบบระมัดระวัง.** context overflow, interruption ของ transport ที่เลือก และ
  การหยุดก่อนเวลาที่เกิดจากการ reasoning อย่างเดียว สามารถ resume ได้อัตโนมัติ แต่ incident
  deduplication และ progress-aware circuit breakers ป้องกันการกู้คืนแบบวนลูป
- **ให้ผู้ใช้เป็นผู้ควบคุม.** Force Compact คือทางช่วยฉุกเฉินแบบ manual ที่มองเห็นได้เสมอ
  ขณะที่ Cloud Mode ยังคง workflow และตัวควบคุมแบบ upstream ไว้

ผลลัพธ์เชิงปฏิบัติคือเรียบง่าย: งาน local-model ที่ยาวควร compact ก่อนที่จะล้ม,
กู้คืนหนึ่งครั้งเมื่อการกู้คืนปลอดภัย และหยุดอย่างสะอาดเมื่อไม่ปลอดภัยแทนที่จะวนลูป

**pi-web สร้างมาสำหรับคนสองกลุ่ม:**

- 🧑‍💻 **สำหรับนักพัฒนา** — คนที่ใช้ชีวิตอยู่ใน terminal แต่ต้องการดำเนินเซสชันต่อจากมือถือ, ส่งงานต่อไปที่ remote server, หรือ monitor งานที่ใช้เวลานานจากทุกที่
- ✨ **สำหรับคนที่ไม่ได้เป็นนักพัฒนา** — คนที่ต้องการเพียงแอป AI ที่สวยงามซึ่งใช้งานได้ดี เปิดมาใช้ พิมพ์ vibe ไปได้เลย ไม่ต้องใช้ terminal ไม่ต้อง SSH ไม่ต้องสับสน เหมือนเครื่องมือ AI ที่ใช้งานง่ายที่สุด แต่พร้อมทางเลือกของโมเดลและอิสระแบบโอเพนซอร์ส

---

## ทำไม pi-web?

คุณกำลังอยู่ใน flow กับ pi ใน terminal อยู่แล้ว pi-web ให้ความต่อเนื่องนั้นต่อไปเมื่อคุณต้องลุกจากโต๊ะทำงาน:

- **Resume จากทุกที่** — ดำเนินเซสชันต่อจากมือถือ แท็บเล็ต หรือคอมพิวเตอร์เครื่องอื่น ไม่ต้อง SSH ไม่ต้อง Termius — แค่เปิดเบราว์เซอร์
- **Multi-session dashboard** — เริ่มงานในเซสชันหนึ่งขณะดูอีกเซสชันหนึ่งที่กำลัง stream ค้นหาข้ามโปรเจกต์ กรองตาม branch หาสิ่งที่ต้องการได้เร็ว
- **พื้นฐานแบบโอเพนซอร์ส** — pi เป็นโอเพนซอร์สทั้งหมดและ provider-agnostic คุณไม่ถูกผูกติดกับโมเดลหรือ vendor เดียวเดียว pi-web ก็เป็นโอเพนซอร์สเช่นกัน
- **การเข้าถึงระยะไกลที่ปลอดภัย** — token auth ในตัว ทำให้คุณเปิดใช้บน LAN หรือ Tailscale ของคุณได้โดยไม่ต้องกังวล
- **แชร์งานของคุณ** — export เซสชันเป็น static snapshots หรือ secret GitHub Gists ในคลิกเดียว

> อยากรู้เรื่องราวเบื้องหลังไหม? [อ่านเหตุผลที่เราสร้าง →](why.md)

---

## pi-web ในฐานะ AI workspace ส่วนตัวของคุณ 🏠

pi-web เป็น PWA (Progressive Web App) ดังนั้นคุณสามารถ **ติดตั้งเหมือนแอป native** บน desktop, laptop, มือถือ หรือแท็บเล็ต — ไม่ต้องใช้ app store บน desktop มันจะเปิดเป็นหน้าต่างของตัวเองโดยไม่มี browser chrome ทำให้มันดูและรู้สึกเหมือนแอป desktop ของจริง

นึกภาพมันในฐานะ **Claude Cowork ของคุณเอง** — AI workspace ส่วนตัวที่อยู่บนเครื่องของคุณ — ยกเว้นมันคือโอเพนซอร์สและ model-agnostic:

- **คุณเป็นเจ้าของ stack.** เลือกโมเดลใดก็ได้ สลับเมื่อไหร่ก็ได้ รันแบบ local แล้วข้อมูลของคุณจะไม่ออกจากเครื่องเลย
- **คนที่ไม่มีความรู้ด้านเทคนิคก็ใช้ได้.** ติดตั้ง pi-web บนเครื่องพวกเขา สอนวิธีใช้สักครั้งเดียว แล้วพวกเขาพร้อมไปต่อ พ่อแม่ คู่ชีวิต เพื่อนที่ไม่ถนัดเทคโนโลยี — ไม่ต้อง terminal ไม่ต้อง SSH แค่หน้า chat ที่คุ้นเคย
- **ติดตั้งครั้งเดียว ใช้ได้หลายคน.** ติดตั้งบน desktop ของคุณและแชร์หน้าจอ หรือเปิดใช้บนเครือข่ายที่บ้าน แล้วให้สมาชิกในครอบครัวเปิดจากอุปกรณ์ของพวกเขาเอง

ต้องการมากกว่าการเขียนโค้ดไหม? เปลี่ยนมันเป็น [personal assistant](personal-assistant.md) เฉพาะทางที่รู้จักว่าคุณเป็นใครและอยู่บนเครื่องของคุณ — เหมือน OpenClaw หรือ Hermes ของคุณเอง

> 💡 **เคล็ดลับ:** ติดตั้ง pi-web เป็น PWA จาก Chrome/Edge (คลิกไอคอน install ใน address bar) หรือ Safari (Share → Add to Dock) มันจะแยกไม่ออกระหว่างแอป native

---

## สิ่งที่ pi-web ทำได้

| | |
|---|---|
| 📱 **PWA** | ติดตั้ง pi-web เป็น Progressive Web App บน desktop, มือถือ หรือแท็บเล็ต เพื่อความรู้สึกแบบ native |
| 🔄 **ดำเนินเซสชันต่อ** | กลับมารับบทบทสนทนาใดก็ได้ตรงที่คุณเคยหยุด — ข้อความ, รูปภาพ, การสลับโมเดล ทั้งหมดจากเบราว์เซอร์ |
| 🆕 **สร้างเซสชันใหม่** | สร้างเซสชันใหม่กับ project path ใดก็ได้ ตรงจาก web UI |
| 📡 **Live streaming** | ดูคำตอบของ pi stream แบบ real-time ด้วย latency ~ms. โหมด Follow ช่วยให้เกาะอยู่ที่ล่าสุด |
| 🌲 **Tree view** | สำรวจ message tree แบบ native ของ pi — ดูโครงสร้างบทสนทนาทั้งหมด กระโดดไปยัง branch ใดก็ได้ และ fork จากจุดใดก็ได้ |
| 🔀 **Fork เซสชัน** | Fork เซสชันจากข้อความใด ๆ หรือแม้แต่ tool call เฉพาะเจาะจง — สำรวจทิศทางต่าง ๆ โดยไม่เสียตำแหน่งของคุณ |
| 🔍 **เรียกดู & ค้นหา** | กรองเซสชันข้ามโปรเจกต์ ค้นหาตามชื่อ สำรวจ branch — history ทั้งหมดของเซสชันในพริบตา |
| 🌿 **Git integration** | ดู branch ปัจจุบันและเปิด GitHub PR ได้จาก session viewer ตรง ๆ |
| 📝 **Scratchpad** | จดโน้ต, todos หรือความคิดสั้น ๆ ข้าง ๆ เซสชันของคุณโดยไม่ต้องสลับแอป |
| 💬 **Annotations** | ไฮไลต์และคอมเมนต์ในจุดใด ๆ ของเซสชัน — ดีสำหรับการ review โค้ด ให้ feedback หรือ bookmark จุดสำคัญ |
| 🎨 **Themes & customization** | สลับระหว่าง dark และ light mode ปรับ UI ตามความชอบ — ทำให้ pi-web รู้สึกเป็น *ของคุณ* |
| 🌐 **Multi-language** | ภาษา built-in 14 ภาษา (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). เพิ่มภาษา custom ของคุณเองจาก Settings |
| 🐱 **Wellness & pomodoro** | vibe coding มากเกินไปไม่ดีต่อสุขภาพ ตัวจับเวลา pomodoro ในตัวพร้อมเพื่อนคู่ใจลูกแมวและคำเตือนการนอน เพื่อรักษาสมดุลของคุณ |
| 📤 **Share & export** | ดาวน์โหลด JSONL, export static snapshots ที่ render ด้วยดีไซน์ `pi.dev` แบบ native ของ pi, หรือแชร์เป็น private GitHub Gists — ทั้งหมด render บน client-side |
| 🔔 **เสียงแจ้งเตือน** | เสียงแจ้งเตือนที่ปรับแต่งได้สำหรับเหตุการณ์ของเซสชัน — อยู่ใน loop แม้ pi-web จะอยู่ใน tab อื่น |
| ⌨️ **Keyboard shortcuts** | การนำทางแบบ Vim, action เร็ว — [คู่มือครบ →](keyboard-shortcuts.md) |
| 🤖 **Personal assistant** | เปลี่ยน pi-web เป็น AI assistant ของคุณเองที่อยู่บนคอมพิวเตอร์ของคุณ — เช่น OpenClaw หรือ Hermes. [ติดตั้ง →](personal-assistant.md) |
| 🗓️ **พูดกับ schedules** | จากเซสชัน pi บอกว่า "เพิ่ม schedule เวลา 2 ทุ่มตามเวลาสิงคโปร์เพื่อ …" — `/skill:pi-web-schedule` |
| 📝 **พูดกับ notes & settings** | "เขียนสิ่งนี้ใน notes" (`/skill:pi-web-notes`) หรือ "สลับเป็น dark mode" (`/skill:pi-web-settings`) |

---

## การนำทางอย่างรวดเร็ว

| หากคุณกำลังหา… | อ่าน |
|---|---|
| วิธีติดตั้ง ปรับ setting และใช้งาน pi-web | [install.md](install.md) |
| ใช้ pi-web เป็น personal assistant | [personal-assistant.md](personal-assistant.md) |
| คู่มือ keyboard shortcuts | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| ทำไม pi-web ถึงมีอยู่ | [why.md](why.md) |
| อะไรจะตามมาต่อ | [roadmap.md](roadmap.md) |
| ติดตั้งมีปัญหาไหม? ให้ LLM ของคุณแก้ — คัดลอก link llm-debug.md ให้มัน | [llm-debug.md](llm-debug.md) |
| การดูแลรักษาวีรชัน local-model นี้ | [development notes](../../docs/dev/local-llm-development.md) |

---

## Screenshots

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

pi-web สร้างขึ้นด้วยความรักและช่วงเวลากลางคืนมากมาย ผมจ่ายค่า coding plans (Claude Code, OpenCode, etc.) จากกระเป๋าตัวเองเพื่อเดินหน้าโปรเจกต์นี้ หากคุณเห็นว่า pi-web เป็นประโยชน์ การสนับสนุนของคุณจะมีความหมายมาก

**วิธีช่วยเหลือ:**

- 💰 **[Sponsor บน GitHub](https://github.com/sponsors/setkyar)** — ช่วยหนุนค่าใช้จ่ายของเครื่องมือที่ทำให้สิ่งนี้เกิดขึ้น
- ☕ **[ซื้อกาแฟให้](https://buymeacoffee.com/setkyar)** — ทุกน้อยนิดก็ช่วยได้
- ⭐ **กด Star repo** — ไม่เสียอะไรเลย และช่วยให้คนค้นพบ pi-web มากขึ้น
- 📢 **แชร์ให้เพื่อน & ครอบครัว** — ถ้าคุณรู้จักใครที่คงจะชอบ pi-web ส่งต่อไปให้

Sponsor ไม่ได้? ไม่มีปัญหาเลย — star และ share จะช่วยได้มาก ๆ ขอบคุณที่แวะมา 🙏

---

ขอให้สนุกกับการเขียนโค้ด! 🚀
