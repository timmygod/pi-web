# ทำไมต้องมี pi-web?

ผมเสพติด Claude Code เป็นอย่างมาก ผมใช้มันอยู่เสมอ ถ้าผมไม่ได้นั่งอยู่หน้าคอมพิวเตอร์ ผมก็กำลังคิดถึงมันอยู่ ผมรู้สึกว่าผมไม่ได้เผาโทเคนให้มากพอ นั่นคือยุคแรก ๆ ของ Claude Code และผมก็คิดว่า ทำไมผมถึงไม่สามารถเชื่อมต่อจากโทรศัพท์ไม่ได้ ผมติดตั้ง Termius ขึ้นมาแต่ก็ไม่ค่อยชอบมันเท่าไหร่

ผมเริ่มสร้างของผมเอง แต่ต้องหยุดเมื่อ Claude เปิดตัวแอป Claude Code สำหรับมือถือ

จากนั้นผมเป็นหมอนรองกระดูกสันหลังยึดกระดูกสันหลังและผมแทบทำอะไรไม่ได้มาก เวลาผ่านไป ผมรู้สึกดีขึ้นบ้างและอยากกลับมาทำโปรเจกต์ Claude Code ผ่านเว็บ/PWA ของผมต่อ

จากนั้น Claude Code เริ่มแบนการใช้งานที่อยู่นอก harness ของพวกเขา ผมรู้สึกว่าไม่คุ้มค่าอีกต่อไป

จากนั้นผมก็เจอ pi.dev และสำรวจไปบ้างแต่ไม่ได้ลงลึก ผมอ่านเกี่ยวกับมัน ดูวิดีโอเกี่ยวกับมัน และตัดสินใจลองอย่างจริงจัง และตอนนี้ผมก็หลงรัก pi ไปเสียแล้ว

เนื่องจากมันเป็นโอเพนซอร์ส ผมรู้สึกว่าน่าลงมือสร้างเพื่อมัน ผมยังมีตัวเลือก provider อื่น ๆ ด้วย ผมยังรู้สึกว่า việcพึ่งพา provider/model เดียวกันเช่น Anthropic/Claude นั้นไม่ยั่งยืน

ดังนั้นผมจึงสร้างมันที่นี่

การ checkout นี้ถูกดูแลรักษาในฐานะ version local-model ของ pi-web มันติดตามโปรเจกต์ upstream สำหรับ shared improvements แต่ยังคง local deployment, context stability และ local-model testing ไว้ใน track ที่ปล่อยแยกต่างหาก

## ทำไม local model ถึงต้องมี operating profile ที่ต่างออกไป

ประสบการณ์เดิมของ pi-web เป็นพื้นฐานที่ยอดเยี่ยม แต่ local inference
มี failure modes ที่ต่างจาก hosted model ทั่วไป local model อาจชะลอ
ตัวลงอย่างมากเมื่อ context เพิ่มขึ้น แบ่งปัน memory จำกัดกับส่วนอื่น
ของเครื่อง หยุดลงหลังจากสร้างแค่ reasoning หรือสูญเสียการ run ที่ยาว
เนื่องจาก local transport failure ชั่วคราว การจัดการกรณีเหล่านั้น
เหมือนกับการ fail แบบ cloud ทำให้ UI ดูเข้ากันได้ แต่ session ที่แท้จริง
ยังคงเปราะบาง

version นี้เข้าหาคำปัญหานี้เป็นชั้น ๆ:

1. **รักษา upstream เป็นอันดับแรก.** Shared UI และ session behavior ยังคง
   มาจาก pi-web; local changes ถูกแยกอยู่ภายใต้ effective Local Mode.
2. **ป้องกันก่อนกู้คืน.** รอยต่อ context แบบเปอร์เซ็นต์ 65% ถูกบังคับใช้
   ก่อน model call ถัดไป รวมถึง call ที่อยู่ภายใน long tool loops.
3. **กู้คืนเฉพาะเมื่อมีหลักฐาน.** การดำเนินการต่อแบบอัตโนมัติถูกจำกัดไว้
   เฉพาะเหตุการณ์ context, transport และ thinking-only ที่ถูกระบุ—not
   authentication, quota หรือ error จาก provider ทั่วไป
4. **กำหนดขอบเขตให้ทุกการดำเนินการอัตโนมัติ.** การกู้คืน incidents ถูก
   deduplicate ต้องมี progress ก่อนกู้คืนอีกครั้ง และการ start
   พิจารณา Local session ที่ active ล่าสุดอย่างมากไม่เกินหนึ่ง
5. **คงทางออกแบบ manual ไว้.** Force Compact สรุปผลแทนการล้าง history,
   ดังนั้นผู้ใช้สามารถกู้คืน session โดยไม่ต้องทำเป็นว่า context ไม่เคยมีอยู่
6. **ปกป้อง cloud compatibility.** Cloud Mode รักษา semantics และ
   controls แบบ upstream; local-model optimizations ไม่ได้ redefine cloud sessions โดยเงียบ

นั่นคือความแตกต่างที่แท้จริงใน fork นี้: มันมองว่า local inference เป็น
operational environment ที่ต่างออกไป ไม่ใช่แค่ชื่อ model อีกตัวใน dropdown
