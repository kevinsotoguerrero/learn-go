package main

import (
	"fmt"
	"learn_go/src/mypackage"
)

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand)

}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo un PC %s con %d de RAM y %d de disco duro", myPc.brand, myPc.ram, myPc.disk)

}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar) //{Ford 2020}

	//otra forma de instanciar un struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar) //{Ferrari 0}

	var myCar2 mypackage.CarPublic
	myCar2.Brand = "Mercedez"
	fmt.Println(myCar2) //{Mercedez 0}

	var myCar3 mypackage.CarPublic2
	myCar3.Year = 2020
	fmt.Println(myCar3) //{ 2020}, no se puede acceder al atributo brand porque es privado

	//PUNTEROS> acceso a la memoria
	a := 50
	b := &a         //b es el punerto de a
	fmt.Println(a)  //50
	fmt.Println(b)  //0xc0000120f0
	fmt.Println(*b) //50

	/* el & llama al espacio en la memoria, * llama al valor que hay en ese espacio de memoria, se utiliza para optimizar asignacion de valores*/

	*b = 100
	fmt.Println(a) //100. se cambio a porque *b y a apuntan al mismo esapcio en memoria

	myPc := pc{ram: 16, disk: 256, brand: "Panasonic"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc.ram) //16
	myPc.duplicateRAM()
	fmt.Println(myPc.ram) //32

	fmt.Println(myPc.String())

}
