package main

import "fmt"

func main() {
	/*
		A anotação `[...]` significa que estou criando um array com número fixo, só que quero que o tamanho
		seja a partir da quantidade de valores que eu inicializa-lo.
		O Compilador que vai contar a quantidade de elementos e atribuir a variável
	*/
	numbers := [...]int{1, 2, 3, 4, 5}

	/*
		O Exemplo abaixo é equivalente ao do foreach
		```java
		var numbers = List.of(1,2,3,4,5)

		for (var number : numbers) {
			System.out.println(number)
		}
		```
	*/
	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i, number)
	}

	//Podemos ignorar o índice colocando o `_` no lugar da variável
	for _, number := range numbers {
		fmt.Println("Valor: ", number)
	}

	// Se não adicionar o elemento do índice `_` ou `i` a variável `number` será o index e não o valor do elemento
	for number := range numbers {
		fmt.Println("Index: ", number)
	}
}
