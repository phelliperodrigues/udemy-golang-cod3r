package main

import "fmt"

// Toda função recursiva precisa ter uma condição de parada para não gerar um StackOverFlow

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
