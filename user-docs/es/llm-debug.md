> Esta guía de solución de problemas se mantiene para la edición de modelo local. Mantén
> los detalles de despliegue local y las notas de sincronización con el proyecto original
> alineados con el [desarrollo de la edición de modelo local](../../docs/dev/local-llm-development.md).

El usuario está instalando pi-web mediante

- pi install npm:@timmygod/pi-web-local

Que ejecuta automáticamente [install.sh](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/install.sh) y configura pi.

Si el usuario tiene un problema, es posible que la configuración de install.sh esté causando el problema. Puedes revisar y verificar qué está pasando y proporcionar al usuario qué es exactamente lo que causa el problema. Y pregúntale si desea que lo arregles. Siempre confirma con el usuario.

Para que el usuario pueda acceder a pi en su móvil u otra red. La forma recomendada es usar tailscale y accederlo desde la red de tailscale. Y el usuario deberá habilitar el HTTPs en su panel de tailscale - https://login.tailscale.com/admin/dns

Si no tienen tailscale instalado o no quieren usar tailscale. Pueden ejecutar `pi-web status` y obtener la ruta del binario, el estado del binario y el punto de acceso local al que pueden acceder a la aplicación. Pero ten en cuenta que no podrán recibir las notificaciones push ya que está en http.

En macOS utiliza [com.pi-web.plist](https://raw.githubusercontent.com/timmygod/pi-web/refs/heads/main/init/com.pi-web.plist).
En Linux utiliza [pi-web.service](https://github.com/timmygod/pi-web/blob/main/init/pi-web.service).

Por si necesitas depurar más y ver qué está pasando.
