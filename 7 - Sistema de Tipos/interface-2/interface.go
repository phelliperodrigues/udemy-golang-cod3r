package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	tuboLigado      bool
	velocidadeAtual int
}

//É preciso passar ferrari como ponteiro para poder de fato alterar o valor da `instacia`
func (f *ferrari) ligarTurbo() {
	f.tuboLigado = true
}

func main() {
	carro1 := ferrari{"f40", false, 0}
	carro1.ligarTurbo()

	//Quando criar uma variavel tipada com a interfaca é preciso referenciar o ponteiro
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro1)
}
