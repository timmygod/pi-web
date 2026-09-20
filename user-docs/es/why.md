# ¿Por qué pi-web?

Me resulta adicto a Claude Code. Lo uso todo el tiempo. Si no estoy sentado frente al ordenador, estoy pensando en él. Siento que no estoy quemando suficientes tokens. Eran los primeros días de Claude Code. Y yo me preguntaba, ¿por qué no puedo reanudar desde mi teléfono? Configuré Termius y no me gustó mucho.

Empecé a crear el mío propio y lo paré cuando Claude presentó su aplicación móvil de Claude Code.

Luego me diagnosticaron un disco herniado y no podía hacer mucho más. Pasó el tiempo, sentí que me había recuperado un poco y quise continuar mi proyecto de Claude Code vía web/pwa.

Entonces Claude Code empezó a prohibir el uso fuera de su propio harness. Y siento que no vale la pena.

Después descubrí pi.dev y exploré un poco, pero no me adentré de verdad. Leí sobre ello, vi vídeos y decidí darle un intento completo, y ahora estoy totalmente enganchado a pi.

Al ser open source, siento que vale la pena construir para ello. También obtengo diferentes opciones de proveedor. Además, siento que depender de un solo proveedor/modelo como Anthropic/Claude no es sostenible.

Así que lo estoy construyendo aquí.

Este checkout se mantiene como una edición de modelo local de pi-web. Sigue al
proyecto upstream para mejoras compartidas, manteniendo la despliegue local,
la estabilidad del contexto y las pruebas de modelo local en un recorrido de
lanzamiento separado.

## Por qué un modelo local necesita un perfil de funcionamiento diferente

La experiencia original de pi-web es una excelente base, pero la inferencia
local tiene modos de fallo diferentes de los de un modelo hospedado típico. Un modelo local puede
ralentizarse bruscamente a medida que crece el contexto, compartir memoria limitada con el resto
de la máquina, detenerse tras producir solo razonamiento, o perder una ejecución larga por un fallo
transitorio del transporte local. Tratar esos casos exactamente como fallos en la nube hace que la UI
parezca compatible mientras la sesión real sigue siendo frágil.

Esta edición aborda el problema por capas:

1. **Preservar upstream primero.** La UI compartida y el comportamiento de sesión
   siguen proviniendo de pi-web; los cambios locales están aislados detrás del
   modo Local efectivo.
2. **Prevenir antes de recuperar.** Se aplica un límite de contexto basado en
   porcentaje del 65% antes de las llamadas posteriores al modelo, incluidas las
   llamadas dentro de bucles de herramientas largos.
3. **Recuperar solo con evidencia.** La continuación automática se limita a incidentes
   reconocidos de contexto, transporte y solo razonamiento; no a autenticación,
   cuota o errores de proveedor arbitrarios.
4. **Limitar cada acción autónoma.** Los incidentes de recuperación se deduplican,
   se requiere progreso antes de otro rescate, y el arranque considera como
   máximo una sesión Local recientemente activa.
5. **Mantener una salida manual.** Force Compact resume en lugar de borrar el
   historial, para que el usuario pueda rescatar una sesión sin fingir que el
   contexto nunca existió.
6. **Proteger la compatibilidad con la nube.** Cloud Mode mantiene las semánticas y
   los controles upstream; las optimizaciones de modelo local no redefinen
   silenciosamente las sesiones de la nube.

Esa es la diferencia real en este fork: trata la inferencia local como un entorno
operativo distinto, no simplemente como otro nombre de modelo en un desplegable.
