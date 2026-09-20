# Roadmap

ဒီ roadmap က pi-web ၏ local-model edition အတွက် ဖြစ်ပါတယ်။ Upstream feature များကို စက္ရတ်လိုက် ယှဉ်ညှိပြီး၊ local deployment နှင့် reliability လုပ်ငန်းများကို ဒီလိုင်းတွင် စစ်ဆေးပြီး ထုတ်ဝေပါသည်။ [Local-model edition development](../../docs/dev/local-llm-development.md) ကို ကြည့်ပါ။

pi-web သည် အသုံးပြုသူအဖွဲ့နှစ်ဖွဲ့အတွက် ဆောက်လုပ်ထားခြင်းဖြစ်ပါသည်။

- **Developers များအတွက်** — terminal တွင်နေထိုင်သော်လည်း mobile အတူတက် sessions ဆက်လက်နိုင်ရန်၊ remote server သို့ hand off လုပ်နိုင်ရန်၊ သို့မဟုတ် နေရာမရွေး long-running tasks များကို စောင့်ကြည့်နိုင်ရန် လိုအပ်သူများ။
- **Non-developers များအတွက်** — ကောင်းမွန်လှပသော AI app တစ်ခုသာ လိုချင်သူများ။ ဖွင့်ရုံ၊ ရိုက်ရုံ၊ vibe ရုံ။ Terminal မလို၊ SSH မလို၊ ရှုပ်ထွေးမှု မရှိ။ အသုံးလွယ်ဆုံး AI tools များအတိုင်းပင်၊ သို့သော် model ရွေးချယ်ခြင်းနှင့် open-source အခွင့်အာ�ာနှင့်အတူ။

အောက်ပါအရာများ ကျရောက်လာပါမည်။

ဒီ edition သည် upstream pi-web ကို မတူညီသော release line တစ်ခုဖြင့် ခံစောသည်။ Upstream feature များကို စက္ရတ်လိုက် ဝင်ရောက်ယူပြီး၊ local-model reliability လုပ်ငန်းများကို upstream release history မပြောင်းလဲဘဲ ဤနေရာတွင် နှစ်ထွေးထားပြီး စစ်ဆေးပါသည်။

---

## Now (shipped)

[the features table](README.md#what-you-can-do-with-pi-web) တွင် ရေးဆွဲထားသမျှ အားလုံးကို ယနေ့တွင် အသုံးပြုနိုင်ပါပြီ။

---

## Next up

| # | Feature | ပါဝင်သည့်အရာ |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | Telegram သို့ Discord မှတစ်ဆင့် pi နှင့် ချက်ချင်းစကားပြောနိုင်သည် — on the go လုပ်ဆောင်မှုများအတွက် personal assistant workflows များအတွက် အထူးကောင်းမွန်သည်။ |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | Token tracking၊ cost estimation၊ session analytics များဖြင့် ခရိုက်သည်အတိုင်းသိပါ။ |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | Sessions အားလုံးတွင် thinking၊ tools၊ tool outputs များအတွက် သင့်ဝင်သော visibility ကို သတ်မှတ်ပါ။ |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | pi အနေ၌ ဆက်လက်လည်ပတ်နေစဉ် follow-up instructions များကို ပေးပို့နိုင်သည် — mid-flight တွင် လမ်းညွှန်ပါ။ |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | Web UI မှတိုက်ရိုက်ဖြင့် ရှည်လျားသော အစကားပြောမှုများကို compact လုပ်နိုင်သည်၊ terminal မလိုအပ်ပါ။ |

---

## Planned

| # | Feature | ပါဝင်သည့်အရာ |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | Project file tree ကို browse လုပ်ပြီး git ပြောင်းလဲမှုများကို pi-web တွင် တိုက်ရိုက်ဖြင့် မြင်နိုင်သည်။ Opt-in ဖြစ်သောကြောင့် သင့်လမ်းတွင် မနှောင့်နှေးပါ။ |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Prompts များကို automatic ဖြင့် ရောင်းချနိုင်သည် — daily standups၊ မနက်စခေါက် summaries၊ ပုံမှန်လုပ်ဆောင်ချက်များ။ Safety အတွက် Admin-gated ဖြစ်သည်။ |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | သင့် muscle memory နှင့် ညီအောင် keyboard shortcuts များအားလုံးကို remap လုပ်နိုင်သည်။ |

---

## Vision

အချိန်အတော်ကြာသော အရည်ဆုံး ဦးပန်း: pi-web သည် **pi ၏ interface** ဖြစ်သင့်သည် — လူအားလုံးအတွက်။

- **Non-devs များ** သည် အခြား app များအတိုင်းပင် ဖွင့်ပါ။ Model တစ်ခုကို ရွေးပါ။ ရိုက်ပါ။ ပြီးပါပြီ။ Command line မလိုအပ်ပါ။
- **Developers များ** အတွက် အနက်ရိသော integration များ ရရှိသည် — remote handoff၊ multi-session dashboards၊ git-aware browsing၊ messaging bots များ။
- **လူအားလုံး** အတွက် model freedom၊ open-source transparency၊ နှင့် တစ်ခုစီတွင် စဉ်းစားမှုရှိသောသဘောရရှိသည့် UI ကို ရရှိပါမည်။

---

> 💡 Idea တစ်ခုရှိပါသလား။ [Open an issue](https://github.com/timmygod/pi-web/issues/new) သို့မဟုတ် ကိစ္စပြောဆိုချက်တွင် ပါဝင်ပါ။
