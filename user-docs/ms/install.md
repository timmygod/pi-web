# Pemasangan & Penggunaan

## Ciri-ciri

### Kawalan jarak jauh

- Teruskan sesi apa jua dari pelayar web dengan lampiran teks atau imej
- Mula sesi baharu sepenuhnya pada mana-mana laluan projek, terus dari UI web
- Pertukaran model di pelayar web dan pemilih tahap pemikiran, bagi setiap sesi
- Status pekerja bagi setiap sesi (sedia / berjalan / ralat) dengan pulih automatik semasa kerosakan
- Pelbagai sesi berjalan serentak — mulakan kerja dalam satu, saksikan yang lain merentaskan
- `PI_WEB_TOKEN` untuk pendedahan LAN yang selamat — diperlukan secara lalai untuk sebarang ikatan bukan-loopback yang dinyatakan

### Membaca sesi

- Semak sesi merentas projek dengan penapis, carian, dan navigasi cawangan penuh
- Kemas kini inkremental langsung sementara pi masih berjalan (melalui fsnotify; latensi ~ms)
- Mod susulan untuk mengikut sesi aktif
- Pautan dalam ke individu mesej
- Muat turun sesi sebagai JSONL
- Kongsi paparan statik sebagai Gist GitHub rahsia
- Sambungan pi `/web`, `/remote`, `/refresh`, `/pi-web token` dan `/pi-web set-token` untuk membuka sesi, kod QR jarak jauh, selari sesi, dan pengurusan token
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) supaya sesebuah sesi boleh mengurus jadual, papan coret projek, dan tetapan dalam bahasa natural

## Pilih mod sesi

Edisi ini menggunakan pembekal dan model yang telah dikonfigurasi dalam pi; Mod Lokal ialah dasar runtime, bukan penjedi model berasingan atau skrin API-key kedua.
Pilih satu mod apabila membuat sesi, atau ubah selepas larian semasa menetapkan:

| Mod | Gunakan apabila | Kelakuan |
|------|-------------|----------|
| **Auto** | Anda mahu pi-web membuat keputusan | Menyelesaikan hujung lokal/LAN daripada metadata pembekal bila boleh; jika tidak, kekal pada laluan normal |
| **Local** | Model sedang berjalan pada mesin ini atau LAN anda | Membolehkan sempadan pemadatan 65%, titik semak terikat, Force Compact, dan pemulihan automatik berpelindung |
| **Cloud** | Model yang dipilih dihoskan dan mesti mengikuti kelakuan hulu | Menjauhkan dasar pemadatan dan pemulihan khusus-lokal daripada sesi |

Pilihan Local atau Cloud secara manual menang atas pengesanan automatik dan kekal merentas
muat semula dan mulakan semula. Sesi yang sedang berjalan menolak perubahan mod sehingga pekerjanya
menetapkan, jadi mod yang dipaparkan dalam UI sentiasa sepadan dengan dasar yang benar-benar digunakan.

## Persyaratan

- [Go](https://go.dev) 1.25+ (hanya untuk membina dari sumber)
- `pi` pada `PATH` anda untuk sembang/pertukaran model pelayar
- Pilihan: `gh` untuk perkongsian
- Pada Windows: pi memerlukan shell bash untuk alatan shellnya — [Git for Windows](https://git-scm.com/download/win) sudah memadai (lihat dokumen Windows pi)

## Pasang

### Pakej Pi (disyorkan)

```bash
pi install npm:@timmygod/pi-web-local
```

Arahan tunggal ini:
- Memasang pakej pi npm di bawah direktori pakej pi
- Menjalankan skrip `postinstall` pakej (`install.sh`, atau `install.ps1` pada Windows)
- Memuat turun binari pi-web yang sepadan untuk versi pakej dan platform anda dari GitHub Releases
- Memasangnya ke `~/.pi/agent/bin/pi-web` (`pi-web.exe` pada Windows)
- Mengatur mulakan automatik semasa log masuk (launchd pada macOS, systemd pada Linux, pelancaran kunci Run pada Windows)
- Mendaftar arahan pi `/web`, `/remote`, `/refresh`, `/pi-web token`, dan `/pi-web set-token`

Pen tajuk automatik sesi terbina dalam pi-web (bukan sambungan) dan dikonfigurasi pada halaman `/settings`. Ia mengaktif secara lalai: pi-web memberi nama sesi secara automatik menggunakan heuristik bina-dalam percuma (tanpa AI), bertajuk semula pada setiap mesej baharu. Anda boleh bertukar ke tajuk sekali per sesi, dan/atau memilih model untuk menulis tajuk yang lebih bijak berbanding heuristik.

Pada Linux, mulakan automatik dikonfigurasi sebagai perkhidmatan systemd pengguna di `~/.config/systemd/user/pi-web.service`. Perpasang menulis semula `ExecStart` kepada laluan binari yang benar-benar dipasang. Jika Tailscale tersedia semasa runtime, pi-web menerbitkan server localhost dengan Tailscale Serve HTTPS. Jika systemd pengguna tidak tersedia, jalankannya secara manual dengan `~/.pi/agent/bin/pi-web -o`.

Untuk memasang hanya untuk projek tertentu (dikongsi dengan pasukan anda melalui `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Kemudian mulakan semula pi (atau jalankan `/reload`), dan gunakan `/web`, `/pi-web`, `/remote`, `/refresh`. Urus token akses anda dengan `/pi-web token` dan `/pi-web set-token`.

Jika npm terhenti dengan `ENOTEMPTY` semasa menaip semula `@timmygod/pi-web-local`, buang direktori sandaran tersembunyi lalai npm dan pasang semula pakej:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Pemasangan pantas (tidak perlu alatan bina)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Ini memuat turun binari pi-web terbaharu, memasangnya ke `/usr/local/bin` (`~/.pi/agent/bin` pada Windows), dan mengatur mulakan automatik semasa log masuk. Tidak perlu Go, Node, atau pi.

### Muat turun binari

Binari ter-bina disertakan pada setiap [GitHub Release](https://github.com/timmygod/pi-web/releases).

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

Kemudah pindahkan ia ke PATH anda:

```bash
cp pi-web ~/.pi/agent/bin/
# atau sistem-lembang:
sudo cp pi-web /usr/local/bin/
```

### Bina dari sumber

Semakan keluar ini ialah edisi model-lokal bagi pi-web. Bina normal menghasilkan
aplikasi web dan latar belakang bersama-sama; penjagaan model-lokal diaktifkan
di runtime oleh Mod Lokal berkesan sesi, bukan oleh binari berasingan.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # membina talian Vite, kemudian membenangkannya ke dalam binari Go

# pilihan: letakkannya pada PATH
cp pi-web ~/.pi/agent/bin/
```

Talian hadapan dibenangi oleh `web/assets_embed.go`, jadi `go build` memerlukan
`web/dist` wujud dahulu. `make build` melakukan kedua-dua langkah mengikut urutan; jika anda bina
dengan tangan, jalankan `npm --prefix web install && npm --prefix web run build` sebelum
`go build ./cmd/pi-web`.

Untuk aliran bercabang yang dikekalkan, selari hulu, dan senarai semak pengesahan Mod Lokal, lihat [catatan pembangunan model-lokal](../../docs/dev/local-llm-development.md).

### Buka bersama-sama dengan contoh yang dipasang

Biarkan contoh yang dipasang terus berjalan pada port `31415`, kemudian mulakan semakan keluar
sumber dalam mod pembangunan:

```bash
make dev
```

Buka `http://127.0.0.1:31416`. `make dev` menetapkan persekitaran pembangunan internal `PI_WEB_DEV=1`,
jadi semakan keluar sumber berkongsi sesi, tetapan, dan data SQLite dengan contoh yang dipasang sambil
memelihara kunci runtime dan fail keadaan pembangunan berasingan. Contoh yang dipasang secara biasa dan
dilancarkan secara manual tidak berubah dan memelihara kelakuan satu-contoh asal.

Untuk mengelakkan kerja autonom yang berganda, mod pembangunan tidak menjalankan
gelung jadual, pengosong gilir sembang, pen tajuk automatik, atau notifikasi tolak. Permintaan terus
yang dibuat melalui UI pembangunan masih berfungsi. Jangan kendalikan sesi sembang yang sama dari kedua-dua
contoh serentak; setiap proses mempunyai pengurus pekerja RPC sendiri.

`make dev` memerlukan [Air](https://github.com/air-verse/air) untuk muat semula panas Go:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` ialah talian rangka pembangunan, bukan mod pelbagai-contoh terbiak
untuk pengeluaran.

## Lepas pasang

```bash
pi remove npm:@timmygod/pi-web-local
```

Ini menjalankan skrip `preuninstall` pakej (`uninstall.sh`, atau `uninstall.ps1`
pada Windows), yang menghentikan contoh yang berjalan dan membuang:

- binari pi-web (`~/.pi/agent/bin/pi-web`, atau `/usr/local/bin/pi-web` untuk pemasangan berdikari)
- fail versi (`~/.pi/agent/pi-web-version`)
- fail keadaan runtime (`~/.pi/agent/pi-web/pi-web-state.json`)
- tetapan mulakan automatik (plist launchd pada macOS, perkhidmatan pengguna systemd pada Linux, entri kunci Run + skrip pelancar pada Windows)

Data anda dipelihara supaya pemasangan semula kemudian sambung dari mana anda berhenti:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, fail sesi
anda di bawah `~/.pi/agent/sessions/`, dan `~/.config/pi-web/env` (termasuk
`PI_WEB_TOKEN`). Buang mereka secara manual jika anda mahu permulaan bersih.

## Penggunaan

```bash
# Mula pada port lalai (31415)
pi-web

# Mula dan buka pelayar
pi-web -o

# Port tersuai
pi-web -p 8080

# Guna ganti hos ikatan (loopback tidak diaktifkan secara lalai)
pi-web --host 127.0.0.1

# Ikatan bukan-loopback memerlukan token — pi-web akan menolak bermula jika tidak
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Secara lalai, pi-web mengikat kepada `127.0.0.1`. Jika Tailscale berjalan dengan MagicDNS **dan `PI_WEB_TOKEN` ditetapkan**, pi-web juga menjalankan `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` dan memaparkan URL HTTPS tailnet. Tanpa token, pi-web kekal khusus-loopback dan skip Tailscale Serve, jadi rakan tailnet tidak dapat mencapai ejen tanpa pengesahan. Sebarang ikatan bukan-loopback yang dinyatakan juga memerlukan `PI_WEB_TOKEN` ditetapkan; luluskan `--insecure` untuk guna ganti untuk ujian lokal.

## Akses Jarak Jauh

Biarkan pi-web mendengar secara lokal, kemudian gunakan URL HTTPS Tailscale yang dipaparkan dari telefon atau komputer riba anda pada tailnet.

Pada macOS, pasang dan buka Tailscale secara interaktif, sahkan prompt pentadbir, dan log masuk. Kemudian jalankan `/pi-web restart`, diikuti dengan `/remote`.

Pada Linux, benarkan pengguna anda mengurus Tailscale sebelum memasang/menjalankan pi-web, jika tidak `tailscale serve` mungkin memerlukan sudo dan mulakan automatik boleh gagal:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Mula pi-web dengan token supaya ia menerbitkan hujung HTTPS Tailscale
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Dari mana-mana peranti Tailscale-terhubung lain, buka
#    URL "Tailscale HTTPS" yang dipaparkan dan masukkan token sekali.
```

> Secara lalai, pi-web menolak mengikat kepada alamat bukan-loopback melainkan `PI_WEB_TOKEN` ditetapkan — sesiapa yang boleh mencapai alamat yang diikat boleh melihat sesi dan menghantar arahan ke pi jika tidak. Untuk guna ganti penjagaan ini untuk ujian rangkaian lokal, luluskan `--insecure`. **Jangan gunakan `--insecure` pada Tailscale atau mana-mana alamat yang boleh dicapai dari luar mesin anda.**
>
> Klien boleh luluskan token melalui pengepala `Authorization: Bearer <token>`, pengepala `X-Pi-Token`, atau sekali melalui `?token=<token>` (yang menetapkan cookie `pi_token` untuk permintaan susulan). Token yang diluluskan melalui `?token=` berakhir dalam sejarah pelayar, log akses pelayan, dan pengepala `Referer` daripada mana-mana pautan pada halaman — lebih baik bentuk pengepala untuk apa-apa di luar penanda buku awal.

## Sembang Pelayar

Buka halaman sesi dan gunakan penggubah di bahagian bawah untuk meneruskan sesi tepat itu.

- `Enter` menghantar, `Shift+Enter` menyisip baris baru
- Seret dan lepaskan atau tampal imej terus ke dalam penggubah
- Pemilih model dan pemilih tahap pemikiran terletak pada pengepala — perubahan digunakan kepada pekerja pi asas serta-merta
- Setiap sesi aktif mendapat pekerja `pi --mode rpc` khusus sendiri, jadi sesi berbeza tidak memblok antara satu sama lain

## Kongsi Sesi

Klik **Kongsi** pada halaman sesi untuk mencipta Gist GitHub rahsia.

Persyaratan:
- `gh` dipasang
- `gh auth login` selesai

Kongsi memulangkan:
- URL gist rahsia
- URL pratinjau pada `https://pi.dev/session/#<gistId>`

Gist berkongsi ialah paparan dan tidak dikemas kini secara langsung.

## Mulakan Automatik semasa Log Masuk

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Pasang perkhidmatan pengguna systemd
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Pilihan: tetapkan PI_WEB_TOKEN anda untuk ikatan bukan-loopback
# (atau gunakan /pi-web set-token <token> dari dalam pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Aktifkan dan mula
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Semak status
systemctl --user status pi-web.service

# Lihat log
journalctl --user -u pi-web.service -f
```

> Untuk perkhidmatan bermula semasa but (sebelum log masuk), gunakan perkhidmatan sistem sebaliknya:
> salin `init/pi-web.service` ke `/etc/systemd/system/` dan gunakan `sudo systemctl`.

### Windows

Perpasang mengonfigurasi ini secara automatik, tanpa memerlukan hak pentadbir: entri
`pi-web` di bawah `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
melancarkan `~/.config/pi-web/pi-web-start.vbs` semasa log masuk, yang memulakan binari
bersembunyi (tiada tetingkap konsol) selepas memuat `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...).

Untuk uruskannya dengan tangan:

```powershell
# Mula / henti
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Buang mulakan automatik
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Tiada pengawasan perkhidmatan pada Windows: jika pi-web rosak ia kekal terhenti
sehingga log masuk seterusnya (launchd/systemd memulakan semula ia secara automatik pada
platform lain).
