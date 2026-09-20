<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · **Bahasa Melayu** · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Jalankan ejen kod [pi](https://pi.dev) anda daripada telefon, tablet, atau laptop — di mana sahaja pada rangkaian anda, atau secara jarak jauh melalui Tailscale.

Ia adalah PWA penuh, jadi anda boleh memasangnya dan menggunakannya seperti aplikasi asli pada sebarang peranti. Anggap ia sebagai ruang kerja AI peribadi anda sendiri — seperti Claude's Cowork, tetapi dengan model yang berbeza — berchat merentas model, menulis kod daripada telefon anda, atau jadikan ia [penolong peribadi](../en/personal-assistant.md) yang kekal pada mesin anda.

Jadikan ia milik anda: tukar tema dan fon, serta gunakan dalam bahasa anda sendiri — pi-web dibekalkan dengan beberapa bahasa dan anda boleh menambah bahasa anda sendiri. Lebih banyak ciri sedang dalam perjalanan, tetapi ia tidak akan menjadi kumbung: apa-apa yang anda tidak perlukan boleh dimatikan dalam tetapan.

</div>

## Mengapa edisi model tempatan ini?

pi-web asal kekal sebagai asas hulu (upstream) untuk ciri dan
pembaikan bersama. Edisi ini mengekalkan pengalaman tersebut, kemudian menambah lapisan kebolehandalan untuk
model yang berjalan pada mesin anda sendiri atau di tempat lain pada LAN anda — di mana penjanaan
biasanya lebih perlahan, memori adalah terhad, dan konteks yang panjang boleh menjgentilkan sesi yang sihat.

| Bidang | pi-web Hulu (Upstream) | Edisi ini |
|------|-----------------|--------------|
| Dasar model/_RUNTIME | Kelakuan pi-web standard | Mod **Auto / Local / Cloud** per-sesi, dengan pengesanan tempatan yang peka terhadap endpoint dan penyahpalingan manual yang kekal |
| Pengendalian konteks panjang | Kelakuan pemampatan pi normal | Mod Tempatan memampat secara proaktif pada **65%** dan memeriksa semula di dalam gelung panggilan alat yang panjang sebelum permintaan model seterusnya |
| Keselamatan pemampatan | Ringkasan standard | Sekatan titik semak bergulung, satu penulisan semula yang lebih ketat untuk keluaran yang tidak sah/dibataskan, dan pengesanan tanpa kemajuan sebagai ganti pemampatan semula tanpa akhir |
| Larian yang terganggu | Pengendalian pekerja dan ralat normal | Pemulihan terhad untuk limpahan konteks, berhenti-terfikir sahaja, dan gangguan penghantaran yang dipilih, dengan pemecah gelung yang kekal |
| Penyelamatan manual | Butiran konteks standard | **Force Compact** kekal tersedia sebagai laluan pemulihan ekspisit tanpa memadamkan perbualan |
| Keupayaan serasi dan pelepasan | Projek dan baris pelepasan asal | Jagaan tempatan sahaja kekal di sebalik Mod Tempatan; Mod Awan mengekalkan kelakuan hulu, dan perubahan hulu diteliti dan dilepaskan di sini secara bebas |

Ini bukan penulisan semula atau pengganti untuk hulu. Ia adalah profil pengendalian yang dikekalkan secara sengaja untuk orang yang mahukan privasi dan kawalan model tempatan tanpa menerima sesi larian panjang yang rapuh. Lihat
[panduan pengguna](../en/README.md) untuk aliran kerja berhadapan pengguna dan
[pembangunan edisi model tempatan](../../docs/dev/local-llm-development.md) untuk
penggunaannya dan dasar penyegerakan.

> [!TIP]
> Baru di sini? **[Baca panduan pengguna →](../en/README.md)** untuk lawatan penuh ciri, langkah pemasangan, dan petua. ([Bahasa lain →](../README.md))

## Tangkapan Skrin

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Meja (Desktop)</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## Bagaimana Ia Bekerja Bersama

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

- **pi** menulis JSONL perbualan ke `~/.pi/agent/sessions/` semasa ia bekerja.
- **pi-web** ialah pelayan Go yang membaca fail-fail tersebut, merentaknya dalam pelayar, dan menghantar kemas kini langsung melalui SSE.
- Pekerja **pi --mode rpc** mengendalikan chat yang dimulakan dari pelayar — satu per sesi, diendapkan selepas 10 minit tidak aktif.
- **fsnotify** memantau direktori sesi supaya pelayar memuat semula dalam masa beberapa milisaat selepas keluaran baharu.
- **Tailscale Serve** menerbitkan pelayan localhost sebagai titik akhir HTTPS pada tailnet anda.

## Pemasangan

```bash
pi install npm:@timmygod/pi-web-local
```

Itu sahaja — ia memuat turun binari yang sepadan, menyediakan auto-pemulaan, dan mendaftarkan arahan `/web`, `/pi-web`, `/remote`, dan `/refresh`.

Setelah dipasang, buka `http://127.0.0.1:31415` dalam pelayar anda. Daripada pi, gunakan `/web` untuk membuka sesi semasa dalam pelayar anda serta-merta. Jika Tailscale sedang berjalan pada mesin anda, pi-web secara automatik menerbitkan titik akhir HTTPS pada tailnet anda — gunakan `/remote` daripada pi untuk mendapatkan kod QR dan URL untuk sebarang peranti pada tailnet anda.

> **Akses jarak jauh macOS:** Pasang dan buka Tailscale secara interaktif, sahkan petikan pengarah, dan log masuk. Kemudian jalankan `/pi-web restart`, diikuti dengan `/remote`.

Untuk pemasangan manual, muat turun binari, atau pembinaan dari sumber, lihat [user-docs/install.md](../en/install.md).

## Integrasi Pi

Setelah `pi install npm:@timmygod/pi-web-local`, anda memperoleh:

| Arahan | Apa yang dilakukannya |
|---------|--------------|
| `/web` | Buka sesi semasa dalam pelayar anda (sedar SSH: langkau pelayar dan tunjuk URL sahaja) |
| `/pi-web` | Tunjuk status, versi, mula/hentikan/mula semula pelayan, atau kemas kini |
| `/remote` | Tunjuk kod QR dan URL untuk akses jarak jauh melalui Tailscale |
| `/refresh` | Tarik pesan baharu yang ditulis dari pelayar jarak jauh kembali ke sesi terminal |

**Pentajutan automatik** sesi terbina ke dalam pi-web sendiri dan dikonfigurasi pada halaman `/settings`. Ia **hidup secara lalai** dan menamakan sesi secara automatik. Anda boleh memilih:

- **Bila tajukan** — sekali per sesi, atau pada setiap pesan baharu (lalai).
- **Model tajuk** — secara lalai, **heuristik perkataan terbina-dalam (tanpa AI) yang percuma dan serta-merta**, atau pilih model (cth. yang kecil/pantas) untuk tajuk yang lebih bijak, ditulis oleh model.

Pakej ini juga memasang binari pi-web ke `~/.pi/agent/bin/pi-web` dan menyediakan auto-pemulaan pada log masuk.

## Auto-Pemulaan pada Log Masuk

Arahan `pi install npm:@timmygod/pi-web-local` menyetel ini secara automatik:

| Sistem Operasi | Mekanisme |
|----|-----------|
| macOS | plist launchd pada `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | perkhidmatan pengguna systemd pada `~/.config/systemd/user/pi-web.service` |
| Windows | item kunci-Run `HKCU` yang melancarkan per Pemula tersembunyi dalam `~/.config/pi-web/` |

Untuk menetapkan token untuk akses jarak jauh, cipta `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Untuk butiran lanjut (penyediaan manual, port tersuai, pengikatan bukan loopback), lihat [user-docs/install.md](../en/install.md).

## Pembangunan

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Untuk penyegerakan hulu, ujian model tempatan, dan aliran kerja
pelepasan selari, lihat [Pembangunan edisi model tempatan](../../docs/dev/local-llm-development.md).
