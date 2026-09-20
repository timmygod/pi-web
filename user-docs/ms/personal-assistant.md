# pi-web sebagai Pembantu Peribadi Anda

Aliran kerja ini disokong oleh edisi model-tempatan. Untuk pembangunan, penyegerakan, dan dasar pelepasan edisi tersebut, lihat [Pemaju edisi model-tempatan](../../docs/dev/local-llm-development.md).

pi-web bukan hanya untuk penulisan kod — anda boleh menjadikannya **pembantu AI peribadi** yang mendiami komputer anda, seperti memiliki OpenClaw atau Hermes anda sendiri.

## Cara ia berfungsi

Anda mewujudkan sebuah folder khusus pada mesin anda — itulah tempat pembantu anda berada. Di dalamnya, anda meletakkan fail `APPEND_SYSTEM.md` yang mentakrifkan siapa pembantu anda, apa yang diketahuinya, dan bagaimana ia berkelakuan. pi-web memberikan anda antara muka perbualan yang kemas untuk bercakap dengannya daripada mana-mana peranti.

## Langkah demi langkah

### 1. Cipta folder pembantu anda

Pilih sebuah folder pada komputer anda. Sesetengah seperti:

```
~/my-assistant/
```

### 2. Takrifkan pembantu anda

Cipta fail `APPEND_SYSTEM.md` di dalam folder tersebut. Ini adalah tempat anda memberitahu pi siapa pembantu anda:

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

pi secara automatik menambah ini ke setiap prompt sistem perbualan, jadi pembantu anda sentiasa tahu siapa anda dan bagaimana untuk membantu.

### 3. Mula sesi di dalam folder tersebut

Dalam pi-web, cipta sesi baharu yang diarahkan ke `~/my-assistant/` (atau apa jua nama yang anda berikan). Itupun cukup — anda sedang bercakap dengan pembantu peribadi anda.

### 4. Gunakan daripada mana-mana tempat

Pasang pi-web sebagai PWA pada telefon, tablet, atau komputer riba anda. Pembantu anda sentiasa di situ — tanya apa sahaja, bila-bila masa.

## Idea untuk pembantu anda

| Peranan | Apa hendak diletakkan dalam APPEND_SYSTEM.md |
|---|---|
| 🧠 **Pentunjuk kehidupan** | Matlamat anda, tabiat yang sedang anda usahakan, prompt penjurnalan |
| 🏠 **Pengurus rumah** | Format senarai bahan mentah, pilihan ahli keluarga, perancangan hidangan |
| 💼 **Rakan sekerja** | Peranan anda, projek semasa, format nota mesyuarat, konteks syarikat |
| 📚 **Rakan belajar** | Apa yang anda sedang belajar, gaya penjelasan yang dipinati, mod uji saya |
| ✍️ **Pembantu penulisan** | Gaya penulisan anda, pilihan nada, format biasa yang anda gunakan |

## Tambah lebih banyak konteks

Anda boleh meletakkan apa sahaja dalam folder pembantu anda yang membantu pi menjadi lebih berguna:

- `notes/` — fail rujukan yang boleh dibaca oleh pembantu anda
- `context.md` — maklumat latar belakang tentang kehidupan atau kerja anda
- `projects.md` — projek semasa dan statusnya

pi boleh membaca fail dalam folder, jadi semakin banyak konteks yang anda beri, semakin bagus ia menjadi.

## Minta pi-web melakukan perkara

Selepas `pi install npm:@timmygod/pi-web-local`, sesi boleh bercakap dengan pi-web itu sendiri.
Cuba:

- “Tambah jadual pada jam 2 petang waktu Singapura untuk merumuskan peti masuk saya”
- “Senaraikan jadual pi-web saya”
- “Jeda jadual peti masuk”
- “Tuliskan ini dalam nota”
- “Tukar pi-web ke mod gelap / matikan tajuk automatik”

Kemahiran **/skill:pi-web-schedule** yang disertakan menukarkannya kepada jadual pi-web sebenar (sama yang anda edit pada `/schedules`). Setiap pelancaran memulakan sesi **baharu**, jadi arahan mesti berdiri sendiri — “ringkaskan e-mel tidak dibaca dalam ~/inbox” berfungsi; “teruskan apa yang kami sedang lakukan” tidak.

Jadual hanya dilaksanakan semasa pi-web berjalan.

---

> 💡 **Petua:** Mula dengan yang ringkas. Hanya beberapa baris tentang siapa anda dan bagaimana anda mahu pembantu berkelakuan. Perbaiki sedikit demi sedikit seiring waktu apabila anda belajar apa yang berfungsi.
