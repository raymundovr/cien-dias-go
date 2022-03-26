# Cien días de código en Golang
# 💯🐭

Los nombres de los proyectos no reflejan los días. Sólo siguen su propia secuencia.

# Uno
El paquete `os` tiene un arreglo `Args` que almacena los argumentos con que el programa fue invocado
```go
import (
    "fmt"
    "os"
)
...

for index, arg := range os.Args {
    fmt.Println(index, arg);
}
```
