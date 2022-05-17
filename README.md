# Go

Es compilado y tipado

## Para correr archivo

`go run nombreArchivo.go`

o

`go build nombreArchivo.go` para compilar generando un archivo nombreArchivo.exe

## Declaracion de variables y constantes

int separa 64bits en maquinas de 64bits y 32 en las de 32bits
int8 separa 8 bits, entonces almacena numeros desde el -128 hasta el 127
uint es el mismo int pero solo para numeros positivos
uint8 almacena numeros desde el 0 hasta 255

Si se declara una variable sin asignarle un valor se le asignan los zero values

### ZERO VALUES

var a int //0
var b float64 //0
var c string //""
var d bool //false

## DEFER, CONTINUE Y BREAK

- **defer** para correr esa linea de codigo al final de la ejecucion pero dentro de la misma funcion o bloque de codigo
- **continue** en un ciclo, para continuar con la siguiente iteracion sin terminar el proceso actual
- **break** en un ciclo, para terminar el ciclo


## ARRAY Y SLICE

Si se requiere almacenar diferentes tipos de datos se usa un slice de interfaces

### Diferencias

- cuando se itera en un for range y se cambian los valores de la estructura, un slice va tomando los que se cambian en el proceso mientras que un array toma los que estaban planteados al inicio del for
- un array se declara con cantidad de ubicaciones definida, mientras que el slice no
- el array es inmutable, asi que no se le pueden agregar mas posiciones

## STRUCTS Y GO MODULES

Struct son las clases en Go
Go modules es el gestor de paquetes en go, para poder utilizar packetes sin necesidad de utilizar la GOPATH
`go mod init` para iniciar el go modules
`go mod tidy` para descargar dependencias
con `go env` vemos las variables que utiliza Go


## Modificadores de acceso

Con la primera letra en mayuscula es publico, en minuscula es privado (exported and unexported)

## PUNTEROS

Referencia a la ubicacion en memoria
`&` hace referencia a la ubicacion en memoria de un valor
`*` hace referencia a al valor de una ubicacion en memoria

Si en un metodo se pasa la referencia al valor se modifica el valor, si se pasa solo el valor, sin referencia, se esta enviando una copia que es cambiada o utilizada solo en el metodo, pero la variable como tal continua igual


## CONCURRENCIA Y GO ROUTINE

La funcion main corre en una go routine
El paralelismo toma la cantidad de tareas y las procesa por hilos de ejecucion de forma paralela cada una, al finalizar una, anuncia la finalizacion y muere el hilo de ejecucion y asi con cada una. La concurrencia contaria con los mismos hilos de ejecucion pero en vez de esperar que el proceso este finalizado pasa a otro, y cuando este finalizado pasa a otro y se anuncia la finalizacion de las tareas en conjunto. 

