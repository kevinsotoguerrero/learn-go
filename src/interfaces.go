package main

import "fmt"

//INTEFACES

/* son funcionalidades donde se agrupan metodos en comun que son utilizados por diferentes structs para su facilidad de uso
 */

type figuras2D interface {
	area() float64 //funcion que tiene esta interface y lo que devuelve
}
type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main4() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	calcular(myCuadrado)
	calcular(myRectangulo)

	//LISTA DE INTERFACES

	/* son utiles para almacenar diferentes tipos de datos en un slice*/

	myInterface := []interface{}{"hola", 12, true, 4.9}
	fmt.Println(myInterface...)

	//CONCURRENCIA

}
