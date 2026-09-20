<h1 align="center">pi-web (Remote Control Your Pi)</h1>

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/timmygod/pi-web?style=flat&logo=github&label=stars&cacheSeconds=86400)](https://github.com/timmygod/pi-web/stargazers)
[![npm downloads](https://img.shields.io/npm/dt/@timmygod/pi-web-local?label=downloads&color=2ea043)](https://www.npmjs.com/package/@timmygod/pi-web-local)
[![license MIT](https://img.shields.io/npm/l/@timmygod/pi-web-local?label=license&color=0a7bbb)](../../LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Join-26A5E4?logo=telegram&logoColor=white)](https://t.me/+NJvFOTTa0wNjNTc9)
![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-555)

[English](../../README.md) · **Español** · [Français](README.fr.md) · [Deutsch](README.de.md) · [中文](README.zh.md) · [日本語](README.ja.md) · [Bahasa Indonesia](README.id.md) · [Bahasa Melayu](README.ms.md) · [Tiếng Việt](README.vi.md) · [ไทย](README.th.md) · [Filipino](README.fil.md) · [မြန်မာ](README.my.md) · [ភាសាខ្មែរ](README.km.md) · [ລາວ](README.lo.md)

</div>

<div align="center">

Controla tu agente de codificación [pi](https://pi.dev) desde tu teléfono, tablet o laptop: en cualquier lugar de tu red, o de forma remota a través de Tailscale.

Es una PWA completa, por lo que puedes instalarla y usarla como una app nativa en cualquier dispositivo. Piensa en ella como tu propio espacio de trabajo de IA personal: como el Cowork de Claude, pero con diferentes modelos; chatea entre modelos, programa desde tu teléfono o conviértela en un [asistente personal](../en/personal-assistant.md) que vive en tu máquina.

Házla tuya: cambia temas y fuentes, y úsala en tu propio idioma; pi-web viene con varios idiomas y puedes añadir los tuyos. Hay más funciones en camino, pero no se volverá rellena: lo que no necesites puedes desactivarlo en la configuración.

</div>

## ¿Por qué esta edición de modelo local?

El pi-web original sigue siendo la base upstream de características y
correcciones compartidas. Esta edición conserva esa experiencia y luego añade una capa de fiabilidad para
modelos que se ejecutan en tu propia máquina o en otro lugar de tu LAN, donde la generación
suele ser más lenta, la memoria es finita y un contexto largo puede detener una sesión que, de lo contrario, estaría sana.

| Área | pi-web upstream | Esta edición |
|------|-----------------|--------------|
| Política de modelo/entorno de ejecución | Comportamiento estándar de pi-web | Modo **Auto / Local / Cloud** por sesión, con detección local consciente del endpoint y una anulación manual persistente |
| Manejo de contexto largo | Comportamiento normal de compactación de pi | El modo Local compacta proactivamente al **65%** y vuelve a comprobar dentro de los bucles largos de llamadas a herramientas antes de la siguiente solicitud al modelo |
| Seguridad de la compactación | Resumen estándar | Puntos de verificación rodados acotados, una reescritura más estricta para salida inválida/limitada y detección de falta de progreso en lugar de recompactaciones interminables |
| Ejecuciones interrumpidas | Manejo normal de workers y errores | Recuperación acotada para desbordamiento de contexto, paradas solo de razonamiento e interrupciones de transporte seleccionadas, con rupturas de bucle persistentes |
| Rescate manual | Detalles estándar de contexto | **Compactación forzada** sigue disponible como ruta de recuperación explícita sin borrar la conversación |
| Compatibilidad y lanzamientos | Proyecto y línea de lanzamientos original | Las salvaguardas solo locales se mantienen detrás del modo Local; el modo Cloud preserva el comportamiento upstream y los cambios upstream se revisan y lanzan aquí de forma independiente |

Esto no es una reescritura ni un reemplazo del upstream. Es un perfil operativo
mantenido deliberadamente para personas que quieren privacidad y control de modelos locales
sin aceptar sesiones de larga duración frágiles. Consulta la
[guía de usuario](../en/README.md) para el flujo de trabajo orientado al usuario y el
[desarrollo de la edición de modelo local](../../docs/dev/local-llm-development.md) para la
implementación y la política de sincronización.

> [!TIP]
> ¿Nueva/o por aquí? **[Lee la guía de usuario →](../en/README.md)** para una visita completa de características, pasos de instalación y consejos. ([Otros idiomas →](../README.md))

## Capturas de pantalla

<div align="center">
  <img src="../assets/pi-web-desktop-screenshot.png" alt="Desktop" width="90%" /><br />
  <em>Escritorio</em>
  <br /><br />
  <img src="../assets/pi-web-mobile-screenshot.png" alt="Mobile" width="90%" /><br />
  <em>Móvil</em>
</div>

## Cómo encajan las piezas

```
 pi (terminal)                 Browser (phone / tablet / laptop)
      │                                │
      │  writes JSONL                  │  HTTP + SSE
      ▼                                ▼
 ~/.pi/agent/sessions/  ←───  pi-web (Go HTTP server)
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
              pi --mode rpc      fsnotify         tailscale serve
            (per‑session       (live reload)      (remote HTTPS
             chat worker)                           via MagicDNS)
```

- **pi** escribe la conversación en JSONL en `~/.pi/agent/sessions/` mientras trabaja.
- **pi-web** es un servidor en Go que lee esos archivos, los renderiza en el navegador y transmite actualizaciones en vivo a través de SSE.
- Los workers de **pi --mode rpc** manejan el chat iniciado desde el navegador: uno por sesión, recolectados tras 10 min inactivo.
- **fsnotify** vigila el directorio de sesiones para que el navegador se recargue en milisegundos desde una nueva salida.
- **Tailscale Serve** publica el servidor de localhost como un endpoint HTTPS en tu tailnet.

## Instalación

```bash
pi install npm:@timmygod/pi-web-local
```

Eso es todo: descarga el binario correspondiente, configura el autoarranque y registra los comandos `/web`, `/pi-web`, `/remote` y `/refresh`.

Una vez instalado, abre `http://127.0.0.1:31415` en tu navegador. Desde pi, usa `/web` para abrir la sesión actual en tu navegador al instante. Si Tailscale está ejecutándose en tu máquina, pi-web publica automáticamente un endpoint HTTPS en tu tailnet: usa `/remote` desde pi para obtener un código QR y una URL para cualquier dispositivo de tu tailnet.

> **Acceso remoto en macOS:** Instala y abre Tailscale de forma interactiva, aprueba el aviso de administrador e inicia sesión. Luego ejecuta `/pi-web restart`, seguido de `/remote`.

Para instalaciones manuales, descargas de binarios o compilación desde el código fuente, consulta [user-docs/install.md](../en/install.md).

## Integración con Pi

Después de `pi install npm:@timmygod/pi-web-local`, obtienes:

| Comando | Qué hace |
|---------|--------------|
| `/web` | Abre la sesión actual en tu navegador (conciencia SSH: omite el navegador y muestra solo la URL) |
| `/pi-web` | Muestra estado, versión, inicia/detiene/reinicia el servidor o actualiza |
| `/remote` | Muestra un código QR y una URL para acceso remoto a través de Tailscale |
| `/refresh` | Recupera los nuevos mensajes escritos desde navegadores remotos de vuelta a la sesión de terminal |

El **autotitulado** de sesiones está integrado en el propio pi-web y se configura en la página `/settings`. Está **activado por defecto** y nombra las sesiones automáticamente. Puedes elegir:

- **Cuándo titular** — una vez por sesión o con cada nuevo mensaje (el valor predeterminado).
- **Modelo de título** — por defecto, una **heurística de palabras integrada (sin IA)**, gratuita e instantánea, o elige un modelo (p. ej., uno pequeño/rápido) para títulos más inteligentes escritos por el modelo.

El paquete también instala el binario pi-web en `~/.pi/agent/bin/pi-web` y configura el autoarranque al iniciar sesión.

## Autoarranque al iniciar sesión

El comando `pi install npm:@timmygod/pi-web-local` configura esto automáticamente:

| Sistema | Mecanismo |
|----|-----------|
| macOS | plist de launchd en `~/Library/LaunchAgents/com.pi-web.plist` |
| Linux | servicio de usuario de systemd en `~/.config/systemd/user/pi-web.service` |
| Windows | entrada de la clave Run de `HKCU` que lanza un iniciador oculto en `~/.config/pi-web/` |

Para establecer un token para acceso remoto, crea `~/.config/pi-web/env`:

```
PI_WEB_TOKEN=your-token-here
```

Para más detalles (configuración manual, puertos personalizados, bindings no-loopback), consulta [user-docs/install.md](../en/install.md).

## Desarrollo

```bash
make setup   # install frontend deps and download Go modules
make check   # frontend test/build + Go test/vet
make build   # setup if needed, build frontend, then build ./pi-web
```

Para la sincronización con upstream, pruebas de modelo local y el flujo de trabajo de
lanzamiento paralelo, consulta [Desarrollo de la edición de modelo local](../../docs/dev/local-llm-development.md).
