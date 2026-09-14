<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dw/@timmygod/pi-web-local?label=downloads/wk&color=2ea043&cacheSeconds=86400)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb&cacheSeconds=86400)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · [Français](README.fr.md) · **Deutsch** · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Steuere deinen [pi](https://pi.dev) Coding Agent von deinem Handy, Tablet oder Laptop — überall in deinem Netzwerk oder remote über Tailscale.

Es ist eine vollständige PWA, sodass du sie installieren und wie eine native App auf jedem Gerät nutzen kannst. Betrachte es als deinen persönlichen KI-Arbeitsbereich — wie Claude's Cowork, aber mit verschiedenen Modellen — chatte modellübergreifend, programmiere von deinem Handy aus oder mach es zu einem [persönlichen Assistenten](../en/personal-assistant.md), der auf deinem Rechner lebt.

Mach es zu deinem: Wechsle Themes und Schriften und nutze es in deiner eigenen Sprache — pi-web wird mit mehreren Sprachen ausgeliefert und du kannst eigene hinzufügen. Weitere Funktionen sind in Arbeit, aber es wird nicht aufgebläht: Alles, was du nicht brauchst, kannst du in den Einstellungen deaktivieren.

</div>

## Warum diese Lokales-Modell-Ausgabe?

Das ursprüngliche pi-web bleibt die Upstream-Grundlage für gemeinsame Funktionen und Korrekturen. Diese Ausgabe behält diese Erfahrung bei und fügt eine Zuverlässigkeitsschicht für Modelle hinzu, die auf Ihrem eigenen Rechner oder anderswo in Ihrem LAN laufen – wo die Generierung oft langsamer ist, der Speicher begrenzt ist und ein langer Kontext eine ansonsten gesunde Sitzung blockieren kann.

| Bereich | Upstream pi-web | Diese Ausgabe |
|------|-----------------|--------------|
| Modell-/Laufzeitrichtlinie | Standard-pi-web-Verhalten | **Auto / Local / Cloud**-Modus pro Sitzung, mit endpunkt-bewusster lokaler Erkennung und einer dauerhaften manuellen Überschreibung |
| Langkontext-Verarbeitung | Normales pi-Kompaktierungsverhalten | Der Local Mode komprimiert proaktiv bei **65%** und prüft innerhalb langer Tool-Call-Schleifen erneut, bevor eine weitere Anbieteranfrage gestellt wird |
| Kompaktierungssicherheit | Standard-Zusammenfassungen | Begrenzte rollierende Prüfpunkte, eine engere Umschreibung für ungültige/begrenzte Ausgaben und keine Fortschritts-Erkennung anstelle endloser Re-Kompaktierung |
| Unterbrochene Läufe | Normale Worker- und Fehlerbehandlung | Begrenzte Wiederherstellung bei Kontextüberlauf, reinen Denk-Stopps und ausgewählten Transportunterbrechungen, mit dauerhaften Schleifenunterbrechern |
| Manuelle Rettung | Standard-Kontextdetails | **Force Compact** bleibt als expliziter Wiederherstellungspfad verfügbar, ohne die Konversation zu löschen |
| Kompatibilität und Releases | Originalprojekt und Release-Linie | Lokale Schutzmaßnahmen bleiben hinter dem Local Mode; der Cloud Mode bewahrt das Upstream-Verhalten, und Upstream-Änderungen werden hier unabhängig überprüft und veröffentlicht |

Dies ist keine Neuschreibung oder ein Ersatz für Upstream. Es ist ein bewusst gepflegtes Betriebsprofil für Personen, die Privatsphäre und Kontrolle lokaler Modelle wünschen, ohne fragile langlaufende Sitzungen zu akzeptieren. Siehe den [Benutzerleitfaden](../en/README.md) für den benutzerorientierten Workflow und [Entwicklung der Lokales-Modell-Ausgabe](../../docs/dev/local-llm-development.md) für die Implementierung und die Synchronisierungsrichtlinie.

> [!WARNING]
> pi-web befindet sich derzeit in der **beta**-Phase. Dinge werden sich ändern und kaputtgehen!

> [!TIP]
> Neu hier? **[Lies die Bedienungsanleitung →](../en/README.md)** für eine vollständige Tour der Funktionen, Installationsschritte und Tipps. ([Andere Sprachen →](../README.md))

## Screenshots

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Desktop</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile PWA" width="90%" /><br />
  <em>Mobile PWA</em>
</div>

## Wie es zusammenpasst

```
 pi (Terminal)                Browser (Handy / Tablet / Laptop)
      │                                │
      │  schreibt JSONL               │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP-Server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (Chat-Worker       (Live-Neuladen)    (Remote-HTTPS
             pro Sitzung)                          via MagicDNS)
```

- **pi** schreibt während der Arbeit Konversations-JSONL nach `~/.pi/agent/sessions/`.
- **pi-web** ist ein Go-Server, der diese Dateien liest, sie im Browser rendert und Live-Updates via SSE streamt.
- **pi --mode rpc** Worker verarbeiten vom Browser initiierte Chats — einer pro Sitzung, werden nach 10 Min. Inaktivität beendet.
- **fsnotify** überwacht das Sitzungsverzeichnis, sodass der Browser innerhalb von Millisekunden nach neuer Ausgabe aktualisiert.
- **Tailscale Serve** veröffentlicht den localhost-Server als HTTPS-Endpunkt in deinem Tailnet.

## Installation

```bash
pi install npm:@timmygod/pi-web-local@beta
```

Das war's — es lädt die passende Binärdatei herunter, richtet den Autostart ein und registriert die Befehle `/web`, `/pi-web`, `/remote` und `/refresh`.

Nach der Installation öffne `http://127.0.0.1:31415` in deinem Browser. Verwende in pi `/web`, um die aktuelle Sitzung sofort im Browser zu öffnen. Wenn Tailscale auf deinem Rechner läuft, veröffentlicht pi-web automatisch einen HTTPS-Endpunkt in deinem Tailnet — verwende `/remote` in pi, um einen QR-Code und eine URL für jedes Gerät in deinem Tailnet zu erhalten.

> **Fernzugriff unter macOS:** Installiere und öffne Tailscale interaktiv, bestätige die Administratorabfrage und melde dich an. Führe anschließend `/pi-web restart` und danach `/remote` aus.

Für manuelle Installationen, Binär-Downloads oder das Bauen aus dem Quellcode siehe [user-docs/install.md](../en/install.md).

## Pi-Integration

Nach `pi install npm:@timmygod/pi-web-local@beta` erhältst du:

| Befehl | Was er macht |
|--------|--------------|
| `/web` | Öffnet die aktuelle Sitzung im Browser (SSH-fähig: überspringt den Browser und zeigt nur die URL) |
| `/pi-web` | Zeigt Status, Version, Start/Stopp/Neustart des Servers oder Update |
| `/remote` | Zeigt einen QR-Code und eine URL für den Remote-Zugriff über Tailscale |
| `/refresh` | Holt neue, von Remote-Browsern geschriebene Nachrichten zurück in die Terminal-Sitzung |

Die automatische **Titelvergabe** für Sitzungen ist in pi-web integriert und wird auf der `/settings`-Seite konfiguriert. Sie ist **standardmäßig aktiviert** und benennt Sitzungen automatisch. Du kannst wählen:

- **Wann ein Titel vergeben wird** — einmal pro Sitzung oder bei jeder neuen Nachricht (Standard).
- **Titelmodell** — standardmäßig eine kostenlose, sofortige **integrierte Wortheuristik (keine KI)** oder wähle ein Modell (z.B. ein kleines/schnelles) für intelligentere, modellgenerierte Titel.

Das Paket installiert außerdem die pi-web-Binärdatei nach `~/.pi/agent/bin/pi-web` und richtet den Autostart bei der Anmeldung ein.

## Autostart bei der Anmeldung

Der Befehl `pi install npm:@timmygod/pi-web-local@beta` richtet dies automatisch ein:

| Betriebssystem | Mechanismus |
|----------------|-------------|
| macOS | launchd plist unter `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | systemd user service unter `~/.config/systemd/user/pi-web.service` |

Um ein Token für den Remote-Zugriff festzulegen, erstelle `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=dein-token-hier
```

Für weitere Details (manuelle Einrichtung, benutzerdefinierte Ports, nicht-loopback-Bindungen) siehe [user-docs/install.md](../en/install.md).

## Entwicklung

```bash
make setup   # Frontend-Abhängigkeiten installieren und Go-Module herunterladen
make check   # Frontend testen/bauen + Go test/vet
make build   # ggf. setup, Frontend bauen, dann ./pi-web bauen
```
