# pi-web သို့ အဆိုထည့်ပါ 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · **မြန်မာ** · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**pi-web ကို စမ်းကြည့်ဖို့ စဉ်းစားနေပါသလား? ဝိုင်းပါ — အရူးသွပ်သွပ် ရှိလာပါလိမ့်မယ်။**

pi-web သည် [pi](https://pi.dev) — open-source AI coding agent အတွက် အလှကောင်းကောင်း web UI နှင့် PWA တစ်ခုဖြစ်ပါသည်။ ကိုယ်တိုင်၏ pi session များကို browser မည်သည့်စက်မှမဆို၊ device မည်သည့်အလက်နဲ့မဆို လေ့လာရန်၊ ဖတ်ရှုရန်၊ ဆက်လက်လုပ်ဆောင်ရန် ကြိုတင်စဉ်းစားထားသော features များဖြင့် ဖြစ်စေပါသည်။

## ဤ version ထဲမှာ ဘာကွဲပြားသလဲ?

ဤ repository သည် upstream pi-web interface နှင့် shared features များကို ထိန်းသိမ်းထားပြီး၊ selected model ကို လက်ရှိပတ်ဝန်းကျင်တွင် သို့မဟုတ် ကိုယ်တိုင်၏ LAN တွင် လည်ပတ်စဉ် session များကို how to protect သည်ကို ပြောင်းလဲထားပါသည်။

- **Session တစ်ခုချင်းစီအတွက် runtime policy ကိုရွေးချယ်နိုင်သည်။** provider metadata က ရှင်းလင်းပါက local/LAN endpoint များကို auto detect ပေးသည်; Local နှင့် Cloud များသည် manual override ပြစ်ပွန်းပွဲအသိများဖြစ်ပါသည်။
- **Context failure များကို ကြိုတင်ကာကွယ်သည်။** Local Mode သည် usage 65% တွင် compact လုပ်ပြီး၊ tool call များကြားတွင် ကြည့်ရှုသည့်အလုပ်၊ အနောက်ဆုံး model request အတွင်း မတိုင်မီ ထပ်မံစစ်ဆေးသည်။
- **Summary များကို သတ်မှတ်ထားသည်။** Rolling checkpoints များသည် old summary ကို ပြီးစီးအောင် ကြီးထွားစေရန် ကာကွယ်ပြီး၊ tighter budget ဖြင့် once ကျော်လွှား retry လုပ်ကာ، compaction က meaningful progress မမီပါက ရပ်ရပ်စနစ်သင့်ရပ်သည်။
- **conservative ဖြင့် ပြန်လည်ရယူသည်။** Context overflow, selected transport interruptions, နှင့် reasoning-only premature stops များသည် auto resume ဖြစ်နိုင်သော်လည်း， incident deduplication နှင့် progress-aware circuit breakers များသည် recovery loops များကို ကာကွယ်သည်။
- **User ကို မိမိထိန်းချုပ်ခွင့် ပေးသည်။** Force Compact မည်သည့်အခါတွင်မဆို visible manual rescue path အဖြစ် ရှိနေပြီး၊ Cloud Mode သည် upstream workflow နှင့် controls များကို ထိန်းသိမ်းထားသည်။

အကျိုးရလဒ်ကလည်း ရိုးရှင်းပါသည် — local model ကြာရှည်သော task တစ်ခုသည် ပြိုကျမယ့်အရင် compact ဖြစ်သင့်ပြီး， recovery သည် safe ဖြစ်နေပါက once ဖြင့် ပြန်ရယူသင့်ကာ， မဖြစ်ပါက loop မနေဘဲ လင်းလင်းလျင်လျင် ရပ်သင့်ပါသည်။

**pi-web သည် လူနှစ်မျိုးအတွက် ဖန်တီးထားပါသည်:**

- 🧑‍💻 **Developers အတွက်** — terminal တွင် နေရင်း mobile မှ sessions ကို ဆက်လက်လုပ်ချင်၊ remote server သို့ hand off လုပ်ချင်၊ မည်သည့်နေရာမှမဆို ကြာရှည်သော tasks များကို monitor လုပ်ချင်သူများ။
- ✨ **Non-developers အတွက်** — လုပ်ဆောင်နိုင်သော လှပသော AI app တစ်ခုပဲ လိုချင်သူများ။ ဖွင့်၊ ရေး， vibe ဖြစ်ပါ။ Terminal မလို၊ SSH မလို၊ ပေါက်ကြောမှု မလို။ Model choice နှင့် open-source freedom ဖြင့် အရည်အသွေးမြင့်လှသော AI tools များနက်တောင့်။

---

## pi-web ဟာ ဘာကြောင့်လဲ?

ကိုယ်တိုင်သည် terminal တွင် pi ဖြင့် flow တွင် ရပ်တန့်နေပါသည်။ pi-web သည် မိမိ၏ desk မှ အနီးစပ်တွင် နေရာရောက်ချိန် momentum ကို ဆက်ရှက်ရွက်စေသည်:

- **မည်သည့်နေရာမှမဆို resume** — phone၊ tablet၊ သို့မဟုတ် ကွန်ပျူတာအခြားမှ session တစ်ခုကို ဆက်လက်လုပ်ဆောင်။ SSH မလို၊ Termius မလို — browser ဖွင့်ပါပဲ။
- **Multi-session dashboard** — session တစ်ခုတွင် work ကို စတင်စဉ်， အခြား session တစ်ခု stream ကို ကြည့်ရှုနိုင်။ Project များကြား search လုပ်၊ branch ဖြင့် filter လုပ်， လိုအပ်သောအရာကို မြန်မြန် မှန်မှန် ရှာဖွေနိုင်။
- **Open-source foundation** — pi သည် fully open source နှင့် provider-agnostic ဖြစ်သည်။ Model တစ်ခုလက်တစ်ခု သို့ vendor တစ်ခုတည်းတွင် lock in မဖြစ်ပါ။ pi-web ကလည်း open source ဖြစ်သည်။
- **Safe remote access** — LAN သို့မဟုတ် Tailscale တွင် expose လုပ်သည့်အခါ စိုးရိမ်စရာမရှိသည့် built-in token auth။
- **မိမိ၏ work ကို share လုပ်ပါ** — session များကို static snapshots သို့မဟုတ် secret GitHub Gists အဖြစ် တစ်ကြိမ် click ဖြင့် export လုပ်ပါ။

> Background story ကို စိတ်ဝင်စားပါသလား? [ဘာကြောင့် build လုပ်တာလဲ ဖတ်ပါ →](why.md)

---

## pi-web ကို ကိုယ်ပိုင် personal AI workspace အဖြစ်သုံးပါ 🏠

pi-web သည် PWA (Progressive Web App) ဖြစ်သောကြောင့်， desktop， laptop， phone၊ tablet တွင် **native app ကဲ့သို့ install** လုပ်နိုင်ပါသည် — app store မလိုပါ။ Desktop တွင် browser chrome မရှိဘဲ window သီးခြားဖြင့် ဖွင့်သည်， ထို့ကြောင့် desktop application လက်ရှိမြင်စွာဖြင့် ခံစားစေသည်။

**မိမိ၏ Claude Cowork** အဖြစ် ယူပါ — ကိုယ်တိုင်၏ machine တွင် နေသော personal AI workspace — သို့သော် open source နှင့် model-agnostic ဖြစ်သည်:

- **Stack ကို မိမိပိုင်သည်။** Model မည်သည့်အရာကိုမဆို pick လုပ်၊ ဘယ်وقتမဆို ချွေ့နိုင်သည်။ Local တစ်ခုကို run လုပ်ပါက ကိုယ်တိုင်၏ data မည်သည့်အရာမှ မထွက်သွားပါ။
- **Non-technical လူများသည် သုံးနိုင်သည်။** သူတို့၏ machine တွင် pi-web ကို setup လုပ်၊ အသုံးပြုပုံ တစ်ကြိမ် ပြသပြီး၊ သူတို့သည် စတင်နိုင်သည်။ မိဘများ， partner， non-tech friends — terminal မလို， SSH မလို， သိချင်သော chat interface တစ်ခုသာ။
- **Setup တစ်ကြိမ်， users အများ။** Desktop တွင် install လုပ်ပြီး screen ကို share လုပ်， သို့မဟုတ် home network တွင် expose လုပ်ပြီး ဘုံကြီးသားများအတွက် ကိုယ်တိုင်၏ device များတွင် ဖွင့်နိုင်အောင် ပေးပါ။

Coding ထက်ပိုအပ်ပါသလား? ကိုယ်တိုင်၏ ပတ်ဝန်းကျင်ကို ကြိုတင်သိထားပြီး machine တွင် နေသော dedicated [personal assistant](personal-assistant.md) အဖြစ် ပြောင်းလဲပါ — မိမိ၏ OpenClaw သို့မဟုတ် Hermes ကဲ့သို့။

> 💡 **Pro tip:** Chrome/Edge တွင် (address bar တွင် install icon ကို click) သို့မဟုတ် Safari တွင် (Share → Add to Dock) မှ pi-web ကို PWA အဖြစ် install လုပ်ပါ။ Native app နှင့် မခွဲခြားနိုင်တော့ပါ။

---

## pi-web ဖြင့် မည်သည့်အရာများ လုပ်နိုင်သလဲ

| | |
|---|---|
| 📱 **PWA** | Native feel အတွက် desktop， phone， tablet တွင် pi-web ကို Progressive Web App အဖြစ် install လုပ်ပါ။ |
| 🔄 **Sessions ဆက်လက်သည်** | ရပ်ခဲ့သည့်နေရာတွင် စကားပြောမှု မည်သည့်အရာကိုမဆို ပြန်စပါ — text， images， model switching， browser ကုန်ကြမ်းပြုခြင်း။ |
| 🆕 **Sessions အသစ်စတင်သည်** | Project path မည်သည့်အရာမှမဆို fresh session များကို web UI တွင် တိုက်ရိုက် ဖန်တီးပါ။ |
| 📡 **Live streaming** | pi responses ကို ~ms latency ဖြင့် real time တွင် stream ဖြင့် ကြည့်ရှုပါ။ Follow mode ကို latest တွင် locked ဖြင့်ထားသည်။ |
| 🌲 **Tree view** | pi ကို native message tree ကို လမ်းကြောင်းရှာပါ — full conversation structure ကို မြင်ပါ၊ branch မည်သည့်အရာကိုမဆို jump လုပ်ပါ၊ မည်သည့်နေရာမှမဆို fork လုပ်ပါ။ |
| 🔀 **Sessions fork လုပ်သည်** | Message မည်သည့်အရာမှ သို့မဟုတ် tool call သီးခြားတစ်ခုမှ session ကို fork လုပ်ပါ — မိမိ၏ position ကို မဆုံးရှုံးဘဲ လမ်းကြောင်း အမျိုးမျိုးကို explore လုပ်ပါ။ |
| 🔍 **Browse & search** | Project များကြား sessions ကို filter လုပ်၊ name ဖြင့် search လုပ်၊ branches ကို လမ်းကြောင်းရှာပါ — ကိုယ်တိုင်၏ full session history ကို ချက်ချင်းမြင်ပါ။ |
| 🌿 **Git integration** | Current branch ကို မြင်ပြီး， session viewer တွင် တိုက်ရိုက် GitHub PR ကို ဖွင့်ပါ။ |
| 📝 **Scratchpad** | Apps ကို ပြောင်းမလုပ်ဘဲ session များနဲ့အတူ notes， todos， ချက်ချင်း thoughts များကို မှတ်တမ်းတင်ပါ။ |
| 💬 **Annotations** | Session ကို မည်သည့်အပိုင်းကိုမဆို highlight လုပ်ပြီး comment ပါ — code review， feedback， သို့မဟုတ် key moments များအတွက် bookmark လုပ်ရာတွင် ကောင်းသည်။ |
| 🎨 **Themes & customization** | Dark နှင့် light mode ကြား ချွေ့ပါ， UI ကို ကိုက်ညီအောင် tweak လုပ်ပါ — pi-web ကို *ကိုယ်တိုင်၏* ကဲ့သို့ ခံစားစေပါ။ |
| 🌐 **Multi-language** | Built-in languages 14 များ (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសာខ្មែរ, ລາວ). Settings တွင် ကိုယ်တိုင်၏ custom language ကို ပေါင်းထည့်ပါ။ |
| 🐱 **Wellness & pomodoro** | Vibe coding လွန်ကဲခြင်းသည် ကျန်းမာမှု မကောင်းပါ။ Balanced ဖြစ်အောင် cat companion နှင့် sleep reminders ဖြင့် built-in pomodoro timer။ |
| 📤 **Share & export** | JSONL ကို download လုပ်， pi ကို native `pi.dev` look ဖြင့် render လုပ်ထားသော static snapshots များကို export လုပ်， သို့မဟုတ် private GitHub Gists အဖြစ် share လုပ်ပါ — အားလုံးသည် client-side တွင် render လုပ်သည်။ |
| 🔔 **Notification sounds** | Session events များအတွက် customizable notification chimes — pi-web က အခြား tab တွင် ဖြစ်နေချိန်တွင် ဖြတ်မသွားအောင်။ |
| ⌨️ **Keyboard shortcuts** | Vim-style navigation， quick actions — [full reference →](keyboard-shortcuts.md) |
| 🤖 **Personal assistant** | pi-web ကို ကိုယ်တိုင်၏ computer တွင် နေသော မိမိ၏ AI assistant အဖြစ် ပြောင်းလဲပါ — OpenClaw သို့မဟုတ် Hermes ကဲ့သို့။ [Setup လုပ်ပါ →](personal-assistant.md) |
| 🗓️ **Schedules နဲ့ တိုင်ပင်သည်** | Pi session တစ်ခုမှ， "add a schedule at 2am Singapore time to …" ဟု ပြောပါ — `/skill:pi-web-schedule`။ |
| 📝 **Notes & settings နဲ့ တိုင်ပင်သည်** | "Write this in the notes" (`/skill:pi-web-notes`) သို့မဟုတ် "switch to dark mode" (`/skill:pi-web-settings`) ဟု ပြောပါ။ |

---

## Quick navigation

| မိမိရှာဖွေနေသည့်အရာ… | ဖတ်ပါ |
|---|---|
| pi-web ကို install， configure， အသုံးပြုပုံ | [install.md](install.md) |
| pi-web ကို personal assistant အဖြစ် အသုံးပြုခြင်း | [personal-assistant.md](personal-assistant.md) |
| Keyboard shortcuts reference | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| pi-web သည် ဘာကြောင့် ရှိတာလဲ | [why.md](why.md) |
| နောက်တစ်ခု ဘာလာမလဲ | [roadmap.md](roadmap.md) |
| Install trouble ရှိပါသလား? LLM ကို fix လုပ်အောင် ပေးပါ — llm-debug.md link ကို သူတို့ထဲသို့ paste လုပ်ပါ | [llm-debug.md](llm-debug.md) |
| ဤ local-model edition ကို maintain လုပ်ခြင်း | [development notes](../../docs/dev/local-llm-development.md) |

---

## Screenshots

| Desktop | Mobile |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

pi-web သည် love နှင့် ညမစောစော များဖြင့် build လုပ်ထားပါသည်။ Coding plans များ (Claude Code， OpenCode， etc.) ကို ကိုယ်တိုင်၏ pocket ထဲမှ ပေးချေပြီး ဤ project ကို ဆက်လက်ရှေ့တန်းသို့ ဆောင်ရွက်စေပါသည်။ pi-web ကို မိမိအတွက် ကြိုက်နှစ်သက်ပါက၊ မိမိ၏ support သည် ကိုယ်တိုင်အတွက် ကမ္ဘာအလျား ဖြစ်ပါမည်။

**Support လုပ်နိုင်သည့်နည်းလမ်းများ:**

- 💰 **[GitHub တွင် Sponsor](https://github.com/sponsors/setkyar)** — ဤအရာကို ဖြစ်စေသည့် tools များအတွက် ကူညီပါ
- ☕ **[ကော်ဖီ တစ်ဇွန်း ကိုယ်တိုင်အတွက် ဝယ်ပါ](https://buymeacoffee.com/setkyar)** — ဟင်းသေးသေး မည်သည့်အရာကပဲ ဖြစ်စေ ကူညီပါသည်
- ⭐ **Repo ကို Star လုပ်ပါ** — အလုံးခြေမသုံးဘဲ， pi-web ကို ပိုမိုအများပြည့်ဖြင့် မြင်စေပါ
- 📢 **Friends နှင့် family နှင့် share လုပ်ပါ** — pi-web ကို ကြိုက်နှစ်သက်ဖို့ ချက်ချင်းရှိသူတစ်ယောက်ကို သိပါက， သူတို့ထဲသို့ send လုပ်ပါ

Sponsor မလုပ်နိုင်ပါသလား? ဘယ်လိုအရာမှ မလိုပါ — star တစ်ခုနှင့် share တစ်ခုသည် အများကြီး သွားပါသည်။ ဤနေရာတွင် ဖြစ်နေသည့်အတွက် ကျေးဇူးတင်ပါသည်။ 🙏

---

Coding ဖြတ်ကြရပါစို့! 🚀
