package main

import (
	"fmt"
	//con un alias para facilitar el llamado
	pk "personalProjects/learn_go/src/structs_punteros/mypackage"
)

func main() {

	fmt.Println("--------------------------STRUCTS----------------------------")
	structs()

	fmt.Println("--------------------------PUNTEROS----------------------------")
	punteros()

}

func structs() {
	myCar := pk.CarPublic{Brand: "Ford", Year: 2020}
	fmt.Println(myCar) //{Ford 2020}

	//otra forma de instanciar un struct
	var otherCar pk.CarPublic
	otherCar.Brand = "Ferrari"
	fmt.Println(otherCar) //{Ferrari 0}

	var myCar2 pk.CarPublic
	myCar2.Brand = "Mercedez"
	fmt.Println(myCar2) //{Mercedez 0}

	var myCar3 pk.CarPublic2
	myCar3.Year = 2020
	fmt.Println(myCar3) //{ 2020}, no se puede acceder al atributo brand porque es privado

	pk.Message()
}

func punteros() {
	//PUNTEROS> acceso a la memoria
	a := 50
	b := &a         //b es el puntero de a
	fmt.Println(a)  //50
	fmt.Println(b)  //0xc0000120f0
	fmt.Println(*b) //50

	/* el & llama al espacio en la memoria, * llama al valor que hay en ese espacio de memoria, se utiliza para optimizar asignacion de valores*/

	*b = 100
	fmt.Println(a) //100. se cambio a porque *b y a apuntan al mismo esapcio en memoria

	myPc := pk.Pc{Ram: 16, Disk: 256, Brand: "Panasonic"}
	fmt.Println(myPc)

	myPc.Ping()

	fmt.Println(myPc.Ram) //16
	myPc.DuplicateRAM2()
	fmt.Println(myPc.Ram) //32

	fmt.Println(myPc) //Tengo un PC Panasonic con 16 de RAM y 256 de disco duro. output personalizado
}
