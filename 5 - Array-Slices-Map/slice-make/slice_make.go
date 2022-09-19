package main

import "fmt"

func main() {
	// Criando com a função `make` um slice do tipo []int que contem 10 posições
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	/*
		Criando um slice do tipo []int que contem 10 posições porem o array interno contem 20 posições
	*/
	s = make([]int, 10, 20)
	// 1 - Elementos, 2 tamanho do slice, 3 Capacidades do array
	fmt.Println(s, len(s), cap(s))

	/*
		A partir da função `append` criamos um novo slice da linha 14, porem adicionamos mais os elementos abaixo
		Com isso atingimos a capacidade maxima do slice/array
	*/
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
	/*
		Quando tentamos appendar mais uma informação no slice para "estourar" a capacidade
		o slice passa a ter 21 elementos e a capacidade passa a ser 40 (o dobro), pq o slice
		tem a capacidade de crescer automaticamente, na pratica ele cria arrays maiores e replica os dados
		para este novos arrays.
	*/
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
