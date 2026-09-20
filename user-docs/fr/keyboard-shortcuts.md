# Raccourcis clavier

Ces raccourcis s'appliquent à l'édition local-model de pi-web. Le comportement
d'exécution propre à chaque édition est documenté dans [Local-model edition development](../../docs/dev/local-llm-development.md).

## Page d'accueil (`/`)

### Défilement de la page (style vim)

Les mêmes raccourcis style vim fonctionnent sur toutes les pages lorsque le focus **n'est pas** dans un champ, une zone de texte ou un élément contenteditable.

| Raccourci | Action |
|----------|--------|
| `j` | Défiler vers le bas de 300px |
| `k` | Défiler vers le haut de 300px |
| `g g` | Défiler jusqu'en haut de la page |
| `G` (Shift+G) | Défiler jusqu'en bas de la page |
| `Escape` | Retirer le focus de l'élément actif pour que la navigation j/k fonctionne |

### Commandes de l'accueil

| Raccourci | Contexte | Action |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Niveau page | Ouvrir la palette de recherche/sessions |
| `⌘⇧L` / `Ctrl+Shift+L` | Niveau page | Basculer le thème système (clair/sombre) |
| `Escape` | Niveau page | Fermer la palette, le menu ou la fenêtre modale |
| `Enter` | Champ du chemin de nouvelle session | Créer une nouvelle session |

> `⌘K` / `Ctrl+K` est aussi le raccourci de Chrome pour « focaliser la barre d'adresse ». Le navigateur peut l'intercepter à moins que le focus ne soit dans un champ de texte.

## Page de détail de session (`/session?id=...`)

### Défilement de la page (style vim)

Ces raccourcis fonctionnent sur les pages d'accueil et de session lorsque le focus **n'est pas** dans un champ, une zone de texte ou un élément contenteditable.

| Raccourci | Action |
|----------|--------|
| `j` | Défiler vers le bas de 300px |
| `k` | Défiler vers le haut de 300px |
| `g g` | Défiler jusqu'en haut de la page |
| `G` (Shift+G) | Défiler jusqu'en bas de la page |
| `I` (Shift+I) | Mettre le focus sur la zone de texte de la zone de composition du chat |
| `Escape` | Retirer le focus de l'élément actif pour que la navigation j/k fonctionne |

### Barre latérale et navigation

| Raccourci | Contexte | Action |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Niveau page | Basculer la visibilité de la barre latérale |
| `⌘K` / `Ctrl+K` | Niveau page | Ouvrir la palette de la liste des sessions |
| `⌘T` / `Ctrl+T` | Niveau page | Nouvelle session |
| `⌘⇧L` / `Ctrl+Shift+L` | Niveau page | Basculer le thème système (clair/sombre) |
| `⌘⇧N` / `Ctrl+Shift+N` | Niveau page | Basculer la barre latérale brouillon / notes |

> `⌘K` et `⌘T` sont aussi des raccourcis du navigateur (focaliser la barre d'adresse / nouvel onglet). Le navigateur peut les intercepter à moins que le focus ne soit dans un champ de texte.

### Zone de composition du chat

| Raccourci | Contexte | Action |
|----------|---------|--------|
| `Enter` | Zone de texte du chat | Envoyer le message |
| `Shift+Enter` | Zone de texte du chat | Insérer un retour à la ligne |
| `Shift+Tab` | Zone de texte du chat | Passer au niveau de réflexion suivant (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Zone de texte du chat | Ouvrir le popup du sélecteur de modèle (taper pour filtrer, Enter pour sélectionner, le focus revient à la zone de texte) |

### Bascules de visibilité des entrées

| Raccourci | Contexte | Action |
|----------|---------|--------|
| `t` | Lorsque le focus **n'est pas** dans un champ/zone de texte | Basculer la visibilité des réflexions |
| `o` | Lorsque le focus **n'est pas** dans un champ/zone de texte | Basculer la visibilité des outils |
| `p` | Lorsque le focus **n'est pas** dans un champ/zone de texte | Basculer les sorties des outils |

### Palettes, menus et feuilles

| Raccourci | Contexte | Action |
|----------|---------|--------|
| `Escape` | Niveau page | Fermer toute palette, menu ou feuille ouverte |
| `⌘K` / `Ctrl+K` | Niveau page | Ouvrir la palette de la liste des sessions |
| `ArrowUp` / `ArrowDown` | Palette de la liste des sessions | Naviguer dans les résultats de sessions |
| `Enter` | Palette de la liste des sessions | Ouvrir la session sélectionnée (ou la première) |
| `ArrowUp` / `ArrowDown` | Popup du sélecteur de modèle | Naviguer dans la liste des modèles |
| `Enter` | Popup du sélecteur de modèle | Sélectionner le modèle mis en surbrillance |
| `ArrowUp` / `ArrowDown` | Fenêtre modale de fourche | Naviguer dans les messages |
| `Enter` | Fenêtre modale de fourche | Créer une fourche à partir du message mis en surbrillance |
| `Tab` | Feuille plein écran | Faire circuler le focus au sein de la feuille |
| `Escape` | Feuille plein écran | Fermer la feuille |
