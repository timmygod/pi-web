# Willkommen bei pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · [Français](../fr/README.md) · **Deutsch** · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Überlegst du, pi-web auszuprobieren? Lass es dir nicht nehmen — du wirst es lieben.**

pi-web ist eine schöne Web-UI und PWA für [pi](https://pi.dev) — den Open-Source-AI-Coding-Agenten. Es ermöglicht dir, deine pi-Sitzungen aus jedem Browser heraus zu durchsuchen, zu lesen und fortzusetzen, auf jedem Gerät, mit durchdachten Funktionen an jeder Ecke.

## Was ist in dieser Edition anders?

Dieses Repository behält die pi-web-Oberfläche des Upstream-Projekts und die gemeinsamen Funktionen bei,
ändert aber, wie Sitzungen geschützt werden, wenn das ausgewählte Modell lokal oder in
deinem LAN ausgeführt wird.

- **Wähle die Laufzeitrichtlinie pro Sitzung.** Wird lokal/LAN-Endpunkte automatisch erkannt, wenn die
  Provider-Metadaten eindeutig sind; Local und Cloud sind dauerhafte manuelle Überschreibungen.
- **Verhindere Kontextfehler frühzeitig.** Local Mode komprimiert bei 65 % Auslastung und prüft
  erneut zwischen Tool-Aufrufen, vor der nächsten Modell-Anfrage.
- **Halte Zusammenfassungen begrenzt.** Rolling Checkpoints verhindern, dass eine alte
  Zusammenfassung endlos wächst, versuchen einmal mit engerem Budget und stoppen sicher, wenn Komprimierung
  keine sinnvolle Fortschritte bringt.
- **Erhole dich konservativ.** Kontextüberlauf, Unterbrechungen der gewählten Übertragungsart
  und vorzeitige Stopps nur bei Reasoning können sich automatisch wiederherstellen, aber die
  Entduplizierung von Vorfällen und fortschrittsbewusste Circuit Breaker verhindern Wiederherstellungs-Schleifen.
- **Gib die Kontrolle dem Nutzer.** Force Compact bleibt immer der sichtbare manuelle
  Rettungsmechanismus, während Cloud Mode den Upstream-Workflow und die Steuerung beibehält.

Das praktische Ergebnis ist einfach: Eine lange lokale Modell-Aufgabe sollte komprimiert werden, bevor sie
umkippt, sich einmal wiederherstellen, wenn dies sicher ist, und sauber aufhören, statt in einer
Schleife zu enden, wenn es das nicht ist.

**pi-web ist für zwei Arten von Menschen gemacht:**

- 🧑‍💻 **Für Entwickler** — die im Terminal leben, aber Sitzungen vom Handy fortsetzen, an einen Remote-Server übergeben oder lange laufende Aufgaben von überall überwachen möchten.
- ✨ **Für Nicht-Entwickler** — die nur eine schöne AI-App wollen, die funktioniert. Öffnen, tippen, vibe. Kein Terminal, kein SSH, keine Verwirrung. Wie die benutzerfreundlichsten AI-Tools, aber mit Modellwahl und Open-Source-Freiheit.

---

## Warum pi-web?

Du bist bereits tief im Flow mit pi in deinem Terminal. pi-web behält diesen Schwung aufrecht, wenn du deinen Platz verlässt:

- **Von überall fortsetzen** — setze eine Sitzung von deinem Handy, Tablet oder einem anderen Computer fort. Kein SSH, kein Termius — öffne einfach deinen Browser.
- **Multi-Sitzungs-Dashboard** — starte Arbeit in einer Sitzung, während du eine andere streamst. Suche über Projekte hinweg, filtere nach Branch und finde, was du brauchst, schnell.
- **Open-Source-Grundlage** — pi ist vollständig Open Source und Provider-unabhängig. Du bist an kein einzelnes Modell oder einen einzelnen Anbieter gebunden. Auch pi-web ist Open Source.
- **Sicherer Remote-Zugriff** — eingebauter Token-Auth, sodass du es auf deinem LAN oder Tailscale veröffentlichen kannst, ohne Sorgen zu haben.
- **Teile deine Arbeit** — exportiere Sitzungen als statische Snapshots oder geheime GitHub-Gists mit einem Klick.

> Lust auf die Hintergrundgeschichte? [Lies, warum wir es gebaut haben →](why.md)

---

## pi-web als dein persönlicher AI-Arbeitsplatz 🏠

pi-web ist eine PWA (Progressive Web App), sodass du es **wie eine Native-App installieren** kannst auf deinem Desktop, Laptop, Handy oder Tablet — ohne App-Store. Auf dem Desktop öffnet es sich in seinem eigenen Fenster ohne Browser-Chrome, sodass es aussieht und sich anfühlt wie eine echte Desktop-Anwendung.

Denk daran wie an **dein eigenes Claude Cowork** — einen persönlichen AI-Arbeitsplatz, der auf deinem Rechner lebt — nur ist er Open Source und modellunabhängig:

- **Du besitzt den Stack.** Wähle jedes Modell und wechsle, wann immer du möchtest. Führe ein lokales aus und deine Daten verlassen deinen Rechner nie.
- **Nicht-technische Menschen können es nutzen.** Richte pi-web auf ihrem Rechner ein, zeige ihnen einmal, wie es funktioniert, und sie sind startklar. Deine Eltern, dein Partner, deine nicht-tech-freundlichen Freunde — kein Terminal, kein SSH, nur eine vertraute Chat-Oberfläche.
- **Eine Einrichtung, viele Nutzer.** Installiere es auf deinem Desktop und teile deinen Bildschirm, oder veröffentliche es in deinem Heimnetzwerk und lass Familienmitglieder es auf ihren eigenen Geräten öffnen.

Mehr als nur Coding? Wandle es in einen dedizierten [persönlichen Assistenten](personal-assistant.md) um, der weiß, wer du bist und auf deinem Rechner lebt — wie dein eigenes OpenClaw oder Hermes.

> 💡 **Pro-Tipp:** Installiere pi-web als PWA von Chrome/Edge (Klicke das Installations-Symbol in der Adressleiste) oder Safari (Teilen → An Dock anheften). Es wird von einer Native-App ununterscheidbar.

---

## Was du mit pi-web machen kannst

| | |
|---|---|
| 📱 **PWA** | Installiere pi-web als Progressive Web App auf Desktop, Handy oder Tablet für ein natives Gefühl. |
| 🔄 **Sitzungen fortsetzen** | Setze jeden Gesprächsteil genau dort fort, wo du aufgehört hast — Text, Bilder, Modellwechsel, alles aus dem Browser. |
| 🆕 **Neue Sitzungen starten** | Erstelle neue Sitzungen gegen jeden Projekt-Pfad direkt aus der Web-UI. |
| 📡 **Live-Streaming** | Beobachte pi-Antworten in Echtzeit mit ~ms Latenz. Follow-Mode hält dich auf dem neuesten Stand. |
| 🌲 **Baumansicht** | Navigiere pi's natives Nachrichtenzweigesystem — sieh die gesamte Gesprächsstruktur, springe zu jedem Ast und forke von jedem Punkt aus. |
| 🔀 **Sitzungen forkten** | Fork eine Sitzung von jeder Nachricht oder sogar einem spezifischen Tool-Aufruf — erkunde verschiedene Richtungen, ohne deinen Ort zu verlieren. |
| 🔍 **Durchsuchen & suchen** | Filtere Sitzungen über Projekte hinweg, suche nach Namen, navigiere durch Branches — deine gesamte Sitzungshistorie auf einen Blick. |
| 🌿 **Git-Integration** | Sieh den aktuellen Branch und öffne ein GitHub-PR direkt aus dem Sitzungs-Viewer. |
| 📝 **Scratchpad** | Notiere Ideen, To-dos oder schnelle Gedanken neben deinen Sitzungen, ohne die App zu wechseln. |
| 💬 **Annotationen** | Markiere und kommentiere jeden Teil einer Sitzung — perfekt für Code-Review, Feedback oder zumBookmarkenwichtigerMomente. |
| 🎨 **Themes & Anpassung** | Wechsle zwischen dunklem und hellem Modus, passe die UI nach deinem Geschmack an — mache pi-web zu *deinem*. |
| 🌐 **Multi-Sprachig** | 14 eingebaute Sprachen (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Füge deine eigene Sprache in den Einstellungen hinzu. |
| 🐱 **Wellness & Pomodoro** | Zu viel Vibe Coding ist nicht gesund. Eingebauter Pomodoro-Timer mit Katzenvorgabe und Schlaf-Erinnerungen, um dich im Gleichgewicht zu halten. |
| 📤 **Teilen & exportieren** | Lade JSONL herunter, exportiere statische Snapshots, die mit pi's nativem `pi.dev`-Look gerendert werden, oder teile sie als private GitHub-Gists — alles clientseitig gerendert. |
| 🔔 **Benachrichtigungstöne** | Anpassbare Benachrichtigungsklänge für Sitzungsevents — bleibe informiert, selbst wenn pi-web in einem anderen Tab ist. |
| ⌨️ **Tastenkürzel** | Vim-Stil-Navigation, schnelle Aktionen — [vollständige Referenz →](keyboard-shortcuts.md) |
| 🤖 **Persönlicher Assistent** | Wandle pi-web in deinen eigenen AI-Assistenten um, der auf deinem Computer lebt — wie OpenClaw oder Hermes. [Einrichtung →](personal-assistant.md) |
| 🗓️ **Gespräche mit Zeitplänen** | Aus einer pi-Sitzung: „füge einen Zeitplan um 2 Uhr Singapur-Zeit hinzu, um …“ — `/skill:pi-web-schedule`. |
| 📝 **Gespräche mit Notizen & Einstellungen** | „Schreibe das in die Notizen“ (`/skill:pi-web-notes`) oder „wechsel in dunkler Modus“ (`/skill:pi-web-settings`). |

---

## Schnelle Navigation

| Wenn du suchst… | Lies |
|---|---|
| Wie man pi-web installiert, konfiguriert und nutzt | [install.md](install.md) |
| Wie man pi-web als persönlichen Assistenten nutzt | [personal-assistant.md](personal-assistant.md) |
| Tastenkürzel-Referenz | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Warum pi-web existiert | [why.md](why.md) |
| Was als Nächstes kommt | [roadmap.md](roadmap.md) |
| Probleme bei der Installation? Lass deinen LLM es beheben — füge ihnen den llm-debug.md-Link ein | [llm-debug.md](llm-debug.md) |
| Diese lokale Modell-Edition warten | [Entwickler-Notizen](../../docs/dev/local-llm-development.md) |

---

## Screenshots

| Desktop | Mobil |
|---|---|
| ![Desktop](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsor

pi-web wird mit Liebe und viel Arbeit in den späten Nächten gebaut. Ich zahle für Coding-Pläne (Claude Code, OpenCode, etc.) aus eigener Tasche, um dieses Projekt voranzubringen. Wenn pi-web für dich nützlich war, bedeutet deine Unterstützung mir die Welt.

**Möglichkeiten, zu helfen:**

- 💰 **[Sponsoring auf GitHub](https://github.com/sponsors/setkyar)** — hilf, die Tools zu finanzieren, die dies möglich machen
- ☕ **[Kauf mir einen Kaffee](https://buymeacoffee.com/setkyar)** — jeder Beitrag hilft
- ⭐ **Repo Sternchen geben** — kostet nichts und hilft mehr Menschen, pi-web zu entdecken
- 📢 **Mit Freunden & Familie teilen** — wenn du jemanden kennst, der pi-web lieben würde, schicke es ihm

Kannst du nicht sponsern? Keine Sorge — ein Sternchen und ein Beitrag tragen weit. Danke, dass du hier bist. 🙏

---

Guten Coding! 🚀
