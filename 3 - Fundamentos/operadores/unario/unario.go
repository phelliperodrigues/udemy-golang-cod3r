package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas forma postfix
	// Ou seja só existe `var++` e não `++var`
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // u -= 1 ou y = y - 1
	fmt.Println(y)

	// não é possível usar o operador unário em alguma operação como o exemplo abaixo
	// fmt.Println(x == y--)
	// fmt.Println(x++ == y)

}
