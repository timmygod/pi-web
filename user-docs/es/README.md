# Bienvenido a pi-web 🖥️

<div align="center">

[English](../en/README.md) · **Español** · [Français](../fr/README.md) · [Deutsch](../de/README.md) · [中文](../zh/README.md) · [日本語](../ja/README.md) · [Bahasa Indonesia](../id/README.md) · [Bahasa Melayu](../ms/README.md) · [Tiếng Việt](../vi/README.md) · [ไทย](../th/README.md) · [Filipino](../fil/README.md) · [မြန်မာ](../my/README.md) · [ភាសាខ្មែរ](../km/README.md) · [ລາວ](../lo/README.md)

</div>

**¿Pensando en probar pi-web? Hazlo — te enamorará.**

pi-web es una hermosa interfaz web y PWA para [pi](https://pi.dev) — el agente de codificación de IA de código abierto. Te permite explorar, leer y continuar tus sesiones de pi desde cualquier navegador, en cualquier dispositivo, con características pensadas en cada paso.

## ¿Qué es diferente en esta edición?

Este repositorio mantiene la interfaz y funciones compartidas de pi-web upstream, pero
cambia cómo se protegen las sesiones cuando el modelo seleccionado se ejecuta localmente o en
tu LAN.

- **Elige la política de ejecución por sesión.** Auto detecta puntos finales locales/LAN cuando
  los metadatos del proveedor están claros; Local y Cloud son superposiciones manuales persistentes.
- **Previene fallos de contexto temprano.** Local Mode compacta al 65% de uso y verifica
  de nuevo entre llamadas a herramientas, antes de la siguiente solicitud al modelo.
- **Mantiene resúmenes acotados.** Los puntos de control rotativos evitan que un resumen viejo
  crezca sin fin, reintenta una vez con un presupuesto más ajustado y detiene con seguridad cuando la compactación
  no hace progreso significativo.
- **Se recupera de forma conservadora.** El desbordamiento de contexto, las interrupciones del transporte
  seleccionado y las paradas prematuras solo de razonamiento pueden reanudarse automáticamente, pero la
  deduplicación de incidentes y los circuit breakers conscientes del progreso previenen bucles de recuperación.
- **Deja el control al usuario.** Force Compact es siempre el camino manual visible de rescate,
  mientras que Cloud Mode conserva el flujo de trabajo y controles upstream.

El resultado práctico es simple: una tarea larga con un modelo local debería compactarse antes de
fallar, recuperarse una vez cuando la recuperación es segura, y detenerse limpiamente en lugar de
quedarse en bucle cuando no lo es.

**pi-web está hecho para dos tipos de personas:**

- 🧑‍💻 **Para desarrolladores** — que viven en la terminal pero quieren continuar sesiones desde el móvil, pasar a un servidor remoto, o monitorear tareas de larga ejecución desde cualquier lugar.
- ✨ **Para no desarrolladores** — que solo quieren una hermosa app de IA que funcione. Ábrela, escribe, vibe. Sin terminal, sin SSH, sin confusión. Como las herramientas de IA más amigables, pero con elección de modelo y libertad de código abierto.

---

## ¿Por qué pi-web?

Ya estás inmerso en el flujo con pi en tu terminal. pi-web mantiene ese impulso cuando te alejas de tu escritorio:

- **Reanuda desde cualquier lugar** — continúa una sesión desde tu teléfono, tableta u otro computador. Sin SSH, sin Termius — solo abre tu navegador.
- **Panel de control multi-sesión** — inicia trabajo en una sesión mientras observas cómo se transmite otra en streaming. Busca entre proyectos, filtra por rama, encuentra lo que necesitas rápido.
- **Base de código abierto** — pi es completamente de código abierto e independiente del proveedor. No estás atado a un solo modelo o proveedor. pi-web también es de código abierto.
- **Acceso remoto seguro** — autenticación de token integrada para que puedas exponerlo en tu LAN o Tailscale sin preocupaciones.
- **Comparte tu trabajo** — exporta sesiones como snapshots estáticos o GitHub Gists con secreto en un clic.

> ¿Curioso por la historia detrás? [Lee por qué lo construimos →](why.md)

---

## pi-web como tu espacio de trabajo personal de IA 🏠

pi-web es una PWA (Progressive Web App), así que puedes **instalarla como una app nativa** en tu escritorio, laptop, teléfono o tableta — sin necesidad de una tienda de aplicaciones. En escritorio se abre en su propia ventana sin los elementos del navegador, por lo que se ve y se siente como una aplicación de escritorio real.

Piensa en ella como **tu propio Claude Cowork** — un espacio de trabajo personal de IA que vive en tu máquina — excepto que es de código abierto e independiente del modelo:

- **Tú posees el stack.** Elige cualquier modelo, cambia cuando quieras. Ejecuta uno local y tus datos nunca salen de tu máquina.
- **Personas no técnicas pueden usarla.** Configura pi-web en su máquina, muéstrales cómo usarla una vez, y están listos. Tus padres, tu pareja, tus amigos no tecnológicos — sin terminal, sin SSH, solo una interfaz de chat familiar.
- **Una configuración, muchos usuarios.** Instálala en tu escritorio y comparte tu pantalla, o expónla en tu red doméstica y deja que los miembros de la familia la abran en sus propios dispositivos.

¿Quieres más que codificar? Conviértela en un [asistente personal](personal-assistant.md) dedicado que sabe quién eres y vive en tu máquina — como tu propio OpenClaw o Hermes.

> 💡 **Consejo pro:** Instala pi-web como PWA desde Chrome/Edge (clic en el icono de instalar en la barra de direcciones) o Safari (Compartir → Añadir al Dock). Se vuelve indistinguible de una app nativa.

---

## Lo que puedes hacer con pi-web

| | |
|---|---|
| 📱 **PWA** | Instala pi-web como Progressive Web App en escritorio, teléfono o tableta para una sensación nativa. |
| 🔄 **Continúa sesiones** | Retoma cualquier conversación justo donde la dejaste — texto, imágenes, cambio de modelo, todo desde el navegador. |
| 🆕 **Inicia nuevas sesiones** | Crea sesiones nuevas contra cualquier ruta de proyecto, directamente desde la interfaz web. |
| 📡 **Streaming en vivo** | Mira las respuestas de pi en streaming en tiempo real con latencia de ~ms. El modo Follow te mantiene enfocado en lo más reciente. |
| 🌲 **Vista de árbol** | Navega el árbol nativo de mensajes de pi — ve la estructura completa de la conversación, salta a cualquier rama y bifúcate desde cualquier punto. |
| 🔀 **Bifurca sesiones** | Bifurca una sesión desde cualquier mensaje o incluso una llamada de herramienta específica — explora diferentes direcciones sin perder tu lugar. |
| 🔍 **Explora & busca** | Filtra sesiones entre proyectos, busca por nombre, navega ramas — todo tu historial de sesiones de un vistazo. |
| 🌿 **Integración con Git** | Ve la rama actual y abre un PR de GitHub directamente desde el visor de sesión. |
| 📝 **Cuaderno** | Anota notas, tareas o ideas rápidas junto a tus sesiones sin cambiar de aplicación. |
| 💬 **Anotaciones** | Resalta y comenta cualquier parte de una sesión — ideal para revisión de código, retroalimentación o marcar momentos clave. |
| 🎨 **Temas & personalización** | Cambia entre modo oscuro y claro, ajusta la UI a tu gusto — haz que pi-web se sienta como *tuya*. |
| 🌐 **Multi-idioma** | 14 idiomas integrados (English, Español, Français, Deutsch, 中文, 日本語, Bahasa Indonesia, Bahasa Melayu, Tiếng Việt, ไทย, Filipino, မြန်မာ, ភាសាខ្មែរ, ລາວ). Añade tu propio idioma personalizado desde Ajustes. |
| 🐱 **Bienestar & pomodoro** | Demasiado vibe coding no es saludable. Temporizador pomodoro integrado con un gato de compañía y recordatorios de sueño para mantenerte equilibrado. |
| 📤 **Comparte & exporta** | Descarga JSONL, exporta snapshots estáticos renderizados con la apariencia nativa de `pi.dev`, o comparte como GitHub Gists privados — todo renderizado en el cliente. |
| 🔔 **Sonidos de notificación** | Avisos de notificación personalizables para eventos de sesión — mantente al tanto incluso cuando pi-web está en otra pestaña. |
| ⌨️ **Atajos de teclado** | Navegación estilo Vim, acciones rápidas — [referencia completa →](keyboard-shortcuts.md) |
| 🤖 **Asistente personal** | Convierte pi-web en tu propio asistente de IA que vive en tu computador — como OpenClaw o Hermes. [Configúralo →](personal-assistant.md) |
| 🗓️ **Habla con agendas** | Desde una sesión de pi, di "agrega una agenda a las 2am hora de Singapur para …" — `/skill:pi-web-schedule`. |
| 📝 **Habla con notas & ajustes** | "Escribe esto en las notas" (`/skill:pi-web-notes`) o "cambia a modo oscuro" (`/skill:pi-web-settings`). |

---

## Navegación rápida

| Si estás buscando… | Lee |
|---|---|
| Cómo instalar, configurar y usar pi-web | [install.md](install.md) |
| Usar pi-web como asistente personal | [personal-assistant.md](personal-assistant.md) |
| Referencia de atajos de teclado | [keyboard-shortcuts.md](keyboard-shortcuts.md) |
| Por qué existe pi-web | [why.md](why.md) |
| Lo que viene después | [roadmap.md](roadmap.md) |
| ¿Tienes problemas de instalación? Deja que tu LLM lo arregle — pégameles el enlace de llm-debug.md | [llm-debug.md](llm-debug.md) |
| Mantener esta edición de modelo local | [notas de desarrollo](../../docs/dev/local-llm-development.md) |

---

## Capturas de pantalla

| Escritorio | Móvil |
|---|---|
| ![Escritorio](../assets/pi-web-desktop-screenshot.png) | ![Móvil](../assets/pi-web-mobile-screenshot.png) |

---

## 💛 Patrocinio

pi-web está hecho con amor y muchas noches en vela. Pago de mi bolsillo los planes de codificación (Claude Code, OpenCode, etc.) para mantener este proyecto avanzando. Si pi-web te ha sido útil, tu apoyo significaría el mundo.

**Formas de ayudar:**

- 💰 **[Patrocinar en GitHub](https://github.com/sponsors/setkyar)** — ayuda a cubrir las herramientas que hacen esto posible
- ☕ **[Cómprame un café](https://buymeacoffee.com/setkyar)** — cada poco ayuda
- ⭐ **Dale star al repo** — no cuesta nada y ayuda a más personas a descubrir pi-web
- 📢 **Comparte con amigos & familia** — si conoces a alguien que le encantaría pi-web, mandásela

¿No puedes patrocinar? No hay ningún problema — una star y un comparten van muy lejos. Gracias por estar aquí. 🙏

---

¡Feliz codificación! 🚀
