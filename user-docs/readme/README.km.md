<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · **ភាសាខ្មែរ** · [ລາວ](README.lo.md)

</div>

<div align="center">

ប្រើប្រាស់ agent កុំព្យូទ័រ [pi](https://pi.dev) របស់អ្នកពីទូរស័ព្ទ ទូរស័ព្ទល្បឿនលឿន (tablet) ឬកុំព្យូទ័រយួរដៃ — គ្រប់ទីកន្លែងក្នុងបណ្តាញរបស់អ្នក ឬពីទីក្រៅតាមរយៈ Tailscale។

វាជា PWA ពេញលេញ ដូច្នេះអ្នកអាចដំឡើង និងប្រើប្រាស់វាដូចជា app ដើមកំណើតនៅលើឧបករណ៍ណាមួយ។ គិតវាជាប្រតិបត្តិការដ៏បុគ្គលសុខ្លីរបស់អ្នក — ដូចជា Cowork របស់ Claude ប៉ុន្តែមាន model ផ្សេងៗ — ជុះទៅវាឆ្លងកាត់ model កូដពីទូរស័ព្ទ ឬបំប្លែងវាជា [ដៃគូបុគ្គល](../en/personal-assistant.md) ដែលមានអេលីម្ភក្នុងម៉ាស៊ីនរបស់អ្នក។

ធ្វើវាជារបស់អ្នក: ប្តូរ theme និង font ហើយប្រើវាក្នុងភាសារបស់អ្នក — pi-web មកជាមួយភាសាច្រើន និងអ្នកអាចបន្ថែមភាសារបស់អ្នកផ្ទាល់។ មុខងារបន្ថែមទៀតកំពុងមកដល់ ប៉ុន្តែវានឹងមិនប្រើក្រអ៊ូន្ទួននោះទេ: អ្វីដែលអ្នកមិនប្រើទេអាចបិទនៅក្នុងការកំណត់។

</div>

## ហេតុអ្វីបានជាបដ្អាល់កំណែ model ក្នុងម៉ាស៊ីននេះ?

pi-web ដើមនៅតែជាមូលដ្ឋានគ្រឹះ upstream សម្រាប់មុខងារ និងការជួសជុលដែលចែករំលែកគ្នា។ កំណែនេះរក្សាការប្រាសាទនោះ បន្ទាប់មកបន្ថែមជាន់សុវត្ថិភាពសម្រាប់ model ដែលដំណើរការនៅលើម៉ាស៊ីនផ្ទាល់របស់អ្នក ឬនៅកន្លែងផ្សេងទៀតក្នុង LAN របស់អ្នក — ដែលការបង្កើត content តែងតែយឺត អង្គចងចាំមានកម្រិត ហើយ context វែងអាចធ្វើឱ្យ session ដែលសុខភាពល្អជាទូទៅដឺរ។

| តំបន់ | Upstream pi-web | កំណែនេះ |
|------|-----------------|--------------|
| គោលការណ៍ Model/runtime | ប្រតិបត្តិការ pi-web ធម្មតា | ម៉ូទ័រ **Auto / Local / Cloud** ក្នុងមួយ session ដោយមានការស្វែងរក local ដោយយល់ដឹង endpoint និងការ override ដោយដៃដែលរក្សាទុក |
| ការដោះស្រាយ Long-context | ប្រតិបត្តិការ compaction របស់ pi ធម្មតា | Local Mode ធ្វើ compaction ជាមុខងារនៅ **65%** និងពិនិត្យម្តងទៀតនៅក្នុងវដ្ត tool-call វែងមុនសំណើ model បន្ទាប់ |
| សុវត្ថិភាព Compaction | Summary ធម្មតា | Rolling checkpoint មានកម្រិត ការសរសេរឡើងវិញតិចតួចមួយសម្រាប់ output ដែលមិនត្រឹមត្រូវ/មានកម្រិត ហើយការរកមើលកម្លាំងមិនទទួលបានជំនួសឱ្យការ compaction ឡើងវិញមិនគ្រប់ចំណុច |
| ការដំណើរការ被打断 | ការគ្រប់គ្រង worker និង error ធម្មតា | ការឆ្លើយតបមានកម្រិតសម្រាប់ context overflow ការឈប់គិតតែមួយគត់ និងការបាត់បង់ការដឹកជញ្ជូនជាក់លាក់ ដោយមាន loop breaker ដែលរក្សាទុក |
| ការជួយដោយដៃ | ព័ត៌មាន context ធម្មតា | **Force Compact** នៅតែមានសម្រាប់ជាផ្លូវឆ្លើយតបជាក់លាក់ដោយមិនបំបិនវាចជម្រុះទេ |
| ភាពត្រឹមត្រូវនិងការចេញផ្សាយ | គម្រោងដើមនិងខ្សែការចេញផ្សាយ | ការការពារ local តែមួយគត់នៅពីក្រោយ Local Mode; Cloud Mode រក្សាប្រតិបត្តិការ upstream ហើយការប្តូរ upstream ត្រូវបានពិនិត្យ និងចេញផ្សាយនៅទីនេះដោយឯករាជ្យ |

នេះមិនមែនជាការសរសេរឡើងវិញ ឬជំនួសឱ្យ upstream ទេ។ នេះជា operating profile ដែលរក្សាទុកដោយប្រុងប្រយ័ត្ន សម្រាប់អ្នកដែលចង់បានសុវត្ថិភាពបុគ្គល និងការគ្រប់គ្រង model ក្នុងម៉ាស៊ីនដោយមិនទទួលយក session ដំណើរការយូរដែលងាយបែកចេញ។ សូមមើល [ឯកសារណែនាំអ្នកប្រើប្រាស់](../en/README.md) សម្រាប់ workflow ដែលបង្ហាញអ្នកប្រើប្រាស់ និង [ការអភិវឌ្ឍកំណែ model ក្នុងម៉ាស៊ីន](../../docs/dev/local-llm-development.md) សម្រាប់ការអនុវត្តន៍ និងគោលការណ៍ទំនាក់ទំនង។

> [!TIP]
> មកថ្មី? **[អានឯកសារណែនាំ →](../en/README.md)** សម្រាប់ការទស្សនាមុខងារ ជាពេញលេញ ជំហានដំឡើង និងគន្លឹះ។ ([ភាសាផ្សេងទៀត →](../README.md))

## រូបថតសេអ្រេណរប៉ុរ

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## វាតំរែតូចគ្នា

```
 pi (terminal)                 Browser (phone / tablet / laptop)
      │                                │
      │  writes JSONL                  │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           via MagicDNS)
```

- **pi** សរសេរ conversation JSONL ទៅ `~/.pi/agent/sessions/` ខណៈដែលវាកំពុងធ្វើការ។
- **pi-web** ជា server Go ដែលអានឯកសារទាំងនោះ បង្ហាញវាក្នុង browser និងផ្គត់ផ្គង់ការបច្ចុប្បន្នភាពជីវស្ម័រតាមរយៈ SSE។
- worker **pi --mode rpc** ដោះស្រាយ chat ដែលចាប់ផ្តើមពី browser — មួយក្នុងមួយ session ហើយត្រូវបានលុបចោលបន្ទាប់ពី 10 នាទីដែលមិនមានសកម្មភាព។
- **fsnotify** មើលថែក្រុមតំបន់ sessions ដើម្បីឱ្យ browser reload ក្នុងរយៈពេលមិនគួរឱ្យសម្រេចបន្ទាប់ពីមាន output ថ្មី។
- **Tailscale Serve** បង្ហាញ server localhost ជា endpoint HTTPS នៅលើ tailnet របស់អ្នក។

## ការដំឡើង

```bash
pi install npm:@timmygod/pi-web-local
```

នោះគឺជាអ្វីដែលអ្នកត្រូវការ — វាទាញយក binary ដែលត្រូវគ្នា កំណត់ auto-start ហើយចុះពាក្យបញ្ជា `/web` `/pi-web` `/remote` និង `/refresh`។

បន្ទាប់ពីដំឡើង ចង `http://127.0.0.1:31415` នៅក្នុង browser របស់អ្នក។ ពី pi ប្រើ `/web` ដើម្បីបើក session បច្ចុប្បន្ននៅក្នុង browser របស់អ្នកភ្លាមៗ។ ប្រសិនបើ Tailscale កំពុងដំណើរការនៅលើម៉ាស៊ីនរបស់អ្នក pi-web ដោយស្វ័យប្រវត្តិបង្ហាញ endpoint HTTPS នៅលើ tailnet របស់អ្នក — ប្រើ `/remote` ពី pi ដើម្បីទទួលបាន QR code និង URL សម្រាប់ឧបករណ៍ណាមួយនៅលើ tailnet របស់អ្នក។

> **ការចូលប្រើឆ្ងាយ macOS:** ដំឡើង និងចង Tailscale ដោយមានអន្តរកម្ម យល់ព្រមនឹងសំណើអ្នកគ្រប់គ្រង និងចូល។ បន្ទាប់មក執 /pi-web restart បន្ទាប់មក `/remote`។

សម្រាប់ការដំឡើងដោយដៃ ការទាញយក binary ឬការបង្កើតពីសូត់ សូមមើល [user-docs/install.md](../en/install.md)។

## ការប្រើប្រាស់ Pi

បន្ទាប់ពី `pi install npm:@timmygod/pi-web-local` អ្នកទទួលបាន៖

| បញ្ជា | វាធ្វើអ្វី |
|---------|--------------|
| `/web` | បើក session បច្ចុប្បន្ននៅក្នុង browser របស់អ្នក (ស្គាល់ SSH: បញ្ឈប់ browser និងបង្ហាញ URL តែមួយគត់) |
| `/pi-web` | បង្ហាញ ស្ថានភាព version បើក/បិទ/ចាប់ផ្តើមឡើងវិញ server ឬ update |
| `/remote` | បង្ហាញ QR code និង URL សម្រាប់ការចូលប្រើឆ្ងាយតាមរយៈ Tailscale |
| `/refresh` | ទាញយកសារថ្មីដែលបានសរសេរពី browser ឆ្ងាយត្រឡប់មកវិញទៅក្នុង session terminal |

ការ **auto-titling** ត្រូវបានបង្កើតក្នុង pi-web ខ្លួនឯង និងកំណត់នៅទំព័រ `/settings`។ វា **បើកដោយស្វ័យប្រវត្តិ** និងដាក់ឈ្មោះ session ដោយស្វ័យប្រវត្តិ។ អ្នកអាចជ្រើសរើស៖

- **ពេលណាដាក់ឈ្មោះ** — ម្តងក្នុងមួយ session ឬនៅពេលមានសារថ្មីនីមួយៗ (លម្អិតនោះ)។
- **Title model** — **built-in word heuristic (no AI)** ដែលមានដោយឥតគិតថ្លៃ និងភ្លាមៗ ជាលម្អិត ឬជ្រើសរើស model (ឧ. តូច/លឿន) សម្រាប់ title ដែលឆ្លាតវៃ និងសរសេរដោយ model។

package នេះក៏ដំឡើង binary pi-web ទៅ `~/.pi/agent/bin/pi-web` និងកំណត់ auto-start នៅពេលចូលដែរ។

## Auto-Start នៅពេលចូល

បញ្ជា `pi install npm:@timmygod/pi-web-local` កំណត់នេះដោយស្វ័យប្រវត្តិ៖

| ប្រព័ន្ធប្រតិបត្តិការ | យន្តការ |
|----|-----------|
| macOS | launchd plist នៅ `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service នៅ `~/.config/systemd/user/pi-web.service` |
| Windows | ចំណុច Run-key នៃ `HKCU` ដែលចាប់ផ្តើម starter ដែលលាក់នៅ `~/.config/pi-web/` |

ដើម្បីកំណត់ token សម្រាប់ការចូលប្រើឆ្ងាយ បង្កើត `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

សម្រាប់ព័ត៌មានបន្ថែម (ការកំណត់ដោយដៃ custom ports non-loopback binds) សូមមើល [user-docs/install.md](../en/install.md)។

## ការអភិវឌ្ឍ

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

សម្រាប់ការទំនាក់ទំនង upstream ការសាកល្បង model ក្នុងម៉ាស៊ីន និង workflow ការចេញផ្សាយស្របគ្នា សូមមើល [ការអភិវឌ្ឍកំណែ model ក្នុងម៉ាស៊ីន](../../docs/dev/local-llm-development.md)។
