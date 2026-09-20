# pi-web sebagai Asisten Pribadi Anda

Workflow ini didukung oleh edisi model-lokal. Untuk pengembangan,
sinkronisasi, dan kebijakan rilis edisi tersebut, lihat
[Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web bukan hanya untuk coding — Anda bisa menjadikannya **asisten AI pribadi** yang tinggal di komputer Anda, seperti memiliki OpenClaw atau Hermes Anda sendiri.

## Cara kerjanya

Anda membuat folder khusus di mesin Anda — di situlah asisten Anda tinggal. Di dalamnya, Anda menaruh file `APPEND_SYSTEM.md` yang mendefinisikan siapa asisten Anda, apa yang ia ketahui, dan bagaimana ia berperilaku. pi-web memberi Anda antarmuka chat yang indah untuk berinteraksi dengannya dari perangkat mana pun.

## Langkah demi langkah

### 1. Buat folder asisten Anda

Pilih folder di komputer Anda. Misalnya:

```
~/my-assistant/
```

### 2. Definisikan asisten Anda

Buat file `APPEND_SYSTEM.md` di dalam folder tersebut. Di sinilah Anda memberi tahu pi siapa asisten Anda:

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

pi secara otomatis menambahkan ini ke system prompt setiap percakapan, sehingga asisten Anda selalu tahu siapa Anda dan bagaimana cara membantu.

### 3. Mulai sesi di folder tersebut

Di pi-web, buat sesi baru yang mengarah ke `~/my-assistant/` (atau nama apa pun yang Anda berikan). Selesai — Anda sedang berbicara dengan asisten pribadi Anda.

### 4. Gunakan dari mana saja

Install pi-web sebagai PWA di ponsel, tablet, atau laptop Anda. Asisten Anda selalu ada — tanyakan apa pun, kapan pun.

## Ide untuk asisten Anda

| Role | What to put in APPEND_SYSTEM.md |
|---|---|
| 🧠 **Life coach** | Your goals, habits you're working on, journaling prompts |
| 🏠 **Home manager** | Grocery list format, family members' preferences, meal planning |
| 💼 **Work buddy** | Your role, current projects, meeting note format, company context |
| 📚 **Study partner** | What you're learning, preferred explanation style, quiz me mode |
| ✍️ **Writing assistant** | Your writing style, tone preferences, common formats you use |

## Tambahkan lebih banyak konteks

Anda bisa menaruh apa pun di folder asisten Anda yang membantu pi menjadi lebih berguna:

- `notes/` — file referensi yang bisa dibaca asisten Anda
- `context.md` — informasi latar belakang tentang kehidupan atau pekerjaan Anda
- `projects.md` — proyek saat ini dan statusnya

pi bisa membaca file di dalam folder, jadi semakin banyak konteks yang Anda berikan, semakin baik kinerjanya.

## Minta pi-web melakukan sesuatu

Setelah `pi install npm:@timmygod/pi-web-local`, sesi bisa berkomunikasi dengan pi-web itu sendiri.
Coba:

- “Add a schedule at 2am Singapore time to summarize my inbox”
- “List my pi-web schedules”
- “Pause the inbox schedule”
- “Write this down in the notes”
- “Switch pi-web to dark mode / turn auto-title off”

Skill **/skill:pi-web-schedule** yang disertakan mengubahnya menjadi jadwal pi-web yang nyata (yang sama yang Anda edit di `/schedules`). Setiap eksekusi memulai sesi **baru**, sehingga instruksi harus berdiri sendiri — “summarize unread mail in ~/inbox” berhasil; “continue what we were doing” tidak.

Jadwal hanya berjalan selama pi-web berjalan.

---

> 💡 **Tip:** Mulai dari yang sederhana. Hanya beberapa baris tentang siapa Anda dan bagaimana Anda ingin asisten berperilaku. Iterasi seiring waktu saat Anda mempelajari apa yang berhasil.
