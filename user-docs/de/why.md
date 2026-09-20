# Warum pi-web?

Ich bin sozusagen süchtig nach Claude Code. Ich nutze es ständig. Wenn ich nicht vor dem Computer sitze, dann denke ich über es nach. Ich habe das Gefühl, dass ich nicht genug Tokens verbruche. Es waren die Anfangstage von Claude Code. Und ich dachte mir, warum kann ich nicht von meinem Telefon aus weitermachen? Ich habe Termius eingerichtet und es hat mir nicht wirklich gefallen.

Ich begann, mir etwas Eigenes zu bauen, und hörte auf, als Claude seine Claude Code mobile app einführte.

Dann hatte ich einen Bandscheibenvorfall und konnte nicht wirklich viel tun. Die Zeit verging und ich fühlte mich etwas erholt und wollte mein Claude Code via web/pwa Projekt fortsetzen.

Dann begann Claude Code, Nutzungen außerhalb ihres eigenen Harness zu blockieren. Und ich fühle, dass es sich nicht lohnt.

Dann fand ich pi.dev und habe mich ein bisschen damit beschäftigt, aber ich bin nicht wirklich tief eingestiegen. Ich habe darüber gelesen, Videos darüber geschaut und mich entschieden, es voll auszuprobieren – und jetzt bin ich total in pi.

Da es Open Source ist, fühle ich, dass es sich zu bauen lohnt. Ich habe auch verschiedene Anbieter zur Auswahl. Ich fühle auch, dass die Abhängigkeit von einem Anbieter/Modell wie Anthropic/Claude nicht nachhaltig ist.

Also baue ich es hier.

Dieses Checkout wird als lokale-Modell-Edition von pi-web gepflegt. Es folgt dem
upstream Projekt für gemeinsame Verbesserungen, während lokale Deployment,
Kontext-Stabilität und lokale-Modell-Tests auf einem separat veröffentlichten
Track bleiben.

## Warum ein lokales Modell ein anderes Betriebsprofil braucht

Das ursprüngliche pi-web Erlebnis ist eine ausgezeichnete Grundlage, aber lokale
Inferenz hat andere Fehlerarten als ein typisches gehostetes Modell. Ein lokales
Modell kann sich mit wachsendem Kontext stark verlangsamen, den begrenzten
Speicher mit dem Rest der Maschine teilen, nach der Erzeugung nur von
Reasoning stoppen oder einen langen Lauf durch einen transienten lokalen
Transportfehler verlieren. Diese Fälle exakt wie Cloud-Fehler zu behandeln,
lässt die UI kompatibel aussehen, während die eigentliche Session fragil bleibt.

Diese Edition geht das Problem in Schichten an:

1. **Upstream zuerst bewahren.** Gemeinsame UI- und Session-Verhalten kommen
   weiterhin von pi-web; lokale Änderungen werden hinter dem effektiven Local
   Mode isoliert.
2. **Vorher verhindern, statt danach wiederherstellen.** Eine prozentbasierte
   65%-Kontextgrenze wird vor nachfolgenden Modellaufrufen erzwungen,
   einschließlich von Aufrufen innerhalb langer Tool-Schleifen.
3. **Nur mit Belegen wiederherstellen.** Automatische Fortsetzung ist auf
   erkannte Kontext-, Transport- und thinking-only-Vorfälle begrenzt – nicht
   auf Authentifizierung, Kontingent oder beliebige Anbieterfehler.
4. **Jede autonome Aktion begrenzen.** Wiederherstellungsvorfälle werden
   dedupliziert, Fortschritt wird vor einem weiteren Rettungsschritt
   verlangt, und beim Start wird höchstens eine kürzlich aktive Local-Session
   berücksichtigt.
5. **Einen manuellen Ausweg beibehalten.** Force Compact fasst zusammen,
   anstatt den Verlauf zu löschen, damit die Nutzenden eine Session retten
   können, ohne zu tun, als hätte der Kontext nie existiert.
6. **Cloud-Kompatibilität schützen.** Cloud Mode behält die Upstream-Semantik
   und -Steuerelemente bei; lokale-Modell-Optimierungen definieren Cloud-
   Sessions nicht stillschweigend neu.

Das ist der eigentliche Unterschied in diesem Fork: Es behandelt lokale
Inferenz als eigenständiges Betriebsumfeld, nicht einfach als einen weiteren
Modellnamen in einer Dropdown-Liste.
