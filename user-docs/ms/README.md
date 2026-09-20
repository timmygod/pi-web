# Selamat datang ke pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · **Bahasa Melayu** · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Bertindakkan untuk mencuba pi-web? Cuba sahaja — anda pasti akan menyukainya.**

pi-web ialah antara muka web (UI) dan PWA yang cantik untuk [pi](https://pi.dev) — ejen pengekodan AI sumber terbuka. Ia membolehkan anda melayari, membaca, dan meneruskan sesi pi anda daripada mana-mana pelayar, pada mana-mana peranti, dengan ciri yang berfikir di setiap langkah.

## Apa yang berbeza dalam edisi ini?

Repositori ini mengekalkan antara muka pi-web hulu dan ciri berkongsi, tetapi
mengubah cara sesi dilindungi apabila model yang dipilih berjalan secara setempat atau pada
LAN anda.

- **Pilih dasar eksekusi (runtime) setiap sesi.** Auto mengesan endpoint setempat/LAN apabila
  metadata provider jelas; Local dan Cloud ialah penindasan (override) manual yang kekal.
- **Mencegah kegagalan konteks lebih awal.** Local Mode mengemas (compact) pada 65% penggunaan dan memeriksa
  sekali lagi di antara panggilan alat, sebelum permohonan model seterusnya.
- **Menekalkan ringkasan terhad.** Punca semak bergulung (rolling checkpoint) mengelakkan ringkasan lama
  daripada membesar tanpa henti, mencuba semula sekali dengan bajet yang lebih ketat, dan berhenti dengan selamat apabila pengemaskinian (compaction) tidak
  membuahkan kemajuan yang bermakna.
- **Pemulihan secara konservatif.** Pelumpuran (overflow) konteks, gangguan penghantaran (transport) yang dipilih,
  dan hentian awal yang hanya berfikir (reasoning) boleh dimulakan semula secara automatik, tetapi pengasingan (deduplication) insiden
  dan pemecah litar (circuit breaker) yang peka terhadap kemajuan menghalang gelung pemulihan.
- **Membiarkan pengguna kekal mengawal.** Force Compact sentiasa ialah laluan
  penyelamatan manual yang kelihatan, manakala Cloud Mode mengekalkan aliran kerja dan kawalan hulu.

Hasil praktikalnya adalah mudah: tugasan model setempat yang panjang sepatutnya mengemas sebelum ia
runtuh, pemulihan sekali apabila pemulihan selamat, dan berhenti dengan kemas apabila ia tidak
dapat dipulihkan, berbanding berputar.

**pi-web dibina untuk dua jenis orang:**

- 🧑‍💻 **Untuk pengatur rebels** — yang hidup dalam terminal tetapi ingin meneruskan sesi dari telefon mudah alih, menyerahkan kepada pelayan jauh, atau memantau tugasan jangka panjang dari mana-mana.
- ✨ **Untuk bukan pengatur rebels** — yang hanya ingin aplikasi AI yang cantik dan berfungsi. Buka, taip, vibe. Tanpa terminal, tanpa SSH, tanpa kekeliruan. Seperti alat AI paling mesra pengguna, tetapi dengan pilihan model dan kebebasan sumber terbuka.

---

## Mengapa pi-web?

Anda sudah jauh dalam aliran dengan pi di terminal anda. pi-web mengekalkan momentum tersebut apabila anda jauh dari meja kerja:

- **Mulakan semula dari mana-mana** — teruskan sesi dari telefon, tablet, atau komputer lain. Tanpa SSH, tanpa Termius — hanya buka pelayar anda.
- **Dasbor berbilang sesi** — mulakan kerja dalam satu sesi sambil memerhati satu lagi yang sedang dialirkan. Cari merentas projek, tapis mengikut branch, cari apa yang anda perlu dengan cepat.
- **Asas sumber terbuka** — pi sepenuhnya sumber terbuka dan tidak bergantung pada provider. Anda tidak dikunci kepada satu model atau pembekal sahaja. pi-web juga sumber terbuka.
- **Akses jauh yang selamat** — pengesahan token terbina dalam (token auth) supaya anda boleh mempamerkannya pada LAN atau Tailscale anda tanpa bimbang.
- **Kongsi kerja anda** — output sesi sebagai snapshot statik atau GitHub Gist rahsia dengan satu klik.

> Tertanya-tanya tentang latar belakang? [Baca mengapa kami membina →](why.md)

---

## pi-web sebagai ruang kerja AI peribadi anda 🏠

pi-web ialah PWA (Progressive Web App), jadi anda boleh **memasangnya seperti aplikasi asal (native)** pada desktop, laptop, telefon, atau tablet — tanpa perlu kedai aplikasi. Pada desktop ia dibuka dalam tetingkapnya sendiri tanpa rangka pelayar (chrome), maka ia kelihatan dan berasa seperti aplikasi desktop sebenar.

Fikirkan ia sebagai **Claude Cowork anda sendiri** — ruang kerja AI peribadi yang berada pada mesin anda — kecuali ia sumber terbuka dan tidak bergantung pada model:

- **Anda memiliki stack tersebut.** Pilih mana-mana model, tukar bilamana yang anda suka. Jalankan satu yang setempat dan data anda tidak akan pernah meninggalkan mesin anda.
- **Orang bukan teknikal boleh menggunakannya.** Setelkan pi-web pada mesin mereka, tunjukkan cara menggunakannya sekali, dan mereka sudah bersedia. Ibu bapa anda, pasangan anda, rakan yang tidak teknikal — tanpa terminal, tanpa SSH, hanya antara muka chat yang biasa.
- **Satu setup, ramai pengguna.** Pasang pada desktop anda dan kongsi skrin anda, atau pameran pada rangkaian rumah dan biarkan ahli keluarga membukanya pada peranti mereka sendiri.

Ingin lebih daripada pengekodan? Tukar ia menjadi [penolong peribadi](personal-assistant.md) khusus yang mengetahui siapa anda dan berada pada mesin anda — seperti OpenClaw atau Hermes anda sendiri.

> 💡 **Petua profesional:** Pasang pi-web sebagai PWA dari Chrome/Edge (klik ikon pasang di bar alamat) atau Safari (Kongsi → Tambah ke Dock). Ia menjadi tidak dapat dibezakan daripada aplikasi asal (native).

---

## Apa yang boleh anda lakukan dengan pi-web

| | |
|---|---|
| 📱 **PWA** | Pasang pi-web sebagai Progressive Web App pada desktop, telefon, atau tablet untuk rasa asal (native). |
| 🔄 **Teruskan sesi** | Ambil sesi semasa anda berhenti — teks, imej, pertukaran model, semuanya dari pelayar. |
| 🆕 **Mulakan sesi baharu** | Cipta sesi baharu terhadap mana-mana laluan projek, terus dari UI web. |
| 📡 **Pengaliran langsung** | Lihat respons pi dialirkan secara masa nyata dengan latensi ~ms. Mod mengikuti (Follow mode) kekal pada terkini. |
| 🌲 **Pandangan pokok** | Navigasi pokok mesej asli pi — lihat struktur perbualan penuh, melompat ke mana-mana branch, dan bercabang (fork) dari mana-mana titik. |
| 🔀 **Cabang (Fork) sesi** | Cabang (Fork) sesi dari mana-mana mesej atau bahkan panggilan alat tertentu — teliti arah berbeza tanpa kehilangan tempat anda. |
| 🔍 **Layar & cari** | Tapis sesi merentas projek, cari mengikut nama, navigasi branch — sejarah sesi anda penuh dengan pandangan sekilas. |
| 🌿 **Integrasi Git** | Lihat branch semasa dan buka PR GitHub terus dari peninjau sesi. |
| 📝 **Scratchpad** | Catat nota, todo, atau idea pantas di sisi sesi anda tanpa menukar aplikasi. |
| 💬 **Anatasi** | Nyatakan dan komen sebahagian sesi — sesuai untuk semakan kod, maklum balas, atau penanda saat penting. |
| 🎨 **Tema & penyesuaian** | Tukar antara mod gelap dan cahaya, sesuaikan UI mengikut cita rasa anda — jadikan pi-web berasa *milik anda*. |
| 🌐 **Berbilang bahasa** | 14 bahasa terbina dalam (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Tambah bahasa kustom anda sendiri dari Tetapan (Settings). |
| 🐱 **Kesihatan & pomodoro** | Terlalu banyak vibe coding tidak sihat. Pemasa pomodoro terbina dalam dengan teman kucing dan peringatan tidur untuk mengekalkan keseimbangan anda. |
| 📤 **Kongsi & output** | Muat turun JSONL, output snapshot statik yang dirender dengan gaya `pi.dev` asli pi, atau kongsi sebagai GitHub Gist persendirian — semuanya dirender di sisi klien. |
| 🔔 **Bunyi notifikasi** | Chime notifikasi yang boleh disesuaikan untuk peristiwa sesi — kekal kemas walaupun pi-web berada di tab lain. |
| ⌨️ **Pintasan papan kekunci** | Navigasi gaya Vim, tindakan pantas — [rujukan penuh →](keyboard-shortcuts.md) |
| 🤖 **Penolong peribadi** | Tukar pi-web kepada penolong AI anda sendiri yang berada pada komputer anda — seperti OpenClaw atau Hermes. [Setelkan →](personal-assistant.md) |
| 🗓️ **Bicarakan jadual** | Dari sesi pi, katakan “tambah jadual pada 2am waktu Singapura untuk …” — `/skill:pi-web-schedule`. |
| 📝 **Bicarakan nota & tetapan** | “Tulis ini dalam nota” (`/skill:pi-web-notes`) atau “tukar ke mod gelap” (`/skill:pi-web-settings`). |

---

## Navigasi pantas

| Jika anda mencari… | Baca |
|---|---|
| Cara pasang, tentuk, dan guna pi-web | [install.md](install.md) |
| Guna pi-web sebagai penolong peribadi | [personal-assistant.md](personal-assistant.md) |
| Rujukan pintasan papan kekunci | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Mengapa pi-web wujud | [why.md](why.md) |
| Apa yang akan datang | [roadmap.md](roadmap.md) |
| Ada masalah pemasangan? Biarkan LLM anda membetulkannya — tampal pautan llm-debug.md kepada mereka | [llm-debug.md](llm-debug.md) |
| Menyelenggara edisi model setempat ini | [development notes](../../docs/dev/local-llm-development.md) |

---

## Screenshot

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Penaja

pi-web dibina dengan cinta dan banyak malam larut. Saya membayar pelan pengekodan (Claude Code, OpenCode, dll.) dari poket sendiri untuk mengekalkan projek ini terus bergerak ke hadapan. Jika pi-web berguna kepada anda, sokongan anda akan bermakna banyak.

**Cara membantu:**

- 💰 **[Taja di GitHub](https://github.com/sponsors/setkyar)** — bantu menampung alat yang membolehkan ini
- ☕ **[Beli saya kopi](https://buymeacoffee.com/setkyar)** — sedikit demi sedikit membantu
- ⭐ **Bintangkan (Star) repositori** — ia percuma dan membantu lebih ramai orang menemui pi-web
- 📢 **Kongsi dengan rakan & keluarga** — jika anda tahu seseorang yang akan menyintai pi-web, hantar ia kepada mereka

Tidak mampu menaja? Tidak mengapa — satu bintang dan perkongsian sudah banyak membantu. Terima kasih kerana berada di sini. 🙏

---

Selamat mengkod! 🚀
