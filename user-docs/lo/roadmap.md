# Roadmap

Roadmap ນີ້ແມ່ນຂອງລາຍເວີຊັນ local-model ຂອງ pi-web. ລັກສະນະຈາກ upstream ຖືກສະໜາມຄວາມສອດຄ່ອງ (synchronize) ມື້ລະເທື່ອ, ສ່ວນວຽກດ້ານການ deploy ໃນ local ແລະ ຄວາມໝັ້ນຄົງ ຖືກກວດສອບ ແລະ ປ່ອຍຜົນອອກມາໃນເສັ້ນນີ້; ເບິ່ງ [Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web ຖືກສ້າງມາສຳລັບກຸ່ມແຈກ 2 ກຸ່ມ:

- **ສຳລັບ developer** — ຜູ້ທີ່ຢູ່ໃນ terminal ແຕ່ຢາກຕໍ່ເພີ່ມ session ຈາກໂທລະສັບ, ສົ່ງຕໍ່ໃຫ້ server ເທັດ, ຫຼື ຝົກສັງເກດວຽກທີ່ໃຊ້ເວລາດົນຈາກທຸກແຖວ.
- **ສຳລັບຜູ້ທີ່ບໍ່ແມ່ນ developer** — ທີ່ພຽງແຕ່ຢາກໄດ້ AI app ທີ່ສວຍງາມ ແລະ ໃຊ້ໄດ້ເລີຍ. ເປີດມັນ, ຕີພິມ, ເລີຍ. ບໍ່ມີ terminal, ບໍ່ມີ SSH, ບໍ່ມີຄວາມສັບສົນ. ຄືເຄື່ອງມື AI ທີ່ໃຊ້ງ່າຍທີ່ສຸດ, ແຕ່ມີທາງເລືອກ model ແລະ ຄວາມອິດສະຫຼະຂອງ open-source.

ນີ້ຄືສິ່ງທີ່ຈະມາ.

Edition ນີ້ຕິດຕາມ upstream pi-web ໃນເສັ້ນ release ແຍກກັນ. ລັກສະນະຈາກ upstream ຖືກນຳເຂົ້າມື້ລະເທື່ອ; ວຽກດ້ານຄວາມໝັ້ນຄົງຂອງ local-model ຖືກຈັດອັນດັບ ແລະ ກວດສອບທາງນີ້ ໂດຍບໍ່ປ່ຽນປະຫວັດ release ຂອງ upstream.

---

## ກຳລັງ (ໃສ່ແລ້ວ)

ທຸກສິ່ງທີ່ລາຍການຢູ່ໃນ [ຕາຕຳລາລັກສະນະ](README.md#what-you-can-do-with-pi-web) ສາມາດໃຊ້ໄດ້ແລ້ວມື້ນີ້.

---

## ຕໍ່ໄປ

| # | ລັກສະນະ | ເຮັດຫຍັງ |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | ຊົ່ວ chat ກັບ pi ຜ່ານ Telegram ຫຼື Discord — ເໝາະສຳລັບ workflow ຂອງໂທລະພັດໃສ່ໃນຂະນະເດີນທາງ. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | ຕິດຕາມ token, ຄວາມແຕ້ມຄ່າໃຊ້ຈ່າຍ, ການວິເຄາະ session — ຫຼຸດວ່າທ່ານໃຊ້ pi ແນວໃດ. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | ຕັ້ງ visibility ທີ່ທ່ານຕ້ອງການສຳລັບ thinking, tools ແລະ ຜົນຂອງ tools ໃນທຸກ session. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | ສົ່ງຄຳສັ່ງຕໍ່ເພີ່ມໃນຂະນະທີ່ pi ກຳລັງ run — ປູນທາງໃຫ້ມັນໃນທາງ. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | ຫຍໍ້ conversation ດົນໄວ້ຈາກ web UI ແລ້ວໂດຍກົງ, ບໍ່ຕ້ອງໃຊ້ terminal. |

---

## ດຳລົມ

| # | ລັກສະນະ | ເຮັດຫຍັງ |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | ເບິ່ງຕົວໂຕ້ໂດຍ project ແລະ ເບິ່ງການປ່ຽນແປງຂອງ git ໂດຍກົງໃນ pi-web. Opt-in, ດັ່ງນັ້ນມັນຈຶ່ງບໍ່ຢູ່ໃນທາງຂອງທ່ານ. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | ເລືອກໄວ້ໃຫ້ prompt ຖືກ run ດ້ວຍຕົນເອງ — daily standups, ສະຫຼຸບຕອນເຊົ້າ, ວຽກທີ່ເກີດຮຽບຮ້ອຍ. ຖືກຈັກກະພັດຜ່ານ admin ສຳລັບຄວາມປອດໄພ. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | ປ່ຽນທຸກ keyboard shortcut ໃຫ້ກົງກັບ muscle memory ຂອງທ່ານ. |

---

## ຈາຍ

ເປົ້າໝາຍໆຍາຍ: pi-web ຄວນເປັນ **ຕົວຕໍ່ກັນສຳລັບ pi** — ສຳລັບທຸກຄົນ.

- **ຜູ້ທີ່ບໍ່ແມ່ນ dev** ເປີດມັນຄື app ອື່ນ. ເລືອກ model. ຕີພິມ. ສຳເລັດ. ບໍ່ມີ command line ນັບແຕ່ເລີ່ມ.
- **Devs** ໄດ້ການ integrate ລະອຽດລະອຽດ — ຂົນສົ່ງ remote, multi-session dashboards, ການເບິ່ງທີ່ຮູ້ git, ແລະ messaging bots.
- **ທຸກຄົນ** ໄດ້ຄວາມອິດສະຫຼະໃນ model, ຄວາມຄົບຖ້ວນຂອງ open-source, ແລະ UI ທີ່ຮູ້ສຶກຄິດສະບາຍໃນທຸກໆມື້.

---

> 💡 ມີຄວາມຄິດ? [ເປີດ issue](https://github.com/timmygod/pi-web/issues/new) ຫຼື ເຂົ້າຮ່ວມໃນການຊົດວາດ.
