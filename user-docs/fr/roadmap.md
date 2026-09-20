# Feuille de route

Cette feuille de route appartient à l'édition local-model de pi-web. Les fonctionnalités amont sont synchronisées périodiquement, tandis que les travaux de déploiement local et de fiabilité sont validés et publiés sur cette ligne ; voir [Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web est conçu pour deux audiences :

- **Pour les développeurs** — qui vivent dans le terminal mais veulent continuer des sessions depuis le mobile, transférer vers un serveur distant, ou garder un œil sur les tâches de longue durée depuis n'importe où.
- **Pour les non-développeurs** — qui veulent simplement une belle app IA qui fonctionne. On l'ouvre, on tape, on vibe. Pas de terminal, pas de SSH, pas de confusion. Comme les outils IA les plus conviviaux, mais avec le choix du modèle et la liberté open source.

Voici ce qui arrive.

Cette édition suit pi-web amont sur une ligne de release séparée. Les fonctionnalités amont sont importées périodiquement ; les travaux de fiabilité local-model sont priorisés et validés ici sans modifier l'historique des releases amont.

---

## Maintenant (déployé)

Tout ce qui est listé dans [le tableau des fonctionnalités](README.md#what-you-can-do-with-pi-web) est disponible aujourd'hui.

---

## À venir

| # | Fonctionnalité | Ce qu'elle fait |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Telegram & Discord bots** | Discuter avec pi via Telegram ou Discord — parfait pour les workflows d'assistant personnel en déplacement. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Usage insights** | Suivi des tokens, estimation des coûts, analytiques de session — savoir comment vous utilisez pi. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Configurable defaults** | Définir votre visibilité préférée pour le raisonnement, les outils et les sorties d'outils sur toutes les sessions. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Steering / queue** | Envoyer des instructions de suivi pendant que pi est encore en cours d'exécution — le guider en plein vol. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **`/compact` command** | Compresser les longues conversations directement depuis l'interface web, sans terminal. |

---

## Prévu

| # | Fonctionnalité | Ce qu'elle fait |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **File Explorer & Git Diff** | Parcourir l'arborescence des fichiers du projet et voir les modifications git directement dans pi-web. Optionnel, pour ne pas vous gêner. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Scheduler** | Planifier l'exécution automatique de prompts — standups quotidiens, résumés du matin, tâches récurrentes. Réservé aux administrateurs pour la sécurité. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Customizable shortcuts** | Reconfigurer chaque raccourci clavier pour correspondre à votre mémoire musculaire. |

---

## Vision

L'objectif à long terme : pi-web doit être **l'interface pour pi** — pour tout le monde.

- **Les non-développeurs** l'ouvrent comme n'importe quelle autre app. Ils choisissent un modèle. Ils tapent. C'est fait. Jamais de ligne de commande.
- **Les développeurs** obtiennent une intégration profonde — transfert distant, tableaux de bord multi-sessions, navigation git-aware, bots de messagerie.
- **Tout le monde** obtient la liberté de choix du modèle, la transparence open source, et une interface qui se sent soignée à chaque tour.

---

> 💡 Une idée ? [Ouvrir un issue](https://github.com/timmygod/pi-web/issues/new) ou rejoindre la discussion.
