package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var prod produto
	prod = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(prod, prod.precoComDesconto())

	// Forma definida, precisa ser criado na ordem que foi declarada na struct
	prod2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(prod2, prod2.precoComDesconto())
}
