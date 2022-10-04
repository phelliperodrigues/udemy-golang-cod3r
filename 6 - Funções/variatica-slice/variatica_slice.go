package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	//Para passar um slice em uma função viriática basta colocar os três pontinhos apos a variável
	/*
		Não pode passar um array como parâmetro, precisa ser um slice
			aprovados := [4]string{"Maria", "Pedro", "Rafael", "Guilherme"}
			ou
			aprovados := [...]string{"Maria", "Pedro", "Rafael", "Guilherme"}
		Não funcionaria
		cannot use aprovados (variable of type [4]string) as type []string in argument to imprimirAprovados

	*/
	imprimirAprovados(aprovados...)
}
