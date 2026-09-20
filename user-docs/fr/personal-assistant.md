# pi-web comme assistant personnel

Ce flux de travail est pris en charge par l'édition local-model. Pour la
politique de développement, de synchronisation et de publication de l'édition,
voir [Local-model edition development](../../docs/dev/local-llm-development.md).

pi-web n'est pas réservé au codage — vous pouvez en faire un **assistant IA personnel** qui vit sur votre ordinateur, comme avoir votre propre OpenClaw ou Hermes.

## Comment ça marche

Vous créez un dossier dédié sur votre machine — c'est là que vit votre assistant. À l'intérieur, vous déposez un fichier `APPEND_SYSTEM.md` qui définit qui est votre assistant, ce qu'il sait et comment il se comporte. pi-web vous offre une belle interface de chat pour lui parler depuis n'importe quel appareil.

## Étape par étape

### 1. Créer votre dossier d'assistant

Choisissez un dossier sur votre ordinateur. Quelque chose comme :

```
~/my-assistant/
```

### 2. Définir votre assistant

Créez un fichier `APPEND_SYSTEM.md` dans ce dossier. C'est là que vous dites à pi qui est votre assistant :

```markdown
# My Personal Assistant

You are Jarvis, my personal AI assistant. You help me with:

- Daily planning and reminders
- Research and summarization
- Drafting emails and messages
- Brainstorming ideas
- Keeping track of things I mention

## About me

- I'm a software engineer who works remotely
- I have a cat named Pixel
- I prefer short, direct answers
- My timezone is PST

## Rules

- Be concise — I value brevity
- If you don't know something, say so
- Proactively remind me of things I asked you to track
```

pi ajoute automatiquement ceci à chaque prompt système de conversation, donc votre assistant sait toujours qui vous êtes et comment vous aider.

### 3. Démarrer une session dans ce dossier

Dans pi-web, créez une nouvelle session pointant vers `~/my-assistant/` (ou le nom que vous lui avez donné). C'est tout — vous parlez à votre assistant personnel.

### 4. L'utiliser depuis n'importe où

Installez pi-web comme PWA sur votre téléphone, tablette ou ordinateur portable. Votre assistant est toujours là — demandez-lui n'importe quoi, à tout moment.

## Idées pour votre assistant

| Rôle | Que mettre dans APPEND_SYSTEM.md |
|---|---|
| 🧠 **Coach de vie** | Vos objectifs, les habitudes sur lesquelles vous travaillez, des prompts de journaling |
| 🏠 **Gestionnaire de maison** | Format de liste de courses, préférences des membres de la famille, planification des repas |
| 💼 **Compagnon de travail** | Votre rôle, projets en cours, format des notes de réunion, contexte de l'entreprise |
| 📚 **Partenaire d'étude** | Ce que vous apprenez, style de présentation préféré, mode quiz |
| ✍️ **Assistant de rédaction** | Votre style d'écriture, préférences de ton, formats courants que vous utilisez |

## Ajouter plus de contexte

Vous pouvez mettre tout ce qui aide pi à être plus utile dans votre dossier d'assistant :

- `notes/` — fichiers de référence que votre assistant peut lire
- `context.md` — informations de contexte sur votre vie ou votre travail
- `projects.md` — projets en cours et leur statut

pi peut lire les fichiers du dossier, donc plus vous lui donnez de contexte, mieux c'est.

## Demandez à pi-web de faire des choses

Après `pi install npm:@timmygod/pi-web-local`, les sessions peuvent parler à pi-web lui-même.
Essayez :

- « Ajoute une planification à 2h heure de Singapour pour résumer ma boîte mail »
- « Liste mes planifications pi-web »
- « Mets la planification inbox en pause »
- « Note ça dans les notes »
- « Passe pi-web en mode sombre / désactive l'auto-titre »

La compétence **/skill:pi-web-schedule** incluse transforme cela en une vraie planification pi-web
(les mêmes que vous modifiez dans `/schedules`). Chaque déclenchement démarre une **nouvelle**
session, donc les instructions doivent être autonomes — « résume les e-mails non lus dans
~/inbox » fonctionne ; « continue ce qu'on faisait » ne fonctionne pas.

Les planifications ne s'exécutent que pendant que pi-web est en cours d'exécution.

---

> 💡 **Astuce :** Commencez simple. Juste quelques lignes sur qui vous êtes et comment vous voulez que l'assistant se comporte. Itérez au fil du temps à mesure que vous découvrez ce qui fonctionne.
