package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	/*
		- O Slice serve geralmente para referenciar parte de um array
		- Slice não é um array
		- Slice É um pedaço de uma array.
		- O exemplo abaixo estamos passando para o slice `s2` os valores da posição 1 e 2
			o valor `3` na declaração `a2[1:3]` significa que vou pegar da posição 1 até a 3 mas não pegarei a posição 3
		- Mas o slice `s2` não gerou uma referencia diferente, ele simplesmente criou uma estrutura que tem o ponteiro dos
			elementos no qual ele aponta no array
	*/
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)
	/*
		- A sentença abaixo cria um novo slice apontando do inicio do array até o index 2, mas sem atribuir o valor do index 2
		- É um novo slice, mas internamente ele continua apontamos para a mesma referencia do array `a2`
	*/
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// Podemos imaginar um slice tendo um tamanho e um ponteiro para um elemento de um array

	//Podemos fazer um slice a partir de outro slice
	//Mesmo assim o slice `s4` aponta para o endereço do array `a2`
	s4 := s2[:1]
	fmt.Println(s2, s4)

}
