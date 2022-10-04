package main

import "fmt"

// Toda função recursiva precisa ter uma condição de parada para não gerar um StackOverFlow
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-5)
	if err != nil {
		fmt.Println(err)
	}
}
