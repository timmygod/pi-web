# pi-web ကို သင့်ကိုယ်ရေးကူညီသူအဖြစ် အသုံးပြုခြင်း

ဤ workflow ကို local-model edition က ထောက်ပံ့ပေးပါသည်။ edition ၏
development၊ synchronization၊ release policy အကြောင်းအတွက်
[Local-model edition development](../../docs/dev/local-llm-development.md) ကို ကြည့်ရှုပါ။

pi-web သည် coding အတွက်သာမက — သင့်computer ပေါ်တွင်နေထိုင်ပြီး သင့်ကိုယ်ပိုင် OpenClaw သို့မဟုတ် Hermes ကဲ့သို့သော **personal AI assistant** အဖြစ် ပြောင်းလဲအသုံးပြုနိုင်ပါသည်။

## ဘယ်လို လည်ပတ်သနည်း

သင့်machine ပေါ်တွင် အဓိကသတ်မှတ်ထားသော folder တစ်ခု create လုပ်ပါ — သင့်assistant ၏ နေရာဖြစ်ပါသည်။ ထိုအတွင်းတွင် သင့်assistant သည်ใคร ဖြစ်သည်၊ ဘာများ သိရှိပြီး၊ ဘယ်လို လုပ်ဆောင်သည် ဆိုသည်ကို define လုပ်ထားသော `APPEND_SYSTEM.md` file တစ်ခု ထည့်ပါ။ pi-web သည် သင့်ကို ဘယ် device မှမဆို ထိုassistant နှင့် စကားပြောနိုင်ရန် chat interface လှပလွန်းသော တစ်ခုကို ပေးထားပါသည်။

## အဆင့်ဆင့် လုပ်ဆောင်နည်း

### 1. သင့် assistant folder ကို create လုပ်ပါ

သင့် computer ပေါ်တွင် folder တစ်ခု ရွေးပါ။ ဥပမာ -

```
~/my-assistant/
```

### 2. သင့် assistant ကို define လုပ်ပါ

ထို folder အတွင်း `APPEND_SYSTEM.md` file တစ်ခု create လုပ်ပါ။ ဤနေရာတွင် သင့် assistant သည် ဘယ်သူ ဖြစ်သည်ကို pi ကို ပြောပြပါ -

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

pi သည် ဤ text ကို နေ့စဉ် conversation တိုင်း၏ system prompt တွင် ထပ်တိုးပါရန် automatically append လုပ်ပါသည်၊ ထို့ကြောင့် သင့် assistant သည် သင်သည် အရင်းအမြစ်နှင့် ဘယ်လို ကူညီနိုင်သည်ကို အမြဲသိရှိနေပါသည်။

### 3. ထို folder တွင် session တစ်ခု start လုပ်ပါ

pi-web တွင် `~/my-assistant/` (သို့မဟုတ် သင်အမည်ပေးထားသည့် folder) အတွက် new session တစ်ခု create လုပ်ပါ။ အဆုံးသတ်ပါ — သင်သည် သင့်ကိုယ်ပိုင် assistant နှင့် စကားပြောနေပြီဖြစ်ပါသည်။

### 4. ဘယ်နေရာမှမဆို အသုံးပြုပါ

သင့် phone၊ tablet သို့မဟုတ် laptop ပေါ်တွင် pi-web ကို PWA အနေဖြင့် install လုပ်ပါ။ သင့် assistant သည် အမြဲအားလုံးရှိပါသည် — ဘယ်အချိန်တွင်မဆို ဘာမေးခွန်းမဆို မေးပါ။

## သင့် assistant အတွက် အယူအဆများ

| ပခန်း | APPEND_SYSTEM.md တွင် ဘာထည့်ရမလဲ |
|---|---|
| 🧠 **Life coach** | သင့် target များ၊ လုပ်ဆောင်နေသော အထွေထွေချက်များ၊ journaling prompts |
| 🏠 **Home manager** | Grocery list format၊ ဘဝိသရာလူများ၏ preference များ၊ meal planning |
| 💼 **Work buddy** | သင့် role၊ ဖြစ်ပေါ်နေသော project များ၊ meeting note format၊ ကုမ္ပဏီ context |
| 📚 **Study partner** | ဘာသင်ယူနေသည်၊ သိမ်းပေးလိုသော explanation style၊ quiz me mode |
| ✍️ **Writing assistant** | သင့် writing style၊ tone preferences၊ သုံးစွဲသော format များ |

## Context ပိုထပ်ထည့်ပါ

သင့် assistant folder တွင် pi ကို ပိုမိုအသုံးဝင်စေရန် ဘာမဆို ထည့်နိုင်ပါသည် -

- `notes/` — သင့် assistant က ဖတ်နိုင်သော reference file များ
- `context.md` — သင့်ဘဝ သို့မဟုတ် အလုပ်အတွက် background information
- `projects.md` — ဖြစ်ပေါ်နေသော project များနှင့် အခြေအနေများ

pi သည် folder ထဲရှိ file များကို ဖတ်နိုင်ပါသည်၊ ထို့ကြောင့် သင် ပိုများ context ပေးသေးတတ်ပါက ပိုမိုအကောင်းဖြစ်လာပါသည်။

## pi-web ကို ပြဿနာဖြေရှင်းပါ

`pi install npm:@timmygod/pi-web-local` တွင် install လုပ်ပြီးသောအခါ session များသည် pi-web ကိုယ်တိုင်နှင့် တုံ့ပြန်ပြောဆိုနိုင်ပါသည်။
စမ်းကြည့်ပါ -

- "Add a schedule at 2am Singapore time to summarize my inbox"
- "List my pi-web schedules"
- "Pause the inbox schedule"
- "Write this down in the notes"
- "Switch pi-web to dark mode / turn auto-title off"

bundled **/skill:pi-web-schedule** skill ကသည် ထိုလုပ်ငန်းစဉ်ကို ဖြစ်အောင် ဖြစ်စေသည် (သင် `/schedules` တွင် edit လုပ်သည့် schedule များနှင့် တူပါသည်။) နေ့တိုင်း firing အသစ်တစ်ခု **new** session ကို စတင်ပါသည်၊ ထို့ကြောင့် instruction များသည် တစ်ကိုယ်တည်း ဆန့်နိုင်ရပါသည် — "summarize unread mail in ~/inbox" ကတော့ သင့်တော်ပါသည်၊ "continue what we were doing" ကတော့မော်။

Schedules များသည် pi-web က ဖြစ်သလိုသာ run ဖြစ်ပါသည်။

---

> 💡 **Tip:** ရိုးရိုးစင်းစင်း စတင်ပါ။ သင်ဘယ်သူ ဖြစ်သည်နှင့် assistant က ဘယ်လို လုပ်ဆောင်ပေးလိုသည် ပါဝင်သည့် ကိန်းဂဏန်းတစ်လေးမျှသာ။ ဘာအလုပ်ဖြစ်သည်ကို သင်သိလာသည်နှင့်အမျှ အချိန်နှင့်အမျှ သိမ်းပါ။
