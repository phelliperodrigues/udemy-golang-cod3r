package main

import "fmt"

/*
Uma função com a anotação defer, garante que sera executado uma sentença de código antes de sair do método
A função abaixo quando receber por exemplo o numero '5000' irá entrar, printar a linha 14, após isso irá printar o defer
da linha 11 e então ai sim devolver o valor no return da linha 15
*/
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")

	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
