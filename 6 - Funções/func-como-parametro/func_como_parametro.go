package main

import "fmt"

/*
Em go podemos utilizar alguns recursos de programação funcional
mesmo não sendo uma linguagem funcional na sua essência.
*/
func multiplacao(a, b int) int {
	return a * b
}

/*
A função `exec()` recebe como parâmetro uma outra função
e para isso devemos passar do tipo como `func(int, int) int`
que a função recebida, tera como parâmetros 2 variáveis int e retorno int
*/
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplacao, 3, 4)
	fmt.Println(resultado)
}
