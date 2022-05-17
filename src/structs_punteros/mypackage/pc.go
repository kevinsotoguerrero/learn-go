package mypackage

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "pong")

}

//si se *pc se envia la variable y los cambios que se realicen se haran a la variable como tal
func (myPc *Pc) DuplicateRAM() {
	myPc.Ram = myPc.Ram * 2
}

//si se pasa pc se envia una copia de la variable, pero no la variable en si, por lo cual el valor cambia pero solo dentro de la funcion
func (myPc Pc) DuplicateRAM2() {
	myPc.Ram = myPc.Ram * 2
	fmt.Println(myPc.Ram)
}

//String customiza la salida cuando se llame la instancia del struct
func (myPc Pc) String() string {
	return fmt.Sprintf("Tengo un PC %s con %d de RAM y %d de disco duro", myPc.Brand, myPc.Ram, myPc.Disk)

}
