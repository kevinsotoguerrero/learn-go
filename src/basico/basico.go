package main //nombre del paquete donde esta guardado
import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("----------------DECLARACION DE VARIABLES Y CONSTANTES----------------")
	declaracionDeVariablesYConstantes()
	fmt.Println("----------------IMPRIMIR CON FORMATO Y TIPOS DE VARIABLES----------------")
	imprimirConFormatoYTipoDeVariable()

	fmt.Println("--------------------FUNCIONES---------------------")

	fmt.Println("Llamando a la funcion con return:", returnValue(2, 5, "Aguas"))
	var1, _, var3 := returnValueMore(2, 5, "Aguas") //con guion bajo se descarta un valor que pasa la funcion
	fmt.Printf("Llamando %d a la funcion con return:%s\n", var1, var3)

	fmt.Printf("Funcion del area de un circulo con el radio: %.2f\n", areaCirculo(12))

	fmt.Println("---------------------CICLOS Y CONDICIONALES-----------------------")
	ciclosYCondicionales()

	fmt.Println("---------------------CAPTURA DE ERROR Y PARSE-----------------------")
	capturaErrorYParse()

	fmt.Println("---------------------DEFER BREAK Y CONTINUE-----------------------")
	usandoDeferBreakYContinue()

	fmt.Println("---------------------ARRAY SLICE-----------------------")
	arraySlice()

	fmt.Println("---------------------PALINDROMO-----------------------")
	esPalindromo("ama")

	fmt.Println("---------------------MAPS-----------------------")
	usandoMaps()

}

func declaracionDeVariablesYConstantes() {
	//DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.1416

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
}

func imprimirConFormatoYTipoDeVariable() {
	var complejo complex128 = 10 + 8i
	base := 12

	//IMPRIMIR CON FORMATO
	nombre := "Mateo"
	/* %d para enteros
	   %s para string
	   %v es una mala practica pero es si no se sabe el tipo */
	fmt.Printf("%s tiene mas de %d piojos\n", nombre, base)
	fmt.Printf("%v tiene mas de %v piojos\n", nombre, complejo)
	/* Mateo tiene mas de 12 piojos
	   Mateo tiene mas de (10+8i) piojoscomplex128 */

	//IMPRIMIR TIPO DE DATO
	fmt.Println("Tipos de dato:")
	fmt.Println(reflect.TypeOf(complejo))        //con la libreria reflect
	fmt.Printf("Tipo de complejo: %T", complejo) //con la libreria fmt

	pruebaGuardarTipofmt := fmt.Sprintf("%T", complejo)
	pruebaGuardarTiporeflect := reflect.TypeOf(complejo)

	fmt.Println(pruebaGuardarTipofmt)                     //complex128
	fmt.Println(pruebaGuardarTiporeflect)                 //complex128
	fmt.Println(reflect.TypeOf(pruebaGuardarTipofmt))     //string
	fmt.Println(reflect.TypeOf(pruebaGuardarTiporeflect)) //*reflect.rtype
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

func ciclosYCondicionales() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	/* counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	} */

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for j, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", j, par)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	//Switch
	modulo := 5 % 2
	switch modulo {
	//switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("no es tan par")
	}

	//Switch sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("mayor que 100")
	case value < 100 && value > 0:
		fmt.Println("mayor que 0 pero menor que 100")
	default:
		fmt.Println("negativo")
	}

}

func capturaErrorYParse() {
	value, err := strconv.Atoi("53")
	//value, err := strconv.Atoi("oewi")
	//cuando no hay error retorna nil en la variable err
	if err != nil {
		//con log.Fatal se muestra el mensaje y se para proceso
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}

func usandoDeferBreakYContinue() {
	//
	defer fmt.Println("hola")
	ejemploDeferAnidado()
	fmt.Println("mundi")
	defer fmt.Println("probando2")
	fmt.Println("munti")

	defer fmt.Println("probando")

	/* si esta funcion solo llegara hasta aqui imprimiria:
	ani1
	ani3
	ani2
	mundi
	munti
	probando
	probando2
	hola */

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

func ejemploDeferAnidado() {
	fmt.Println("ani1")
	defer fmt.Println("ani2")
	fmt.Println("ani3")
}

func arraySlice() {
	// ARRAY

	var array [4]int
	array[0] = 1
	fmt.Println(array, len(array), cap(array))

	//esto imprime [1 0 0 0] 4 4
	//queda declarado con 4 valores solo se ingresa el primero en la posicion cero los otros quedan con los zero values, len da la cantidad de datos que hay en este momento que son 4 y cap la capacidad de datos que se pueden guardar en ese array.

	//SLICE

	var slice []int
	slice = append(slice, 2)
	//array = append(array, "r")
	fmt.Println(slice, len(slice), cap(slice)) //[2] 1 1
	fmt.Println(reflect.TypeOf(array))         //con la libreria reflect
	fmt.Println(reflect.TypeOf(slice))         //con la libreria reflect

	//EJEMPLO  de la diferencia entre slice y array: cuando se itera en un for range y se cambian los valores de la estructura, un slice va tomando los que se cambian en el proceso mientras que un array toma los que estaban planteados al inicio del for
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		if index == len(numbers)-1 {
			numbers[0] += value

		} else {
			numbers[index+1] += value
		}
	}
	fmt.Println("numbers con slice")
	fmt.Println(numbers)

	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers2 {
		if index == len(numbers2)-1 {
			numbers2[0] += value
		} else {
			numbers2[index+1] += value
		}
	}
	fmt.Println("numbers con array")
	fmt.Println(numbers2)

	//los slice se diferencian en a los array en que no se dice cuantos valores van a almacenar en la declaracion
	slice3 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice3, len(slice3), cap(slice3)) // imprime [0,1,2,3,4,5,6] 7 7
	//Metodos del slice3
	fmt.Println(slice3[0])   //0
	fmt.Println(slice3[:3])  //0,1,2
	fmt.Println(slice3[2:4]) //2,3
	//estos metodos tambien se pueden usar en los array

	//los array son inmutables los slice son mutables, por lo tanto no se puede agregar mas elementos al array

	slice3 = append(slice3, 7)
	newSlice3 := []int{8, 9, 10}
	slice3 = append(slice3, newSlice3...) //los tres puntos significan que va a tomar cada elemendo de newSlice3 y lo pasa como argumento, es como si se pusiera append(slice3, 8,9,10)

	//como recorrer los elementos

	slice2 := []string{"hola", "que", "hace"}
	for indice, valor := range slice2 {
		fmt.Println(indice, valor)
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

}

func esPalindromo(palabra string) bool {

	es := false
	lenPalabra := len(palabra)
	limite := lenPalabra / 2
	for i := 0; i < limite; i++ {
		fmt.Println(palabra[i])         //97
		fmt.Println(string(palabra[i])) //a
		//un string se puede iterar caracter por caracter pero al llamarse una posicion devuelve el codigo ascci, con la funcion string() devuelve la letra
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
