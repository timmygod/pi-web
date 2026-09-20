# Instalasi & Penggunaan

## Fitur

### Kontrol jarak jauh

- Lanjutkan sesi apa pun dari browser dengan lampiran teks atau gambar
- Mulai sesi baru sepenuhnya di jalur proyek mana pun, langsung dari antarmuka web
- Pemilih model di browser dan pemilih tingkat berpikir, per sesi
- Status worker per sesi (idle / running / error) dengan pemulihan otomatis saat crash
- Beberapa sesi berjalan paralel — mulai pekerjaan di satu, saksikan sesi lain yang sedang streaming
- `PI_WEB_TOKEN` untuk eksposur LAN yang aman — diwajibkan secara bawaan untuk semua bind non-loopback eksplisit

### Membaca sesi

- Jelajahi sesi antar proyek dengan filter, pencarian, dan navigasi cabang penuh
- Pembaruan inkremental langsung saat pi masih berjalan (via fsnotify; latensi ~ms)
- Mode ikuti untuk memantau sesi aktif
- Tautan langsung ke pesan individu
- Unduh sesi sebagai JSONL
- Bagikan snapshot statis sebagai Gist GitHub rahasia
- Ekstensi pi `/web`, `/remote`, `/refresh`, `/pi-web token` dan `/pi-web set-token` untuk membuka sesi, QR jarak jauh, sinkronisasi sesi, dan manajemen token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) sehingga sesi dapat mengelola jadwal, catatan sementara proyek, dan pengaturan dalam bahasa alami

## Pilih mode sesi

Edisi ini menggunakan penyedia dan model yang sudah dikonfigurasi di pi; Local Mode
adalah kebijakan runtime, bukan installer model terpisah atau layar API-key kedua.
Pilih mode saat membuat sesi, atau ubah setelah jalannya saat ini mereda:

| Mode | Gunakan saat | Perilaku |
|------|-------------|----------|
| **Auto** | Anda ingin pi-web yang memutuskan | Menyelesaikan endpoint lokal/LAN dari metadata penyedia bila memungkinkan; jika tidak, mempertahankan jalur normal |
| **Local** | Model berjalan di mesin ini atau LAN Anda | Mengaktifkan batas kompresi 65%, checkpoint terbatas, Force Compact, dan pemulihan otomatis berjaga |
| **Cloud** | Model yang dipilih di-host dan harus mengikuti perilaku upstream | Menjaga kebijakan kompresi dan pemulihan hanya-lokal keluar dari sesi |

Pemilihan Local atau Cloud manual menang atas deteksi otomatis dan bertahan
melalui reload dan restart. Sesi yang sedang berjalan menolak perubahan mode hingga worker-nya
mereda, sehingga mode yang ditampilkan di UI selalu sesuai dengan kebijakan yang benar-benar digunakan.

## Persyaratan

- [Go](https://go.dev) 1.25+ (hanya untuk membangun dari sumber)
- `pi` di `PATH` Anda untuk pemilih model chat browser
- Opsional: `gh` untuk berbagi
- Di Windows: pi memerlukan shell bash untuk alat shell-nya — [Git for Windows](https://git-scm.com/download/win) sudah cukup (lihat dokumentasi Windows pi)

## Instalasi

### Paket Pi (disarankan)

```bash
pi install npm:@timmygod/pi-web-local
```

Perintah tunggal ini:
- Menginstal paket pi npm di bawah direktori paket pi
- Menjalankan skrip `postinstall` paket (`install.sh`, atau `install.ps1` di Windows)
- Mengunduh biner pi-web yang sesuai dengan versi paket dan platform Anda dari GitHub Releases
- Menginstalnya ke `~/.pi/agent/bin/pi-web` (`pi-web.exe` di Windows)
- Menyiapkan mulai-otomatis saat login (launchd di macOS, systemd di Linux, peluncur kunci Run di Windows)
- Mendaftarkan perintah pi `/web`, `/remote`, `/refresh`, `/pi-web token`, dan `/pi-web set-token`

Penamaan otomatis sesi dibangun ke dalam pi-web (bukan ekstensi) dan dikonfigurasi di halaman `/settings`. Aktif secara bawaan: pi-web memberi nama sesi secara otomatis menggunakan heuristik kata bawaan gratis (tanpa AI), menamai ulang di setiap pesan baru. Anda dapat beralih ke penamaan sekali per sesi, dan/atau memilih model untuk menulis judul yang lebih cerdas alih-alih heuristik.

Di Linux, mulai-otomatis dikonfigurasi sebagai layanan systemd pengguna di `~/.config/systemd/user/pi-web.service`. Penginstal menulis ulang `ExecStart`-nya ke jalur biner terpasang yang sebenarnya. Jika Tailscale tersedia saat runtime, pi-web menerbitkan server localhost dengan Tailscale Serve HTTPS. Jika systemd pengguna tidak tersedia, jalankan secara manual dengan `~/.pi/agent/bin/pi-web -o`.

Untuk menginstal hanya untuk proyek tertentu (dibagikan dengan tim Anda melalui `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Lalu restart pi (atau jalankan `/reload`), dan gunakan `/web`, `/pi-web`, `/remote`, `/refresh`. Kelola token akses Anda dengan `/pi-web token` dan `/pi-web set-token`.

Jika npm berhenti dengan `ENOTEMPTY` saat mengganti nama `@timmygod/pi-web-local`, hapus direktori cadangan tersembunyi npm yang usang dan instal ulang paketnya:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Instalasi cepat (tidak memerlukan alat build)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Ini mengunduh biner pi-web terbaru, menginstalnya ke `/usr/local/bin` (`~/.pi/agent/bin` di Windows), dan menyiapkan mulai-otomatis saat login. Tidak memerlukan Go, Node, atau pi.

### Unduh biner

Biner pra-bangun terlampir pada setiap [GitHub Release](https://github.com/timmygod/pi-web/releases).

```bash
# macOS (Apple Silicon)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-arm64
chmod +x pi-web

# macOS (Intel)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-amd64
chmod +x pi-web

# Linux (amd64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-amd64
chmod +x pi-web

# Linux (arm64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-arm64
chmod +x pi-web
```

```powershell
# Windows (x64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-amd64.exe

# Windows (ARM64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-arm64.exe
```

Lalu pindahkan ke PATH Anda:

```bash
cp pi-web ~/.pi/agent/bin/
# atau sistem-secara:
sudo cp pi-web /usr/local/bin/
```

### Bangun dari sumber

Checkout ini adalah edisi model-lokal dari pi-web. Build normal menghasilkan
aplikasi web dan backend bersama-sama; penjagaan model-lokal diaktifkan
di runtime oleh Local Mode efektif dari sesi, bukan oleh biner terpisah.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # membangun bundle Vite, lalu menyematkannya ke dalam biner Go

# opsional: taruh di PATH
cp pi-web ~/.pi/agent/bin/
```

Bundle frontend disematkan oleh `web/assets_embed.go`, sehingga `go build` memerlukan
`web/dist` ada terlebih dahulu. `make build` melakukan kedua langkah secara berurutan; jika Anda membangun
dengan tangan, jalankan `npm --prefix web install && npm --prefix web run build` sebelum
`go build ./cmd/pi-web`.

Untuk alur kerja fork yang dipelihara, sinkronisasi upstream, dan daftar periksa
verifikasi Local Mode, lihat [catatan pengembangan model-lokal](../../docs/dev/local-llm-development.md).

### Berkembang berdampingan dengan instance terpasang

Biarkan instance terpasang tetap berjalan di port `31415`, lalu mulai checkout
sumber dalam mode pengembangan:

```bash
make dev
```

Buka `http://127.0.0.1:31416`. `make dev` menetapkan lingkungan pengembangan
internal `PI_WEB_DEV=1`, sehingga checkout sumber berbagi sesi, pengaturan, dan
data SQLite dengan instance terpasang sambil mempertahankan kunci runtime pengembangan
dan file status terpisah. Instance yang terpasang secara biasa dan diluncurkan secara manual
tidak berubah dan mempertahankan perilaku instance-tunggal aslinya.

Untuk mencegah pekerjaan otonom yang duplikat, mode pengembangan tidak menjalankan
loop jadwal, pemvakum antrean chat, penamaan otomatis, atau notifikasi push. Permintaan
langsung yang dilakukan melalui UI pengembangan tetap berfungsi. Jangan mengendarai sesi
chat yang sama dari kedua instance sekaligus; setiap proses memiliki manajer
worker RPC-nya sendiri.

`make dev` memerlukan [Air](https://github.com/air-verse/air) untuk hot reload Go:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` adalah perpipaan kerangka pengembangan, bukan mode
multi-instance produksi yang didukung.

## Penghapusan

```bash
pi remove npm:@timmygod/pi-web-local
```

Ini menjalankan skrip `preuninstall` paket (`uninstall.sh`, atau `uninstall.ps1`
di Windows), yang menghentikan instance yang berjalan dan menghapus:

- biner pi-web (`~/.pi/agent/bin/pi-web`, atau `/usr/local/bin/pi-web` untuk instalasi mandiri)
- file versi (`~/.pi/agent/pi-web-version`)
- file status runtime (`~/.pi/agent/pi-web/pi-web-state.json`)
- konfigurasi mulai-otomatis (plist launchd di macOS, layanan pengguna systemd di Linux, entri kunci Run + skrip peluncur di Windows)

Data Anda dipertahankan sehingga instalasi ulang berikutnya mengambil tempat di mana Anda berhenti:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, file sesi
Anda di bawah `~/.pi/agent/sessions/`, dan `~/.config/pi-web/env` (termasuk
`PI_WEB_TOKEN`). Hapus itu secara manual jika Anda ingin mulai dari nol.

## Penggunaan

```bash
# Mulai di port bawaan (31415)
pi-web

# Mulai dan buka browser
pi-web -o

# Port khusus
pi-web -p 8080

# Timpa host bind (loopback tidak terautentikasi secara bawaan)
pi-web --host 127.0.0.1

# Bind non-loopback memerlukan token — pi-web menolak untuk memulai jika tidak
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Secara bawaan, pi-web mengikat ke `127.0.0.1`. Jika Tailscale berjalan dengan MagicDNS **dan `PI_WEB_TOKEN` ditetapkan**, pi-web juga menjalankan `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` dan mencetak URL HTTPS tailnet. Tanpa token, pi-web tetap hanya-loopback dan melompati Tailscale Serve, sehingga peer tailnet tidak dapat mencapai agen tanpa autentikasi. Bind non-loopback eksplisit apa pun juga memerlukan `PI_WEB_TOKEN` ditetapkan; lewati `--insecure` untuk menimpanya untuk pengujian lokal.

## Akses Jarak Jauh

Biarkan pi-web mendengarkan secara lokal, lalu gunakan URL HTTPS Tailscale yang tercetak dari ponsel atau laptop Anda di tailnet.

Di macOS, instal dan buka Tailscale secara interaktif, setujui prompt administrator, dan masuk. Lalu jalankan `/pi-web restart`, diikuti dengan `/remote`.

Di Linux, izinkan pengguna Anda mengelola Tailscale sebelum menginstal/menjalankan pi-web, jika tidak `tailscale serve` mungkin memerlukan sudo dan mulai-otomatis dapat gagal:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Mulai pi-web dengan token agar menerbitkan endpoint HTTPS Tailscale
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Dari perangkat terkoneksi Tailscale lainnya, buka
#    URL "Tailscale HTTPS" yang tercetak dan masukkan token sekali.
```

> Secara bawaan, pi-web menolak mengikat ke alamat non-loopback kecuali `PI_WEB_TOKEN` ditetapkan — siapa pun yang dapat mencapai alamat yang terikat dapat dengan demikian melihat sesi dan mengirim instruksi ke pi. Untuk menimpanya penjagaan ini untuk pengujian jaringan-lokal, lewati `--insecure`. **Jangan gunakan `--insecure` di Tailscale atau alamat apa pun yang dapat dicapai dari luar mesin Anda.**
>
> Klien dapat melampirkan token melalui header `Authorization: Bearer <token>`, header `X-Pi-Token`, atau sekali melalui `?token=<token>` (yang menetapkan cookie `pi_token` untuk permintaan berikutnya). Token yang dilewatkan melalui `?token=` berakhir di riwayat browser, log akses server, dan header `Referer` dari tautan apa pun di halaman — lebih suka bentuk header untuk apa pun di luar bookmark awal.

## Chat Browser

Buka halaman sesi dan gunakan composer di bagian bawah untuk melanjutkan sesi yang tepat itu.

- `Enter` mengirim, `Shift+Enter` menyisipkan baris baru
- Tarik-lepas atau tempel gambar langsung ke dalam composer
- Pemilih model dan pemilih tingkat berpikir berada di header — perubahan diterapkan ke worker pi yang mendasari secara langsung
- Setiap sesi aktif mendapat worker `pi --mode rpc` khususnya sendiri, sehingga sesi yang berbeda tidak saling memblokir

## Berbagi Sesi

Klik **Share** di halaman sesi untuk membuat Gist GitHub rahasia.

Persyaratan:
- `gh` terinstal
- `gh auth login` selesai

Berbagi mengembalikan:
- URL gist rahasia
- URL pratinjau di `https://pi.dev/session/#<gistId>`

Gist yang dibagikan adalah snapshot dan tidak diperbarui secara langsung.

## Mulai-Otomatis saat Login

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Instal layanan pengguna systemd
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Opsional: set PI_WEB_TOKEN Anda untuk bind non-loopback
# (atau gunakan /pi-web set-token <token> dari dalam pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Aktifkan dan mulai
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Periksa status
systemctl --user status pi-web.service

# Lihat log
journalctl --user -u pi-web.service -f
```

> Agar layanan mulai saat boot (sebelum login), gunakan layanan sistem alih-alih:
> salin `init/pi-web.service` ke `/etc/systemd/system/` dan gunakan `sudo systemctl`.

### Windows

Penginstal mengonfigurasi ini secara otomatis, tanpa memerlukan hak admin: entri
`pi-web` di bawah `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
meluncurkan `~/.config/pi-web/pi-web-start.vbs` saat login, yang memulai biner
tersembunyi (tanpa jendela konsol) setelah memuat `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...).

Untuk mengelolanya dengan tangan:

```powershell
# Mulai / hentikan
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Hapus mulai-otomatis
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Tidak ada pengawasan layanan di Windows: jika pi-web crash, ia tetap mati
hingga login berikutnya (launchd/systemd menjalankannya ulang secara otomatis di platform
lainnya).
