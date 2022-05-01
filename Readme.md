# Cien días de código en Golang
# 💯🐭

Los nombres de los proyectos no reflejan los días. Sólo siguen su propia secuencia. Están ordenados de más reciente a más antiguo.

## Nueve
Go no tiene en estos momentos un buen soporte de genéricos. Por ejemplo no es posible definir una estructura que tenga un valor genérico que permita hacer comparaciones con todos los operadores: ==, !=, <, > <=, >=.

La interfaz `comparable` sólo opera con == y !=.

Para poder crear una estructura que pueda soportar un valor con múltiples tipos se puede declarar una interfaz.
```go
type Valuable interface {
	int | float64 | string
}
```
Posteriormente se usa dicha interfaz para declarar una `struct` usando la misma sintaxis como en `generic`.
```go
type BinaryTree[T Valuable] struct {
	value		T
	left,right	*BinaryTree[T]
}
```


## Ocho

Si los archivos .go de una carpeta no corresponden todos al mismo `package` hay errores.

Para organizarlo se pueden crear subcarpetas, cada una con archivos que declaran un `package` diferente.

Si no hay un `package main` no se puede ejecutar usando `go run .`. Esto significa que paquetes diferentes a main sólo pueden ser librerías.

Es posible sobrecargar la representación de un tipo mediante una función, por ejemplo aquí se hace con `string`
```go

func (f Farenheit) String() string {
	return fmt.Sprintf("%.2f °F", f)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}
```

De la misma manera
```go
func (f Farenheit) IntoCelsius() Celsius {
	return Celsius( (f - 32) / 1.8000 )
}

func (k Kelvin) IntoCelsius() Celsius {
	return Celsius(k - 273.15)
}
```

Para usar la sobrecarga
```go
temp := Celsius(37)
conv := temp.IntoFarenheit()
convStr = conv.String()
```
## Siete

Declarar constantes que asemejan a `enum`s
```go
type EventType string

const (
	one		EventType = "one"
	two				  = "two"		
	three			  = "three"
)

type EventTypeNumber uint8

const (
	oneNumber	EventTypeNumber = iota
	twoNumber
	threeNumber
)

```
`iota` genera autonumeración comenzando en cero.

## Seis

Siguiendo https://www.mongodb.com/languages/golang y https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures


## Cinco

Declaraciones de estructuras son
```go
type nombre struct {
    ...
}
```

Los campos van dentro de la declaración indicando su tipo de datos
```go
type person struct {
    ID string
}
```

Las estructuras soportan anotaciones para los campos usando \`\`, por ejemplo para convertir a un campo JSON
```go
type person struct {
    ID string `json:"id"` 
}
```

Declarar variables con una estructura como tipo se da de la forma
```go
var onePerson = person{ ID: "1" }

var persons = person[]{
    {ID: "1"},
    {ID: "2"},
}
```

El paquete [gin](https://gin-gonic.com) es un web framework. Opcionalmente se necesita el paquete `net/http` para ciertas constantes estándares en HTTP.
```go
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Run("localhost:8080")
}
```

"""
gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)
"""

Las funciones que realizan acciones reciben un apuntador a `gin.Context`.

```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
```

Gin también puede transformar los datos enviados en una nueva entidad a partir de la struct.

Para esto usa `c.BindJSON()`
```go
func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return ;
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

Se observa que cuando se envían datos vacíos la entidad se crea con datos por default.
```json
...
  {
    "id": "",
    "title": "",
    "artist": "",
    "price": 0
  }
...
```

También que `gin` envía automáticamente un error 400 si falla el parse y no se manipula la respuesta del servidor.

Una respuesta puede enviar un objeto a partir de la struct `H` de `gin`.
```go
c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
```

Remover un elemento en un slice se realiza mendiante la composición de los elementos excluyentes

```go
    for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.Status(200)
			return
		}
	}
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

Para separar funcionalidades en diferentes archivos dentro del mismo paquete sólo hace falta declarar el mismo en los diferentes archivos
```go

// main.go
package main
...


// filters.go
package main
```

El paquete `image` tiene colores en el rango [0, 65535]. La explicación está [en el blog del paquete](https://go.dev/blog/image). Para usarlo es necesario convertilo a `uint8` ejemplo `ru8 := uint8(r / 0x101)`.

Para escribir una imagen a archivo es necesario utilizar un descriptor y un codificador dependiendo del tipo de imagen

```go
    if imgFormat == "png" {
		saveErr := png.Encode(writer, img)
		if saveErr != nil {
			writer.Close()
			return false;
		}
	}
```

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

