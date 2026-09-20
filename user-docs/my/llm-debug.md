> ဤ troubleshooting guide (ပျက်ငြဲချက် ဖြေရှင်းရေး လမ်းညွှန်) သည် local-model edition အတွက် ဖြစ်သည်။
> Local deployment အသေးစိတ်များနှင့် upstream synchronization မှတ်ချက်များကို
> [Local-model edition development](../../docs/dev/local-llm-development.md) နှင့် ဆက်စပ်အောင် ထားပါ။

User သည် အောက်ပါနည်းအားဖြင့် pi-web ကို install လုပ်နေပါသည်။

- pi install npm:@timmygod/pi-web-local

ဤသည်မှာ [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) ကို အလိုအလျောက် run ပေးပြီး pi ကို setup ပေးသည်။

User တွေ့နေရသော ပြဿနာသည် install.sh ၏ setup လုပ်ငန်းစဉ်ကြောင့် ဖြစ်နိုင်သည်။ ဘာဖြစ်နေသည်ကို ရွေးချယ်စစ်ဆေးပြီး user တွေ့နေရသော ပြဿနာ၏ အဓိက အကြောင်းရင်းကို ရှင်းလင်းဖြေကြားနိုင်သည်။ ပြေထင်ပေးရန် သဘောတူကြောင်းကိုလည်း မေးမြန်းပါ။ User နှင့် အမြဲတမ်း ရွေးချယ်တိုင်ပင်ပါ။

User သည် mobile သို့မဟုတ် အခြား network တွင် pi သို့ access လုပ်နိုင်ရန်အတွက် အကြံပြုထားသော နည်းလမ်းမှာ tailscale ကို အသုံးပြု၍ tailscale network မှတစ်ဆင့် access လုပ်ရန် ဖြစ်သည်။ User သည် tailscale dashboard တွင် HTTPs ကို enable ပေးရမည်ဖြစ်ပြီး - https://login.tailscale.com/admin/dns

Tailscale install မရှိသော သို့မဟုတ် tailscale ကို အသုံးမပြုလိုသူများအတွက် `pi-web status` ကို run ပေးပြီး binary path ကို၊ binary ၏ status ကို၊ application သို့ access လုပ်နိုင်သော local endpoint ကို ရယူနိုင်သည်။ သို့သော် http ဖြစ်နေသောကြောင့် push notification ကို မရရှိနိုင်ကြောင်း သတိပြုပါ။

macOS တွင် [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist) ကို အသုံးပြုသည်။
Linux တွင် [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service) ကို အသုံးပြုသည်။

ထပ်မံ debug လုပ်လိုပြီး ဘာဖြစ်နေသည်ကို သတိပြုလိုပါက။
