# Pourquoi pi-web ?

Je suis un peu accro à Claude Code. Je l'utilise tout le temps. Si je ne suis pas assis devant l'ordinateur, je y pense. J'ai l'impression de ne pas assez consommer de jetons. C'était les premiers jours de Claude Code. Et je me disais, pourquoi ne puis-je pas reprendre depuis mon téléphone ? J'ai configuré Termius et je n'ai vraiment pas aimé ça.

J'ai commencé à créer le mien et j'ai arrêté quand Claude a introduit son application mobile Claude Code.

Ensuite, j'ai eu un disque hernié et je ne pouvais vraiment pas faire grand-chose. Le temps a passé et j'ai senti que je me rétablissais un peu, et j'avais envie de continuer mon projet Claude Code via web/pwa.

Ensuite, Claude Code a commencé à bannir l'utilisation en dehors de leur propre harnais. Et j'ai l'impression que ce n'est plus rentable.

Ensuite, j'ai découvert pi.dev et j'ai exploré un peu, mais je n'avais pas vraiment creusé. J'ai lu à ce sujet, regardé des vidéos et j'ai décidé d'y consacrer un essai complet, et maintenant je suis totalement pris par pi.

Puisque c'est open source, j'ai l'impression que ça vaut le coup d'y travailler dessus. J'ai aussi le choix de différents fournisseurs. Je ressens aussi que compter sur un seul fournisseur/modèle comme Anthropic/Claude n'est pas durable.

Donc, je le construis ici.

Ce checkout est maintenu en tant qu'édition local-model de pi-web. Il suit le
projet amont pour les améliorations partagées, tout en conservant le déploiement local,
la stabilité du contexte et le test des local-models sur une piste de publication séparée.

## Pourquoi un local-model nécessite un profil opérationnel différent

L'expérience pi-web originale est une excellente base, mais l'inférence locale
a des modes d'échec différents d'un modèle hébergé typique. Un local-model peut
ralentir brusquement à mesure que le contexte grandit, partager une mémoire limitée
avec le reste de la machine, s'arrêter après avoir produit uniquement du raisonnement,
ou perdre une longue exécution à cause d'une défaillance transitoire du transport local.
Traiter ces cas exactement comme des échecs cloud rend l'UI apparemment compatible
alors que la session réelle reste fragile.

Cette édition aborde le problème par couches :

1. **Préserver l'amont en priorité.** Les comportements d'UI et de session partagés
   continuent de provenir de pi-web ; les modifications locales sont isolées derrière un Local Mode effectif.
2. **Prévenir avant de récupérer.** Une limite de contexte basée sur un pourcentage de 65% est
   appliquée avant les appels modèle suivants, y compris les appels au sein de longues boucles d'outils.
3. **Ne récupérer qu'avec des preuves.** La continuation automatique est limitée aux incidents
   reconnus de contexte, de transport et de thinking-only — et non aux erreurs d'authentification, de quota
   ou de fournisseur arbitraires.
4. **Bonder toute action autonome.** Les incidents de récupération sont dédupliqués,
   un progrès est exigé avant un autre sauvetage, et le démarrage prend en compte au plus une
   session Local récemment active.
5. **Conserver une sortie manuelle.** Force Compact résume plutôt qu'il n'efface l'historique,
   afin que l'utilisateur puisse sauver une session sans prétendre que le contexte n'a jamais existé.
6. **Protéger la compatibilité cloud.** Le Cloud Mode conserve les sémantiques et
   les contrôles de l'amont ; les optimisations pour local-models ne redéfinissent pas silencieusement les sessions cloud.

Voilà la vraie différence dans ce fork : il traite l'inférence locale comme un
environnement opérationnel distinct, et non simplement comme un autre nom de modèle dans un menu déroulant.
