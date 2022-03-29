# Cien días de código en Golang
# 💯🐭

Los nombres de los proyectos no reflejan los días. Sólo siguen su propia secuencia.

# Uno
El paquete `os` tiene un `slice`, `Args`, que almacena los argumentos con que el programa fue invocado
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

El paquete `strconv` contiene funciones para hacer conversiones de strings
```go
// val == 0, err != nil
val, err := strconv.Atoi("x")


// val == 10, err == nil
val, err := strconv.Atoi("10")
```

Para hacer un join the strings se puede usar el paquete `strings`
```go
strings.Join(os.Args[1:], " ")
```

El estilo recomendado para nombrar en Go es usar `camelCase` para los identificadores, excepto los que son de acceso público, que utilizan `PascalCase`.


Se pueden declarar constantes usando `const`.