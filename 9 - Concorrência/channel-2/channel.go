package main

import (
	"fmt"
	"time"
)

// Channel (nanal) - é a forma de comunicação entre gorotines
// channel é um tipo
func doisTresQuatroVezes(base int, c chan int) {

	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)

	// se tentar receber outro channel vai dar um erro de `fatal error: all gorotines ares asleep - deadlock!` por que
	// não tem existe mais ninguem para enviar algo ao channel
	// fmt.Println(<-c)
}
