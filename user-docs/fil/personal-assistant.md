# pi-web bilang Personal Assistant mo

Suporta ng local-model edition ang workflow na ito. Para sa development, synchronization, at release policy ng edition, tingnan ang
[Local-model edition development](../../docs/dev/local-llm-development.md).

Hindi lang para sa coding ang pi-web — maaari mong gawin itong **personal AI assistant** na nakatira sa computer mo, parang may sarili kang OpenClaw o Hermes.

## Paano ito gumagana

Gumagawa ka ng dedicated folder sa iyong machine — doon nakatira ang iyong assistant. Loob nito, inilagay ang isang `APPEND_SYSTEM.md` file na nagdedediba kung sino ang iyong assistant, ano ang alam nito, at kung paano ito gumagana. Binibigyan ka ng pi-web ng magandang chat interface para makausap ito mula sa alinmang device.

## Step by step

### 1. Lumikha ng iyong assistant folder

Pumili ng folder sa iyong computer. Maaaring:

```
~/my-assistant/
```

### 2. I-define ang iyong assistant

Gumawa ng `APPEND_SYSTEM.md` file sa loob ng folder na iyon. Dito pangsasabi mo sa pi kung sino ang iyong assistant:

```markdown
# My Personal Assistant

You are Jarvis, my personal AI assistant. You help me with:

- Daily planning and reminders
- Research and summarization
- Drafting emails and messages
- Brainstorming ideas
- Keeping track of things I mention

## About me

- I'm a software engineer who works remotely
- I have a cat named Pixel
- I prefer short, direct answers
- My timezone is PST

## Rules

- Be concise — I value brevity
- If you don't know something, say so
- Proactively remind me of things I asked you to track
```

Awtomatikong idinadagdag ng pi ito sa system prompt ng bawat conversation, kaya laging alam ng iyong assistant kung ikaw sino at kung paano makakatulong.

### 3. I-start ang isang session sa folder na iyon

Sa pi-web, gumawa ng bagong session na nakapunta sa `~/my-assistant/` (o ano pa itong nakatawag mo). Tapos na yan — nagsasalita ka na sa iyong personal assistant.

### 4. Gamitin ito sa alinman

I-install ang pi-web bilang PWA sa iyong phone, tablet, o laptop. Laging andiyan ang iyong assistant — itanong ang alinman, anumang oras.

## Mga ideya para sa iyong assistant

| Role | Ano ang ilalagay sa APPEND_SYSTEM.md |
|---|---|
| 🧠 **Life coach** | Ang mga goals mo, ang mga gawi na tinutulak mo, ang mga prompt para sa journaling |
| 🏠 **Home manager** | Format ng grocery list, preferences ng mga miyembro ng pamilya, meal planning |
| 💼 **Work buddy** | Ang role mo, kasalukuyang mga proyekto, format ng meeting note, konteksto ng kompanya |
| 📚 **Study partner** | Ano ang natututunan mo, paboritong istilo ng paliwanag, quiz me mode |
| ✍️ **Writing assistant** | Ang istilo ng pagsulat mo, preferences ng tono, karaniwang format na ginagamit mo |

## Magdagdag pa ng konteksto

Maaari mong ilagay sa assistant folder mo ang alinmang bagay na makakatulong upang maging mas makapangyarihan ang pi:

- `notes/` — reference files na mababasa ng iyong assistant
- `context.md` — impormasyon sa background tungkol sa buhay o trabaho mo
- `projects.md` — kasalukuyang mga proyekto at ang kanilang status

Mababasa ng pi ang mga files sa folder, kaya't ang mas maraming konteksto na ibinibigay mo, mas nasisiyahan ito.

## Hayaan ang pi-web na gawin ang mga bagay

Pagkatapos ng `pi install npm:@timmygod/pi-web-local`, maaari na ng mga session makausap ang pi-web mismo.
Subukan:

- “Magdagdag ng iskedyul at 2am oras ng Singapore para i-summarize ang inbox ko”
- “Ilista ang mga pi-web schedules ko”
- “I-pause ang inbox schedule”
- “Isulat ito sa notes”
- “Palitan ang pi-web sa dark mode / itikim ang auto-title”

Ang bundled na **/skill:pi-web-schedule** skill ay nagpapalit nito sa tunay na pi-web
schedule (parehong mga inii-edit mo sa `/schedules`). Bawat pag-firing ay nagsisimula ng **bagong**
session, kaya ang mga instructions ay kailangang mag-isa — “i-summarize ang hindi nabasang mail sa
~/inbox” ay gumagana; “ituloy ang ginagawa natin” ay hindi.

Tumatakbo lamang ang mga schedule habang naka-on ang pi-web.

---

> 💡 **Tip:** Magsimula nang simple. Isang ilang linya lang tungkol sa ikaw at kung paano mo gusto na gumana ang assistant. Palayain sa paglipas ng oras habang natututo ka kung ano ang gumagana.
