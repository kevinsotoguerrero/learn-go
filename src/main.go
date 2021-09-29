package main //nombre del paquete donde esta guardado
import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main2() {
	//DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//DECLARACION DE VARIABLES
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area) //12 14 0
	/* int separa 64bits en maquinas de 64bits y 32 en las de 32bits
	   int8 separa 8 bits, entonces almacena numeros desde el -128 hasta el 127
	   uint es el mismo int pero solo para numeros positivos
	   uint8 almacena numeros desde el 0 hasta 255 */

	//ZERO VALUES
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d) //0 0  false

	//OPERADORES
	x := 12
	y := 50

	suma := x + y
	modulo := x % y
	x++
	y--

	fmt.Println(suma, modulo, x, y)

	fmt.Println(math.Pi)

	//NUMEROS COMPLEJOS
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64

	var complejo complex128 = 10 + 8i

	fmt.Println(complejo)

	//IMPRIMIR CON FORMATO
	nombre := "Mateo"
	/* %d para enteros
	   %s para string
	   %v es una mala practica pero es si no se sabe el tipo */
	fmt.Printf("%s tiene mas de %d piojos\n", nombre, base)
	fmt.Printf("%v tiene mas de %v piojos", nombre, complejo)
	/* Mateo tiene mas de 12 piojos
	   Mateo tiene mas de (10+8i) piojoscomplex128 */

	//IMPRIMIR TIPO DE DATO
	fmt.Println(reflect.TypeOf(complejo))        //con la libreria reflect
	fmt.Printf("Tipo de complejo: %T", complejo) //con la libreria fmt

	pruebaGuardarTipofmt := fmt.Sprintf("%T", complejo)
	pruebaGuardarTiporeflect := reflect.TypeOf(complejo)

	fmt.Println(pruebaGuardarTipofmt)
	fmt.Println(pruebaGuardarTiporeflect)
	fmt.Println(reflect.TypeOf(pruebaGuardarTipofmt))
	fmt.Println(reflect.TypeOf(pruebaGuardarTiporeflect))

	/* complex128
	   complex128
	   string
	   *reflect.rtype */

	fmt.Println("Llamando a la funcion con return:", returnValue(2, 5, "Aguas"))
	var1, var2, var3 := returnValueMore(2, 5, "Aguas")
	fmt.Printf("Llamando %d a la funcion %d con return:%s\n", var1, var2, var3)

	fmt.Printf("Funcion del area de un circulo con el radio: %.2f\n", areaCirculo(12))

	// ARRAY

	var array [4]int
	array[0] = 1
	fmt.Println(array, len(array), cap(array))
	//esto imprime [1 0 0 0] 4 4
	//queda declarado con 4 valores solo se ingresa el primero en la posicion cero los otros quedan con los zero values, len da la cantidad de datos que hay en este momento que son 4 y cap la capacidad de datos que se pueden guardar en ese array.

	// SLICE

	//los slice se diferencian en a los array en que no se dice cuantos valores van a almacenar en la declaracion
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice)) // imprime [0,1,2,3,4,5,6] 7 7
	//Metodos del slice
	fmt.Println(slice[0])   //0
	fmt.Println(slice[:3])  //0,1,2
	fmt.Println(slice[2:4]) //2,3
	//estos metodos tambien se pueden usar en los array

	//los array son inmutables los slice son mutables, por lo tanto no se puede agregar mas elementos al array

	slice = append(slice, 7)
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) //los tres puntos significan que va a tomar cada elemendo de newSlice y lo pasa como argumento, es como si se pusiera append(slice, 8,9,10)

	//como recorrer los elementos

	slice2 := []string{"hola", "que", "hace"}
	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}
	//0 hola
	//1 que
	//2 hace

	//si se quiere omitir el indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	//si se quiere omitir el valor
	for i := range slice {
		fmt.Println(i)
	}

	usandoMaps()

}

//FUNCIONES

//VAN POR FUERA DE LA FUNCION MAIN
func returnValue(a, b int, c string) string {
	return fmt.Sprintf("%s es igual a %d", c, a*b)
}
func returnValueMore(a, b int, c string) (d, e int, f string) {
	return a * b, a - b, fmt.Sprintf("%s es igual a %d", c, a*b)
}
func areaCirculo(radio float64) float64 {
	return 2 * math.Pi * radio
}

//CICLOS

func ciclos() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for j, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", j, par)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

func usandoBreakContinue() {
	for i := 0; i < 10; i++ {
		fmt.Println("oes mi so")
		if i == 2 {
			fmt.Println("oes mi so en el 2")
			continue //hace que pase al siguiente valor del i sin terminar el ciclo
		}

		if i == 6 {
			fmt.Println("oes mi so siendo 6")

			break //hace que se termine el ciclo
		}

		fmt.Println("oes mi so desde afuera")

	}
}

//DEFER

func usandoDefer() {
	//defer hace que esa linea de codigo se ejecute al final pero en la misma funcion. Entonces si si se llama algo antes esta funcion despues y despues otra funcion imprimiria la anterior en este imprimiria mundi, despues hola y despues la siguiente funcion. Se puede usar al llamar la funcion en el main y seria la ultima funcion que se ejecutaria. se puede usar mas de un defer y se imprime de ultimo el primero que se puso
	defer fmt.Println("hola")
	fmt.Println("mundi")
	defer fmt.Println("probando2")
	fmt.Println("munti")

	defer fmt.Println("probando")

	/* imprime: mundi
	munti
	probando
	probando2
	hola */
}

func esPalindromo(palabra string) bool {

	es := false
	lenPalabra := len(palabra)
	limite := lenPalabra / 2
	for i := 0; i < limite; i++ {
		if strings.ToUpper(string(palabra[i])) == strings.ToUpper(string(palabra[lenPalabra-1-i])) {
			es = true
		} else {
			return false
		}

	}
	return es
}

func usandoMaps() {
	//Llave valor con maps
	m := make(map[string]int)
	m["jose"] = 14
	m["felipe"] = 20
	fmt.Println(m) //map[felipe:20 jose:14]

	//recorrer map

	for i, valor := range m {
		fmt.Println(i, valor)
	}
	/* jose 14
	felipe 20 */

	//los maps son concurrentes, asi que pueden imprimir los valores de forma aleatorea, si se requiere un orden en la impresion es mejor usar array o slice. son mas agiles por la implementacion de concurrencia de forma nativa

	//encontrar un valor
	valor, ok := m["mario"]
	fmt.Println(valor, ok) //0 false devuelve esto ya que mario no es una llave en el diccionario entonces devuelve el  zero value, con el segundo valor se puede capturar si la llave que se ingresa existe en el diccionario
}
