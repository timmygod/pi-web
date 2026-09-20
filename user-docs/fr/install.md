# Installation & Utilisation

## Fonctionnalités

### Télécommande

- Reprenez n'importe quelle session depuis le navigateur, avec des pièces jointes texte ou image
- Lancez une session totalement neuve sur n'importe quel chemin de projet, directement depuis l'interface web
- Sélecteur de modèle dans le navigateur et de niveau de réflexion, par session
- Statut du worker par session (inactif / en cours / erreur) avec récupération automatique en cas de crash
- Plusieurs sessions tournent en parallèle — lancez du travail dans l'une, observez l'autre en direct
- `PI_WEB_TOKEN` pour une exposition LAN sécurisée — requis par défaut pour tout bind explicite non-loopback

### Lecture des sessions

- Parcourez les sessions à travers les projets avec des filtres, une recherche et une navigation complète dans les branches
- Mises à jour incrémentielles en direct pendant que pi est encore en cours d'exécution (via fsnotify ; latence ~ms)
- Mode follow pour suivre les sessions actives
- Liens profonds vers des messages individuels
- Téléchargez une session au format JSONL
- Partagez des instantanés statiques en tant que Gists GitHub secrets
- Les extensions pi `/web`, `/remote`, `/refresh`, `/pi-web token` et `/pi-web set-token` pour ouvrir des sessions, un QR à distance, la synchronisation des sessions et la gestion du jeton
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) pour qu'une session puisse gérer les échéanciers, le brouillon du projet et les paramètres en langage naturel

## Choisissez le mode de session

Cette édition utilise les fournisseurs et les modèles déjà configurés dans pi ; le mode Local
est une politique d'exécution, et non un installateur de modèles séparé ni un second écran de clé API.
Choisissez un mode lors de la création d'une session, ou modifiez-le après que l'exécution en cours se soit stabilisée :

| Mode | À utiliser quand | Comportement |
|------|-------------|----------|
| **Auto** | Vous voulez que pi-web décide | Résout les points d'accès locaux/LAN à partir des métadonnées du fournisseur quand c'est possible ; sinon conserve le chemin normal |
| **Local** | Le modèle tourne sur cette machine ou sur votre LAN | Active la limite de compactation à 65 %, des points de contrôle bornés, Force Compact et la récupération automatique protégée |
| **Cloud** | Le modèle sélectionné est hébergé et doit suivre le comportement en amont | Garde la politique de compactation et de récupération locales uniquement à l'écart de la session |

La sélection manuelle de Local ou Cloud prime sur la détection automatique et persiste à travers
les rechargements et les redémarrages. Une session en cours refuse les changements de mode jusqu'à ce que son
worker se stabilise, donc le mode affiché dans l'interface correspond toujours à la politique réellement en cours d'utilisation.

## Prérequis

- [Go](https://go.dev) 1.25+ (uniquement pour la compilation à partir des sources)
- `pi` sur votre `PATH` pour le chat et le changement de modèle dans le navigateur
- Optionnel : `gh` pour le partage
- Sur Windows : pi a besoin d'une shell bash pour son outil shell — [Git for Windows](https://git-scm.com/download/win) suffit (voir la documentation de pi pour Windows)

## Installation

### Paquet Pi (recommandé)

```bash
pi install npm:@timmygod/pi-web-local
```

Cette seule commande :
- Installe le paquet pi npm sous le répertoire de paquets de pi
- Exécute le script `postinstall` du paquet (`install.sh`, ou `install.ps1` sur Windows)
- Télécharge le binaire pi-web correspondant à la version de votre paquet et à votre plateforme depuis les Releases GitHub
- L'installe dans `~/.pi/agent/bin/pi-web` (`pi-web.exe` sur Windows)
- Configure le démarrage automatique à la connexion (launchd sur macOS, systemd sur Linux, un lanceur de clé Run sur Windows)
- Enregistre les commandes pi `/web`, `/remote`, `/refresh`, `/pi-web token` et `/pi-web set-token`

Le titrage automatique des sessions est intégré à pi-web (pas à l'extension) et se configure sur la page `/settings`. Il est activé par défaut : pi-web titre automatiquement les sessions à l'aide d'une heuristique de mots intégrée gratuite (sans IA), en retitrant à chaque nouveau message. Vous pouvez passer au titrage une fois par session, et/ou choisir un modèle pour rédiger des titres plus intelligents au lieu de l'heuristique.

Sur Linux, le démarrage automatique est configuré comme un service systemd utilisateur dans `~/.config/systemd/user/pi-web.service`. L'installeur réécrit son `ExecStart` vers le chemin réel du binaire installé. Si Tailscale est disponible à l'exécution, pi-web publie le serveur localhost avec Tailscale Serve HTTPS. Si systemd utilisateur n'est pas disponible, exécutez-le manuellement avec `~/.pi/agent/bin/pi-web -o`.

Pour installer uniquement pour un projet spécifique (partagé avec votre équipe via `.pi/settings.json`) :

```bash
pi install -l npm:@timmygod/pi-web-local
```

Puis redémarrez pi (ou exécutez `/reload`), et utilisez `/web`, `/pi-web`, `/remote`, `/refresh`. Gérez votre jeton d'accès avec `/pi-web token` et `/pi-web set-token`.

Si npm s'arrête avec `ENOTEMPTY` lors du renommage de `@timmygod/pi-web-local`, supprimez les répertoires de sauvegarde cachés obsolètes de npm et réinstallez le paquet :

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Installation rapide (aucun outil de compilation nécessaire)

macOS / Linux :

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell) :

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Ceci télécharge le binaire pi-web le plus récent, l'installe dans `/usr/local/bin` (`~/.pi/agent/bin` sur Windows) et configure le démarrage automatique à la connexion. Aucun Go, Node ni pi requis.

### Télécharger le binaire

Les binaires précompilés sont joints à chaque [Release GitHub](https://github.com/timmygod/pi-web/releases).

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

Puis déplacez-le sur votre PATH :

```bash
cp pi-web ~/.pi/agent/bin/
# ou system-wide :
sudo cp pi-web /usr/local/bin/
```

### Compiler à partir des sources

Ce dépôt est l'édition local-model de pi-web. La compilation normale produit
l'application web et le backend ensemble ; les garde-fous local-model sont activés
à l'exécution par le mode Local effectif de la session, et non par un binaire séparé.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # compile le bundle Vite, puis l'intègre dans le binaire Go

# optionnel : le mettre sur le PATH
cp pi-web ~/.pi/agent/bin/
```

Le bundle du frontend est intégré par `web/assets_embed.go`, donc `go build` a besoin
que `web/dist` existe d'abord. `make build` fait les deux étapes dans l'ordre ; si vous
compilez à la main, exécutez `npm --prefix web install && npm --prefix web run build` avant
`go build ./cmd/pi-web`.

Pour le flux de travail du fork maintenu, la synchronisation en amont et la liste de vérification du Local Mode, voir [les notes de développement local-model](../../docs/dev/local-llm-development.md).

### Développer à côté d'une instance installée

Laissez l'instance installée tourner sur le port `31415`, puis lancez le
dépôt source en mode développement :

```bash
make dev
```

Ouvrez `http://127.0.0.1:31416`. `make dev` définit l'environnement interne `PI_WEB_DEV=1`
de développement, donc le dépôt source partage les sessions, les paramètres et les
données SQLite avec l'instance installée tout en conservant un verrou d'exécution de développement
et un fichier d'état séparés. Les instances installées normalement et lancées manuellement
ne sont pas modifiées et conservent le comportement d'instance unique d'origine.

Pour éviter un travail autonome en double, le mode développement n'exécute pas la
boucle d'échéancier, le vidage de la file de chat, le titrage automatique ni les notifications push. Les
demandes directes faites via l'interface de développement fonctionnent toujours. Ne pilotez pas la même
session de chat depuis les deux instances en même temps ; chaque processus a son propre
gestionnaire de worker RPC.

`make dev` nécessite [Air](https://github.com/air-verse/air) pour le rechargement à chaud Go :

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` est un outillage de harnais de développement, et non un mode
multi-instance de production pris en charge.

## Désinstallation

```bash
pi remove npm:@timmygod/pi-web-local
```

Ceci exécute le script `preuninstall` du paquet (`uninstall.sh`, ou `uninstall.ps1`
sur Windows), qui arrête l'instance en cours et supprime :

- le binaire pi-web (`~/.pi/agent/bin/pi-web`, ou `/usr/local/bin/pi-web` pour les installations autonomes)
- le fichier de version (`~/.pi/agent/pi-web-version`)
- le fichier d'état d'exécution (`~/.pi/agent/pi-web/pi-web-state.json`)
- la configuration de démarrage automatique (plist launchd sur macOS, service systemd utilisateur sur Linux, entrée de clé Run + scripts lanceur sur Windows)

Vos données sont conservées afin qu'une réinstallation ultérieure reprenne là où vous vous étiez arrêté :
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, vos fichiers de
session sous `~/.pi/agent/sessions/`, et `~/.config/pi-web/env` (y compris
`PI_WEB_TOKEN`). Supprimez-les manuellement si vous voulez une page blanche.

## Utilisation

```bash
# Démarrer sur le port par défaut (31415)
pi-web

# Démarrer et ouvrir un navigateur
pi-web -o

# Port personnalisé
pi-web -p 8080

# Overrider l'hôte de bind (le loopback est non authentifié par défaut)
pi-web --host 127.0.0.1

# Un bind non-loopback requiert un jeton — pi-web refuse de démarrer sinon
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Par défaut, pi-web se lie à `127.0.0.1`. Si Tailscale est en cours d'exécution avec MagicDNS **et que `PI_WEB_TOKEN` est défini**, pi-web exécute également `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` et affiche l'URL HTTPS tailnet. Sans jeton, pi-web reste uniquement en loopback et passe Tailscale Serve, donc les pairs tailnet ne peuvent pas accéder à l'agent non authentifié. Tout bind explicite non-loopback requiert également que `PI_WEB_TOKEN` soit défini ; passez `--insecure` pour overridé cela pour des tests locaux.

## Accès à distance

Laissez pi-web écouter localement, puis utilisez l'URL HTTPS Tailscale affichée depuis votre téléphone ou votre ordinateur portable sur le tailnet.

Sur macOS, installez et ouvrez Tailscale en interactif, approuvez l'invite d'administrateur et connectez-vous. Puis exécutez `/pi-web restart`, suivi de `/remote`.

Sur Linux, autorisez votre utilisateur à gérer Tailscale avant d'installer/executer pi-web, sinon `tailscale serve` peut exiger sudo et le démarrage automatique peut échouer :

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Démarrer pi-web avec un jeton afin qu'il publie le point d'accès HTTPS Tailscale
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Depuis n'importe quel autre appareil connecté à Tailscale, ouvrez l'URL
#    « Tailscale HTTPS » affichée et saisissez le jeton une fois.
```

> Par défaut, pi-web refuse de se lier à une adresse non-loopback sauf si `PI_WEB_TOKEN` est défini — quiconque peut atteindre l'adresse liée pourrait sinon consulter les sessions et envoyer des instructions à pi. Pour overridé ce garde-fou pour des tests sur le réseau local, passez `--insecure`. **N'utilisez pas `--insecure` sur Tailscale ou sur une adresse accessible depuis l'extérieur de votre machine.**
>
> Les clients peuvent transmettre le jeton via l'en-tête `Authorization: Bearer <token>`, l'en-tête `X-Pi-Token`, ou une fois via `?token=<token>` (qui définit un cookie `pi_token` pour les requêtes suivantes). Les jetons transmis via `?token=` finissent dans l'historique du navigateur, les journaux d'accès du serveur et les en-têtes `Referer` de tous les liens sur la page — préférez la forme en-tête pour tout ce qui dépasse le marque-page initial.

## Chat dans le navigateur

Ouvrez une page de session et utilisez la zone de composition en bas pour reprendre exactement cette session.

- `Enter` envoie, `Shift+Enter` insère un retour à la ligne
- Glisser-déposer ou coller des images directement dans la zone de composition
- Le sélecteur de modèle et le sélecteur de niveau de réflexion se trouvent dans l'en-tête — les changements s'appliquent immédiatement au worker pi sous-jacent
- Chaque session active possède son propre worker dédié `pi --mode rpc`, donc les différentes sessions ne se bloquent pas entre elles

## Partage des sessions

Cliquez sur **Partager** sur une page de session pour créer un Gist GitHub secret.

Prérequis :
- `gh` installé
- `gh auth login` terminé

Le partage renvoie :
- l'URL du gist secret
- une URL d'aperçu dans `https://pi.dev/session/#<gistId>`

Les gists partagés sont des instantanés et ne se mettent pas à jour en direct.

## Démarrage automatique à la connexion

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Installer le service systemd utilisateur
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Optionnel : définir votre PI_WEB_TOKEN pour les binds non-loopback
# (ou utilisez /pi-web set-token <token> depuis pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Activer et démarrer
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Vérifier l'état
systemctl --user status pi-web.service

# Afficher les journaux
journalctl --user -u pi-web.service -f
```

> Pour que le service démarre au démarrage (avant la connexion), utilisez un service système à la place :
> copiez `init/pi-web.service` vers `/etc/systemd/system/` et utilisez `sudo systemctl`.

### Windows

L'installeur configure cela automatiquement, sans avoir besoin de droits administrateur : une
entrée `pi-web` sous `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
lanche `~/.config/pi-web/pi-web-start.vbs` à la connexion, qui démarre le binaire
masqué (sans fenêtre de console) après avoir chargé `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...).

Pour le gérer manuellement :

```powershell
# Démarrer / arrêter
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Supprimer le démarrage automatique
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

Il n'y a pas de supervision de service sur Windows : si pi-web crash, il reste arrêté
jusqu'à la prochaine connexion (launchd/systemd le redémarre automatiquement sur les autres
plateformes).
