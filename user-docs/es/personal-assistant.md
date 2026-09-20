# pi-web como tu asistente personal

Este flujo de trabajo es compatible con la edición de modelo local. Para conocer la política de desarrollo, sincronización y publicación de la edición, consulta [Desarrollo de la edición de modelo local](../../docs/dev/local-llm-development.md).

pi-web no es solo para programar: puedes convertirlo en un **asistente de IA personal** que vive en tu computadora, como tener tu propio OpenClaw o Hermes.

## Cómo funciona

Creas una carpeta dedicada en tu máquina: ahí vive tu asistente. Dentro, colocas un archivo `APPEND_SYSTEM.md` que define quién es tu asistente, qué sabe y cómo se comporta. pi-web te ofrece una hermosa interfaz de chat para hablar con él desde cualquier dispositivo.

## Paso a paso

### 1. Crea tu carpeta de asistente

Elige una carpeta en tu computadora. Algo como:

```
~/my-assistant/
```

### 2. Define tu asistente

Crea un archivo `APPEND_SYSTEM.md` dentro de esa carpeta. Aquí le dices a pi quién es tu asistente:

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

pi agrega automáticamente esto al system prompt de cada conversación, así que tu asistente siempre sabe quién eres y cómo ayudarte.

### 3. Inicia una sesión en esa carpeta

En pi-web, crea una nueva sesión apuntando a `~/my-assistant/` (o como le hayas llamado). Eso es todo: estás hablando con tu asistente personal.

### 4. Úsalo desde cualquier lugar

Instala pi-web como PWA en tu teléfono, tablet o laptop. Tu asistente siempre está ahí: pregúntale lo que sea, en cualquier momento.

## Ideas para tu asistente

| Rol | Qué poner en APPEND_SYSTEM.md |
|---|---|
| 🧠 **Coach de vida** | Tus metas, hábitos en los que estás trabajando, prompts para journaling |
| 🏠 **Gestor del hogar** | Formato de la lista de compras, preferencias de los miembros de la familia, planificación de comidas |
| 💼 **Compañero de trabajo** | Tu puesto, proyectos actuales, formato de notas de reuniones, contexto de la empresa |
| 📚 **Compañero de estudio** | Lo que estás aprendiendo, estilo de explicación preferido, modo "quiz me" |
| ✍️ **Asistente de escritura** | Tu estilo de escritura, preferencias de tono, formatos comunes que usas |

## Añade más contexto

Puedes poner cualquier cosa en tu carpeta de asistente que ayude a pi a ser más útil:

- `notes/` — archivos de referencia que tu asistente puede leer
- `context.md` — información de fondo sobre tu vida o trabajo
- `projects.md` — proyectos actuales y su estado

pi puede leer los archivos de la carpeta, así que cuanta más le des de contexto, mejor se vuelve.

## Pídele a pi-web que haga cosas

Después de `pi install npm:@timmygod/pi-web-local`, las sesiones pueden hablar con pi-web mismo.
Prueba:

- "Añade una programación a las 2 am hora de Singapur para resumir mi bandeja de entrada"
- "Lista mis programaciones de pi-web"
- "Pausa la programación de la bandeja de entrada"
- "Anota esto en las notas"
- "Cambia pi-web a modo oscuro / desactiva el auto-titulo"

El **/skill:pi-web-schedule** skill incluido convierte eso en una programación real de pi-web (las mismas que editas en `/schedules`). Cada ejecución inicia una **nueva** sesión, así que las instrucciones tienen que ser independientes: "resumir el correo no leído en ~/inbox" funciona; "continuar con lo que estábamos haciendo" no.

Las programaciones solo se ejecutan mientras pi-web esté corriendo.

---

> 💡 **Consejo:** Empieza simple. Solo unas cuantas líneas sobre quién eres y cómo quieres que el asistente se comporte. Itera con el tiempo a medida que descubres qué funciona.
