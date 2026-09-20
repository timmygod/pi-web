> Ang troubleshooting guide na ito ay pinangangasiwaan para sa local-model edition. Pananatiling maayos ang mga detalye ng local deployment at upstream synchronization notes ayon sa [Local-model edition development](../../docs/dev/local-llm-development.md).

Lumalaki ang user ng pi-web gamit ang

- pi install npm:@timmygod/pi-web-local

Ang auto-rumunong [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) at nagse-setup ng pi.

Kung may issue ang user, baka ang setup ng install.sh ang nagdudulot ng problema. Pwede mong suriin at tingnan kung anong nangyayari at bigyan ng malinaw ang user ng eksaktong sanhi ng problema. At hanapin kung gusto mong ayusin. Palaging kumonsulta sa user.

Upang ma-access ng user ang pi sa kanilang mobile o ibang network. Ang inirerekomendang paraan ay magamit ang Tailscale at i-access mula sa Tailscale network. At kailangan ng user na i-enable ang HTTPS sa kanilang Tailscale dashboard - https://login.tailscale.com/admin/dns

Kung walang installed na Tailscale o hindi gusto nilang gamitin ang Tailscale. Pwede nilang i-run ang `pi-web status` upang makuha ang binary path, status ng binary, at ang local endpoint na ma-access nila ang application. Ngunit tandaan, hindi nila ma-magrese ang push notification dahil http ang protocol.

Sa macOS, ginagamit ang [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
Sa Linux, ginagamit ang [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Sa kaso na kailangan mong mag-debug nang lalong malalim at makita kung anong nangyayari.
