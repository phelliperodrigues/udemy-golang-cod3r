package main

import (
	"fmt"
	m "math" //o `m` é como um alias para o import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	/* forma reduzida de criar uma var
	o simbulo `:=` declara e inicializa a variavel
	*/
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	//Bloco de contrução de constantes
	const (
		a = 1
		b = 2
	)
	//Bloco de construção de variaveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//Declaração em uma mesma linha varias variaveis do mesmo valor
	var e, f bool = true, false
	fmt.Println(e, f)

	//Declaração em uma mesma linha varias variaveis com tipos diferentes
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)

	//Todas as variaveis precisam ser usadas se não o programa não compila

}
