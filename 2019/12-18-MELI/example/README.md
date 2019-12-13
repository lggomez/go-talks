# Meetup MELI 2019/12/18: Trace - Ejemplo

En este directorio hay 2 proyectos (requieren go modules, se recomienda 1.13+):

- *server:* servidor web de prueba que responde pings, produce logs, popula un contador sobre un mutex y dado un timer se cierra a los 15 segundos
- *client:* cliente que le pega de manera constante al ping del servidor, se ejecuta durante 13 segundos (para amortiguar el startup y finalizacion del servidor)

El código de los proyectos se puede editar libremente para cambiar estas restricciones y experimentar con los handlers

## Instrucciones

### Automático

- Ejecutar run.sh: este script va a compilar los proyectos, ejecutarlos en orden y obtener el trace. *Opcional:* make.sh va solamente a compilar los proyectos y mover los binarios al directorio /example

*PD:* en caso de no cerrarse automáticamente, se puede suspender la ejecución del mismo una vez que se haya cerrado el servidor web. Si el trace no se obtuvo con éxito, volver a ejecutarlo

### Manual

- En client/main/server: `go run main.go`
- En client/main/client: `go run main.go`
- `curl -XGET "localhost:5050/debug/pprof/trace?seconds=5" --output example_trace`
- go tool trace example_trace