# Roadmap

Ang roadmap na ito ay karapatan ng local-model edition ng pi-web. Iis同步 (synchronized) nang regular ang mga feature mula sa upstream, habang ang trabaho sa local deployment at reliability ay pinapatunayan at inililabas sa linya na ito; tingnan ang [Pag-unlad ng local-model edition](../../docs/dev/local-llm-development.md).

Ang pi-web ay itinatayo para sa dalawang grupo:

- **Para sa mga developer** — ang naninirahan sa terminal ngunit nais magpadala ng session mula sa mobile, mag-hand off sa isang remote server, o magtanong ng mata sa mga long-running tasks mula kahit saan.
- **Para sa mga hindi developer** — ang gustong lang ng isang magandang AI app na gumagana. Buksan, mag-type, vibe. Walang terminal, walang SSH, walang kalituhan. Katulad ng pinaka user-friendly na AI tools, ngunit may choice ng model at kalayaan ng open-source.

Narito ang darating.

Sinusundan ng edisyon na ito ang upstream pi-web sa isang magkakaibang release line. Iniiimport nang regular ang mga upstream feature; pinaprioridad at pinapatunayan dito ang local-model reliability work nang walang pagbabago sa release history ng upstream.

---

## Ngayon (inilabas)

Lahat ng nakalista sa [features table](README.md#what-you-can-do-with-pi-web) ay buhay na ngayon.

---

## Susunod

| # | Feature | Ano ang ginagawa nito |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | Makipag-usap sa pi sa pamamagitan ng Telegram o Discord — perpekto para sa mga personal assistant workflow sa labas. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | Pagbabantay ng token, pagtataya ng gastos, analytics ng session — alam mo kung paano ginagamit mo ang pi. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | Itakda ang preferred visibility mo para sa thinking, tools, at tool outputs sa lahat ng session. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | Ipadala ang follow-up instructions habang tumatakas pa rin ang pi — gabayan ito sa gitna ng landas. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | I-compact ang mahahabang usapan direktang mula sa web UI, walang terminal na kailangan. |

---

## Plano

| # | Feature | Ano ang ginagawa nito |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | Mag-browse sa file tree ng project at tingnan ang git changes direkta sa pi-web. Opt-in, para manatili itong malayo sa daan mo. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Iskedyul ang mga prompt na magaganap nang automatic — araw-araw na standups, morning summaries, paulit-ulit na mga gawain. Admin-gated para sa kaligtasan. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | I-remap ang bawat keyboard shortcut para tumugma sa iyong muscle memory. |

---

## Vision

Ang pangmatagalang layunin: ang pi-web ay dapat maging **ang interface ng pi** — para sa lahat.

- **Hindi devs** ay buksan itong tulad ng anumang app. Pumili ng model. Mag-type. tapos. Walang command line kailanman.
- **Devs** ay makakakuha ng malalim na integration — remote handoff, multi-session dashboards, git-aware browsing, messaging bots.
- **Lahat** ay makakakuha ng kalayaan ng model, open-source transparency, at isang UI na pakiramdam ay maingat sa bawat hakbang.

---

> 💡 May idea? [Buksan ang isang issue](https://github.com/timmygod/pi-web/issues/new) o sumali sa talakayan.
