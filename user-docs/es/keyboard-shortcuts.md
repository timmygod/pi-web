# Atajos de teclado

Estos atajos se aplican a la edición local-model de pi-web. El comportamiento de
ejecución específico de la edición está documentado en [Desarrollo de la edición local-model](../../docs/dev/local-llm-development.md).

## Página de índice (`/`)

### Desplazamiento de página (estilo vim)

Los mismos atajos de estilo vim funcionan en todas las páginas cuando el foco **no** está en un input, textarea o elemento contenteditable.

| Atajo | Acción |
|----------|--------|
| `j` | Desplazar hacia abajo 300px |
| `k` | Desplazar hacia arriba 300px |
| `g g` | Desplazar al inicio de la página |
| `G` (Shift+G) | Desplazar al final de la página |
| `Escape` | Quitar el foco del input activo para que la navegación con j/k funcione |

### Comandos del índice

| Atajo | Contexto | Acción |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Nivel de página | Abrir el panel de búsqueda/sesiones |
| `⌘⇧L` / `Ctrl+Shift+L` | Nivel de página | Alternar el tema del sistema (claro/oscuro) |
| `Escape` | Nivel de página | Cerrar el panel, menú o modal |
| `Enter` | Input de ruta de nueva sesión | Crear nueva sesión |

> `⌘K` / `Ctrl+K` también es el atajo de Chrome para "enfocar la barra de direcciones". El navegador puede interceptarlo a menos que el foco esté dentro de un input de texto.

## Página de detalle de sesión (`/session?id=...`)

### Desplazamiento de página (estilo vim)

Estos funcionan tanto en la página de índice como en la de sesión cuando el foco **no** está en un input, textarea o elemento contenteditable.

| Atajo | Acción |
|----------|--------|
| `j` | Desplazar hacia abajo 300px |
| `k` | Desplazar hacia arriba 300px |
| `g g` | Desplazar al inicio de la página |
| `G` (Shift+G) | Desplazar al final de la página |
| `I` (Shift+I) | Enfocar el textarea del compositor de chat |
| `Escape` | Quitar el foco del input activo para que la navegación con j/k funcione |

### Barra lateral y navegación

| Atajo | Contexto | Acción |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Nivel de página | Alternar la visibilidad de la barra lateral |
| `⌘K` / `Ctrl+K` | Nivel de página | Abrir el panel de lista de sesiones |
| `⌘T` / `Ctrl+T` | Nivel de página | Nueva sesión |
| `⌘⇧L` / `Ctrl+Shift+L` | Nivel de página | Alternar el tema del sistema (claro/oscuro) |
| `⌘⇧N` / `Ctrl+Shift+N` | Nivel de página | Alternar la barra lateral de borrador / notas |

> `⌘K` y `⌘T` también son atajos del navegador (enfocar la barra de direcciones / nueva pestaña). El navegador puede interceptarlos a menos que el foco esté dentro de un input de texto.

### Compositor de chat

| Atajo | Contexto | Acción |
|----------|---------|--------|
| `Enter` | Textarea del chat | Enviar mensaje |
| `Shift+Enter` | Textarea del chat | Insertar salto de línea |
| `Shift+Tab` | Textarea del chat | Alternar al siguiente nivel de pensamiento (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Textarea del chat | Abrir la ventana emergente del selector de modelo (escribir para filtrar, Enter para seleccionar, el foco vuelve al textarea) |

### Alternadores de visibilidad de entradas

| Atajo | Contexto | Acción |
|----------|---------|--------|
| `t` | Cuando el foco **no** está en un input/textarea | Alternar la visibilidad del pensamiento |
| `o` | Cuando el foco **no** está en un input/textarea | Alternar la visibilidad de las herramientas |
| `p` | Cuando el foco **no** está en un input/textarea | Alternar las salidas de herramientas |

### Paneles, menús y hojas

| Atajo | Contexto | Acción |
|----------|---------|--------|
| `Escape` | Nivel de página | Cerrar cualquier panel, menú o hoja abierta |
| `⌘K` / `Ctrl+K` | Nivel de página | Abrir el panel de lista de sesiones |
| `ArrowUp` / `ArrowDown` | Panel de lista de sesiones | Navegar por los resultados de sesiones |
| `Enter` | Panel de lista de sesiones | Abrir la sesión seleccionada (o la primera) |
| `ArrowUp` / `ArrowDown` | Ventana emergente del selector de modelo | Navegar por la lista de modelos |
| `Enter` | Ventana emergente del selector de modelo | Seleccionar el modelo resaltado |
| `ArrowUp` / `ArrowDown` | Modal de bifurcación | Navegar por los mensajes |
| `Enter` | Modal de bifurcación | Bifurcar desde el mensaje resaltado |
| `Tab` | Hoja a pantalla completa | Alternar el foco dentro de la hoja |
| `Escape` | Hoja a pantalla completa | Cerrar la hoja |
