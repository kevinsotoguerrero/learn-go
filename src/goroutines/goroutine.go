package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
func main() {

	//los WaitGroup no son tan usados ya que son dificiles de mantener, es mejor usar los channels para las go rutine
	var wg sync.WaitGroup //el paquete sync interactua primitivamente con las go rutine

	fmt.Println("hello")
	wg.Add(1)            //esto significa que se va a agregar 1 go rutine y que debe esperar a que se termine
	go say("world", &wg) // se le pone go al inicio para que sea concurrente. pero en este caso solo imprime hello porque no le da tiempo de imprimir el world en la concurrencia. se intenta poner un time.Sleep para que espere que termine y de la respuesta completa pero es algo contraproducente porque se demoraria mas que no siendo concurrente

	//time.Sleep(time.Second * 1)//mala practica ya que espera un tiempo independiente de la finalizacion de los hilos de ejecucion

	go func(text string) { //funcion anonima, funcion que se crea dentro del main para usar en el momento
		fmt.Println(text)
	}("adios")
	wg.Wait()
}
