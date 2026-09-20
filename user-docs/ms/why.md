# Kenapa pi-web?

Saya agak ketagih dengan Claude Code. Saya selalu menggunakannya. Kalau saya tidak duduk di hadapan komputer, saya pun berfikir tentangnya. Saya rasa saya tidak membakar token yang mencukupi. Ini adalah zaman awal Claude Code. Dan saya berfikir, kenapa saya tidak dapat menyambung semakan dari telefon saya? Saya menetapkan Termius dan saya tidak begitu menyukainya.

Saya mula mencipta aplikasi saya sendiri dan berhenti apabila Claude memperkenalkan aplikasi mudah alih Claude Code mereka.

Kemudian saya mendapat sumbingan cakera dan saya tidak benar-benar dapat membuat apa yang banyak. Masa berlalu dan saya rasa pulih sedikit dan saya mahu meneruskan projek Claude Code saya melalui web/pwa.

Kemudian Claude Code mula mengharamkan penggunaan di luar harness mereka sendiri. Dan saya rasa ia tidak berbaloi.

Kemudian saya menemui pi.dev dan menerangkanya sedikit tetapi belum benar-benar mendalami. Saya membaca tentangnya, menonton video tentangnya dan memutuskan untuk mencuba sepenuhnya dan kini saya benar-benar tenggelam dalam pi.

Memandangkan ia sumber terbuka, saya rasa ia berbaloi untuk dibina. Saya juga mendapat pilihan pembekal yang berbeza. Saya juga merasakan bergantung kepada satu pembekal/model seperti Anthropic/Claude adalah tidak mampan.

Jadi saya membina ia di sini.

Semakan keluaran ini dikekalkan sebagai edisi model tempatan bagi pi-web. Ia mengikuti projek hulu untuk penambahbaikan bersama, pada masa yang sama mengekalkan penebangan tempatan, kestabilan konteks, dan pengujian model tempatan pada jalan relak yang dikeluarkan secara berasingan.

## Mengapa model tempatan memerlukan profil operasi yang berbeza

Pengalaman pi-web asal adalah tapak asas yang ممتاز, tetapi inferens tempatan mempunyai mod kegagalan yang berbeza daripada model berhos tipikal. Model tempatan mungkin melambat dengan ketara apabila konteks membesar, berkongsi memori terhad dengan bahagian lain mesin, berhenti selepas menghasilkan penalaran sahaja, atau hilang satu larian panjang akibat kegagalan pengangkutan tempatan yang sementara. Memperlakukan kes-kes tersebut tepat seperti kegagalan awan menjadikan UI kelihatan serasi walaupun sesi sebenar kekal rapuh.

Edisi ini mendekati masalah tersebut secara berlapis:

1. **Preserve upstream first.** UI bersama dan kelakuan sesi terus datang
   dari pi-web; perubahan tempatan diasingkan di sebalik Local Mode yang berkesan.
2. **Prevent before recovering.** Sempadan konteks berdasarkan peratusan 65% dikenakan
   sebelum panggilan model susulan, termasuk panggilan di dalam gelung alat yang panjang.
3. **Recover only with evidence.** Pentyambung automatik dihadkan kepada insiden konteks,
   pengangkutan, dan penalaran sahaja yang dikenali—bukan pengesahan, kuota, atau
   ralat pembekal yang sewenang-wenangnya.
4. **Bound every autonomous action.** Insiden pulih digem-bin,
   kemajuan diperlukan sebelum penyelamatan lain, dan permulaan mempertimbangkan satu
   sesi Local yang baru-baru ini aktif pada paling banyak.
5. **Keep a manual exit.** Force Compact meringkaskan dan bukan menghapus sejarh, jadi
   pengguna dapat menyelamatkan sesi tanpa berpura-pura konteks itu tidak pernah wujud.
6. **Protect cloud compatibility.** Cloud Mode mengekalkan semantik dan kawalan
   upstream; pengoptimuman model tempatan tidak mensila-mentakrifkan semula sesi awan secara senyap.

Itulah perbezaan sebenar dalam forka ini: ia memperlakukan inferens tempatan sebagai persekitaran operasi yang berasingan, dan bukan sekadar nama model lain dalam senarai turun.
