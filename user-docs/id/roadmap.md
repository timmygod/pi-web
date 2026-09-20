# Roadmap

Roadmap ini milik edisi local-model dari pi-web. Fitur upstream
disinkronkan secara berkala, sementara pekerjaan deployment lokal dan keandalan
divalidasi dan dirilis di jalur ini; lihat [Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web dibuat untuk dua audiens:

- **Untuk developer** — yang hidup di terminal tetapi ingin melanjutkan sesi dari mobile, meneruskan ke server remote, atau mengawasi tugas yang berjalan lama dari mana saja.
- **Untuk non-developer** — yang hanya ingin aplikasi AI yang indah dan berfungsi. Buka, ketik, vibe. Tanpa terminal, tanpa SSH, tanpa kebingungan. Seperti tools AI paling ramah pengguna, tetapi dengan pilihan model dan kebebasan open-source.

Berikut yang akan datang.

Edisi ini melacak pi-web upstream di jalur rilis yang terpisah. Fitur
upstream diimpor secara berkala; pekerjaan keandalan local-model diprioritaskan
dan divalidasi di sini tanpa mengubah riwayat rilis upstream.

---

## Sekarang (sudah dirilis)

Semua yang tercantum di [tabel fitur](README.md#what-you-can-do-with-pi-web) sudah tersedia hari ini.

---

## Berikutnya

| # | Fitur | Fungsinya |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Bot Telegram & Discord** | Ngobrol dengan pi lewat Telegram atau Discord — cocok untuk alur kerja asisten pribadi saat bepergian. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | Pelacakan token, estimasi biaya, analitik sesi — ketahui cara Anda menggunakan pi. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | Atur visibilitas yang Anda sukai untuk thinking, tools, dan output tools di semua sesi. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | Kirim instruksi lanjutan saat pi masih berjalan — pandu di tengah jalan. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **Perintah `/compact`** | Ringkas percakapan panjang langsung dari web UI, tanpa terminal. |

---

## Perencanaan

| # | Fitur | Fungsinya |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | Jelajahi pohon file proyek dan lihat perubahan git langsung di pi-web. Opt-in, sehingga tidak mengganggu Anda. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Jadwalkan prompt untuk berjalan otomatis — standup harian, ringkasan pagi, tugas berulang. Admin-gated demi keamanan. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Shortcut yang bisa dikustomisasi** | Remap semua shortcut keyboard agar sesuai memori otot Anda. |

---

## Visi

Tujuan jangka panjang: pi-web harus menjadi **antarmuka untuk pi** — untuk semua orang.

- **Non-dev** membukanya seperti aplikasi lain. Pilih model. Ketik. Selesai. Tanpa command line sama sekali.
- **Dev** mendapat integrasi mendalam — remote handoff, dashboard multi-sesi, browsing yang memahami git, messaging bots.
- **Semua orang** mendapat kebebasan model, transparansi open-source, dan UI yang terasa cermat di setiap langkah.

---

> 💡 Punya ide? [Buat issue](https://github.com/timmygod/pi-web/issues/new) atau bergabunglah ke diskusi.
