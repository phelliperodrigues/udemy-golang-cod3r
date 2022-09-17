package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	// Uma única expressão para executar o for
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	// For tradicional, inicializa variável, Condição, Incremento
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("Misturando ...")

	//Aninamento (não recomendado) smell code 💩
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")

		}
	}

	// for infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

}
