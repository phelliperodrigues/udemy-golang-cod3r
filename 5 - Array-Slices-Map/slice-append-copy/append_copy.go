package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	/*
		A tentativa de appendar em um array gera o seguinte erro:
		first argument to append must be a slice; have array1
	*/
	// array1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	/*
		A sentença abaixo copia tudo que tem no slice1 para o slice2
		porem como o slice1 tem 3 elementos e o slice2 é de 2 elementos
		o Go copia somente os 2 primeiros elementos do slice1
		não é possível copiar os valores de um array para o slice : copy(slice2, array1)

	*/
	copy(slice2, slice1)
	fmt.Println(slice2)
}
