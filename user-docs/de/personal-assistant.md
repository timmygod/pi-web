# pi-web als Ihr persönlicher Assistent

Dieser Workflow wird von der Local-Model-Edition unterstützt. Für die
Entwicklung, Synchronisation und das Release-Verfahren dieser Edition siehe
[Local-Model-Edition-Entwicklung](../../docs/dev/local-llm-development.md).

pi-web eignet sich nicht nur für das Programmieren – Sie können es zu einem **persönlichen KI-Assistenten** machen, der auf Ihrem Computer lebt, wie Ihre eigene OpenClaw- oder Hermes-Instanz.

## So funktioniert es

Sie erstellen einen dedizierten Ordner auf Ihrem Computer – dort lebt Ihr Assistent. Darin legen Sie eine `APPEND_SYSTEM.md`-Datei ab, die definiert, wer Ihr Assistent ist, was er weiß und wie er sich verhält. pi-web bietet Ihnen eine schöne Chat-Oberfläche, um von jedem Gerät mit ihm zu sprechen.

## Schritt für Schritt

### 1. Ihren Assistenten-Ordner erstellen

Wählen Sie einen Ordner auf Ihrem Computer. Etwas wie:

```
~/my-assistant/
```

### 2. Ihren Assistenten definieren

Erstellen Sie in diesem Ordner eine `APPEND_SYSTEM.md`-Datei. Hier geben Sie pi an, wer Ihr Assistent ist:

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

pi hängt diese Datei automatisch an den System-Prompt jedes Gesprächs an, sodass Ihr Assistent immer weiß, wer Sie sind und wie er helfen kann.

### 3. Eine Sitzung in diesem Ordner starten

Erstellen Sie in pi-web eine neue Sitzung, die auf `~/my-assistant/` zeigt (oder auf welchen Namen Sie ihn auch immer benannt haben). Und das war's – Sie sprechen mit Ihrem persönlichen Assistenten.

### 4. Von überall aus nutzen

Installieren Sie pi-web als PWA auf Ihrem Smartphone, Tablet oder Laptop. Ihr Assistent ist immer dort – stellen Sie ihm jederzeit alles, was Sie möchten.

## Ideen für Ihren Assistenten

| Rolle | Was in APPEND_SYSTEM.md stehen soll |
|---|---|
| 🧠 **Life Coach** | Ihre Ziele, Gewohnheiten, an denen Sie arbeiten, Journaling-Prompts |
| 🏠 **Hausmanager** | Format der Einkaufsliste, Vorlieben der Familienmitglieder, Essensplanung |
| 💼 **Arbeitspartner** | Ihre Rolle, aktuelle Projekte, Format der Meeting-Notizen, Unternehmenskontext |
| 📚 **Lernpartner** | Was Sie lernen, bevorzugter Erklärungsstil, Quiz-me-Modus |
| ✍️ **Schreibassistent** | Ihr Schreibstil, Tonfall-Vorlieben, häufig verwendete Formate |

## Mehr Kontext hinzufügen

Sie können alles in Ihren Assistenten-Ordner legen, was pi nützlicher macht:

- `notes/` – Referenzdateien, die Ihr Assistent lesen kann
- `context.md` – Hintergrundinformationen über Ihr Leben oder Ihre Arbeit
- `projects.md` – aktuelle Projekte und deren Status

pi kann Dateien im Ordner lesen, je mehr Kontext Sie ihm geben, desto besser wird er.

## pi-web Aufgaben delegieren

Nach `pi install npm:@timmygod/pi-web-local` können Sitzungen mit pi-web selbst sprechen.
Provozieren Sie Folgendes:

- „Füge einen Zeitplan für 2 Uhr morgens Singapur-Zeit hinzu, um mein Postfach zusammenzufassen"
- „Liste meine pi-web-Zeitpläne auf"
- „Setze den Postfach-Zeitplan auf Pause"
- „Schreibe das in die Notizen"
- „Wechsle pi-web in den dunklen Modus / schalte Auto-Titel aus"

Die mitgelieferte **/skill:pi-web-schedule**-Skill macht daraus einen echten pi-web
Zeitplan (dieselben, die Sie unter `/schedules` bearbeiten). Jeder Lauf startet eine **neue**
Sitzung, daher müssen die Anweisungen für sich allein stehen – „ungelesene Mails in
~/inbox zusammenfassen" funktioniert; „setz das fort, was wir taten" nicht.

Zeitpläne laufen nur, wenn pi-web läuft.

---

> 💡 **Tipp:** Fangen Sie einfach an. Nur ein paar Zeilen darüber, wer Sie sind und wie der Assistent sich verhalten soll. Iterieren Sie mit der Zeit, wenn Sie herausfinden, was funktioniert.
