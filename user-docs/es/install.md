# Instalación y uso

## Funciones

### Control remoto

- Continúa cualquier sesión desde el navegador con archivos adjuntos de texto o imagen
- Inicia una sesión totalmente nueva contra cualquier ruta de proyecto, directamente desde la interfaz web
- Cambio de modelo en el navegador y selector de nivel de razonamiento, por sesión
- Estado del trabajador de cada sesión (inactivo / en ejecución / error) con recuperación automática ante un fallo
- Múltiples sesiones ejecutadas en paralelo — inicia trabajo en una y observa cómo otra transmite datos
- `PI_WEB_TOKEN` para exponer de forma segura en la LAN — requerido de forma predeterminada para cualquier enlace explícito no en lazo (loopback)

### Lectura de sesiones

- Explora sesiones entre proyectos con filtros, búsqueda y navegación completa de ramas
- Actualizaciones incrementales en vivo mientras pi sigue en ejecución (vía fsnotify; latencia de ~ms)
- Modo seguimiento para observar sesiones activas
- Enlaces directos a mensajes individuales
- Descarga una sesión como JSONL
- Comparte instantáneas estáticas como Gists secretos de GitHub
- Extensiones de pi `/web`, `/remote`, `/refresh`, `/pi-web token` y `/pi-web set-token` para abrir sesiones, QR remoto, sincronización de sesiones y gestión de tokens
- `/skill:pi-web-schedule`, `/skill:pi-web-notes`, `/skill:pi-web-settings` (`pi-web-ctl`) para que una sesión pueda gestionar horarios, el bloc de notas del proyecto y la configuración con lenguaje natural

## Elige el modo de sesión

Esta edición usa los proveedores y modelos ya configurados en pi; el Local Mode
es una política de tiempo de ejecución, no un instalador de modelos separado ni una segunda pantalla de clave de API.
Elige un modo al crear una sesión, o cámbialo después de que la ejecución actual se estabilice:

| Modo | Úsalo cuando | Comportamiento |
|------|-------------|----------|
| **Auto** | Quieres que pi-web decida | Resuelve endpoints locales/LAN desde los metadatos del proveedor cuando es posible; en caso contrario mantiene la ruta normal |
| **Local** | El modelo se ejecuta en este equipo o en tu LAN | Habilita el límite de compactación del 65%, puntos de control limitados, Force Compact y recuperación automática protegida |
| **Cloud** | El modelo seleccionado está alojado y debería seguir el comportamiento de la fuente | Mantiene fuera de la sesión la política de compactación y recuperación solo local |

La selección manual de Local o Cloud tiene prioridad sobre la detección automática y persiste entre
recargas y reinicios. Una sesión en ejecución rechaza los cambios de modo hasta que su
trabajador se estabilice, por lo que el modo mostrado en la interfaz siempre coincide con la política realmente en uso.

## Requisitos

- [Go](https://go.dev) 1.25+ (solo para compilar desde el código fuente)
- `pi` en tu `PATH` para el chat/cambio de modelo en el navegador
- Opcional: `gh` para compartir
- En Windows: pi necesita un shell bash para su herramienta de shell — [Git for Windows](https://git-scm.com/download/win) es suficiente (consulta la documentación de Windows de pi)

## Instalación

### Paquete de Pi (recomendado)

```bash
pi install npm:@timmygod/pi-web-local
```

Este único comando:
- Instala el paquete npm de pi en el directorio de paquetes de pi
- Ejecuta el guion `postinstall` del paquete (`install.sh`, o `install.ps1` en Windows)
- Descarga el binario de pi-web correspondiente a tu versión del paquete y plataforma desde GitHub Releases
- Lo instala en `~/.pi/agent/bin/pi-web` (`pi-web.exe` en Windows)
- Configura el inicio automático al iniciar sesión (launchd en macOS, systemd en Linux, un lanzador de la clave Run en Windows)
- Registra los comandos pi `/web`, `/remote`, `/refresh`, `/pi-web token` y `/pi-web set-token`

El titulado automático de sesiones está integrado en pi-web (no en la extensión) y se configura en la página `/settings`. Está activado por defecto: pi-web nombra las sesiones automáticamente usando una heurística gratuita de palabras integrada (sin IA), re-titulando con cada mensaje nuevo. Puedes cambiar al titulado una vez por sesión, y/o elegir un modelo para escribir títulos más inteligentes en lugar de la heurística.

En Linux, el inicio automático se configura como un servicio de usuario systemd en `~/.config/systemd/user/pi-web.service`. El instalador reescribe su `ExecStart` con la ruta real del binario instalado. Si Tailscale está disponible en tiempo de ejecución, pi-web publica el servidor de localhost con Tailscale Serve HTTPS. Si el systemd de usuario no está disponible, ejecútalo manualmente con `~/.pi/agent/bin/pi-web -o`.

Para instalar solo para un proyecto específico (compartido con tu equipo vía `.pi/settings.json`):

```bash
pi install -l npm:@timmygod/pi-web-local
```

Después reinicia pi (o ejecuta `/reload`) y usa `/web`, `/pi-web`, `/remote`, `/refresh`. Gestiona tu token de acceso con `/pi-web token` y `/pi-web set-token`.

Si npm se interrumpe con `ENOTEMPTY` al renombrar `@timmygod/pi-web-local`, elimina los directorios de copia de seguridad ocultos obsoletos de npm y reinstala el paquete:

```bash
rm -rf ~/.pi/agent/npm/node_modules/@timmygod/.pi-web-local-*
pi install npm:@timmygod/pi-web-local
```

### Instalación rápida (no requiere herramientas de compilación)

macOS / Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/timmygod/pi-web/main/install.sh | bash
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/timmygod/pi-web/main/install.ps1 | iex
```

Esto descarga el último binario de pi-web, lo instala en `/usr/local/bin` (`~/.pi/agent/bin` en Windows) y configura el inicio automático al iniciar sesión. No requiere Go, Node ni pi.

### Descargar el binario

Los binarios precompilados están adjuntos a cada [GitHub Release](https://github.com/timmygod/pi-web/releases).

```bash
# macOS (Apple Silicon)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-arm64
chmod +x pi-web

# macOS (Intel)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-darwin-amd64
chmod +x pi-web

# Linux (amd64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-amd64
chmod +x pi-web

# Linux (arm64)
curl -L -o pi-web https://github.com/timmygod/pi-web/releases/latest/download/pi-web-linux-arm64
chmod +x pi-web
```

```powershell
# Windows (x64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-amd64.exe

# Windows (ARM64)
irm -OutFile pi-web.exe https://github.com/timmygod/pi-web/releases/latest/download/pi-web-windows-arm64.exe
```

Después muévelo a tu PATH:

```bash
cp pi-web ~/.pi/agent/bin/
# o a nivel de sistema:
sudo cp pi-web /usr/local/bin/
```

### Compilar desde el código fuente

Esta copia es la edición de modelo local de pi-web. La compilación normal produce
la aplicación web y el backend juntas; las salvaguardas de modelo local se habilitan
en tiempo de ejecución por el Local Mode efectivo de la sesión, no por un binario separado.

```bash
git clone https://github.com/timmygod/pi-web.git
cd pi-web
make build   # compila el bundle de Vite y luego lo incrusta en el binario de Go

# opcional: ponlo en el PATH
cp pi-web ~/.pi/agent/bin/
```

El bundle del frontend se incrusta mediante `web/assets_embed.go`, por lo que `go build` necesita
que `web/dist` exista primero. `make build` realiza ambas pasos en orden; si compilas
manualmente, ejecuta `npm --prefix web install && npm --prefix web run build` antes de
`go build ./cmd/pi-web`.

Para el flujo de trabajo de la copia mantenida, la sincronización con la fuente original y la lista de verificación
de verificación del Local Mode, consulta [las notas de desarrollo del modelo local](../../docs/dev/local-llm-development.md).

### Desarrollar junto a una instancia instalada

Deja la instancia instalada en ejecución en el puerto `31415` y después inicia la copia
del código fuente en modo desarrollo:

```bash
make dev
```

Abre `http://127.0.0.1:31416`. `make dev` establece la variable de entorno interna `PI_WEB_DEV=1`
de desarrollo, por lo que la copia del código fuente comparte sesiones, configuración y
datos de SQLite con la instancia instalada mientras mantiene un archivo de bloqueo de tiempo de ejecución
de desarrollo y un archivo de estado separados. Las instancias instaladas normalmente y lanzadas manualmente
no cambian y conservan el comportamiento original de instancia única.

Para evitar trabajo autónomo duplicado, el modo de desarrollo no ejecuta el
bucle de horarios, el vaciador de la cola de chat, el titulado automático ni las notificaciones push. Las
solicitudes directas realizadas a través de la interfaz de desarrollo siguen funcionando. No conduzcas la misma
sesión de chat desde ambas instancias a la vez; cada proceso tiene su propio gestor
de trabajadores RPC.

`make dev` requiere [Air](https://github.com/air-verse/air) para la recarga en caliente de Go:

```bash
go install github.com/air-verse/air@latest
```

`PI_WEB_DEV` es un mecanismo de arnés de desarrollo, no un modo de múltiples instancias de producción soportado.

## Desinstalación

```bash
pi remove npm:@timmygod/pi-web-local
```

Esto ejecuta el guion `preuninstall` del paquete (`uninstall.sh`, o `uninstall.ps1`
en Windows), que detiene la instancia en ejecución y elimina:

- el binario de pi-web (`~/.pi/agent/bin/pi-web`, o `/usr/local/bin/pi-web` para instalaciones independientes)
- el archivo de versión (`~/.pi/agent/pi-web-version`)
- el archivo de estado de tiempo de ejecución (`~/.pi/agent/pi-web/pi-web-state.json`)
- la configuración de inicio automático (plist de launchd en macOS, servicio de usuario systemd en Linux, entrada de la clave Run + guiones del lanzador en Windows)

Tus datos se conservan para que una reinstalación posterior continúe donde lo dejaste:
`~/.pi/agent/pi-web.sqlite`, `~/.pi/agent/pi-web-memory.sqlite`, tus archivos de
sesión bajo `~/.pi/agent/sessions/`, y `~/.config/pi-web/env` (incluido
`PI_WEB_TOKEN`). Elimínalos manualmente si quieres empezar de cero.

## Uso

```bash
# Iniciar en el puerto predeterminado (31415)
pi-web

# Iniciar y abrir un navegador
pi-web -o

# Puerto personalizado
pi-web -p 8080

# Anular el host de enlace (loopback no requiere autenticación por defecto)
pi-web --host 127.0.0.1

# Un enlace no en lazo (loopback) requiere un token — pi-web se niega a iniciar en caso contrario
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web --host 192.168.1.50
```

Por defecto, pi-web se enlaza a `127.0.0.1`. Si Tailscale está en ejecución con MagicDNS **y `PI_WEB_TOKEN` está establecido**, pi-web también ejecuta `tailscale serve --bg --https=<port> http://127.0.0.1:<port>` e imprime la URL HTTPS del tailnet. Sin un token, pi-web se mantiene solo en loopback y omite Tailscale Serve, por lo que los pares del tailnet no pueden alcanzar al agente sin autenticación. Cualquier enlace explícito no en lazo (loopback) también requiere que `PI_WEB_TOKEN` esté establecido; pasa `--insecure` para anularlo en pruebas locales.

## Acceso remoto

Deja pi-web escuchando localmente y usa la URL HTTPS de Tailscale impresa desde tu teléfono o portátil en el tailnet.

En macOS, instala y abre Tailscale de forma interactiva, aprueba el cuadro de diálogo del administrador e inicia sesión. Después ejecuta `/pi-web restart`, seguido de `/remote`.

En Linux, permite a tu usuario gestionar Tailscale antes de instalar/ejecutar pi-web, de lo contrario `tailscale serve` puede requerir sudo y el inicio automático puede fallar:

```bash
sudo tailscale set --operator=$USER
```

```bash
# 1. Inicia pi-web con un token para que publique el endpoint HTTPS de Tailscale
PI_WEB_TOKEN=$(openssl rand -hex 16) pi-web

# 2. Desde cualquier otro dispositivo conectado a Tailscale, abre la
#    URL "Tailscale HTTPS" impresa e introduce el token una vez.
```

> Por defecto, pi-web se niega a enlazarse a una dirección no en lazo (loopback) a menos que `PI_WEB_TOKEN` esté establecido — cualquiera que pueda alcanzar la dirección enlazada podría ver las sesiones y enviar instrucciones a pi. Para anular esta salvaguarda en pruebas de red local, pasa `--insecure`. **No uses `--insecure` con Tailscale ni con ninguna dirección alcanzable desde fuera de tu equipo.**
>
> Los clientes pueden pasar el token a través de la cabecera `Authorization: Bearer <token>`, la cabecera `X-Pi-Token`, o una vez mediante `?token=<token>` (que establece una cookie `pi_token` para las solicitudes posteriores). Los tokens pasados mediante `?token=` acaban en el historial del navegador, los registros de acceso del servidor y en las cabeceras `Referer` de cualquier enlace de la página — prefiere la forma de cabecera para todo lo que no sea el marcador inicial.

## Chat en el navegador

Abre una página de sesión y usa el compositor en la parte inferior para continuar esa sesión exacta.

- `Enter` envía, `Shift+Enter` inserta un salto de línea
- Arrastra y suelta o pega imágenes directamente en el compositor
- El selector de modelo y el selector de nivel de razonamiento están en la cabecera — los cambios se aplican inmediatamente al trabajador de pi subyacente
- Cada sesión activa obtiene su propio trabajador `pi --mode rpc` dedicado, por lo que las distintas sesiones no se bloquean entre sí

## Compartir sesiones

Haz clic en **Compartir** en una página de sesión para crear un Gist secreto de GitHub.

Requisitos:
- `gh` instalado
- `gh auth login` completado

Al compartir se devuelven:
- la URL del gist secreto
- una URL de vista previa en `https://pi.dev/session/#<gistId>`

Los gists compartidos son instantáneas y no se actualizan en vivo.

## Inicio automático al iniciar sesión

### macOS

```bash
cp init/com.pi-web.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.pi-web.plist
```

### Linux (systemd)

```bash
# Instala el servicio de usuario systemd
mkdir -p ~/.config/systemd/user
cp init/pi-web.service ~/.config/systemd/user/

# Opcional: establece tu PI_WEB_TOKEN para enlaces no en lazo (loopback)
# (o usa /pi-web set-token <token> desde dentro de pi)
mkdir -p ~/.config/pi-web
echo 'PI_WEB_TOKEN=your-token-here' > ~/.config/pi-web/env

# Habilita y arranca
systemctl --user daemon-reload
systemctl --user enable --now pi-web.service

# Comprueba el estado
systemctl --user status pi-web.service

# Visualiza los registros
journalctl --user -u pi-web.service -f
```

> Para que el servicio se inicie al arrancar (antes de iniciar sesión), usa un servicio del sistema en su lugar:
> copia `init/pi-web.service` a `/etc/systemd/system/` y usa `sudo systemctl`.

### Windows

El instalador configura esto automáticamente, sin necesidad de derechos de administrador: una
entrada `pi-web` bajo `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
inicia `~/.config/pi-web/pi-web-start.vbs` al iniciar sesión, que inicia el binario
oculto (sin ventana de consola) después de cargar `~/.config/pi-web/env`
(`PI_WEB_TOKEN`, `PATH`, ...).

Para gestionarlo manualmente:

```powershell
# Iniciar / detener
wscript.exe "$HOME\.config\pi-web\pi-web-start.vbs"
taskkill /IM pi-web.exe /F

# Eliminar el inicio automático
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'pi-web'
```

No hay supervisión de servicio en Windows: si pi-web se cae, permanece detenido
hasta el siguiente inicio de sesión (launchd/systemd lo reinician automáticamente en las otras
plataformas).
