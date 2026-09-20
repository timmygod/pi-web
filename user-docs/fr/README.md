# Bienvenue sur pi-web 🖥️

<div align="center">

[English](../en/README.md) · [Español](../es/README.md) · **Français** · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**Vous envisagez d'essayer pi-web ? Foncez — vous allez l'adorer.**

pi-web est une belle interface web et PWA pour [pi](https://pi.dev) — l'agent de codage IA open source. Il vous permet de parcourir, lire et reprendre vos sessions pi depuis n'importe quel navigateur, sur n'importe quel appareil, avec des fonctionnalités réfléchies à chaque étape.

## Qu'est-ce qui est différent dans cette édition ?

Ce dépôt conserve l'interface et les fonctionnalités partagées de pi-web en amont, mais
modifie la manière dont les sessions sont protégées lorsque le modèle sélectionné fonctionne
en local ou sur votre LAN.

- **Choisissez la politique d'exécution par session.** Le mode Auto détecte les points de
  terminaison local/LAN lorsque les métadonnées du fournisseur sont claires ; Local et Cloud
  sont des remplacements manuels persistants.
- **Prévenir les échecs de contexte tôt.** Le mode Local compresse à 65 % d'utilisation et
  vérifie à nouveau entre les appels d'outils, avant la prochaine demande au modèle.
- **Conserver des résumés bornés.** Des points de contrôle roulants évitent qu'un ancien
  résumé ne grossisse sans fin, réessaient une fois avec un budget plus serré, et s'arrêtent
  en toute sécurité lorsque la compaction n'apporte pas de progrès significatifs.
- **Récupérer de manière prudente.** Le débordement de contexte, les interruptions du
  transport sélectionné, et les arrêts prématurés limités au raisonnement peuvent reprendre
  automatiquement, mais la déduplication des incidents et les disjoncteurs sensibles aux
  progrès empêchent les boucles de récupération.
- **Laisser l'utilisateur en contrôle.** Force Compact est toujours le chemin de secours
  manuel visible, tandis que le mode Cloud conserve le flux de travail et les contrôles en amont.

Le résultat pratique est simple : une longue tâche avec un modèle local devrait compacter
avant de s'effondrer, récupérer une fois lorsque la récupération est sûre, et s'arrêter
proprement plutôt que de boucler lorsque ce n'est pas le cas.

**pi-web est conçu pour deux types de personnes :**

- 🧑‍💻 **Pour les développeurs** — qui vivent dans le terminal mais veulent reprendre leurs sessions depuis leur mobile, transférer vers un serveur distant, ou surveiller les tâches de longue durée depuis n'importe où.
- ✨ **Pour les non-développeurs** — qui veulent simplement une belle application IA qui fonctionne. Ouvrez-la, tapez, vibrez. Pas de terminal, pas de SSH, pas de confusion. Comme les outils IA les plus conviviables, mais avec le choix du modèle et la liberté open source.

---

## Pourquoi pi-web ?

Vous êtes déjà en pleine concentration avec pi dans votre terminal. pi-web maintient cet élan lorsque vous vous éloignez de votre bureau :

- **Reprendre de n'importe où** — continuez une session depuis votre téléphone, tablette ou un autre ordinateur. Pas de SSH, pas de Termius — ouvrez simplement votre navigateur.
- **Tableau de bord multi-sessions** — lancez du travail dans une session tout en suivant le streaming d'une autre. Recherchez entre les projets, filtrez par branche, trouvez ce dont vous avez besoin rapidement.
- **Base open source** — pi est entièrement open source et agnostique par rapport aux fournisseurs. Vous n'êtes pas enfermé dans un seul modèle ou fournisseur. pi-web est aussi open source.
- **Accès distant sécurisé** — authentification par jeton intégrée pour l'exposer sur votre LAN ou Tailscale sans vous soucier.
- **Partagez votre travail** — exportez les sessions sous forme de snapshots statiques ou de Gists GitHub secrets en un clic.

> Curieux de connaître l'histoire ? [Lisez pourquoi nous l'avons construit →](why.md)

---

## pi-web en tant qu'espace de travail IA personnel 🏠

pi-web est une PWA (Progressive Web App), vous pouvez donc **l'installer comme une application native** sur votre ordinateur de bureau, portable, téléphone ou tablette — sans boutique d'applications. Sur ordinateur, il s'ouvre dans sa propre fenêtre sans les éléments du navigateur, et il a l'air et la sensation d'une véritable application de bureau.

Pensez-y comme **votre propre Claude Cowork** — un espace de travail IA personnel qui vit sur votre machine — sauf qu'il est open source et agnostique par rapport aux modèles :

- **Vous possédez la pile.** Choisissez n'importe quel modèle, changez quand vous voulez. Exécutez un modèle local et vos données ne quittent jamais votre machine.
- **Les non-techniciens peuvent l'utiliser.** Configurez pi-web sur leur machine, montrez-leur comment l'utiliser une fois, et ils sont bons pour partir. Vos parents, votre partenaire, vos amis non-techniques — pas de terminal, pas de SSH, juste une interface de chat familière.
- **Une configuration, de nombreux utilisateurs.** Installez-le sur votre ordinateur de bureau et partagez votre écran, ou exposez-le sur votre réseau domestique et laissez les membres de votre famille l'ouvrir sur leurs propres appareils.

Vous voulez plus que le codage ? Transformez-le en [assistant personnel](personal-assistant.md) dédié qui sait qui vous êtes et qui vit sur votre machine — comme votre propre OpenClaw ou Hermes.

> 💡 **Astuce :** Installez pi-web en tant que PWA depuis Chrome/Edge (cliquez sur l'icône d'installation dans la barre d'adresse) ou Safari (Partager → Ajouter au Dock). Il devient indistinguable d'une application native.

---

## Ce que vous pouvez faire avec pi-web

| | |
|---|---|
| 📱 **PWA** | Installez pi-web en tant que Progressive Web App sur ordinateur, téléphone ou tablette pour une sensation native. |
| 🔄 **Continuer les sessions** | Reprenez n'importe quelle conversation là où vous l'avez laissée — texte, images, changement de modèle, le tout depuis le navigateur. |
| 🆕 **Commencer de nouvelles sessions** | Créez de nouvelles sessions pour n'importe quel chemin de projet, directement depuis l'interface web. |
| 📡 **Streaming en direct** | Regardez les réponses de pi se streamer en temps réel avec une latence de ~ms. Le mode Suivi vous garde accroché à la dernière. |
| 🌲 **Vue arborescente** | Naviguez dans l'arborescence native de messages de pi — voyez la structure complète de la conversation, sautez à n'importe quelle branche, et fourchez à partir de n'importe quel point. |
| 🔀 **Fourcher les sessions** | Fourchez une session à partir de n'importe quel message ou même d'un appel d'outil spécifique — explorez différentes directions sans perdre votre place. |
| 🔍 **Parcourir & rechercher** | Filtrez les sessions entre les projets, recherchez par nom, naviguez dans les branches — votre historique complet de sessions en un coup d'œil. |
| 🌿 **Intégration Git** | Voyez la branche actuelle et ouvrez une PR GitHub directement depuis le visualiseur de session. |
| 📝 **Bloc-notes** | Notez des mémos, des to-dos ou des idées rapides à côté de vos sessions sans changer d'application. |
| 💬 **Annotations** | Mettez en surbrillance et commentez n'importe quelle partie d'une session — idéal pour la revue de code, le retour ou le signalement des moments clés. |
| 🎨 **Thèmes & personnalisation** | Basculez entre le mode sombre et clair, ajustez l'UI à votre goût — faites de pi-web quelque chose qui vous ressemble. |
| 🌐 **Multi-langues** | 14 langues intégrées (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Ajoutez votre propre langue personnalisée depuis les Paramètres. |
| 🐱 **Bien-être & pomodoro** | Trop de vibe coding n'est pas sain. Minuteur pomodoro intégré avec un compagnon chat et des rappels de sommeil pour garder l'équilibre. |
| 📤 **Partager & exporter** | Téléchargez du JSONL, exportez des snapshots statiques rendus avec l'aspect natif `pi.dev` de pi, ou partagez en tant que Gists GitHub privés — le tout rendu côté client. |
| 🔔 **Sons de notification** | Carillons de notification personnalisables pour les événements de session — restez dans la boucle même lorsque pi-web est dans un autre onglet. |
| ⌨️ **Raccourcis clavier** | Navigation style Vim, actions rapides — [référence complète →](keyboard-shortcuts.md) |
| 🤖 **Assistant personnel** | Transformez pi-web en votre propre assistant IA qui vit sur votre ordinateur — comme OpenClaw ou Hermes. [Configurez-le →](personal-assistant.md) |
| 🗓️ **Parler aux plannings** | Depuis une session pi, dites « ajoute un planning à 2h heure de Singapour pour … » — `/skill:pi-web-schedule`. |
| 📝 **Parler aux notes & paramètres** | « Écris ça dans les notes » (`/skill:pi-web-notes`) ou « passe en mode sombre » (`/skill:pi-web-settings`). |

---

## Navigation rapide

| Si vous cherchez… | Lisez |
|---|---|
| Comment installer, configurer et utiliser pi-web | [install.md](install.md) |
| Utiliser pi-web comme assistant personnel | [personal-assistant.md](personal-assistant.md) |
| Référence des raccourcis clavier | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Pourquoi pi-web existe | [why.md](why.md) |
| Ce qui arrive à la suite | [roadmap.md](roadmap.md) |
| Des problèmes d'installation ? Laissez votre LLM les réparer — collez-leur le lien llm-debug.md | [llm-debug.md](llm-debug.md) |
| Entretenir cette édition de modèle local | [notes de développement](../../docs/dev/local-llm-development.md) |

---

## Captures d'écran

| Ordinateur | Mobile |
|---|---|
| ![Ordinateur](../assets/pi-web-desktop-screenshot.png) | ![Mobile](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Sponsoring

pi-web est construit avec amour et beaucoup de nuits blanches. Je paie de mes poches les plans de codage (Claude Code, OpenCode, etc.) pour faire avancer ce projet. Si pi-web vous a été utile, votre soutien signifierait beaucoup.

**Manières d'aider :**

- 💰 **[Sponsoriser sur GitHub](https://github.com/sponsors/setkyar)** — aidez à couvrir les outils qui rendent cela possible
- ☕ **[Offrez-moi un café](https://buymeacoffee.com/setkyar)** — chaque petite aide compte
- ⭐ **Mettre une étoile au dépôt** — ça ne coûte rien et aide plus de gens à découvrir pi-web
- 📢 **Partager avec vos amis & votre famille** — si vous connaissez quelqu'un qui aimerait pi-web, envoyez-le-lui

Vous ne pouvez pas sponsoriser ? Aucun souci du tout — une étoile et un partage vont déjà loin. Merci d'être là. 🙏

---

Bon codage ! 🚀
