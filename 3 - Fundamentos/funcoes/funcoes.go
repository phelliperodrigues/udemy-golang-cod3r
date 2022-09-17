package main

import "fmt"

// Uma função pode ter 0 ou N valores de retorno
// Uma função pode ter 0 ou N valores de parâmetros
// A função abaixo retorna um valor `int`
func somar(a int, b int) int {
	return a + b
}

// A função abaixo é do tipo `void` pois não retorna nenhum valor
func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	//Declarando variável `resultado` a partir da chamada da função `somar()`
	resultado := somar(3, 4)

	imprimir(resultado)
}
