package main

import "fmt"

func main() {
	//Criando um slice
	s1 := make([]int, 10, 20)
	//Criando segundo slice e apendando os valores 1,2,3
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	//Alterando o valor do index 0 do slice s1 deve também alterar o valor do index 0 do slice s2, pois os dois apontam para o mesmo array interno
	s1[0] = 7
	fmt.Println(s1, s2)

}
