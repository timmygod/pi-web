# Maligayang Pagdating sa pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · **Filipino** · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Isinusulat mo bang pagsubok ng pi-web? Gawin mo na — magmamahal ka rito.**

Ang pi-web ay isang magandang web UI at PWA para sa [pi](https://pi.dev) — ang open-source na AI coding agent. Pinapayagan ito mong mag-browse, basahin, at ipagpatuloy ang iyong mga pi session mula sa alinmang browser, sa alinmang device, may mapag-isiling na mga feature sa bawat paglipat.

## Ano ang iba't ibang sa edisyong ito?

Itinututkad ng repository na ito ang upstream na pi-web interface at shared features, ngunit
pinapalitan nito kung paano pinoprotektahan ang mga session kapag ang napiling model ay nag-e-execute nang local o sa
iyong LAN.

- **Piliin ang runtime policy bawat session.** Awtomatikong naaawa ang local/LAN endpoints kapag
  malinaw na ang provider metadata; ang Local at Cloud ay permanenteng manual na overrides.
- **Iwasan ang context failures nang maagang.** Ang Local Mode ay nag-c-compact sa 65% ng usage at tinatantya
  muli sa pagitan ng mga tool calls, bago ang susunod na model request.
- **Panatilihin ang limit ng mga summary.** Iwasan ng mga rolling checkpoint ang walang-hanggang paglago ng lumang
  summary, subukan ulit isang beses na may mas makitid na budget, at huminto nang ligtas kapag walang
  makabuluhang pag-unlad ang compaction.
- **Makinig na pag-anyayaya.** Ang context overflow, mga pigil ng napiling transport,
  at mga reasoning-only na maagang paghinto ay maaaring mag-restart nang awtomatiko, ngunit
  ang deduplication ng insidente at mga progress-aware circuit breaker ay humaharang sa mga loop ng recovery.
- **Manatili ang kontrol sa user.** Ang Force Compact ay laging ang makikiting manual na rescue
  path, habang ang Cloud Mode ay nananatili ang upstream na workflow at controls.

Ang praktikal na resulta ay simpleng: ang mahabang local-model task ay dapat mag-compact bago ito
mahulog, mag-recover isang beses kapag ligtas ang recovery, at huminto nang malinis sa halip na
humalo nang walang dibdib kapag hindi.

**Gawa ang pi-web para sa dalawang uri ng tao:**

- 🧑‍💻 **Para sa mga developer** — na nakatira sa terminal ngunit nais magpatuloy ng session mula sa mobile, maghand-off sa isang remote server, o monitorin ang mga mahabang-running task mula sa anumang lugar.
- ✨ **Para sa mga hindi-developer** — na gustong-gusto lang ng magandang AI app na gumagana. Buksan, i-type, mag-vibe. Walang terminal, walang SSH, walang lilito. Katulad ng pinaka-user-friendly na AI tools, ngunit may pagpipilian ng model at open-source na kalayaan.

---

## Bakit pi-web?

Naka-lalim ka na sa flow na may pi sa iyong terminal. Pinapanatili ng pi-web ang momentum na iyon kapag naglalayo ka sa iyong desk:

- **Magpatuloy mula sa kahit saan** — ipagpatuloy ang session mula sa iyong phone, tablet, o ibang computer. Walang SSH, walang Termius — buksan mo lang ang iyong browser.
- **Multi-session dashboard** — ilantad ang gawa sa isang session habang nanonood ng stream ng isa pa. Maghanap sa mga proyekto, i-filter ayon sa branch, malutas kung ano ang kinakailangan nang mabilis.
- **Open-source na pundasyon** — ang pi ay buong-buso na open source at hindi pinapayak sa provider. Walang kinakailangan na lumipat sa isang model o vendor. Ang pi-web ay open source din.
- **Ligtas na remote access** — built-in na token auth para ipahintulot ito sa iyong LAN o Tailscale nang walang pag-aalala.
- **Magbahagi ng gawa mo** — i-export ang mga session bilang static snapshots o secret na GitHub Gist sa iisang pag-click.

> Kuriyoso sa backdrop? [Basahin kung bakit namin ito ginawa →](why.md)

---

## Ang pi-web bilang iyong sariling AI workspace 🏠

Ang pi-web ay isang PWA (Progressive Web App), kaya maaari mo itong **i-install katulad ng native app** sa iyong desktop, laptop, phone, o tablet — walang app store na kailangan. Sa desktop, bukas ito sa sariling window na walang browser chrome, kaya parang totoong desktop application ang dating nito.

Isipin itong **sarili mong Claude Cowork** — sariling AI workspace na nakatira sa iyong makina — ngunit open source at walang kinakailangan sa model:

- **Ikaw ang may-ari ng stack.** Pumili ng alinmang model, palitan kapag gusto mo. Manatili ang local na isa at hindi na ang data mo ayaw sa iyong makina.
- **Maaari gamitan ng mga hindi-teknikal.** Itakda ang pi-web sa kanilang makina, ipaalam kung paano gamitin ito isang beses, at handa na sila. Ang iyong magulang, ang iyong partner, ang iyong hindi-teknikal na kaibigan — walang terminal, walang SSH, kung ano ang natural na chat interface.
- **Isang setup, maraming user.** I-install sa iyong desktop at i-share ang iyong screen, o ipahintulot sa iyong tahanan na network at payagan ang mga miyembro ng pamilya na buksan ito sa kanilang sariling device.

Gusto mo pa ng higit sa coding? Huwag, para sa isang dedicated na [personal assistant](personal-assistant.md) na nakakakilala sa iyo at nakatira sa iyong makina — katulad ng sarili mong OpenClaw o Hermes.

> 💡 **Pro tip:** I-install ang pi-web bilang PWA mula sa Chrome/Edge (i-click ang install icon sa address bar) o Safari (Share → Add to Dock). Walang magkakaiba ito mula sa native app.

---

## Mga magagawa sa pi-web

| | |
|---|---|
| 📱 **PWA** | I-install ang pi-web bilang Progressive Web App sa desktop, phone, o tablet para sa native na dating. |
| 🔄 **Ipatuloy ang mga session** | Ituloy ang alinmang usapan sa kung saan mo ito iwan — teksto, larawan, paglipat ng model, lahat mula sa browser. |
| 🆕 **Mag-umpisa ng bagong session** | Lumikha ng bagong session laban sa alinmang project path, direkta mula sa web UI. |
| 📡 **Live streaming** | Manood ng stream ng mga sagot ng pi sa real time na may ~ms na latency. Ang Follow mode ay tinatago ang pinakabago. |
| 🌲 **Tree view** | Mag-navigate sa native na message tree ng pi — makita ang buong estrukturang usapan, lumipat sa alinmang branch, at mag-fork mula sa alinmang punto. |
| 🔀 **Mag-fork ng session** | Mag-fork ng session mula sa alinmang message o kahit na specific na tool call — subukin ang iba't ibang direksyon nang hindi nawawalan ng iyong lugar. |
| 🔍 **Mag-browse & hanap** | I-filter ang mga session sa mga proyekto, hanapin ayon sa pangalan, mag-navigate ng mga branch — ang buong session history mo sa isang tingin. |
| 🌿 **Git integration** | Makita ang kasalukuyang branch at buksan ang isang GitHub PR direkta mula sa session viewer. |
| 📝 **Scratchpad** | Isulat ang mga tala, to-do, o maikling ideya katabi ng mga session nang walang paglipat ng app. |
| 💬 **Annotations** | I-highlight at mag-komento sa alinmang bahagi ng session — maganda para sa code review, feedback, o pag-mark ng mahalagang moment. |
| 🎨 **Themes & customization** | Lumipat sa dark o light mode, ayusin ang UI ayon sa kaibigan — gawing katulad ng *iyong sarili* ang dating ng pi-web. |
| 🌐 **Multi-language** | 14 na built-in na wika (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Magdagdag ng sarili mong custom na wika mula sa Settings. |
| 🐱 **Wellness & pomodoro** | Sobrang lalong vibe coding ay hindi maganda. Built-in na pomodoro timer na may cat companion at sleep reminders para manatili ka balanced. |
| 📤 **Magbahagi & i-export** | I-download ang JSONL, i-export ang static snapshots na rendered ng native na `pi.dev` look ng pi, o ibahagi bilang private na GitHub Gist — lahat ng ito ay rendered client-side. |
| 🔔 **Notification sounds** | Mag-customize na mga notification chime para sa mga session event — manatili ka na loop kahit nasa ibang tab ang pi-web. |
| ⌨️ **Keyboard shortcuts** | Vim-style na pag-navigate, mabilis na mga aksyon — [buong reference →](keyboard-shortcuts.md) |
| 🤖 **Personal assistant** | Huwag ang pi-web sa iyong sariling AI assistant na nakatira sa iyong computer — katulad ng OpenClaw o Hermes. [Itakda →](personal-assistant.md) |
| 🗓️ **Usapanin ang mga schedule** | Mula sa pi session, sabihing "idagdag ang schedule at 2am na oras ng Singapore para sa …" — `/skill:pi-web-schedule`. |
| 📝 **Usapanin ang mga tala & settings** | "Isulat ito sa mga tala" (`/skill:pi-web-notes`) o "lilipat sa dark mode" (`/skill:pi-web-settings`). |

---

## Mabilisang navigation

| Kapag hinahanap mo… | Basahin |
|---|---|
| Paano i-install, itakda, at gamitin ang pi-web | [install.md](install.md) |
| Gumamit ng pi-web bilang personal assistant | [personal-assistant.md](personal-assistant.md) |
| Mga keyboard shortcuts reference | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Kung bakit bumubuhay ang pi-web | [why.md](why.md) |
| Ano ang susunod na darating | [roadmap.md](roadmap.md) |
| May problema sa install? Ipasa ang llm-debug.md link sa iyong LLM para ayusin nito | [llm-debug.md](llm-debug.md) |
| Pagmamaintain ng local-model edisyong ito | [development notes](../../docs/dev/local-llm-development.md) |

---

## Mga screenshot

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

Ginawa ang pi-web nang may pag-ibig at maraming gabi na hindi natutulog. Babayaran ng coding plans (Claude Code, OpenCode, atbp.) sa aking sarili ang gastos upang patuloy ang proyekto na ito. Kung tulong ng pi-web, mahalaga ang iyong suporta.

**Mga paraan ng tulong:**

- 💰 **[Sponsoron sa GitHub](https://github.com/sponsors/setkyar)** — tumulong sa mga tool na nagpapagawa nito
- ☕ **[Bilhin ang aking kape](https://buymeacoffee.com/setkyar)** — bawat maliit na tulong ay nakakatulong
- ⭐ **I-star ang repo** — walang ginaginhawa at nakakatulong sa mas maraming tao na makita ang pi-web
- 📢 **Ibahagi sa mga kaibigan & pamilya** — kung may kilala mong gusto ng pi-web, ipadala sa kanya

Hindi maaaring sumponsor? Walang problema — isang star at isang share ay malaki ang lakas. Salamat sa pagiging nandito. 🙏

---

Maligayang coding! 🚀
