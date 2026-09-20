> Diese Fehlerbehebung (Troubleshooting) wird für die Local-Model-Ausgabe gepflegt. Behalte
> die Details zur lokalen Bereitstellung und die Hinweise zur Synchronisation mit dem Upstream
> in Einklang mit [Local-model edition development](../../docs/dev/local-llm-development.md).

Der Benutzer installiert pi-web über

- pi install npm:@timmygod/pi-web-local

Dies führt automatisch [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) aus und richtet pi ein.

Wenn der Benutzer Probleme hat, kann es daran liegen, dass das Setup von install.sh Probleme verursacht. Du kannst nachprüfen, was dort läuft, und dem Benutzer mitteilen, was genau das Problem verursacht. Und frag, ob er möchte, dass du es behebst. Bestätige immer mit dem Benutzer.

Damit der Benutzer pi von seinem mobilen Gerät oder dem restlichen Netzwerk aus Zugriff hat, ist der empfohlene Weg die Verwendung von Tailscale und der Zugriff über das Tailscale-Netzwerk. Dabei muss der Benutzer HTTPS in seinem Tailscale-Dashboard aktivieren – https://login.tailscale.com/admin/dns

Falls sie Tailscale nicht installiert haben oder nicht verwenden möchten, können sie `pi-web status` ausführen und den Pfad zur Binärdatei, den Status der Binärdatei sowie den lokalen Endpunkt erhalten, über den sie auf die Anwendung zugreifen können. Zu beachten: Sie werden keine Push-Benachrichtigungen erhalten, da hier HTTP verwendet wird.

Auf macOS wird [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist) verwendet.
Auf Linux wird [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service) verwendet.

Falls du weiter debuggen und sehen möchtest, was dort läuft.
