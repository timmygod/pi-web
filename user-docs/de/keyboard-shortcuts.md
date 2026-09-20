# Tastenkürzel

Diese Kürzel gelten für die Local-Model-Edition von pi-web. Die editionspezifische
Laufzeit ist in [Local-Model-Edition Entwicklung](../../docs/dev/local-llm-development.md) dokumentiert.

## Index-Seite (`/`)

### Seitennavigation (vim-Stil)

Die gleichen vim-Stil-Kürzel funktionieren auf allen Seiten, wenn der Fokus **nicht** in einem Eingabefeld, Textbereich oder contenteditable-Element liegt.

| Kürzel | Aktion |
|----------|--------|
| `j` | 300px nach unten scrollen |
| `k` | 300px nach oben scrollen |
| `g g` | Zum Seitenanfang scrollen |
| `G` (Shift+G) | Zum Seitenende scrollen |
| `Escape` | Aktives Eingabefeld entfokusieren, damit j/k-Navigation funktioniert |

### Index-Befehle

| Kürzel | Kontext | Aktion |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Seitenebene | Such-/Sitzungspalette öffnen |
| `⌘⇧L` / `Ctrl+Shift+L` | Seitenebene | Systemthema wechseln (hell/dunkel) |
| `Escape` | Seitenebene | Palette, Menü oder Modal schließen |
| `Enter` | Neues-Sitzung-Eingabefeld | Neue Sitzung erstellen |

> `⌘K` / `Ctrl+K` ist zugleich das Chrome-Kürzel für „Fokus auf Adressleiste". Der Browser kann es abfangen, solange der Fokus nicht in einem Texteingabefeld liegt.

## Sitzungs-Detailseite (`/session?id=...`)

### Seitennavigation (vim-Stil)

Diese funktionieren sowohl auf der Index- als auch auf der Sitzungsseite, wenn der Fokus **nicht** in einem Eingabefeld, Textbereich oder contenteditable-Element liegt.

| Kürzel | Aktion |
|----------|--------|
| `j` | 300px nach unten scrollen |
| `k` | 300px nach oben scrollen |
| `g g` | Zum Seitenanfang scrollen |
| `G` (Shift+G) | Zum Seitenende scrollen |
| `I` (Shift+I) | Fokus auf das Chat-Autoren-Textfeld setzen |
| `Escape` | Aktives Eingabefeld entfokusieren, damit j/k-Navigation funktioniert |

### Seitenleiste & Navigation

| Kürzel | Kontext | Aktion |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Seitenebene | Seitenleiste ein-/ausblenden |
| `⌘K` / `Ctrl+K` | Seitenebene | Sitzungslisten-Palette öffnen |
| `⌘T` / `Ctrl+T` | Seitenebene | Neue Sitzung |
| `⌘⇧L` / `Ctrl+Shift+L` | Seitenebene | Systemthema wechseln (hell/dunkel) |
| `⌘⇧N` / `Ctrl+Shift+N` | Seitenebene | Scratchpad-/Notizen-Seitenleiste ein-/ausblenden |

> `⌘K` und `⌘T` sind zugleich Browser-Kürzel (Fokus auf Adressleiste / neuer Tab). Der Browser kann sie abfangen, solange der Fokus nicht in einem Texteingabefeld liegt.

### Chat-Autor (Composer)

| Kürzel | Kontext | Aktion |
|----------|---------|--------|
| `Enter` | Chat-Textfeld | Nachricht senden |
| `Shift+Enter` | Chat-Textfeld | Zeilenumbruch einfügen |
| `Shift+Tab` | Chat-Textfeld | Zum nächsten Denk-Stufe wechseln (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Chat-Textfeld | Modellauswahl-Popup öffnen (eingeben zum Filtern, Enter zum Auswählen, Fokus kehrt zum Textfeld zurück) |

### Eintrags-Sichtbarkeits-Schalter

| Kürzel | Kontext | Aktion |
|----------|---------|--------|
| `t` | Wenn der Fokus **nicht** in einem Eingabefeld/Textbereich liegt | Denk-Sichtbarkeit umschalten |
| `o` | Wenn der Fokus **nicht** in einem Eingabefeld/Textbereich liegt | Werkzeug-Sichtbarkeit umschalten |
| `p` | Wenn der Fokus **nicht** in einem Eingabefeld/Textbereich liegt | Werkzeug-Ausgaben umschalten |

### Paletten, Menüs & Blätter

| Kürzel | Kontext | Aktion |
|----------|---------|--------|
| `Escape` | Seitenebene | Jede geöffnete Palette, jedes Menü oder jeden Blatt schließen |
| `⌘K` / `Ctrl+K` | Seitenebene | Sitzungslisten-Palette öffnen |
| `ArrowUp` / `ArrowDown` | Sitzungslisten-Palette | Durch Sitzungs-Ergebnisse navigieren |
| `Enter` | Sitzungslisten-Palette | Ausgewählte (oder erste) Sitzung öffnen |
| `ArrowUp` / `ArrowDown` | Modellauswahl-Popup | Durch Modellliste navigieren |
| `Enter` | Modellauswahl-Popup | Hervorgehobenes Modell auswählen |
| `ArrowUp` / `ArrowDown` | Fork-Modal | Durch Nachrichten navigieren |
| `Enter` | Fork-Modal | Ab der hervorgehobenen Nachricht forken |
| `Tab` | Vollbild-Blatt | Fokus innerhalb des Blatts wechseln |
| `Escape` | Vollbild-Blatt | Blatt schließen |
