# ເປັນຫຍັງ pi-web?

ຂ້ອຍມັກ Claude Code ຈົນເກືອບຢາເກົ່າ. ຂ້ອຍໃຊ້ມັນສະໝ່ຮ່ວມຢ່າງຕັດຂວາຍ. ເວົ້າມັນຢູ່ໜ້າຄອມພິວເຕີກໍ່ຄິດເຖິງມັນ. ຂ້ອຍຄິດວ່າຕົນເອງໃຊ້ tokens ບໍ່ພໍພໍ. ມັນຍັງເປັນມື້ຕົ້ນຕ້ອງໂອກຂອງ Claude Code. ແລະ ຂ້ອຍກໍ່ຄິດວ່າ ເປັນຫຍັງຂ້ອຍຈຶ່ງຕໍ່ເນື່ອງຈາກໂທລະສັບບໍ່ໄດ້? ຂ້ອຍຕັ້ງ Termius ແຕ່ກໍ່ບໍ່ໄດ້ມັກມັນຢ່າງຈຳເປັນ.

ຂ້ອຍເລີ່ມສ້າງຂອງຂ້ອຍເອງ ແຕ່ຕັດການອອກແບບຖ້າ Claude ໄດ້ນຳພາມື່ອ app ໂທລະສັບຂອງ Claude Code ຂອງເຂົາເຈົ້າ.

ຕໍ່ມາຂ້ອຍໄດ້ປວດທາງຖັງປິດເປັດ ແລະ ອອກຈາກສະພາບການບໍ່ໄດ້ເຮັດຫຼາຍປັດ. ລາຍປະຫວັນເປັນກະທ້າຍ ແລະ ຂ້ອຍສະທ້ອນວ່າຕົນເອງເບີ່ງມາ ແລະ ຕ້ອງການຕໍ່ສືບສາທິດຄືນຮອງໃຫ້ Claude Code ຜ່ານໂປຣເຈັກ web/pwa ຂອງຕົນເອງ.

ຕໍ່ມາ Claude Code ໄດ້ຢ້ານຂ້າມຄວາມປ່ອງຂອງຜູ້ໃຊ້ນອກມາດຕາມການ of their own harness. ແລະ ຂ້ອຍຄິດວ່າມັນບໍ່ສະທ້ອນມີຈຳເປັນສະເພາະອັນໃຊ້ຢູ່.

ຕໍ່ມາຂ້ອຍພົບ pi.dev ແລະ ໄດ້ພົບພ້າງຄືນສ່ວນຫນຶ່ງ ແຕ່ຍັງບໍ່ໄດ້ຈັດຖ້າຈາກເຂົ້າໄປ. ຂ້ອຍອ່ານກ່ຽວກັບມັນ, ກິດສະພາບວິດີໂອກ່ຽວກັບມັນ ແລະ ຕັດສິນໃຈຈະໃຫ້ຍິນຄືນ ແລະ ຂອຍຍິ້ງສະຕ້ອນມັກ pi ແລ້ວ.

ເພາະວ່າມັນເປັນ open source ຂ້ອຍຄິດວ່າມັນມີຄ່າທີ່ຈະສ້າງສຳລັບ. ຂ້ອຍຍັງໄດ້ເລືອກ provider ອື່ນກົງກໍ. ຂ້ອຍຍັງຄິດວ່າຂອງຢືນຢາງເພາະສາທິດຊຸດຜ່ອນແລະຜ່ອນຈາກໜຶ່ງ provider/model ເຊັ່ນ Claude ບໍ່ສາມາດອອກຜະໂມຊັດໄດ້ສະເໝີໄປ.

ດັ່ງນັ້ນ ຂ້ອຍກຳລັງສ້າງມັນຢູ່ນີ້.

This checkout is maintained as a local-model edition of pi-web. It follows the
upstream project for shared improvements, while keeping local deployment,
context stability, and local-model testing on a separately released track.

## ເປັນຫຍັງ local model ຈຳເປັນໂມດການທົດສະໝ່ຮ່ວມທີ່ແຕກຕ່າງ

ການປະສົບການແຕ່ງອິສລັມຂອງ pi-web ແມ່ນພື້ນຖານທີ່ດີ, ແຕ່ local inference
ມີຮອຍລາຍທີ່ແຕກຕ່າງຈາກ a typical hosted model. A local model may
ມີຊ້າລົງຢ່າງໄວເພື່ອ context ເຕີບ ແລະ ແລະ ແລະ ແລະ ແລະ ແລະ ແລະ ແລະ

This edition approaches the problem in layers:

1. **Preserve upstream first.** Shared UI and session behavior continue to come
   from pi-web; local changes are isolated behind effective Local Mode.
2. **Prevent before recovering.** A percentage-based 65% context boundary is
   enforced before subsequent model calls, including calls inside long tool loops.
3. **Recover only with evidence.** Automatic continuation is limited to recognized
   context, transport, and thinking-only incidents—not authentication, quota, or
   arbitrary provider errors.
4. **Bound every autonomous action.** Recovery incidents are deduplicated,
   progress is required before another rescue, and startup considers at most one
   recently active Local session.
5. **Keep a manual exit.** Force Compact summarizes rather than wipes history, so
   the user can rescue a session without pretending the context never existed.
6. **Protect cloud compatibility.** Cloud Mode keeps the upstream semantics and
   controls; local-model optimizations do not silently redefine cloud sessions.

That is the real difference in this fork: it treats local inference as a distinct
operational environment, not merely another model name in a dropdown.
