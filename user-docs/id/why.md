# Kenapa pi-web?

Saya cukup kecanduan Claude Code. Saya selalu menggunakannya. Kalau saya tidak duduk di depan komputer, saya memikirkannya. Saya merasa seperti saya tidak membakar cukup token. Itu hari-hari awal Claude Code. Dan saya berpikir, kenapa saya tidak bisa melanjutkan dari ponsel saya? Saya mengatur Termius dan saya tidak benar-benar menyukainya.

Saya mulai membuat yang milik saya sendiri dan berhenti ketika Claude memperkenalkan aplikasi seluler Claude Code mereka.

Lalu saya menderita herniated disc dan saya tidak benar-benar bisa melakukan banyak hal. Waktu berlalu dan saya merasa pulih sedikit dan saya ingin melanjutkan proyek Claude Code via web/pwa saya.

Lalu Claude Code mulai melarang penggunaan di luar harness mereka sendiri. Dan saya merasa itu tidak sepadan.

Lalu saya menemukan pi.dev dan menjelajah sedikit tetapi belum benar-benar mendalami. Saya membacanya, menonton video tentangnya dan memutuskan untuk mencoba sepenuhnya dan sekarang saya benar-benar menyukai pi.

Karena open source, saya merasa itu layak untuk dibangun. Saya juga mendapat pilihan provider yang berbeda. Saya juga merasa seperti mengandalkan satu provider/model seperti Anthropic/Claude tidak berkelanjutan.

Jadi saya membangunnya di sini.

Checkout ini dipelihara sebagai edisi model lokal dari pi-web. Ini mengikuti proyek upstream untuk peningkatan bersama, sambil menjaga deployment lokal, stabilitas konteks, dan pengujian model lokal di jalur rilis terpisah.

## Kenapa model lokal memerlukan profil operasi yang berbeda

Pengalaman pi-web asli adalah fondasi yang sangat baik, tetapi inferensi lokal memiliki mode kegagalan yang berbeda dari model hosted tipikal. Model lokal mungkin melambat tajam seiring konteks tumbuh, berbagi memori terbatas dengan sisa mesin, berhenti setelah hanya menghasilkan reasoning, atau kehilangan jalanan panjang karena kegagalan transport lokal transien. Memperlakukan kasus-kasus tersebut tepat seperti kegagalan cloud membuat UI terlihat kompatibel sementara sesi aktual tetap rapuh.

Edisi ini menangani masalah secara berlapis:

1. **Jaga upstream terlebih dahulu.** UI bersama dan perilaku sesi terus berasal dari pi-web; perubahan lokal terisolasi di balik Local Mode yang efektif.
2. **Mencegah sebelum pulih.** Batas konteks berbasis persentase 65% dipaksakan sebelum pemanggilan model berikutnya, termasuk pemanggilan di dalam loop tool yang panjang.
3. **Pulih hanya dengan bukti.** Lanjutan otomatis dibatasi untuk insiden konteks, transport, dan thinking-only yang dikenali—bukan autentikasi, kuota, atau error provider sewenang-wenang.
4. **Batasi setiap aksi otonom.** Insiden pemulihan dideduplikasi, kemajuan diperlukan sebelum penyelamatan berikutnya, dan startup mempertimbangkan paling banyak satu sesi Local yang baru saja aktif.
5. **Simpan jalan keluar manual.** Force Compact merangkum daripada menghapus riwayat, sehingga pengguna dapat menyelamatkan sesi tanpa berpura-pura konteks tidak pernah ada.
6. **Lindungi kompatibilitas cloud.** Cloud Mode mempertahankan semantik dan kontrol upstream; optimasi model lokal tidak mendefinisikan ulang sesi cloud secara diam-diam.

Itulah perbedaan nyata di fork ini: ia memperlakukan inferensi lokal sebagai lingkungan operasional yang berbeda, bukan sekadar nama model lain di dropdown.
