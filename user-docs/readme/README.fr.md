<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · [Español](README.es.md) · **Français** · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Pilotez votre agent de code [pi](https://pi.dev) depuis votre téléphone, votre tablette ou votre ordinateur portable — n'importe où sur votre réseau, ou à distance via Tailscale.

C'est une PWA complète, vous pouvez donc l'installer et l'utiliser comme une application native sur n'importe quel appareil. Pensez-y comme à votre espace de travail IA personnel — comme Cowork de Claude, mais avec d'autres modèles — discutez entre plusieurs modèles, codez depuis votre téléphone, ou transformez-le en [assistant personnel](../en/personal-assistant.md) qui vit sur votre machine.

Faites-en ce que vous voulez : changez de thèmes et de polices, et utilisez-le dans votre propre langue — pi-web est livré avec plusieurs langues et vous pouvez en ajouter d'autres. D'autres fonctionnalités arrivent, mais ça ne va pas devenir encombrant : tout ce dont vous n'avez pas besoin peut être désactivé dans les paramètres.

</div>

## Pourquoi cette édition pour modèles locaux ?

Le pi-web d'origine demeure le socle amont pour les fonctionnalités et correctifs partagés. Cette édition conserve cette expérience, puis y ajoute une couche de fiabilité pour les modèles qui tournent sur votre propre machine ou ailleurs sur votre LAN — là où la génération est souvent plus lente, la mémoire est limitée, et un long contexte peut bloquer une session autrement saine.

| Domaine | pi-web amont | Cette édition |
|------|-----------------|--------------|
| Politique modèle/runtime | Comportement standard de pi-web | Mode **Auto / Local / Cloud** par session, avec détection locale sensible aux points de terminaison et une surcharge manuelle persistante |
| Gestion des longs contextes | Comportement de compaction pi standard | Le Mode Local compacte proactivement à **65%** et vérifie à nouveau dans les boucles de tool-calls longs avant la demande modèle suivante |
| Sécurité de la compaction | Résumés standard | Checkpoints roulants bornés, une réécriture plus stricte pour les sorties invalides/limitées, et une détection d'absence de progression au lieu d'une re-compaction infinie |
| Exécutions interrompues | Gestion normale des workers et des erreurs | Récupération bornée pour les débordements de contexte, les arrêts uniquement en « thinking » et certaines interruptions de transport, avec des coupe-boucles persistants |
| Secours manuel | Détails de contexte standard | **Force Compact** reste disponible comme chemin de récupération explicite sans effacer la conversation |
| Compatibilité et versions | Projet et ligne de versions d'origine | Les garde-fous local-only restent derrière le Mode Local ; le Mode Cloud préserve le comportement amont, et les changements amont sont examinés et publiés ici indépendamment |

Ce n'est ni une réécriture ni un remplacement de l'amont. C'est un profil d'exploitation volontairement maintenu pour ceux qui veulent la confidentialité et le contrôle des modèles locaux sans accepter des sessions de longue durée fragiles. Voir le
[guide utilisateur](../en/README.md) pour le workflow côté utilisateur et le
[développement de l'édition local-model](../../docs/dev/local-llm-development.md) pour
l'implémentation et la politique de synchronisation.

> [!TIP]
> Nouveau ici ? **[Lisez le guide utilisateur →](../en/README.md)** pour une visite complète des fonctionnalités, des étapes d'installation et des astuces. ([Autres langues →](../README.md))

## Captures d'écran

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Bureau</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Mobile</em>
</div>

## Comment tout s'articule

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

- **pi** écrit la conversation en JSONL dans `~/.pi/agent/sessions/` au fur et à mesure.
- **pi-web** est un serveur Go qui lit ces fichiers, les affiche dans le navigateur et transmet les mises à jour en direct via SSE.
- Les workers **pi --mode rpc** gèrent la conversation initiée par le navigateur — un par session, récollés après 10 min d'inactivité.
- **fsnotify** surveille le répertoire des sessions afin que le navigateur se recharge en quelques millisecondes après une nouvelle sortie.
- **Tailscale Serve** expose le serveur localhost comme un point de terminaison HTTPS sur votre tailnet.

## Installation

```bash
pi install npm:@timmygod/pi-web-local
```

C'est tout — il télécharge le binaire correspondant, configure le démarrage automatique et enregistre les commandes `/web`, `/pi-web`, `/remote` et `/refresh`.

Une fois installé, ouvrez `http://127.0.0.1:31415` dans votre navigateur. Depuis pi, utilisez `/web` pour ouvrir la session actuelle immédiatement dans votre navigateur. Si Tailscale tourne sur votre machine, pi-web publie automatiquement un point de terminaison HTTPS sur votre tailnet — utilisez `/remote` depuis pi pour obtenir un code QR et une URL pour n'importe quel appareil de votre tailnet.

> **Accès distant macOS :** Installez et ouvrez Tailscale en interactif, approuvez l'invite d'administrateur et connectez-vous. Puis lancez `/pi-web restart`, suivi de `/remote`.

Pour les installations manuelles, les téléchargements de binaires ou la compilation depuis le source, voir [user-docs/install.md](../en/install.md).

## Intégration Pi

Après `pi install npm:@timmygod/pi-web-local`, vous obtenez :

| Commande | Ce qu'elle fait |
|---------|--------------|
| `/web` | Ouvrir la session actuelle dans votre navigateur (conscient du SSH : passe le navigateur et n'affiche que l'URL) |
| `/pi-web` | Afficher l'état, la version, démarrer/arrêter/redémarrer le serveur, ou mettre à jour |
| `/remote` | Afficher un code QR et une URL pour l'accès distant via Tailscale |
| `/refresh` | Récupérer les nouveaux messages écrits depuis les navigateurs distants et les renvoyer dans la session terminal |

L'**auto-intitulage** des sessions est intégré à pi-web lui-même et configuré sur la page `/settings`. Il est **activé par défaut** et intitule les sessions automatiquement. Vous pouvez choisir :

- **Quand intituler** — une fois par session, ou à chaque nouveau message (le défaut).
- **Modèle d'intitulage** — par défaut une **heuristique de mots intégrée, gratuite et instantanée (sans IA)**, ou choisissez un modèle (par ex. un petit/rapide) pour des intitulés plus intelligents, écrits par le modèle.

Le paquet installe également le binaire pi-web dans `~/.pi/agent/bin/pi-web` et configure le démarrage automatique à la connexion.

## Démarrage automatique à la connexion

La commande `pi install npm:@timmygod/pi-web-local` configure cela automatiquement :

| OS | Mécanisme |
|----|-----------|
| macOS | plist launchd dans `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | service utilisateur systemd dans `~/.config/systemd/user/pi-web.service` |
| Windows | entrée Run-key `HKCU` lançant un lanceur caché dans `~/.config/pi-web/` |

Pour définir un jeton d'accès distant, créez `~/.config/pi-web/env` :

```
PI_WEB_TOKEN=your-token-here
```

Pour plus de détails (configuration manuelle, ports personnalisés, liaisons non-loopback), voir [user-docs/install.md](../en/install.md).

## Développement

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Pour la synchronisation amont, les tests local-model et le workflow de publication en parallèle, voir [développement de l'édition local-model](../../docs/dev/local-llm-development.md).
