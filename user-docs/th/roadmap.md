# Roadmap

Roadmap นี้เป็นของบริษัท edition ของ local model ใน pi-web คุณสมบัติ upstream จะถูกซิงโครไนซ์เป็นระยะ ขณะที่งานด้าน local deployment และความน่าเชื่อถือจะถูกตรวจสอบและปล่อยปล่อยบนสายนี้ ดู [Local-model edition development](../../docs/dev/local-llm-development.md)

pi-web สร้างขึ้นสำหรับสองกลุ่มเป้าหมาย:

- **สำหรับนักพัฒนา** — ผู้ที่ใช้ชีวิตอยู่ใน terminal แต่ต้องการต่อเนื่อง session จากมือถือ ส่งต่อไปยัง remote server หรือจับตางานที่ทำงานนานจากที่ไหนก็ได้
- **สำหรับคนที่ไม่ได้เป็นนักพัฒนา** — ผู้ที่ต้องการแค่แอป AI ที่สวยงามและใช้งานง่าย เปิดขึ้น พิมพ์ vibe ไร้ terminal ไร้ SSH ไร้ความสับสน เหมือนเครื่องมือ AI ที่ใช้ง่ายที่สุด แต่พร้อมตัวเลือก model และเสรีภาพแบบ open-source

นี่คือสิ่งที่กำลังจะมาถึง

edition นี้ track upstream pi-web บนสาย release แยก คุณสมบัติ upstream จะถูก import เป็นระยะ งานด้านความน่าเชื่อถือของ local model จะถูกให้ความสำคัญและตรวจสอบที่นี่โดยไม่เปลี่ยนประวัติการ release ของ upstream

---

## Now (shipped)

ทุกอย่างที่ระบุใน [ตารางคุณสมบัติ](README.md#what-you-can-do-with-pi-web) ใช้งานได้แล้ววันนี้

---

## Next up

| # | Feature | What it does |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | แชทกับ pi ผ่าน Telegram หรือ Discord — เหมาะสมกับ workflow แบบ personal assistant ในยามเดินทาง |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | การติดตาม Token การประมาณการค่าใช้จ่าย และ session analytics — รู้ว่าคุณใช้ pi อย่างไร |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | ตั้งค่า visibility ที่คุณต้องการสำหรับ thinking, tools และ tool outputs ทั้งหมดข้าม session |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | ส่งคำสั่งเพิ่มเติมขณะที่ pi ยังทำงานอยู่ — นำทางกลางทาง |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | บีบอัดบทสนทนาที่ยาวตั้งแต่ใน web UI ไม่ต้องใช้ terminal |

---

## Planned

| # | Feature | What it does |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | ดูโครงสร้างไฟล์โปรเจกต์และดู git changes โดยตรงใน pi-web Opt-in เพื่อให้ไม่กวนการใช้งาน |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | ตั้งเวลา prompt ให้ทำงานอัตโนมัติ — daily standups, morning summaries และ recurring tasks ถูก gate ไว้โดย Admin เพื่อความปลอดภัย |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | ปรับ keyboard shortcuts ทั้งหมดให้เข้ากับ muscle memory ของคุณ |

---

## Vision

เป้าหมายระยะยาว: pi-web ควรเป็น **อินเทอร์เฟซสำหรับ pi** — สำหรับทุกคน

- **Non-devs** เปิดเหมือนแอปอื่นๆ เลือก model พิมพ์ จบ ไม่ต้องใช้ command line เลย
- **Devs** ได้การบูรณาการลึก — remote handoff, multi-session dashboards, git-aware browsing และ messaging bots
- **Everyone** ได้เสรีภาพในการเลือก model ความโปร่งใสแบบ open-source และ UI ที่รู้สึกเป็นมิตรในทุกลมหายใจ

---

> 💡 มีไอเดีย? [เปิด issue](https://github.com/timmygod/pi-web/issues/new) หรือเข้าร่วมการสนทนา
