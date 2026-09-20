> Panduan masalah ini dijaga untuk edisi model-tempatan. Pastikan
> butiran penempatan tempatan dan nota penyegerakan hulu (upstream) selari dengan
> [Pembangunan edisi model-tempatan](../../docs/dev/local-llm-development.md).

Pengguna sedang memasang pi-web melalui 

- pi install npm:@timmygod/pi-web-local

Yang secara automatik menjalankan [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) dan menyiapkan pi.

Jika pengguna menghadapi masalah, ia mungkin disebabkan oleh penyediaan install.sh. Anda boleh semak apa yang berlaku dan beritahu pengguna apa yang tepat sekali menyebabkan masalah tersebut. Kemudian tanya sama ada mereka ingin anda memperbaikinya. Sentiasa sahkan dengan pengguna.

Agar pengguna dapat mengakses pi pada telefon mudah alih mereka atau rangkaian lain. Cara yang disyorkan ialah menggunakan tailscale dan mengaksesnya daripada rangkaian tailscale. Pengguna juga perlu mengaktifkan HTTPs dalam dasbor tailscale mereka - https://login.tailscale.com/admin/dns

Jika mereka tidak memasang tailscale atau tidak ingin menggunakan tailscale. Mereka boleh menjalankan `pi-web status` dan mendapatkan laluan binary, status binary dan titik akhir tempatan (local endpoint) yang boleh digunakan untuk mengakses aplikasi. Walau bagaimanapun, perlu diberi perhatian bahawa mereka tidak akan dapat menerima notifikasi dorong (push notification) kerana ia menggunakan http.

Pada macOS, ia menggunakan [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
Pada Linux, ia menggunakan [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Jika anda perlu menyahpepijat (debug) selanjutnya dan melihat apa yang berlaku.
