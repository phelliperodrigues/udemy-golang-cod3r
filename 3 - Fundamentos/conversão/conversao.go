package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// Para realizar uma operação como tipos de variáveis diferentes é preciso fazer a conversão
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// A conversão `string()` não converte o valor do inteiro em string e sim ele pega o valor da tabela ASC
	fmt.Println("Teste " + string(97))

	// int para string
	//srtconv = String Conversão
	//Itoa = Int to String
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	//primeiro valor `num`, o segundo valor é um erro, quando usamos `_` estamos ignorando o erro
	num, _ := strconv.Atoi("123")
	fmt.Println("num é do tipo", reflect.TypeOf((num)))

	// boolean pode ser convertido com a string "true" para verdadeiro ou "1", qualquer outro valor será false
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
