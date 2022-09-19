package main

import "fmt"

func main() {
	// Inicializando o map de forma literal
	funcionariosAndSalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}
	// Adicionando novo valor no MAP
	funcionariosAndSalarios["Rafael Filho"] = 1350.0

	// Não tem problema tentar excluir um elemento inexistente
	delete(funcionariosAndSalarios, "Inexistente")

	for nome, salario := range funcionariosAndSalarios {
		fmt.Println(nome, salario)
	}
}
