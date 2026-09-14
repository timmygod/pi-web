# Warum pi-web?

Ich bin irgendwie süchtig nach Claude Code. Ich benutze es ständig. Wenn ich nicht vor dem Computer sitze, denke ich daran. Ich habe das Gefühl, ich verbrenne nicht genug Tokens. Es war die Anfangszeit von Claude Code. Und ich dachte mir, warum kann ich nicht vom Handy aus weitermachen? Ich habe Termius eingerichtet und es hat mir nicht wirklich gefallen.

Ich fing an, mein eigenes zu entwickeln und hörte auf, als Claude seine Claude Code Mobile-App vorstellte.

Dann bekam ich einen Bandscheibenvorfall und konnte nicht wirklich viel machen. Die Zeit verging und ich fühlte mich etwas erholt und wollte mein Claude Code via Web/PWA-Projekt fortsetzen.

Dann begann Claude Code, die Nutzung außerhalb ihres eigenen Harness zu verbieten. Und ich habe das Gefühl, es ist es nicht wert.

Dann fand ich pi.dev und erkundete es ein wenig, aber tauchte nicht wirklich ein. Ich las darüber, schaute Videos dazu und beschloss, es richtig auszuprobieren, und jetzt bin ich voll und ganz bei pi.

Da es Open Source ist, finde ich, es lohnt sich, dafür zu entwickeln. Ich habe auch verschiedene Anbieter zur Auswahl. Ich habe auch das Gefühl, dass es nicht nachhaltig ist, sich auf einen Anbieter/ein Modell wie Anthropic/Claude zu verlassen.

Also baue ich es hier.

## Warum ein lokales Modell ein anderes Betriebsprofil benötigt

Die ursprüngliche pi-web-Erfahrung ist eine hervorragende Grundlage, aber lokale Inferenz weist andere Fehlermodi auf als ein typisches gehostetes Modell. Ein lokales Modell kann sich stark verlangsamen, wenn der Kontext wächst, sich begrenzten Speicher mit dem Rest des Systems teilen, nach der Erzeugung von nur Reasoning stoppen oder einen langen Durchlauf aufgrund eines vorübergehenden lokalen Transportfehlers verlieren. Diese Fälle genau wie Cloud-Fehler zu behandeln, lässt die Benutzeroberfläche kompatibel erscheinen, während die tatsächliche Sitzung fragil bleibt.

Diese Ausgabe geht das Problem schichtweise an:

1. **Upstream zuerst bewahren.** Gemeinsame UI- und Sitzungsverhalten stammen weiterhin von pi-web; lokale Änderungen sind hinter effektivem Local Mode isoliert.
2. **Vorbeugen vor Wiederherstellen.** Eine prozentuale 65%-Kontextgrenze wird vor späteren Provider-Aufrufen durchgesetzt, einschließlich Aufrufen innerhalb langer Tool-Schleifen.
3. **Nur mit Nachweis wiederherstellen.** Automatische Fortsetzung ist auf erkannte Kontext-, Transport- und Thinking-only-Vorfälle beschränkt, nicht auf Authentifizierungs-, Quota- oder beliebige Provider-Fehler.
4. **Jede autonome Aktion begrenzen.** Wiederherstellungsvorfälle werden dedupliziert, Fortschritt ist vor einer weiteren Rettung erforderlich, und der Start berücksichtigt höchstens eine kürzlich aktive Local-Sitzung.
5. **Manuellen Ausstieg bewahren.** Force Compact fasst zusammen, statt den Verlauf zu löschen, sodass der Benutzer eine Sitzung retten kann, ohne so zu tun, als hätte der Kontext nie existiert.
6. **Cloud-Kompatibilität schützen.** Cloud Mode behält die Upstream-Semantik und -Steuerungen bei; lokale Modell-Optimierungen definieren Cloud-Sitzungen nicht stillschweigend neu.

Das ist der wahre Unterschied in diesem Fork: Er behandelt lokale Inferenz als eine eigenständige operative Umgebung, nicht einfach als einen weiteren Modellnamen in einem Dropdown-Menü.
