# ¿Por qué pi-web?

Soy un poco adicto a Claude Code. Siempre lo estoy usando. Si no estoy sentado frente a la computadora, estoy pensando en él. Siento que no estoy quemando suficientes tokens. Eran los primeros días de Claude Code. Y pensaba, ¿por qué no puedo retomarlo desde mi teléfono? Configuré Termius y realmente no me gustó.

Empecé a crear el mío y me detuve cuando Claude lanzó su aplicación móvil de Claude Code.

Luego sufrí una hernia de disco y realmente no podía hacer mucho. Pasó el tiempo y me sentí un poco recuperado y quería continuar mi proyecto de Claude Code vía web/pwa.

Luego Claude Code comenzó a prohibir el uso fuera de su propio entorno. Y siento que ya no vale la pena.

Entonces encontré pi.dev y exploré un poco pero no me había sumergido realmente. Leí sobre ello, vi videos sobre ello y decidí probarlo a fondo y ahora estoy totalmente metido en pi.

Como es de código abierto, siento que vale la pena construir para ello. También tengo diferentes opciones de proveedores. También siento que depender de un solo proveedor/modelo como Anthropic/Claude no es sostenible.

Así que lo estoy construyendo aquí.

## Por qué un modelo local necesita un perfil operativo diferente

La experiencia original de pi-web es una base excelente, pero la inferencia local tiene modos de fallo diferentes a los de un modelo alojado típico. Un modelo local puede ralentizarse drásticamente a medida que crece el contexto, compartir memoria limitada con el resto de la máquina, detenerse después de producir solo razonamiento, o perder una ejecución larga debido a un fallo transitorio de transporte local. Tratar esos casos exactamente como fallos en la nube hace que la interfaz de usuario parezca compatible, mientras que la sesión real sigue siendo frágil.

Esta edición aborda el problema en capas:

1. **Preservar upstream primero.** La interfaz de usuario compartida y el comportamiento de la sesión continúan viniendo de pi-web; los cambios locales están aislados detrás del modo Local Mode efectivo.
2. **Prevenir antes de recuperar.** Se aplica un límite de contexto del 65% basado en porcentaje antes de las llamadas posteriores al proveedor, incluidas las llamadas dentro de bucles de herramientas largos.
3. **Recuperar solo con evidencia.** La continuación automática se limita a incidentes reconocidos de contexto, transporte y solo-pensamiento, no a errores de autenticación, cuota o errores arbitrarios del proveedor.
4. **Limitar cada acción autónoma.** Los incidentes de recuperación se deduplican, se requiere progreso antes de otro rescate, y el inicio considera como máximo una sesión Local activa recientemente.
5. **Mantener una salida manual.** Force Compact resume en lugar de borrar el historial, por lo que el usuario puede rescatar una sesión sin fingir que el contexto nunca existió.
6. **Proteger la compatibilidad con la nube.** Cloud Mode mantiene la semántica y los controles de upstream; las optimizaciones del modelo local no redefinen silenciosamente las sesiones en la nube.

Esa es la verdadera diferencia en este fork: trata la inferencia local como un entorno operativo distinto, no simplemente como otro nombre de modelo en un menú desplegable.
