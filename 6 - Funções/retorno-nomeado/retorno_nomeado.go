package main

import "fmt"

/*
Em go podemos nomear o retorno, no exemplo abaixo
Temos 2 retorno do tipo `int` com nomes de variáveis
quando atribuímos valores para essas variáveis basta
colocarmos a anotação `return` no final da função
que o go ja entende que deve retornar essas variáveis
*/
func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // isso é chamado "retorno limpo"
}

func main() {
	x, y := troca(2, 3)

	fmt.Println(x, y)

	segundo, primeiro := troca(1, 2)

	fmt.Println(segundo, primeiro)
}
