package main

import "fmt"

func imprimirResultado(nota float64) {

	//If em go não é preciso utilizar os `()` na expressão, a não seja algum expressão com precedência como `a > c || (a + b > 0)`
	//Obrigatoriamente mesmo que seja uma única sentença de código precisa ter as `{}`
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {

	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
