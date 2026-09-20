# Roadmap

Dieses Roadmap-Dokument gehört zur Local-Model-Edition von pi-web. Upstream-Funktionen
werden regelmäßig synchronisiert, während die Arbeit an lokaler Bereitstellung und Zuverlässigkeit
auf dieser Linie validiert und veröffentlicht wird; siehe [Entwicklung der Local-Model-Edition](../../docs/dev/local-llm-development.md).

pi-web ist für zwei Zielgruppen konzipiert:

- **Für Entwickler** — die in der Terminal-Umgebung leben, aber Sessions vom Mobilgerät aus fortsetzen, an einen Remote-Server übergeben oder laufende Aufgaben von überall im Auge behalten möchten.
- **Für Nicht-Entwickler** — die einfach eine schöne KI-App wollen, die funktioniert. Aufmachen, tippen, viben. Keine Terminal, kein SSH, keine Verwirrung. Wie die nutzerfreundlichsten KI-Tools, aber mit Modellwahl und Open-Source-Freiheit.

Hier ist, was kommt.

Diese Edition folgt dem Upstream pi-web auf einer separaten Release-Linie. Upstream-Funktionen
werden regelmäßig importiert; die Arbeit an der Zuverlässigkeit der Local-Model wird hier priorisiert
und validiert, ohne die Upstream-Release-Geschichte zu verändern.

---

## Jetzt (bereits verfügbar)

Alles, was in [der Funktions-Übersicht](README.md#what-you-can-do-with-pi-web) aufgeführt ist, ist heute live.

---

## Als Nächstes

| # | Funktion | Was es bewirkt |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram- und Discord-Bots** | Mit pi über Telegram oder Discord chatten — perfekt für persönliche Assistenten-Workflows unterwegs. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Nutzungseinsichten** | Token-Tracking, Kostenabschätzung, Session-Analysen — erfahre, wie du pi nutzt. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Konfigurierbare Standards** | Lege deine bevorzugte Sichtbarkeit für Denken, Tools und Tool-Ausgaben über alle Sessions fest. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steuerung / Warteschlange** | Sende Folgeanweisungen, während pi noch läuft — leite es auf dem Weg. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact`-Befehl** | Kompactiere lange Konversationen direkt aus der Web-UI, ohne Terminal. |

---

## Geplant

| # | Funktion | Was es bewirkt |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **Datei-Explorer und Git-Diff** | Durchsuche den Projektdateibaum und sieh Git-Änderungen direkt in pi-web. Opt-in, damit es aus dem Weg bleibt. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Plane Prompts, die automatisch ausgeführt werden — tägliche Standups, Morgen-Zusammenfassungen, wiederkehrende Aufgaben. Aus Sicherheitsgründen durch Admins geschützt. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Anpassbare Shortcuts** | Ordne jede Tastenkombination neu zu, um deiner Muskel记忆 zu entsprechen. |

---

## Vision

Das Langzeitziel: pi-web sollte **die Oberfläche für pi** sein — für alle.

- **Nicht-Entwickler** öffnen es wie jede andere App. Modell wählen. Tippen. Fertig. Keine Kommandozeile jemals.
- **Entwickler** erhalten tiefe Integration — Remote-Handoff, Multi-Session-Dashboards, Git-bewusste Durchsuchung, Messaging-Bots.
- **Alle** erhalten Modellfreiheit, Open-Source-Transparenz und eine UI, die an jeder Stelle durchdacht wirkt.

---

> 💡 Hast du eine Idee? [Eröffne ein Issue](https://github.com/timmygod/pi-web/issues/new) oder schließe dich der Diskussion an.
