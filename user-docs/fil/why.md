# Bakit pi-web?

Nag-aaddict ako sa Claude Code. Laging ginagamit ko ito. Kung hindi ako nakaupo sa harap ng computer, inuusisa ko ito. Nakakaramdam akong hindi ko sapat na binubuhos ang mga token. Noong mga unang araw pa ng Claude Code. At nangungulugod ako, bakit hindi ko ma-resume mula sa phone ko? Na-set up ko ang Termius at hindi naman ako talaga nagustuhan.

Nagsimula akong gumawa ng sarili ko at tumigil nang ipinalanag ng Claude ang kanilang Claude Code mobile app.

Pagkatapos, may herniated disc ako at hindi na ako talaga makagawa ng marami. Lumipas ang oras at nakaramdam akong kaunti nang nakapag-recover at gusto kong ituloy ang Claude Code sa pamamagitan ng web/pwa project ko.

Pagkatapos, nagsimulang ipagbabawal ng Claude Code ang gamit sa labas ng kanilang sariling harness. At kinikilusan kong hindi na worth it.

Pagkatapos, natagpuan ko ang pi.dev at tinananginig ko nang konti pero wala pang talagang malalim na pagsusuri. Binasa ko tungkol dito, tiningnan ko ang mga video tungkol dito, at nagdesisyon akong ibigyan ito ng buong pagsubok—ngayon ay buong-buo na ako sa pi.

Dahil open source ito, akala ko ay worth na worth it na gawin. May iba't ibang provider choices din ako. Nakaramdam din akong hindi sustainable na magtiwala sa iisang provider/model tulad ng Anthropic/Claude.

Kaya ngayon kong itinataayo dito.

Inaangkin itong checkout bilang isang local-model edition ng pi-web. Sundo nito ang upstream project para sa shared improvements, habang pinapanatili ang local deployment, context stability, at local-model testing sa hiwalay na released track.

## Bakit kailangan ng local model ng magkaibang operating profile

Ang orihinal na pi-web experience ay isang napakahusay na pundasyon, ngunit ang local inference ay may magkaibang failure modes kaysa sa isang tipikal na hosted model. Maaaring mabagabagagap ang local model habas lumalaki ang context, makipagbahagi ng limitadong memory sa buong machine, huminto pagkatapos lumikha lamang ng reasoning, o mawala ang isang mahabang run dahil sa transient local transport failure. Ang pagpapanggap na eksaktong pag-aaralan ang mga kaso na ito na parang cloud failures ay nagpapakita sa UI ng compatibility ngunit ang aktwal na session ay nananatiling mahina.

Inaabot ng edisyong ito ang problema sa pamamagitan ng mga antas:

1. **Linagilin muna ang upstream.** Ang shared UI at session behavior ay patuloy na galing sa pi-web; ang local na mga pagbabago ay naa-isolate sa likod ng effective Local Mode.
2. **Iwasan bago makarecover.** May isang percentage-based na 65% context boundary na ipinapatupad bago ang mga sumunod na model calls, kasama ang mga calls sa loob ng mahahabang tool loops.
3. **Makarecover lamang may ebidensya.** Ang automatic na pagpapatuloy ay limitado sa kilalang context, transport, at thinking-only na mga insidente—hindi authentication, quota, o anumang provider errors.
4. **Limitahin ang bawat autonomous na aksyon.** Ang mga recovery incidents ay deduplikado, kinakailangan ng progress bago ang isa pang rescue, at ang startup ay binibigyan ng pangako na iisang recently active Local session lamang.
5. **Panatilihin ang manual na exit.** Ang Force Compact ay nagsasaari bilang halip na wapasin ang history, kaya't maaari ng user na i-rescue ang isang session nang hindi nagpapanggap na wala kailanman ang context.
6. **Protektahan ang cloud compatibility.** Ang Cloud Mode ay nagpapanatili ng upstream semantics at controls; ang local-model na mga optimization ay hindi nakaibig ng silent redefinition sa cloud sessions.

Ito ang tunay na pagkakaiba sa fork na ito: itinakdang isang hiwalay na operational environment ang local inference, hindi lamang isa pang pangalan ng model sa dropdown.
