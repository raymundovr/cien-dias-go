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


Se pueden declarar constantes usando `const`. Sólo pueden ser `number`, `string` o `boolean`

## Tres

El paquete `io/ioutil` soporta carga de archivos

```go
import "io/ioutil"

func main() {
    ...
    data, err := ioutil.ReadFile(path)
}
```

`data` es un arreglo de `byte`, para convertirlo en `string`

```go
...
dataString := string(data)
...
```

## Cuatro
Go tiene manipulación de imágenes en la librería estandar mediante el paquete [img](https://pkg.go.dev/image).
```go
import "img"
```

Para abrir una imágen es necesario primero crear un `io.Reader`
```go
reader, err := os.Open(filePath)
img, s, err := image.Decode(reader)
```
En este caso `s` es el formato de la imagen (png, jpeg, gif), `img` contiene una estructura con referencia a la imagen.

Es posible cargar una librería sólo por sus efectos secundarios mediante `import _`, ejemplos:

```go
import _ "image/png"
import _ "image/jpeg"

// juntos

import (
    _ "image/png"
    _ "image/jpeg"
)
```

En este caso después de importar es posible operar con imágenes png y jpeg.

El paquete `os` también permite leer archivos, aunque en este caso retorna un `reader`
```go
err, reader := os.Open(filePath)
```

Usar `defer` permite detener la ejecución de una función hasta que todas las llamadas o manipulaciones bajo ella hayan sido completadas

```go

reader, err := os.Open(filePath)
if err != nil {
	fmt.Printf("Error reading file %s: %v", filePath, err)
	os.Exit(1)
}
	
defer reader.Close()
```