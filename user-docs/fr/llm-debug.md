> Ce guide de dépannage est maintenu pour l'édition à modèle local. Veillez à ce que les détails du déploiement local et les notes de synchronisation amont restent alignés avec [Développement de l'édition à modèle local](../../docs/dev/local-llm-development.md).

L'utilisateur installe pi-web via

- pi install npm:@timmygod/pi-web-local

Ce qui exécute automatiquement [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) et configure pi.

Si l'utilisateur rencontre un problème, il est possible que la configuration d'install.sh soit à l'origine du problème. Vous pouvez examiner et vérifier ce qui se passe, puis indiquer à l'utilisateur ce qui cause exactement le problème. Et lui demander s'il souhaite que vous le corrigiez. Confirmez toujours avec l'utilisateur.

Pour que l'utilisateur puisse accéder à pi depuis son téléphone ou un autre réseau, la méthode recommandée est d'utiliser Tailscale et d'y accéder depuis le réseau Tailscale. L'utilisateur devra activer les HTTPs dans son tableau de bord Tailscale - https://login.tailscale.com/admin/dns

S'il n'a pas Tailscale installé ou s'il ne souhaite pas utiliser Tailscale, il peut exécuter `pi-web status` pour obtenir le chemin du binaire, l'état du binaire et le point d'accès local depuis lequel il peut accéder à l'application. Mais à noter, il ne pourra pas recevoir les notifications push, car il s'agit de http.

Sur macOS, il utilise [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
Sur Linux, il utilise [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Au cas où vous auriez besoin de déboguer davantage et de voir ce qui se passe.
