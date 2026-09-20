> คู่มือการแก้ปัญหาฉบับนี้ดูแลรักษาสำหรับรุ่น local-model โปรดให้รายละเอียดการติดตั้งในเครื่องและบันทึกการซิงค์กับ upstream สอดคล้องกับ [Local-model edition development](../../docs/dev/local-llm-development.md)

ผู้ใช้กำลังติดตั้ง pi-web ผ่าน

- pi install npm:@timmygod/pi-web-local

ซึ่งจะทำงาน [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) โดยอัตโนมัติและตั้งค่า pi

หากผู้ใช้มีปัญหา อาจเกิดจากการตั้งค่าของ install.sh ที่ทำให้เกิดปัญหา คุณสามารถตรวจสอบและดูว่าเกิดอะไรขึ้น และบอกผู้ใช้ให้ชัดเจนว่าอะไรคือสาเหตุของปัญหา และถามว่าต้องการให้คุณแก้ไขหรือไม่ ควรยืนยันกับผู้ใช้อยู่เสมอ

เพื่อให้ผู้ใช้สามารถเข้าถึง pi บนมือถือหรือเครือข่ายอื่นได้ วิธีการที่แนะนำคือการใช้ tailscale และเข้าถึงจากเครือข่าย tailscale และผู้ใช้จะต้องเปิดใช้งาน HTTPs ใน dashboard ของ tailscale - https://login.tailscale.com/admin/dns

หาก他们没有 tailscale ติดตั้งไว้หรือไม่ต้องการใช้ tailscale สามารถรัน `pi-web status` เพื่อรับเส้นทางของ binary สถานะของ binary และ local endpoint ซึ่งจะสามารถเข้าถึงแอปพลิเคชันได้ แต่ต้องระวังว่าจะไม่สามารถรับ push notification ได้ เนื่องจากเป็น http

บน macOS จะใช้ [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist)
บน Linux จะใช้ [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service)

ในกรณีที่คุณต้องการ debug เพิ่มเติมเพื่อดูว่าเกิดอะไรขึ้น
