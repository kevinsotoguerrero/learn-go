package main

import (
	"fmt"
)

//la `<-` es opcional, pero es buena practica en caso de que el channel solo vaya a recibir informacion. `c <-chan string` en caso de que solo vaya a entregar informacion
func say(text string, c chan<- string) {
	c <- text //indica que se va a ingresar ese dato en ese canal
}
func say2(text string, c chan string) {
	c <- text //indica que se va a ingresar ese dato en ese canal
}

func message(message string, c chan<- string) {
	c <- message
}
func main() {
	c := make(chan string, 1) //el parametro del final no es necesario pero es buena practica ponerlo. Es el numero de go routines. Se le dice buffer
	fmt.Println("oess")

	go say("Hello", c)

	fmt.Println(<-c)

	c2 := make(chan string, 2)
	c2 <- "mensaje1"
	c2 <- "mensaje2"
	fmt.Println(len(c2), cap(c2))

	close(c2) //para cerrar el canal para que no le ingresen mas datos, asi tenga capacidad

	//antes de hacer el for toca cerrar el canal o sino sigue esperando que le ingrese mas informacion
	for message := range c2 {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje en email1", email1)
	go message("mensaje en email2", email2)

	for i := 0; i < 2; i++ {
		//para tomar una decision cuando un hilo de ejecucion ya haya terminado
		select {
		case m1 := <-email1:
			fmt.Println(m1)
		case m2 := <-email2:
			fmt.Println(m2)
		}
	}

}
