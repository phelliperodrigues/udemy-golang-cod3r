package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// Pegar a data atual e pegar o nano segundo
	s := rand.NewSource(time.Now().UnixNano())
	// pegando nano segundo e gerar numero aleatorio
	r := rand.New(s)
	// Pega numero aleatório até 10
	return r.Intn(10)
}

func main() {
	// A variável `i` fica visível somente dentro do IF ou ELSE
	// Inicializa variável no primeiro bloco antes do `;` e valida a condição após do `;`
	// Também é suportando no switch
	// Lembra bastante o for, só que não contem o incremento
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}
}
