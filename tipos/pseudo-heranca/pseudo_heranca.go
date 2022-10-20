package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anónimos - Tudo que é de carro vai estar disponivel em ferrari
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s esta com turbo lidado? %v\n", f.nome, f.turboLigado)
}
