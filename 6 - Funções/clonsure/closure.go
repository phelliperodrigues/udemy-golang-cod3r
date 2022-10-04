package main

import "fmt"

/*
	O conceito de closure, alem do fato que a função tem seu próprio escopo onde ela foi definida, e quando
	colocamos a função para um outro contexto diferente, ela tem essa memória, do local onde ela foi definida, e de que escopo ela foi definida
	para que essa possa ler as variáveis do escopo correto
*/

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {
	x := 20

	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
