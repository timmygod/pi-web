# Mengapa pi-web?

Saya agak kecanduan Claude Code. Saya selalu menggunakannya. Jika saya tidak sedang duduk di depan komputer, saya memikirkannya. Saya merasa seperti tidak membakar cukup token. Itu adalah masa-masa awal Claude Code. Dan saya berpikir, mengapa saya tidak bisa melanjutkan dari ponsel saya? Saya menyiapkan Termius dan saya tidak terlalu menyukainya.

Saya mulai membuat milik saya sendiri dan berhenti ketika Claude memperkenalkan aplikasi seluler Claude Code mereka.

Kemudian saya mengalami herniated disc dan saya benar-benar tidak bisa melakukan banyak hal. Waktu berlalu dan saya merasa sedikit pulih dan saya ingin melanjutkan proyek Claude Code via web/pwa saya.

Lalu Claude Code mulai melarang penggunaan di luar harness mereka sendiri. Dan saya merasa itu tidak sepadan.

Kemudian saya menemukan pi.dev dan menjelajah sedikit tetapi belum benar-benar mendalaminya. Saya membaca tentangnya, menonton video tentangnya dan memutuskan untuk mencobanya sepenuhnya dan sekarang saya sepenuhnya menyukai pi.

Karena ini open source, saya merasa ini layak untuk dibangun. Saya juga mendapatkan berbagai pilihan provider. Saya juga merasa bahwa bergantung pada satu provider/model seperti Anthropic/Claude tidaklah berkelanjutan.

Jadi saya membangunnya di sini.

## Mengapa model lokal memerlukan profil operasi yang berbeda

Pengalaman pi-web asli adalah fondasi yang sangat baik, tetapi inferensi lokal memiliki mode kegagalan yang berbeda dari model yang di-hosting secara tipikal. Model lokal dapat melambat secara tajam saat konteks bertambah, berbagi memori terbatas dengan sisa mesin, berhenti setelah hanya menghasilkan penalaran, atau kehilangan proses panjang karena kegagalan transportasi lokal yang sementara. Memperlakukan kasus-kasus tersebut persis seperti kegagalan cloud membuat UI terlihat kompatibel sementara sesi aktual tetap rapuh.

Edisi ini mendekati masalah secara berlapis:

1. **Pertahankan upstream terlebih dahulu.** UI bersama dan perilaku sesi terus berasal dari pi-web; perubahan lokal diisolasi di balik Local Mode yang efektif.
2. **Cegah sebelum memulihkan.** Batas konteks 65% berbasis persentase ditegakkan sebelum panggilan penyedia berikutnya, termasuk panggilan di dalam loop alat yang panjang.
3. **Pulihkan hanya dengan bukti.** Kelanjutan otomatis dibatasi pada insiden konteks, transportasi, dan hanya-pemikiran yang dikenali, bukan kesalahan autentikasi, kuota, atau kesalahan penyedia sembarangan.
4. **Batasi setiap tindakan otonom.** Insiden pemulihan dideduplikasi, kemajuan diperlukan sebelum penyelamatan lain, dan startup mempertimbangkan paling banyak satu sesi Local yang aktif baru-baru ini.
5. **Pertahankan jalan keluar manual.** Force Compact merangkum alih-alih menghapus riwayat, sehingga pengguna dapat menyelamatkan sesi tanpa berpura-pura konteks tidak pernah ada.
6. **Lindungi kompatibilitas cloud.** Cloud Mode mempertahankan semantik dan kontrol upstream; optimasi model lokal tidak secara diam-diam mendefinisikan ulang sesi cloud.

Itulah perbedaan nyata dalam fork ini: ia memperlakukan inferensi lokal sebagai lingkungan operasional yang berbeda, bukan sekadar nama model lain dalam dropdown.
