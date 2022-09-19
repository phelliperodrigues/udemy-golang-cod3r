package main

import "fmt"

func main() {
	//var aprovados map[int] string
	//Maps devem ser inicializados
	aprovados := make(map[int]string)
	// A chave não pode ser duplicada, caso vc "add" uma chave que ja exista vc sobrescrevera o valor anterior
	aprovados[12345678989] = "Maria"
	aprovados[98765432123] = "Pedro"
	aprovados[12345678909] = "Ana"
	fmt.Println(aprovados)

	// Percorrendo map e exibindo valores
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Buscando valor pela chave
	fmt.Println(aprovados[12345678909])

	// Deletando valor pela chave
	delete(aprovados, 12345678909)
	fmt.Println("=>", aprovados[12345678909])
}
