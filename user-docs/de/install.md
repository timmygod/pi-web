# Installation & Nutzung

## Funktionen

### Fernsteuerung

- Setze jede Sitzung aus dem Browser mit Text- oder Bildanhängen fort
- Starte eine brandneue Sitzung für jeden Projektpfad, direkt aus der Web-Oberfläche
- Modellschalter und Denkstufe-Selektor im Browser, pro Sitzung
- Worker-Status pro Sitzung (idle / running / error) mit automatischer Wiederherstellung bei Abstürzen
- Mehrere Sitzungen laufen parallel — starte Arbeit in einer, beobachte wie eine andere streamt
- `PI_WEB_TOKEN` für sichere LAN-Exposition — standardmäßig erforderlich für jede explizite Nicht-Loopback-Bindung

### Sitzungen lesen

- Durchsuche Sitzungen über Projekte hinweg mit Filtern, Suche und vollständiger Branch-Navigation
- Live-Inkrementelle Aktualisierungen, während pi noch läuft (via fsnotify; ~ms Latenz)
- Follow-Modus zum Verfolgen aktiver Sitzungen
- Deep Links zu einzelnen Nachrichten
- Lade eine Sitzung als JSONL herunter
- Teile statische Snapshots als geheime GitHub Gists
- `/web`, `/remote`, `/refresh`, `/pi-web token` und `/pi-web set-token` pi-Erweiterungen zum Öffnen von Sitzungen, Remote QR, Session-Sync und Token-Verwaltung
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`), damit eine Sitzung Pläne, das Projekt-Skriptblock (scratchpad) und Einstellungen in natürlicher Sprache verwalten kann

## Wähle den Sitzungsmodus

Diese Edition verwendet die in pi bereits konfigurierten Provider und Modelle; Local Mode
ist eine Laufzeit-Richtlinie, kein separater Modell-Installer und kein zweiter API-Key-Screen.
Wähle beim Erstellen einer Sitzung einen Modus, oder ändere ihn, nachdem der aktuelle Lauf abgeschlossen ist:

| Modus | Wähle ihn, wenn | Verhalten |
|------|-------------|----------|
| **Auto** | Du willst, dass pi-web entscheidet | Löst lokale/LAN-Endpunkte aus den Provider-Metadaten auf, wo möglich; andernfalls behält den normalen Weg bei |
| **Local** | Das Modell läuft auf dieser Maschine oder in deinem LAN | Aktiviert die 65 %-Kompaktierungsgrenze, begrenzte Checkpoints, Force Compact und geschützte automatische Wiederherstellung |
| **Cloud** | Das gewählte Modell gehostet ist und das Upstream-Verhalten beibehalten sollte | Behält lokale Kompaktierungs- und Wiederherstellungspolitik außerhalb der Sitzung |

Manuelle Local- oder Cloud-Auswahl schlägt die automatische Erkennung und bleibt über
Reloads und Neustarts hinweg erhalten. Eine laufende Sitzung lehnt Modusänderungen ab, bis ihr
Worker abgeschlossen ist, sodass der in der UI angezeigte Modus immer mit der tatsächlich verwendeten Politik übereinstimmt.

## Voraussetzungen

- [Go](https://go.dev) 1.25+ (nur für den Build aus dem Quellcode)
- `pi` in deinem `PATH` für Browser-Chat/Modellschaltung
- Optional: `gh` zum Teilen
- Unter Windows: pi benötigt eine bash-Shell für sein Shell-Tool — [Git for Windows](https://git-scm.com/download/win) genügt (siehe pi's Windows-Dokumentation)

## Installation

### Pi-Paket (empfohlen)

```bash
pi install npm:@timmygod/pi-web-local
```

Dieser einzige Befehl:
- Installiert das npm pi-Paket unter pi's Paketverzeichnis
- Führt das Paket-`postinstall`-Skript aus (`install.sh` oder `install.ps1` unter Windows)
- Lädt das passende pi-web-Binary für deine Paketversion und Plattform von GitHub Releases herunter
- Installiert es nach `~/.pi/agent/bin/pi-web` (`pi-web.exe` unter Windows)
- Richtet Auto-Start beim Login ein (launchd unter macOS, systemd unter Linux, einen Run-Key-Launcher unter Windows)
- Registriert die `/web`, `/remote`, `/refresh`, `/pi-web token` und `/pi-web set-token` pi-Befehle

Das automatische Benennen von Sitzungen ist in pi-web eingebaut (nicht in der Erweiterung) und wird auf der `/settings`-Seite konfiguriert. Es ist standardmäßig aktiv: pi-web benennt Sitzungen automatisch mit einer eingebauten, kostenlosen Wort-Heuristik (ohne KI) und benennt sie bei jeder neuen Nachricht neu. Du kannst auf einmaliges Benennen pro Sitzung umschalten und/oder ein Modell auswählen, das intelligentere Titel statt der Heuristik schreibt.

Unter Linux wird Auto-Start als Benutzer-systemd-Service unter `~/.config/systemd/user/pi-web.service` konfiguriert. Der Installer schreibt dessen `ExecStart` mit dem tatsächlichen installierten Binary-Pfad neu. Wenn Tailscale zur Laufzeit verfügbar ist, publiziert pi-web den localhost-Server über Tailscale Serve HTTPS. Wenn Benutzer-systemd nicht verfügbar ist, führe es manuell aus: `~/.pi/agent/bin/pi-web -o`.

Um nur für ein bestimmtes Projekt zu installieren (mit deinem Team über `.pi/settings.json` geteilt):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Dann starte pi neu (oder führe `/reload` aus) und benutze `/web`, `/pi-web`, `/remote`, `/refresh`. Verwalte dein Zugangs-Token mit `/pi-web token` und `/pi-web set-token`.

Falls npm beim Umbenennen von `@timmygod/pi-web-local` mit `ENOTEMPTY` abbricht, entferne npm's veraltete versteckte Backup-Verzeichnisse und installiere das Paket erneut:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Schnelle Installation (keine Build-Tools erforderlich)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Dies lädt das neueste pi-web-Binary herunter, installiert es nach `/usr/local/bin` (`~/.pi/agent/bin` unter Windows) und richtet Auto-Start beim Login ein. Kein Go, Node oder pi erforderlich.

### Binary herunterladen

Fertig kompilierte Binaries sind an jedem [GitHub Release](https://github.com/timmygod/pi-web/releases) angehängt.

```bash
# macOS (Apple Silicon)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-arm64
chmod +x pi-web

# macOS (Intel)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-amd64
chmod +x pi-web

# Linux (amd64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-amd64
chmod +x pi-web

# Linux (arm64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-arm64
chmod +x pi-web
```

```powershell
# Windows (x64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-amd64.exe

# Windows (ARM64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-arm64.exe
```

Dann verschiebe es in deinen PATH:

```bash
cp pi-web ~/.pi/agent/bin/
# oder systemweit:
sudo cp pi-web /usr/local/bin/
```

### Aus dem Quellcode bauen

Diese Checkout ist die lokalmodellierte Edition von pi-web. Der normale Build erzeugt
die Web-Anwendung und den Backend zusammen; die Sicherheitsmaßnahmen für lokale Modelle werden
zur Laufzeit durch den effektiven Local Mode der Sitzung aktiviert, nicht durch ein separates Binary.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # baut das Vite-Bundle und bindet es dann in das Go-Binary ein

# optional: in den PATH legen
cp pi-web ~/.pi/agent/bin/
```

Das Frontend-Bundle wird von `web/assets_embed.go` eingebettet, daher benötigt `go build`
zunächst `web/dist`. `make build` führt beide Schritte nacheinander aus; wenn du manuell
baust, führe vor `go build ./cmd/pi-web` aus: `npm --prefix web install && npm --prefix web run build`.

Für den gepflegten Fork-Workflow, die Upstream-Synchronisierung und die Local Mode
Verifikations-Checkliste, siehe [die lokalmodellen Entwickler-Notizen](../../docs/dev/local-llm-development.md).

### Neben einer installierten Instanz entwickeln

Lass die installierte Instanz auf Port `31415` laufen und starte dann die
Quellcode-Checkout im Entwicklungsmodus:

```bash
make dev
```

Öffne `http://127.0.0.1:31416`. `make dev` setzt die interne `PI_WEB_DEV=1`
Entwicklungsumgebung, sodass die Quellcode-Checkout Sitzungen, Einstellungen und
SQLite-Daten mit der installierten Instanz teilt, während sie eine separate Entwicklungs-
Laufzeit-Sperre und Statusdatei beibehält. Regelmäßig installierte und manuell gestartete
Instanzen bleiben unverändert und behalten das ursprüngliche Einzelinstanz-Verhalten bei.

Um doppelte autonome Arbeit zu verhindern, führt der Entwicklungsmodus die
Pläne-Schleife, Chat-Queue-Drainer, automatisches Benennen oder Push-Benachrichtigungen nicht aus.
Direkte Anfragen über die Entwickler-Oberfläche funktionieren weiterhin. Steuere dieselbe
Chat-Sitzung nicht gleichzeitig von beiden Instanzen; jeder Prozess hat seinen eigenen RPC Worker
Manager.

`make dev` benötigt [Air](https://github.com/air-verse/air) für Go-Hot-Reload:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` ist Entwicklungs-Harness-Verdrahtung, kein unterstützter Produktions-
Multi-Instanz-Modus.

## Deinstallation

```bash
pi remove npm:@timmygod/pi-web-local
```

Dies führt das Paket-`preuninstall`-Skript aus (`uninstall.sh` oder `uninstall.ps1`
unter Windows), welches die laufende Instanz stoppt und entfernt:

- das pi-web-Binary (`~/.pi/agent/bin/pi-web` oder `/usr/local/bin/pi-web` für eigenständige Installationen)
- die Versionsdatei (`~/.pi/agent/pi-web-version`)
- die Laufzeit-Statusdatei (`~/.pi/agent/pi-web/pi-web-state.json`)
- die Auto-Start-Konfiguration (launchd-Plist unter macOS, systemd-Benutzerservice unter Linux, Run-Key-Eintrag + Launcher-Skripte unter Windows)

Deine Daten werden beibehalten, damit eine spätere Neuinstallation dort weitermacht, wo du aufgehört hast:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, deine
Sitzungsdateien unter `~/.pi/agent/sessions/` und `~/.config/pi-web/env` (einschließlich
`PI_WEB_TOKEN`). Entferne diese manuell, wenn du einen sauberen Neustart möchtest.

## Nutzung

```bash
# Auf dem Standardport starten (31415)
pi-web

# Starten und einen Browser öffnen
pi-web -o

# Custom Port
pi-web -p 8080

# Bind-Host überschreiben (Loopback ist standardmäßig nicht authentifiziert)
pi-web --host 127.0.0.1

# Nicht-Loopback-Bindung erfordert ein Token — pi-web verweigert sonst den Start
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Standardmäßig bindet pi-web an `127.0.0.1`. Wenn Tailscale mit MagicDNS läuft **und `PI_WEB_TOKEN` gesetzt ist**, führt pi-web zusätzlich `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` aus und zeigt die HTTPS-Tailnet-URL an. Ohne Token bleibt pi-web nur auf Loopback und überspringt Tailscale Serve, sodass Tailnet-Peers den Agenten nicht unauthentifiziert erreichen können. Jede explizite Nicht-Loopback-Bindung erfordert ebenfalls `PI_WEB_TOKEN` gesetzt; übergebe `--insecure`, um dies für lokale Tests zu überschreiben.

## Fernzugriff

Lass pi-web lokal lauschen und verwende die angezeigte Tailscale-HTTPS-URL von deinem Telefon oder Laptop im Tailnet.

Unter macOS installiere und öffne Tailscale interaktiv, bestätige den Administrator-Hinweis und melde dich an. Führe dann `/pi-web restart` aus, gefolgt von `/remote`.

Unter Linux erlaube deinem Benutzer, Tailscale zu verwalten, bevor du pi-web installierst/aufstartest, andernfalls kann `tailscale serve` sudo erfordern und Auto-Start fehlschlagen:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Starte pi-web mit einem Token, damit es den Tailscale-HTTPS-Endpunkt publiziert
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Von einem anderen Tailscale-verbindeten Gerät aus die angezeigte
#    "Tailscale HTTPS" URL öffnen und das Token einmalig eingeben.
```

> Standardmäßig verweigert pi-web die Bindung an eine Nicht-Loopback-Adresse, wenn `PI_WEB_TOKEN` nicht gesetzt ist — jeder, der die gebundene Adresse erreichen kann, könnte andernfalls Sitzungen ansehen und Anweisungen an pi senden. Um diese Schutzmaßnahme für lokale Netzwerktests zu überschreiben, übergebe `--insecure`. **Nutze `--insecure` nicht auf Tailscale oder jeder Adresse, die von außerhalb deiner Maschine erreichbar ist.**
>
> Clients können das Token über den `Authorization: Bearer <token>`-Header, den `X-Pi-Token`-Header oder einmalig über `?token=<token>` (setzt ein `pi_token`-Cookie für nachfolgende Anfragen) übergeben. Tokens, die über `?token=` übergeben werden, landen im Browserverlauf, in Server-Access-Logs und in `Referer`-Headern von allen Links auf der Seite — bevorzuge die Header-Form für alles außer dem initialen Lesezeichen.

## Browser-Chat

Öffne eine Sitzungsseite und benutze die Composer unten, um genau diese Sitzung fortzusetzen.

- `Enter` sendet, `Shift+Enter` fügt einen Zeilenumbruch ein
- Ziehe und lasse Bilder fallen oder füge sie direkt in die Composer ein
- Die Modellauswahl und die Denkstufe-Auswahl befinden sich im Header — Änderungen wirken sich sofort auf den zugrunde liegenden pi-Worker aus
- Jede aktive Sitzung erhält einen eigenen dedizierten `pi --mode rpc`-Worker, sodass verschiedene Sitzungen sich gegenseitig nicht blockieren

## Sitzungen teilen

Klicke auf einer Sitzungsseite auf **Share**, um einen geheimen GitHub Gist zu erstellen.

Voraussetzungen:
- `gh` installiert
- `gh auth login` abgeschlossen

Das Teilen liefert:
- die geheime Gist-URL
- eine Preview-URL unter `https://pi.dev/session/#<gistId>`

Geteilte Gists sind Snapshots und aktualisieren sich nicht live.

## Auto-Start beim Login

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Installiere den systemd-Benutzerservice
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optional: setze dein PI_WEB_TOKEN für Nicht-Loopback-Bindungen
# (oder benutze /pi-web set-token <token> aus pi heraus)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Aktivieren und starten
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Status prüfen
systemctl --user status pi-web.service

# Logs anzeigen
journalctl --user -u pi-web.service -f
```

> Damit der Service beim Boot startet (vor dem Login), benutze stattdessen einen Systemservice:
> kopiere `init/pi-web.service` nach `/etc/systemd/system/` und benutze `sudo systemctl`.

### Windows

Der Installer richtet dies automatisch ein, ohne Administratorrechte zu benötigen: ein
`pi-web`-Eintrag unter `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
startet `~/.config/pi-web/pi-web-start.vbs` beim Login, welches das Binary
versteckt startet (kein Konsolenfenster), nachdem es `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...) geladen hat.

Um es manuell zu verwalten:

```powershell
# Starten / stoppen
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Auto-Start entfernen
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Es gibt keine Service-Aufsicht unter Windows: Wenn pi-web abstürzt, bleibt es aus,
bis zum nächsten Login (launchd/systemd starten es auf den anderen
Plattformen automatisch neu).
