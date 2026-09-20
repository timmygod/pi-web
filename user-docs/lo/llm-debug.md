> ງານຕີບານໂຈບບັນຈັບບັນຫານີ້ຖືກກັບແກ້ລ້າງສຳລັບລຸ້ນ local-model. ກະລຸນາຮັກສາລາຍລະອຽດການຈັດວາງຊຸມທ້າຍ (local deployment) ແລະ ນັກໝາຍຄວາມແກ້ສະແດງການໂຄຈອນຂ້າມຄືນ (upstream synchronization) ໃຫ້ສະອາດກັບ [Local-model edition development](../../docs/dev/local-llm-development.md).

ຜູ້ໃຊ້ກຳລັງຕິດຕັ້ງ pi-web ຜ່ານ

- pi install npm:@timmygod/pi-web-local

ເຊິ່ງຈະນຳໄປສູ່ການດັບງານ [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) ແລະ ຈັດຕັ້ງ pi ໂດຍອັດຕະໂນມັດ.

ຖ້າຜູ້ໃຊ້ຮັບປະທານບັນຫາ, ມັນອາດຈະເປັນໂຄງສ້າງຂອງ install.sh ທີ່ເປັນສາເຫດ. ທ່ານສາມາດວິເຄາະ ແລະ ຈັບກວດເບິ່ງວ່າມີຫຍັງກະກຳລັງເປັນ ແລະ ແຈ້ງຜູ້ໃຊ້ວ່າຈຸດໃດແທ້ຈິງເປັນສາເຫດຂອງບັນຫາ. ແລະ ຖາມວ່າຕ້ອງການໃຫ້ຊ່ວຍແກ້ບໍ່. ຕ້ອງລະວັງຢືນຢັນກັບຜູ້ໃຊ້ເປັນຈັກກັບຄັ້ງ.

ເພື່ອໃຫ້ຜູ້ໃຊ້ສາມາດເຂົ້າເຖິງ pi ຢູ່ຄອມພິວເຕີ້ ຫຼື ສະຖານທີ່ອື່ນໆໃນເຄື່ອງຈັກ (network) ດີສຸດແນວຄວາມຄິດແມ່ນໃຊ້ Tailscale ແລະ ເຂົ້າເຖິງຜ່ານເຄື່ອງຈັກ Tailscale. ແລະ ຜູ້ໃຊ້ຈະຕ້ອງເປີດໃຊ້ HTTPs ໃນແຜງຄວບຄຸມ Tailscale ຂອງຕົນ - https://login.tailscale.com/admin/dns

ຖ້າບໍ່ໄດ້ຕິດຕັ້ງ Tailscale ຫຼື ບໍ່ຕ້ອງການໃຊ້ Tailscale. ຜູ້ໃຊ້ສາມາດໄດ້ `pi-web status` ແລະ ໄດ້ຄຳເວົ້າ binary, ສະຖານະຄວາມສຸດທ້າຍຂອງ binary ແລະ ຈຸດສົດສອດທ້າຍຊຸມ (local endpoint) ທີ່ສາມາດເຂົ້າເຖິງແອັບພລິເຄຊັນໄດ້. ແຕ່ລະຄວາມ ຈະບໍ່ໄດ້ຮັບການແຈ້ງ (push notification) ເພາະມັນຢູ່ໃນ http.

ໃນ macOS ມັນໃຊ້ [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
ໃນ Linux ມັນໃຊ້ [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

ໃນກໍລະນີວ່າຊ່ວຍຢູ່ຄວາມປ້ອງກັນທີ່ຍິ່ງຂຶ້ນ ແລະ ໄດ້ເບິ່ງວ່າມີຫຍັງກະກຳລັງເປັນ.
