package main

import "fmt"

// Em Go lang não possui operadores ternarios

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// A Operação ternaria ficaria
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
	fmt.Println(obterResultado(5.2))
}
