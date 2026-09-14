# Kenapa pi-web?

Saya agak ketagih dengan Claude Code. Saya sentiasa menggunakannya. Jika saya tidak duduk di hadapan komputer, saya memikirkannya. Saya rasa seperti tidak membakar token yang mencukupi. Ia adalah hari-hari awal Claude Code. Dan saya berfikir, kenapa saya tidak boleh menyambung dari telefon saya? Saya sediakan Termius dan saya tidak begitu menyukainya.

Saya mula mencipta sendiri dan berhenti apabila Claude memperkenalkan aplikasi mudah alih Claude Code mereka.

Kemudian saya mendapat cakera hernia dan saya tidak boleh melakukan apa-apa sangat. Masa berlalu dan saya berasa pulih sedikit dan saya mahu meneruskan projek Claude Code melalui web/pwa saya.

Kemudian Claude Code mula melarang penggunaan di luar harness mereka sendiri. Dan saya rasa ia tidak berbaloi.

Kemudian saya menemui pi.dev dan meneroka sedikit tetapi belum benar-benar mendalaminya. Saya membaca tentangnya, menonton video tentangnya dan memutuskan untuk mencuba sepenuhnya dan sekarang saya benar-benar minat dengan pi.

Oleh kerana ia sumber terbuka, saya rasa ia berbaloi untuk dibangunkan. Saya mendapat pilihan pembekal yang berbeza juga. Saya juga rasa bergantung kepada satu pembekal/model seperti Anthropic/Claude adalah tidak mampan.

Jadi saya membinanya di sini.

## Mengapa model tempatan memerlukan profil operasi yang berbeza

Pengalaman pi-web asal adalah asas yang sangat baik, tetapi inferens tempatan mempunyai mod kegagalan yang berbeza daripada model hos yang tipikal. Model tempatan mungkin menjadi perlahan dengan ketara apabila konteks berkembang, berkongsi memori terhad dengan bahagian lain mesin, berhenti selepas hanya menghasilkan penaakulan, atau kehilangan larian panjang akibat kegagalan pengangkutan tempatan sementara. Melayan kes-kes tersebut tepat seperti kegagalan cloud menjadikan UI kelihatan serasi sementara sesi sebenar kekal rapuh.

Edisi ini menghampiri masalah secara berlapis:

1. **Kekalkan upstream dahulu.** UI bersama dan tingkah laku sesi terus datang dari pi-web; perubahan tempatan diasingkan di belakang Local Mode yang berkesan.
2. **Cegah sebelum pulih.** Sempadan konteks 65% berasaskan peratusan dikuatkuasakan sebelum panggilan pembekal seterusnya, termasuk panggilan dalam gelung alat yang panjang.
3. **Pulih hanya dengan bukti.** Sambungan automatik terhad kepada insiden konteks, pengangkutan, dan fikiran-sahaja yang diiktiraf, bukan ralat pengesahan, kuota, atau ralat pembekal sewenang-wenangnya.
4. **Hadkan setiap tindakan autonomi.** Insiden pemulihan dideduplikasi, kemajuan diperlukan sebelum penyelamatan lain, dan permulaan mempertimbangkan paling banyak satu sesi Local yang aktif baru-baru ini.
5. **Kekalkan keluar manual.** Force Compact merumuskan dan bukan memadam sejarah, supaya pengguna boleh menyelamatkan sesi tanpa berpura-pura konteks tidak pernah wujud.
6. **Lindungi keserasian cloud.** Cloud Mode mengekalkan semantik dan kawalan upstream; pengoptimuman model tempatan tidak secara senyap mendefinisikan semula sesi cloud.

Itulah perbezaan sebenar dalam fork ini: ia melayan inferens tempatan sebagai persekitaran operasi yang berbeza, bukan sekadar satu lagi nama model dalam senarai jatuh.
