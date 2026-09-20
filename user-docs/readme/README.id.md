<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · **Bahasa Indonesia** · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Kendalikan agen coding [pi](https://pi.dev) Anda dari ponsel, tablet, atau laptop — di mana pun dalam jaringan Anda, atau secara jarak jauh melalui Tailscale.

Ini adalah PWA penuh, sehingga Anda bisa menginstalnya dan menggunakannya seperti aplikasi native di perangkat apa pun. Anggap saja sebagai ruang kerja AI pribadi Anda — seperti Cowork milik Claude, tetapi dengan model yang berbeda — lakukan percakapan lintas model, coding dari ponsel, atau ubah menjadi [asisten pribadi](../en/personal-assistant.md) yang berjalan di mesin Anda.

Jadikan milik Anda: ganti tema dan font, serta gunakan dalam bahasa Anda sendiri — pi-web hadir dengan beberapa bahasa dan Anda bisa menambah bahasa Anda sendiri. Lebih banyak fitur sedang dalam pengembangan, tetapi tidak akan menjadi berlebihan: apa pun yang tidak Anda butuhkan dapat dimatikan di pengaturan.

</div>

## Mengapa edisi model lokal ini?

pi-web asli tetap menjadi fondasi upstream untuk fitur dan perbaikan bersama. Edisi ini mempertahankan pengalaman tersebut, kemudian menambahkan lapisan keandalan untuk model yang berjalan di mesin Anda sendiri atau di tempat lain dalam LAN Anda—di mana proses generasi sering kali lebih lambat, memori terbatas, dan konteks panjang dapat menghentikan sesi yang sebenarnya sehat.

| Area | Upstream pi-web | Edisi ini |
|------|-----------------|--------------|
| Kebijakan model/runtime | Perilaku pi-web standar | Mode **Auto / Local / Cloud** per sesi, dengan deteksi lokal yang sadar endpoint dan override manual yang persisten |
| Penanganan konteks panjang | Perilaku kompaksi pi normal | Local Mode mengompak secara proaktif pada **65%** dan mengecek lagi di dalam loop pemanggilan tool yang panjang sebelum permintaan model berikutnya |
| Keamanan kompaksi | Ringkasan standar | Checkpoint rolling yang dibatasi, satu penulisan ulang yang lebih ketat untuk output tidak valid/terpotong, dan deteksi tanpa kemajuan alih-alih kompaksi ulang yang tak berujung |
| Jalankan yang terputus | Penanganan worker dan error normal | Pemulihan yang dibatasi untuk overflow konteks, pemberhentian hanya-thinking, dan gangguan transport yang dipilih, dengan pemutus loop yang persisten |
| Penyelamatan manual | Detail konteks standar | **Force Compact** tetap tersedia sebagai jalur pemulihan eksplisit tanpa menghapus percakapan |
| Kompatibilitas dan rilis | Proyek dan garis rilis asli | Pengaman khusus-lokal tetap di balik Local Mode; Cloud Mode mempertahankan perilaku upstream, dan perubahan upstream ditinjau dan dirilis di sini secara independen |

Ini bukan penulisan ulang atau pengganti upstream. Ini adalah profil operasi yang dipelihara secara sengaja untuk orang-orang yang menginginkan privasi dan kendali model lokal tanpa menerima sesi jangka panjang yang rapuh. Lihat [panduan pengguna](../en/README.md) untuk alur kerja pengguna dan [pengembangan edisi model-lokal](../../docs/dev/local-llm-development.md) untuk implementasi dan kebijakan sinkronisasi.

> [!TIP]
> Baru di sini? **[Baca panduan pengguna →](../en/README.md)** untuk tur lengkap fitur, langkah instalasi, dan tips. ([Bahasa lain →](../README.md))

## Screenshot

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## Bagaimana Ini Terintegrasi

```
 pi (terminal)                 Browser (phone / tablet / laptop)
      │                                │
      │  writes JSONL                  │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           via MagicDNS)
```

- **pi** menuliskan JSONL percakapan ke `~/.pi/agent/sessions/` saat bekerja.
- **pi-web** adalah server Go yang membaca file-file tersebut, merendernya di browser, dan mengirimkan pembaruan langsung melalui SSE.
- Worker **pi --mode rpc** menangani chat yang diinisiasi browser — satu per sesi, dihapus setelah 10 menit idle.
- **fsnotify** memantau direktori sesi sehingga browser dimuat ulang dalam hitungan milidetik dari output baru.
- **Tailscale Serve** menerbitkan server localhost sebagai endpoint HTTPS di tailnet Anda.

## Instalasi

```bash
pi install npm:@timmygod/pi-web-local
```

Sekian — ini mengunduh binary yang sesuai, mengatur auto-start, dan mendaftarkan perintah `/web`, `/pi-web`, `/remote`, dan `/refresh`.

Setelah diinstal, buka `http://127.0.0.1:31415` di browser Anda. Dari pi, gunakan `/web` untuk membuka sesi saat ini di browser Anda secara instan. Jika Tailscale berjalan di mesin Anda, pi-web secara otomatis menerbitkan endpoint HTTPS di tailnet Anda — gunakan `/remote` dari pi untuk mendapatkan kode QR dan URL untuk perangkat apa pun di tailnet Anda.

> **Akses jarak jauh macOS:** Instal dan buka Tailscale secara interaktif, setujui prompt administrator, dan masuk. Kemudian jalankan `/pi-web restart`, diikuti dengan `/remote`.

Untuk instalasi manual, unduhan binary, atau build dari source, lihat [user-docs/install.md](../en/install.md).

## Integrasi Pi

Setelah `pi install npm:@timmygod/pi-web-local`, Anda mendapatkan:

| Perintah | Fungsinya |
|---------|--------------|
| `/web` | Buka sesi saat ini di browser Anda (SSH-aware: lewati browser dan hanya tampilkan URL) |
| `/pi-web` | Tampilkan status, versi, mulai/berhenti/mulai ulang server, atau update |
| `/remote` | Tampilkan kode QR dan URL untuk akses jarak jauh melalui Tailscale |
| `/refresh` | Ambil pesan baru yang dituliskan dari browser jarak jauh kembali ke sesi terminal |

**Auto-titling** sesi dibangun di dalam pi-web itu sendiri dan dikonfigurasi di halaman `/settings`. Fitur ini **aktif secara default** dan memberi nama sesi secara otomatis. Anda bisa memilih:

- **Kapan memberi judul** — sekali per sesi, atau pada setiap pesan baru (default).
- **Model judul** — secara default heuristik kata bawaan yang gratis dan instan (tanpa AI), atau pilih model (misalnya yang kecil/cepat) untuk judul yang lebih cerdas dan ditulis oleh model.

Paket ini juga menginstal binary pi-web ke `~/.pi/agent/bin/pi-web` dan mengatur auto-start saat login.

## Auto-Start Saat Login

Perintah `pi install npm:@timmygod/pi-web-local` mengatur ini secara otomatis:

| OS | Mekanisme |
|----|-----------|
| macOS | plist launchd di `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | layanan user systemd di `~/.config/systemd/user/pi-web.service` |
| Windows | entri Run-key `HKCU` yang menjalankan starter tersembunyi di `~/.config/pi-web/` |

Untuk mengatur token untuk akses jarak jauh, buat `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Untuk detail lebih lanjut (setup manual, port kustom, bind non-loopback), lihat [user-docs/install.md](../en/install.md).

## Pengembangan

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Untuk sinkronisasi upstream, pengujian model-lokal, dan alur kerja rilis paralel, lihat [Pengembangan edisi model-lokal](../../docs/dev/local-llm-development.md).
