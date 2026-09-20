# Selamat Datang di pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · **Bahasa Indonesia** · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Mempertimbangkan untuk mencoba pi-web? Silakan — Anda akan menyukainya.**

pi-web adalah antarmuka web dan PWA yang indah untuk [pi](https://pi.dev) — agen coding AI open-source. Anda bisa menelusuri, membaca, dan melanjutkan sesi pi Anda dari browser mana pun, di perangkat mana pun, dengan fitur yang dipikirkan matang-matang di setiap langkah.

## Apa yang berbeda di edisi ini?

Repositori ini mempertahankan antarmuka pi-web upstream dan fitur-fitur bersama, tetapi mengubah cara sesi dilindungi ketika model yang dipilih dijalankan secara lokal atau di LAN Anda.

- **Pilih kebijakan runtime per sesi.** Deteksi otomatis endpoint lokal/LAN ketika metadata provider jelas; Local dan Cloud adalah override manual yang persisten.
- **Cegah kegagalan konteks sedini mungkin.** Local Mode melakukan kompresi pada 65% pemakaian dan mengecek ulang antara panggilan tool, sebelum permintaan model berikutnya.
- **Jaga ringkasan tetap terbatas.** Rolling checkpoints menghindari ringkasan lama yang terus membesar, mencoba sekali lagi dengan anggaran yang lebih ketat, dan berhenti dengan aman ketika kompresi tidak memberikan kemajuan berarti.
- **Pulihkan secara konservatif.** Overflow konteks, gangguan transport yang dipilih, dan penghentian prematur hanya-reasoning dapat melanjutkan secara otomatis, tetapi deduplikasi insiden dan circuit breaker yang sadar progres mencegah loop pemulihan.
- **Biarkan pengguna memegang kendali.** Force Compact selalu menjadi jalur penyelamatan manual yang terlihat, sementara Cloud Mode mempertahankan alur kerja dan kontrol upstream.

Hasil praktisnya sederhana: tugas model-model lokal yang panjang seharusnya dikompres sebelum jatuh, dipulihkan sekali ketika pulihan aman, dan berhenti dengan bersih alih-alih melingkar ketika tidak aman.

**pi-web dibuat untuk dua jenis orang:**

- 🧑‍💻 **Untuk developer** — yang hidup di terminal tetapi ingin melanjutkan sesi dari ponsel, menyerahkan ke server jarak jauh, atau memantau tugas yang berjalan lama dari mana saja.
- ✨ **Untuk non-developer** — yang hanya ingin aplikasi AI yang indah dan berfungsi. Buka, ketik, vibe. Tanpa terminal, tanpa SSH, tanpa kebingungan. Seperti alat AI yang paling ramah pengguna, tetapi dengan pilihan model dan kebebasan open-source.

---

## Kenapa pi-web?

Anda sudah dalam alur yang dalam dengan pi di terminal Anda. pi-web mempertahankan momentum itu ketika Anda meninggalkan meja:

- **Lanjutkan dari mana saja** — lanjutkan sesi dari ponsel, tablet, atau komputer lain. Tanpa SSH, tanpa Termius — cukup buka browser Anda.
- **Dasbor multi-sesi** — mulai pekerjaan dalam satu sesi sambil mengamati sesi lain yang streaming. Cari melintasi proyek, filter per branch, temukan apa yang Anda butuhkan dengan cepat.
- **Dasar open-source** — pi sepenuhnya open source dan tidak terikat pada provider. Anda tidak dikunci pada satu model atau vendor. pi-web juga open source.
- **Akses remote yang aman** — autentikasi token bawaan sehingga Anda bisa mengeksposnya di LAN atau Tailscale tanpa khawatir.
- **Bagikan pekerjaan Anda** — ekspor sesi sebagai snapshot statis atau GitHub Gist rahasia dengan satu klik.

> Penasaran dengan cerita di baliknya? [Baca kenapa kami membuatnya →](why.md)

---

## pi-web sebagai ruang kerja AI pribadi Anda 🏠

pi-web adalah PWA (Progressive Web App), sehingga Anda bisa **menginstalnya seperti aplikasi native** di desktop, laptop, ponsel, atau tablet Anda — tanpa perlu app store. Di desktop, ia membuka di jendelanya sendiri tanpa kerangka browser, sehingga terlihat dan terasa seperti aplikasi desktop sungguhan.

Pikirkan sebagai **Claude Cowork Anda sendiri** — ruang kerja AI pribadi yang hidup di mesin Anda — kecuali ia open source dan tidak terikat pada model:

- **Anda yang memiliki stack.** Pilih model apa pun, ganti kapan pun Anda mau. Jalankan satu yang lokal dan data Anda tidak pernah meninggalkan mesin Anda.
- **Orang non-teknis bisa menggunakannya.** Pasang pi-web di mesin mereka, tunjukkan cara memakainya sekali, dan mereka siap berlayar. Orang tua Anda, pasangan Anda, teman-teman non-tech Anda — tanpa terminal, tanpa SSH, hanya antarmuka chat yang sudah familiar.
- **Satu pemasangan, banyak pengguna.** Instal di desktop Anda dan bagikan layar, atau ekspos di jaringan rumah Anda dan biarkan anggota keluarga membukanya di perangkat mereka sendiri.

Ingin lebih dari coding? Ubah menjadi [personal assistant](personal-assistant.md) khusus yang tahu siapa Anda dan hidup di mesin Anda — seperti OpenClaw atau Hermes Anda sendiri.

> 💡 **Tips pro:** Instal pi-web sebagai PWA dari Chrome/Edge (klik ikon instal di bilah alamat) atau Safari (Share → Add to Dock). Ia menjadi tidak dapat dibedakan dari aplikasi native.

---

## Apa yang bisa Anda lakukan dengan pi-web

| | |
|---|---|
| 📱 **PWA** | Instal pi-web sebagai Progressive Web App di desktop, ponsel, atau tablet untuk sensasi native. |
| 🔄 **Lanjutkan sesi** | Ambil alih percakapan apa pun tepat dari tempat Anda berhenti — teks, gambar, pergantian model, semuanya dari browser. |
| 🆕 **Mulai sesi baru** | Buat sesi baru terhadap jalur proyek apa pun, langsung dari antarmuka web. |
| 📡 **Streaming langsung** | Amati respons pi streaming secara real time dengan latensi ~ms. Mode Follow menjaga Anda terkunci pada yang terbaru. |
| 🌲 **Tampilan pohon** | Navigasi pohon pesan native pi — lihat struktur percakapan lengkap, lompat ke branch mana pun, dan fork dari titik mana pun. |
| 🔀 **Fork sesi** | Fork sesi dari pesan apa pun atau bahkan panggilan tool tertentu — jelajahi arah yang berbeda tanpa kehilangan tempat Anda. |
| 🔍 **Jelajah & cari** | Filter sesi melintasi proyek, cari berdasarkan nama, navigasi branch — seluruh riwayat sesi Anda dalam sekilas pandang. |
| 🌿 **Integrasi Git** | Lihat branch saat ini dan buka PR GitHub langsung dari viewer sesi. |
| 📝 **Scratchpad** | Catat catatan, todo, atau pikiran singkat di samping sesi Anda tanpa mengganti aplikasi. |
| 💬 **Anotasi** | Sorot dan beri komentar pada bagian mana pun dari sesi — bagus untuk code review, umpan balik, atau menandai momen penting. |
| 🎨 **Tema & kustomisasi** | Beralih antara mode gelap dan terang, sesuaikan antarmuka sesuai selera — buat pi-web terasa seperti *milik Anda*. |
| 🌐 **Multi-bahasa** | 14 bahasa bawaan (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Tambahkan bahasa kustom Anda sendiri dari Settings. |
| 🐱 **Wellness & pomodoro** | Terlalu banyak vibe coding tidak sehat. Timer pomodoro bawaan dengan teman kucing dan pengingat tidur untuk menjaga keseimbangan Anda. |
| 📤 **Bagikan & ekspor** | Unduh JSONL, ekspor snapshot statis yang dirender dengan tampilan native `pi.dev`, atau bagikan sebagai GitHub Gist privat — semuanya dirender di sisi klien. |
| 🔔 **Suara notifikasi** | Chime notifikasi yang dapat dikustomisasi untuk event sesi — tetap mengikuti kabar bahkan ketika pi-web ada di tab lain. |
| ⌨️ **Shortcut keyboard** | Navigasi gaya Vim, aksi cepat — [referensi lengkap →](keyboard-shortcuts.md) |
| 🤖 **Personal assistant** | Ubah pi-web menjadi asisten AI Anda sendiri yang hidup di komputer Anda — seperti OpenClaw atau Hermes. [Pasang →](personal-assistant.md) |
| 🗓️ **Bicara ke jadwal** | Dari sesi pi, katakan “tambahkan jadwal pukul 2 pagi waktu Singapura untuk …” — `/skill:pi-web-schedule`. |
| 📝 **Bicara ke catatan & pengaturan** | “Tulis ini di catatan” (`/skill:pi-web-notes`) atau “ganti ke mode gelap” (`/skill:pi-web-settings`). |

---

## Navigasi cepat

| Jika Anda mencari… | Baca |
|---|---|
| Cara menginstal, mengonfigurasi, dan menggunakan pi-web | [install.md](install.md) |
| Gunakan pi-web sebagai personal assistant | [personal-assistant.md](personal-assistant.md) |
| Referensi keyboard shortcuts | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Kenapa pi-web ada | [why.md](why.md) |
| Apa yang akan datang berikutnya | [roadmap.md](roadmap.md) |
| Punya masalah instal? Biarkan LLM Anda memperbaikinya — tempelkan tautan llm-debug.md kepada mereka | [llm-debug.md](llm-debug.md) |
| Memelihara edisi model lokal ini | [development notes](../../docs/dev/local-llm-development.md) |

---

## Tangkapan Layar

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

pi-web dibangun dengan cinta dan banyak malam yang larut. Saya membayar paket coding (Claude Code, OpenCode, dll.) dari saku sendiri untuk menjaga proyek ini terus maju. Jika pi-web bermanfaat bagi Anda, dukungan Anda akan berarti banyak.

**Cara membantu:**

- 💰 **[Sponsor di GitHub](https://github.com/sponsors/setkyar)** — bantu tutupi alat-alat yang memungkinkan ini
- ☕ **[Beli saya secangkir kopi](https://buymeacoffee.com/setkyar)** — setiap sedikit sangat membantu
- ⭐ **Star repo** — tidak memakan biaya dan membantu lebih banyak orang menemukan pi-web
- 📢 **Bagikan dengan teman & keluarga** — jika Anda kenal seseorang yang akan menyukai pi-web, kirimkan ke mereka

Tidak bisa sponsor? Tidak masalah sama sekali — sebuah star dan share sangat berarti. Terima kasih sudah ada di sini. 🙏

---

Selamat coding! 🚀
