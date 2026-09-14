# Bakit pi-web?

Medyo adik ako sa Claude Code. Lagi ko itong ginagamit. Kapag hindi ako nakaupo sa harap ng computer, iniisip ko ito. Pakiramdam ko hindi ako nakakasunog ng sapat na tokens. Noong mga unang araw ito ng Claude Code. At iniisip ko, bakit hindi ako makapagpatuloy mula sa aking telepono? Nag-set up ako ng Termius at hindi ko talaga ito nagustuhan.

Nagsimula akong gumawa ng sarili ko at tumigil nang ipinakilala ng Claude ang kanilang Claude Code mobile app.

Pagkatapos ay nagkaroon ako ng herniated disc at wala talaga akong masyadong magawa. Lumipas ang panahon at medyo gumaling ako at gusto kong ipagpatuloy ang aking Claude Code via web/pwa project.

Pagkatapos ay sinimulan ng Claude Code na i-ban ang paggamit sa labas ng kanilang sariling harness. At pakiramdam ko hindi na ito sulit.

Pagkatapos ay natagpuan ko ang pi.dev at nag-explore nang kaunti pero hindi pa talaga sumisid. Binasa ko ang tungkol dito, nanood ng mga video tungkol dito at nagpasya na subukan nang buo at ngayon ay lubos na akong nahuhumaling sa pi.

Dahil ito ay open source, pakiramdam ko sulit itong pagbuuan. Nakakakuha rin ako ng iba't ibang pagpipilian ng provider. Pakiramdam ko rin na ang pag-asa sa iisang provider/model tulad ng Anthropic/Claude ay hindi sustainable.

Kaya ginagawa ko ito dito.

## Bakit kailangan ng isang lokal na modelo ng ibang profile ng operasyon

Ang orihinal na karanasan sa pi-web ay isang mahusay na pundasyon, ngunit ang lokal na paghuhusga ay may iba't ibang paraan ng pagkabigo kumpara sa karaniwang naka-host na modelo. Maaaring mabagal nang husto ang isang lokal na modelo habang lumalaki ang konteksto, magbahagi ng limitadong memorya sa natitirang bahagi ng makina, huminto pagkatapos ng paggawa lamang ng pag-iisip, o mawalan ng mahabang pagtakbo dahil sa pansamantalang pagkabigo sa lokal na transportasyon. Ang pagtrato sa mga kaso na iyon nang eksakto bilang mga pagkabigo sa cloud ay nagpapakita ng UI na tila compatible habang ang aktwal na sesyon ay nananatiling fragile.

Lapitan ng edisyong ito ang problema sa mga layer:

1. **Panatilihin muna ang upstream.** Ang ibinahaging UI at pag-uugali ng sesyon ay patuloy na galing sa pi-web; ang mga lokal na pagbabago ay nakahiwalay sa likod ng epektibong Local Mode.
2. **Iwasan bago ibalik.** Ang isang 65% hangganan ng konteksto na batay sa porsyento ay ipinatupad bago ang mga susunod na tawag sa provider, kabilang ang mga tawag sa loob ng mahahabang tool loop.
3. **Ibalik lamang may ebidensya.** Ang awtomatikong pagpapatuloy ay limitado sa mga kinikilalang insidente ng konteksto, transportasyon, at thinking-only, hindi sa mga error sa authentication, quota, o arbitrary na error ng provider.
4. **Limitahan ang bawat awtonomong aksyon.** Ang mga insidente ng pagbabalik ay deduplicated, kinakailangan ang progreso bago ang isa pang pagligtas, at ang startup ay nag-iisip ng hindi hihigit sa isang kamakailang aktibong Local session.
5. **Panatilihin ang manual na paglabas.** Ang Force Compact ay nagbubuod sa halip na burahin ang history, kaya magagawa ng user na iligtas ang isang sesyon nang hindi nagpapanggap na wala nang umiiral na konteksto.
6. **Protektahan ang compatibility ng cloud.** Ang Cloud Mode ay pinapanatili ang upstream na semantika at kontrol; ang mga optimization ng lokal na modelo ay hindi tahimik na tinutukoy muli ang mga cloud session.

Iyan ang tunay na pagkakaiba sa fork na ito: itinuturing nito ang lokal na paghuhusga bilang isang hiwalay na kapaligiran ng operasyon, hindi lamang isa pang pangalan ng modelo sa dropdown.
