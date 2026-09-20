<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · **Deutsch** · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Steuere deinen [pi](https://pi.dev) Coding-Agenten von deinem Smartphone, Tablet oder Laptop aus — an jedem Ort in deinem Netzwerk oder remote über Tailscale.

Es ist eine vollständige PWA, die du installieren und wie eine native App auf jedem Gerät verwenden kannst. Verstehe sie als deinen persönlichen KI-Arbeitsplatz — ähnlich wie Closes Cowork, aber mit anderen Modellen — chatte über verschiedene Modelle hinweg, codest von deinem Smartphone oder verwandle sie in einen [persönlichen Assistenten](../en/personal-assistant.md), der auf deinem Rechner lebt.

Gestalte sie nach deinem Geschmack: Wechsle Themes und Schriftarten und verwende sie in deiner eigenen Sprache — pi-web wird mit mehreren Sprachen ausgeliefert und du kannst deine eigene hinzufügen. Weitere Funktionen sind in Arbeit, aber es wird nicht aufgebläht: Alles, was du nicht brauchst, kannst du in den Einstellungen deaktivieren.

</div>

## Warum diese lokale-Modell-Edition?

Das ursprüngliche pi-web bleibt die Grundlage im Upstream für gemeinsame Funktionen und
Korrekturen. Diese Edition behält diese Erfahrung bei und fügt dann eine Zuverlässigkeitsschicht für
Modelle hinzu, die auf deinem eigenen Rechner oder woanders in deinem LAN laufen — dort, wo die
Generierung oft langsamer ist, der Speicher begrenzt ist und ein langer Kontext eine ansonsten gesunde
Sitzung zum Stillstand bringen kann.

| Bereich | Upstream pi-web | Diese Edition |
|------|-----------------|--------------|
| Modell-/Runtime-Regelung | Standard-Verhalten von pi-web | Pro-Sitzungs-Modus **Auto / Local / Cloud** mit Endpunkt-bewusster lokaler Erkennung und einer persistenten manuellen Übersteuerung |
| Umgang mit langem Kontext | Normales pi-Compaction-Verhalten | Local Mode komprimiert proaktiv bei **65 %** und prüft erneut innerhalb langer Tool-Call-Loops, bevor die nächste Modell-Anfrage gesendet wird |
| Compaction-Sicherheit | Standard-Zusammenfassungen | Begrenzte rollende Checkpoints, eine engere Neuformulierung für ungültige/abgekappte Ausgaben und Erkennung von Nicht-Fortschritt statt endlosem erneuten Kompaktieren |
| Unterbrochene Abläufe | Normale Worker- und Fehlerbehandlung | Begrenzte Wiederherstellung für Kontextüberlauf, nur-Denken-Stops und bestimmte Transport-Unterbrechungen, mit persistenten Loop-Brechern |
| Manuelle Rettungsaktion | Standard-Kontextdetails | **Force Compact** bleibt als expliziter Wiederherstellungsweg verfügbar, ohne das Gespräch zu löschen |
| Kompatibilität und Releases | Originelles Projekt und Release-Linie | Nur-lokale Schutzmaßnahmen bleiben hinter Local Mode; Cloud Mode bewahrt das Upstream-Verhalten, und Upstream-Änderungen werden hier unabhängig geprüft und freigegeben |

Dies ist kein Rewrite und keine Ersetzung für Upstream. Es ist ein bewusst
gewartetes Betriebsprofil für Menschen, die Datenschutz und Kontrolle durch lokale Modelle wollen,
ohne empfindliche langlaufende Sitzungen in Kauf zu nehmen. Siehe die
[Benutzeranleitung](../en/README.md) für den nutzerseitigen Arbeitsablauf und
[lokale-Modell-Edition-Entwicklung](../../docs/dev/local-llm-development.md) für die
Implementierung und die Synchronisationsrichtlinie.

> [!TIP]
> Neu hier? **[Benutzeranleitung lesen →](../en/README.md)** für einen vollständigen Rundgang durch Funktionen, Installationsschritte und Tipps. ([Andere Sprachen →](../README.md))

## Screenshots

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## Wie alles zusammenpasst

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

- **pi** schreibt während der Arbeit Konversations-JSONL nach `~/.pi/agent/sessions/`.
- **pi-web** ist ein Go-Server, der diese Dateien liest, sie im Browser rendert und Live-Updates per SSE streamt.
- Worker von **pi --mode rpc** verarbeiten im Browser initiierten Chat — einer pro Sitzung, nach 10 Min. Inaktivität beendet.
- **fsnotify** beobachtet den Sitzungsverzeichnis, damit der Browser innerhalb von Millisekunden nach neuer Ausgabe neu lädt.
- **Tailscale Serve** veröffentlicht den localhost-Server als HTTPS-Endpunkt in deinem Tailnet.

## Installation

```bash
pi install npm:@timmygod/pi-web-local
```

Das war's — es lädt das passende Binary herunter, richtet den Autostart ein und registriert die Befehle `/web`, `/pi-web`, `/remote` und `/refresh`.

Nach der Installation öffne `http://127.0.0.1:31415` in deinem Browser. Verwende aus pi heraus `/web`, um die aktuelle Sitzung sofort in deinem Browser zu öffnen. Wenn Tailscale auf deinem Rechner läuft, veröffentlicht pi-web automatisch einen HTTPS-Endpunkt in deinem Tailnet — verwende aus pi heraus `/remote`, um einen QR-Code und eine URL für jedes Gerät in deinem Tailnet zu erhalten.

> **Remote-Zugriff unter macOS:** Installiere und öffne Tailscale interaktiv, bestätige die Administrator-Aufforderung und melde dich an. Führe dann `/pi-web restart` aus, gefolgt von `/remote`.

Für manuelle Installationen, Binary-Downloads oder das Kompilieren aus dem Quellcode siehe [user-docs/install.md](../en/install.md).

## Pi-Integration

Nach `pi install npm:@timmygod/pi-web-local` erhältst du:

| Befehl | Funktion |
|---------|--------------|
| `/web` | Öffne die aktuelle Sitzung in deinem Browser (SSH-bewusst: überspringt den Browser und zeigt nur die URL) |
| `/pi-web` | Zeigt Status, Version, startet/stoppne/neustartet den Server oder aktualisiert |
| `/remote` | Zeigt einen QR-Code und eine URL für den Remote-Zugriff über Tailscale |
| `/refresh` | Zieht neue Nachrichten, die aus Remote-Browsern geschrieben wurden, zurück in die Terminal-Sitzung |

Das **automatische Benennen** von Sitzungen ist direkt in pi-web integriert und wird auf der `/settings`-Seite konfiguriert. Es ist **standardmäßig aktiviert** und benennt Sitzungen automatisch. Du kannst wählen:

- **Wann benannt wird** — einmal pro Sitzung oder bei jeder neuen Nachricht (Standard).
- **Titel-Modell** — standardmäßig ein kostenloser, sofortiger **eingebauter Wort-Heuristik (keine KI)**, oder wähle ein Modell (z. B. ein kleines/schnelles) für intelligentere, vom Modell geschriebene Titel.

Das Paket installiert außerdem das pi-web-Binary nach `~/.pi/agent/bin/pi-web` und richtet den Autostart bei der Anmeldung ein.

## Autostart bei der Anmeldung

Der Befehl `pi install npm:@timmygod/pi-web-local` richtet dies automatisch ein:

| Betriebssystem | Mechanismus |
|----|-----------|
| macOS | launchd plist unter `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd-Benutzerdienst unter `~/.config/systemd/user/pi-web.service` |
| Windows | `HKCU`-Lauf-Schlüsseleintrag, der einen versteckten Starter in `~/.config/pi-web/` startet |

Um einen Token für den Remote-Zugriff zu setzen, erstelle `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Für weitere Details (manuelle Einrichtung, eigene Ports, nicht-Loopback-Binds) siehe [user-docs/install.md](../en/install.md).

## Entwicklung

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Für Upstream-Synchronisierung, lokale-Modell-Tests und den parallelen
Release-Arbeitsablauf siehe [lokale-Modell-Edition-Entwicklung](../../docs/dev/local-llm-development.md).
