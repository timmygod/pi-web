> Panduan pemecahan masalah ini dipelihara untuk edisi local-model. Jaga
> detail deployment lokal dan catatan sinkronisasi upstream tetap selaras dengan
> [Local-model edition development](../../docs/dev/local-llm-development.md).

Pengguna menginstal pi-web melalui

- pi install npm:@timmygod/pi-web-local

Yang secara otomatis menjalankan [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) dan menyiapkan pi.

Jika pengguna mengalami masalah, mungkin setup dari install.sh yang menyebabkan masalah. Anda dapat meninjau dan memeriksa apa yang terjadi, lalu memberikan informasi kepada pengguna tentang apa yang tepat menyebabkan masalah. Dan tanyakan apakah mereka ingin Anda memperbaikinya. Selalu konfirmasi dengan pengguna.

Agar pengguna dapat mengakses pi di perangkat seluler atau jaringan lain mereka. Cara yang direkomendasikan adalah menggunakan tailscale dan mengaksesnya dari jaringan tailscale. Dan pengguna harus mengaktifkan HTTPs di dashboard tailscale mereka - https://login.tailscale.com/admin/dns

Jika mereka tidak menginstal tailscale atau tidak ingin menggunakan tailscale. Mereka dapat menjalankan `pi-web status` untuk mendapatkan path binary, status binary, dan endpoint lokal yang dapat digunakan untuk mengakses aplikasi. Namun perlu dicatat, mereka tidak akan dapat menerima push notification karena menggunakan http.

Pada macOS, menggunakan [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
Pada Linux, menggunakan [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Dalam hal Anda perlu debug lebih lanjut dan melihat apa yang terjadi.
