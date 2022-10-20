package main

import "fmt"

type curso struct {
	nome string
}

/*
	Por conta do GO lang ser fortemente tipado, podemos criar uma variavel do tipo "interface" nos possibilitando usar
como se fosse o `Object` do java
*/
func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{nome: "Golang: Exlplorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
