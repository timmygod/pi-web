# Pourquoi pi-web ?

Je suis un peu accro à Claude Code. Je l’utilise tout le temps. Si je ne suis pas assis devant l’ordinateur, j’y pense. J’ai l’impression de ne pas brûler assez de jetons. C’était les débuts de Claude Code. Et je me disais, pourquoi ne pas reprendre depuis mon téléphone ? J’ai configuré Termius et je n’ai pas vraiment aimé.

J’ai commencé à créer le mien et je me suis arrêté quand Claude a lancé son application mobile Claude Code.

Puis j’ai eu une hernie discale et je ne pouvais pas vraiment faire grand-chose. Le temps a passé, je me suis senti un peu remis et j’ai voulu poursuivre mon projet Claude Code via web/PWA.

Ensuite, Claude Code a commencé à interdire l’utilisation en dehors de son propre environnement. Et je me dis que ça n’en vaut plus la peine.

Puis j’ai découvert pi.dev, j’ai exploré un peu mais sans vraiment plonger. J’ai lu à son sujet, regardé des vidéos et décidé de m’y mettre sérieusement, et maintenant je suis à fond dans pi.

Comme c’est open source, je trouve que ça vaut le coup de construire dessus. J’ai aussi accès à différents fournisseurs. Je pense aussi que dépendre d’un seul fournisseur/modèle comme Anthropic/Claude n’est pas viable.

Alors je le construis ici.

## Pourquoi un modèle local nécessite un profil de fonctionnement différent

L'expérience pi-web originale est une excellente base, mais l'inférence locale présente des modes de défaillance différents de ceux d'un modèle hébergé typique. Un modèle local peut ralentir considérablement à mesure que le contexte s'agrandit, partager une mémoire limitée avec le reste de la machine, s'arrêter après avoir produit uniquement du raisonnement, ou perdre une longue exécution en raison d'une panne de transport locale transitoire. Traiter ces cas exactement comme des pannes cloud rend l'interface utilisateur compatible en apparence, tandis que la session réelle reste fragile.

Cette édition aborde le problème par couches :

1. **Préserver upstream en premier.** L'interface utilisateur partagée et le comportement de session continuent de provenir de pi-web ; les modifications locales sont isolées derrière un Local Mode effectif.
2. **Prévenir avant de récupérer.** Une limite de contexte de 65% basée sur un pourcentage est appliquée avant les appels ultérieurs au fournisseur, y compris les appels à l'intérieur de longues boucles d'outils.
3. **Récupérer uniquement avec des preuves.** La continuation automatique est limitée aux incidents reconnus de contexte, de transport et de réflexion uniquement, et non aux erreurs d'authentification, de quota ou aux erreurs arbitraires du fournisseur.
4. **Limiter chaque action autonome.** Les incidents de récupération sont dédupliqués, des progrès sont requis avant un nouveau sauvetage, et le démarrage prend en compte au plus une session Local active récemment.
5. **Garder une sortie manuelle.** Force Compact résume au lieu d'effacer l'historique, afin que l'utilisateur puisse sauver une session sans prétendre que le contexte n'a jamais existé.
6. **Protéger la compatibilité cloud.** Cloud Mode conserve la sémantique et les contrôles upstream ; les optimisations du modèle local ne redéfinissent pas silencieusement les sessions cloud.

C'est la véritable différence dans ce fork : il traite l'inférence locale comme un environnement opérationnel distinct, et non simplement comme un autre nom de modèle dans un menu déroulant.
