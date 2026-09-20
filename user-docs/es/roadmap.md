# Plan de evolución (Roadmap)

Este plan de evolución pertenece a la edición de modelos locales de pi-web. Las funciones
de la versión principal (upstream) se sincronizan periódicamente, mientras que el trabajo de
despliegue local y fiabilidad se valida y publica en esta línea; consulta [Desarrollo de la edición de modelos locales](../../docs/dev/local-llm-development.md).

pi-web está diseñado para dos audiencias:

- **Para desarrolladores** — que viven en la terminal pero quieren continuar sesiones desde el móvil, pasar el trabajo a un servidor remoto, o mantener un ojo en tareas de larga duración desde cualquier lugar.
- **Para no desarrolladores** — que solo quieren una aplicación de IA hermosa que funcione. Ábrela, escribe y listo. Sin terminal, sin SSH, sin confusión. Como las herramientas de IA más fáciles de usar, pero con elección de modelo y libertad de código abierto.

Esto es lo que está por llegar.

Esta edición sigue al pi-web principal (upstream) en una línea de versiones separada. Las
funciones del upstream se importan periódicamente; el trabajo de fiabilidad de modelos locales
se prioriza y valida aquí sin cambiar el historial de versiones del upstream.

---

## Ahora (publicado)

Todo lo que se lista en [la tabla de funciones](README.md#what-you-can-do-with-pi-web) está disponible hoy.

---

## A continuación

| # | Función | Qué hace |
|---|---|---|
| [#50](https://github.com/timmygod/pi-web/issues/50) | **Bots de Telegram y Discord** | Chatea con pi a través de Telegram o Discord — perfecto para flujos de trabajo de asistente personal en movimiento. |
| [#49](https://github.com/timmygod/pi-web/issues/49) | **Informes de uso** | Seguimiento de tokens, estimación de costos, analítica de sesiones — sabrás cómo usas pi. |
| [#48](https://github.com/timmygod/pi-web/issues/48) | **Valores predeterminados configurables** | Establece tu visibilidad preferida para el razonamiento (thinking), las herramientas y las salidas de herramientas en todas las sesiones. |
| [#46](https://github.com/timmygod/pi-web/issues/46) | **Dirección / cola** | Envía instrucciones de seguimiento mientras pi aún está ejecutándose — guíalo en pleno vuelo. |
| [#41](https://github.com/timmygod/pi-web/issues/41) | **Comando `/compact`** | Compacta conversaciones largas directamente desde la interfaz web, sin necesidad de terminal. |

---

## Planificado

| # | Función | Qué hace |
|---|---|---|
| [#47](https://github.com/timmygod/pi-web/issues/47) | **Explorador de archivos y Git Diff** | Explora el árbol de archivos del proyecto y ver los cambios de git directamente en pi-web. Opcional (opt-in), para que no te estorbe. |
| [#44](https://github.com/timmygod/pi-web/issues/44) | **Programador (Scheduler)** | Programar que los prompts se ejecuten automáticamente — reuniones diarias, resúmenes matutinos, tareas recurrentes. Con restricción de administrador por seguridad. |
| [#43](https://github.com/timmygod/pi-web/issues/43) | **Atajos personalizables** | Reasigna cada atajo de teclado para que se adapte a tu memoria muscular. |

---

## Visión

El objetivo a largo plazo: pi-web debería ser **la interfaz de pi** — para todos.

- **Los no desarrolladores** la abren como cualquier otra aplicación. Eligen un modelo. Escriben. Listo. Nunca usan la línea de comandos.
- **Los desarrolladores** obtienen una integración profunda — transferencia remota, paneles de múltiples sesiones, exploración con conciencia de git, bots de mensajería.
- **Todos** obtienen libertad de modelos, transparencia de código abierto y una interfaz que se siente bien pensada en cada giro.

---

> 💡 ¿Tienes una idea? [Abre un issue](https://github.com/timmygod/pi-web/issues/new) o únete a la discusión.
