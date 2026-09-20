# ยินดีต้อนรับสู่ pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · **ไทย** · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**กำลังคิดว่าจะลองใช้ pi-web อยู่หรือเปล่า? ลองเลย — คุณจะหลงรักมัน**

pi-web คือ UI บนเว็บและ PWA ที่สวยงามสำหรับ [pi](https://pi.dev) — AI coding agent แบบ open-source ช่วยให้คุณสามารถดู อ่าน และต่อยอดเซสชัน pi ของคุณได้จากเบราว์เซอร์ใด ๆ บนอุปกรณ์ใด ๆ พร้อมฟีเจอร์ที่คิดมาอย่างดีในทุกขั้นตอน

## อะไรที่ต่างออกไปในเวอร์ชันนี้?

รีโพซิทอรีนี้ยังคงอินเทอร์เฟซและฟีเจอร์ร่วมกันจาก pi-web ดั้งเดิม (upstream) ไว้ แต่
เปลี่ยนวิธีการปกป้องเซสชันเมื่อโมเดลที่เลือกทำงานบนเครื่อง local หรือบน LAN ของคุณ

- **เลือก runtime policy แบบรายเซสชัน** ระบบจะตรวจจับ endpoint ที่เป็น local/LAN
  อัตโนมัติเมื่อนี้ metadata ของ provider ชัดเจน; ส่วน Local และ Cloud คือการ overrides
  แบบ manual ที่จะถูกบันทึกไว้
- **ป้องกัน context failures ตั้งแต่เนิ่นๆ** Local Mode จะ compact ที่ระดับการใช้ 65% และตรวจ
  อีกครั้งระหว่าง tool calls ก่อนการส่ง request ไปยังโมเดลครั้งถัดไป
- **จำกัดขนาดของ summaries** rolling checkpoints ป้องกันไม่ให้ summary เก่าโตไปเรื่อย ๆ,
  retry หนึ่งครั้งด้วย budget ที่กระชับขึ้น และหยุดอย่างปลอดภัยเมื่อการ compact ไม่สร้างความก้าวหน้า
  ที่ชัดเจน
- **กู้คืนแบบระมัดระวัง** context overflow, การหยุดชะงักของ transport ที่เลือกไว้, และการหยุด
  กลางคันที่เกิดจากการ reason-only สามารถ resume ได้อัตโนมัติ แต่การ dedupe incident และ
  circuit breakers ที่คำนึงถึงความก้าวหน้าจะป้องกันไม่ให้ recovery วนซ้ำ
- **让用户เป็นผู้ควบคุม** Force Compact คือเส้นทางกู้ภัยแบบ manual ที่มองเห็นได้เสมอ
  ขณะเดียวกัน Cloud Mode ยังคง workflow และ controls จาก upstream ไว้

ผลลัพธ์ในทางปฏิบัติเรียบง่าย: งาน task ยาว ๆ ของ local-model ควร compact ก่อนที่จะล้ม,
recover หนึ่งครั้งเมื่อปลอดภัย และหยุดอย่างสะอาดแทนที่จะวนซ้ำเมื่อไม่ปลอดภัย

**pi-web สร้างขึ้นสำหรับคนสองแบบ:**

- 🧑‍💻 **สำหรับนักพัฒนา** — ผู้ที่ทำงานอยู่ใน terminal แต่ต้องการต่อเซสชันจากมือถือ, ส่งมอบงานไปยัง server ไกล หรือมอนิเตอร์งานที่รันนานจากที่ไหนก็ได้
- ✨ **สำหรับคนที่ไม่ได้เป็นนักพัฒนา** — ผู้ที่แค่อยากได้ AI app ที่สวยงามและใช้งานได้ เปิดมา พิมพ์ vibe ได้เลย ไม่ต้องใช้ terminal ไม่ต้องใช้ SSH ไม่สับสน เหมือนเครื่องมือ AI ที่ user-friendly ที่สุด แต่มีตัวเลือกโมเดลและความเสรีแบบ open-source

---

## ทำไมต้องมี pi-web?

คุณกำลังอยู่ใน flow กับ pi ใน terminal อยู่แล้ว pi-web จะรักษา momentum นั้นไว้เมื่อคุณลุกจากโต๊ะทำงาน:

- **Resume จากที่ไหนก็ได้** — ต่อเซสชันจากมือถือ แท็บเล็ต หรือคอมพิวเตอร์เครื่องอื่น ไม่ต้องใช้ SSH ไม่ต้องใช้ Termius — แค่เปิดเบราว์เซอร์
- **Multi-session dashboard** — เริ่มงานในเซสชันหนึ่งขณะดูอีกเซสชันที่กำลัง stream ค้นหาข้ามโปรเจกต์ กรองตาม branch หาสิ่งที่คุณต้องการได้เร็ว
- **ฐานแบบ open-source** — pi เป็น open-source เต็มรูปแบบและ provider-agnostic คุณไม่ถูกผูกมัดกับโมเดลหรือ vendor เดียว pi-web ก็เป็น open-source เช่นกัน
- **เข้าถึงระยะไกลอย่างปลอดภัย** — มี token auth ในตัว จึงสามารถ expose บน LAN หรือ Tailscale ของคุณได้โดยไม่ต้องกังวล
- **แบ่งปันงานของคุณ** — export เซสชันเป็น static snapshots หรือ GitHub Gists แบบ secret ด้วยคลิกเดียว

> สงสัยเรื่องที่มาที่ไป? [อ่านว่าทำไมเราถึงสร้างมัน →](why.md)

---

## pi-web ในฐานะ personal AI workspace ของคุณ 🏠

pi-web เป็น PWA (Progressive Web App) ดังนั้นคุณจึงสามารถ **ติดตั้งมันเหมือนแอป native** บน desktop, laptop, มือถือ หรือแท็บเล็ต — ไม่ต้องไปแอปสโตร์ บน desktop จะเปิดใน window ของตัวเองโดยไม่มี browser chrome ทำให้มีทั้งหน้าตาและความรู้สึกเหมือนแอป desktop จริง ๆ

คิดภาพว่ามันคือ **Claude Cowork ของตัวคุณเอง** — personal AI workspace ที่อยู่บนเครื่องของคุณ — ยกเว้นว่ามันเป็น open-source และ model-agnostic:

- **คุณเป็นเจ้าของ stack นี้** เลือกโมเดลใดก็ได้ เปลี่ยนเมื่อไหร่ก็ได้ รันแบบ local และข้อมูลของคุณจะไม่ออกจากเครื่อง
- **คนที่ไม่ได้มีสกิลด้านเทคนิคก็ใช้ได้** ติดตั้ง pi-web บนเครื่องพวกเขา แสดงวิธีใช้ให้ดูครั้งเดียว แล้วพวกเขาก็ใช้ได้เลย พ่อแม่ คู่ชีวิต เพื่อนที่ไม่ได้สาย tech — ไม่ต้องใช้ terminal ไม่ต้องใช้ SSH แค่ chat interface ที่คุ้นเคย
- **ติดตั้งครั้งเดียว ใช้งานหลายคน** ติดตั้งบน desktop แล้ว share screen หรือ expose บนเครือข่ายที่บ้านเพื่อให้สมาชิกในครอบครัวเปิดบนอุปกรณ์ของตัวเองได้

ต้องการมากกว่า coding ไหม? ทำให้มันเป็น [personal assistant](personal-assistant.md) ที่ ded icated ที่รู้จักว่าคุณเป็นใครและอยู่บนเครื่องของคุณ — เหมือน OpenClaw หรือ Hermes ของตัวคุณเอง

> 💡 **Pro tip:** ติดตั้ง pi-web เป็น PWA จาก Chrome/Edge (คลิกไอคอนติดตั้งในแถบ address) หรือ Safari (Share → Add to Dock) มันจะแยกไม่ออกจากแอป native

---

## สิ่งที่คุณทำได้ที่ pi-web

| | |
|---|---|
| 📱 **PWA** | ติดตั้ง pi-web เป็น Progressive Web App บน desktop, มือถือ หรือแท็บเล็ต เพื่อความรู้สึกแบบ native |
| 🔄 **Continue sessions** | ต่อบทสนทนาใด ๆ จากจุดที่ค้างไว้ — ข้อความ รูปภาพ การสลับโมเดล ทั้งหมดจากเบราว์เซอร์ |
| 🆕 **Start new sessions** | สร้างเซสชันใหม่กับ path โปรเจกต์ใดก็ได้ ตรงจาก web UI |
| 📡 **Live streaming** | ดู response ของ pi streaming แบบ real time ด้วย latency ~ms Follow mode จะพาคุณติดตามตัวล่าสุด |
| 🌲 **Tree view** | เยี่ยมชม message tree แบบ native ของ pi — ดูโครงสร้างบทสนทนาทั้งหมด กระโดดไปยัง branch ใด ๆ และ fork จากจุดใดก็ได้ |
| 🔀 **Fork sessions** | Fork เซสชันจากข้อความใด ๆ หรือแม้แต่จาก tool call เฉพาะเจาะจง — สำรวจทิศทางต่าง ๆ โดยไม่เสียจุดที่ค้างไว้ |
| 🔍 **Browse & search** | กรองเซสชันข้ามโปรเจกต์ ค้นหาตามชื่อ เยี่ยมชม branch — ประวัติเซสชันทั้งหมดของคุณในรวดเดียว |
| 🌿 **Git integration** | ดู branch ปัจจุบันและเปิด GitHub PR ได้โดยตรงจาก session viewer |
| 📝 **Scratchpad** | จดบันทึก todos หรือไอเดียอย่างรวดเร็วข้าง ๆ เซสชันของคุณโดยไม่ต้องสลับแอป |
| 💬 **Annotations** | ไฮไลต์และคอมเมนต์ในส่วนใดของเซสชันก็ได้ — เหมาะกับ code review, feedback หรือ bookmark ช่วงสำคัญ |
| 🎨 **Themes & customization** | สลับระหว่าง dark และ light mode ปรับ UI ให้ถูกใจ — ทำให้ pi-web รู้สึกเป็นของ *คุณ* |
| 🌐 **Multi-language** | มี 14 ภาษาในตัว (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ) เพิ่มภาษา custom ของคุณเองจาก Settings |
| 🐱 **Wellness & pomodoro** | vibe coding มากเกินไปไม่เป็นไรต่อสุขภาพ มี pomodoro timer ในตัว พร้อมแมวคู่ใจและการเตือนให้พักผ่อน เพื่อสมดุลของคุณ |
| 📤 **Share & export** | โหลด JSONL, export static snapshots ที่ render ด้วย look แบบ native ของ `pi.dev`, หรือแชร์เป็น GitHub Gists แบบ private — ทั้งหมด render ที่ฝั่ง client |
| 🔔 **Notification sounds** | เสียง notification ปรับแต่งได้สำหรับ session events — ติดตามสถานะแม้ pi-web จะอยู่ที่แท็บอื่น |
| ⌨️ **Keyboard shortcuts** | การนำทางแบบ Vim, quick actions — [คู่มือเต็ม →](keyboard-shortcuts.md) |
| 🤖 **Personal assistant** | เปลี่ยน pi-web ให้เป็น AI assistant ของคุณเองที่อยู่ที่เครื่องของคุณ — เหมือน OpenClaw หรือ Hermes. [ตั้งค่า →](personal-assistant.md) |
| 🗓️ **Talk to schedules** | จากเซสชัน pi พูดว่า "เพิ่ม schedule เวลา 2 ทุ่มเวลาสิงคโปร์ไปที่ …" — `/skill:pi-web-schedule`. |
| 📝 **Talk to notes & settings** | "เขียนนี่ลง notes" (`/skill:pi-web-notes`) หรือ "สลับเป็น dark mode" (`/skill:pi-web-settings`). |

---

## นำทางแบบย่อ

| หากคุณกำลังมองหา… | อ่าน |
|---|---|
| วิธีติดตั้ง configure และใช้ pi-web | [install.md](install.md) |
| ใช้ pi-web เป็น personal assistant | [personal-assistant.md](personal-assistant.md) |
| คู่มือ keyboard shortcuts | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| ทำไม pi-web จึงมีอยู่ | [why.md](why.md) |
| สิ่งที่จะมาในลำดับถัดไป | [roadmap.md](roadmap.md) |
| มีปัญหาการติดตั้ง? ให้ LLM ของคุณแก้ไข — คัดลอกลิงก์ llm-debug.md ไปให้เขา | [llm-debug.md](llm-debug.md) |
| การดูแลรักษาเวอร์ชัน local-model นี้ | [development notes](../../docs/dev/local-llm-development.md) |

---

## ภาพหน้าจอ

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

pi-web สร้างขึ้นด้วยความรักและค่ำคืนที่ดึกมากมาย ผมจ่ายค่า coding plans (Claude Code, OpenCode, ฯลฯ) ด้วยเงินของตัวเองเพื่อให้โปรเจกต์นี้ไปต่อ หากคุณได้ประโยชน์จาก pi-web การสนับสนุนของคุณจะมีความหมายอย่างมาก

**วิธีช่วยเหลือ:**

- 💰 **[Sponsor on GitHub](https://github.com/sponsors/setkyar)** — ช่วยอุดหนุนเครื่องมือที่ทำให้สิ่งนี้เป็นไปได้
- ☕ **[Buy me a coffee](https://buymeacoffee.com/setkyar)** — ทุกความช่วยเหลือมีค่า
- ⭐ **Star reponี้** — ไม่เสียอะไรเลยและช่วยให้คนค้นพบ pi-web ได้มากขึ้น
- 📢 **แชร์ให้เพื่อนและครอบครัว** — หากคุณรู้จักใครที่想必จะรัก pi-web ส่งมันไปให้เขา

ไม่สามารถ sponsor ได้ไหม? ไม่ต้องกังวลเลย — star และ share ครั้งเดียวมีคุณค่ามาก ขอบคุณที่อยู่นี่ 🙏

---

มีความสุขกับการ code ครับ! 🚀
