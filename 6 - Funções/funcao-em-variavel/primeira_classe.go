package main

import "fmt"

// Podemos atribuir uma função à uma variável
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))
	//Podemos criar uma função local
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
