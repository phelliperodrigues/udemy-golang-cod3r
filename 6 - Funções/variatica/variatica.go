package main

import "fmt"

/*
O fato de usarmos os "..." nos parâmetros indicão que esta função é viriática
Esta função pode receber 1 ... N parâmetros
*/
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(1, 2, 3, 4, 5, 6, 100, 8.9))
}
