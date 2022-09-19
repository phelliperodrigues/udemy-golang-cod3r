package main

import "fmt"

func main() {
	/* Arrays são estruturas homogêneas ou seja, eles tem o mesmo tipo de dados dentro dele,
	e é uma estrutura estáticas, que significa que é fixo, ou seja, se definir um array com 10 posições ele
	nunca mudará, o máximo que podemos fazer é pegar a variável e apontar para uma outra estrutura de tamanho diferente
	*/
	var notas [3]float64
	// o array por padrão é inicializado com os valores `zeros` do tipo de atributo, no caso abaixo o valor `0` do float64
	fmt.Println(notas)

	/* O array é uma estrutura indexada, significa que cada elemento tem um índice, conseguimos acessar o elemento a partir do índice,
	começando com índice `0`
	*/
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// Se tentar atribuir algo no índice `3` teremos um erro de compilação `invalid argument: index 3 out of bounds [0:3]`
	// notas[3] = 10
	fmt.Println(notas)

	// Calcular a média utilizando o for tradicional
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)

}
