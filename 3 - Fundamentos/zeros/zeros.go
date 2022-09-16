package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
	/*
		Valores numéricos são atribuído por padrão 0 (zero)
		Valores booleano são atribuído `false` por padrão
		Valores String é atribuído como vazias como padrão
		E ponteiro de memórias são atribuídos nulos
	*/
}
