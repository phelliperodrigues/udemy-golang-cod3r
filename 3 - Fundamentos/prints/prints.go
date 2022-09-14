package main

import "fmt"

func main() {

	//A função Print() imprime o valor sem quebrar a linha no final
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//A função Println() imprime o valor e depois quebra uma linha no final
	fmt.Println(" Nova")
	fmt.Println(" linha.")

	x := 3.141516
	// O print abaixo não executa, em Go não é possível concatenar uma variável float64 com uma string
	//fmt.Println("O valor de x é " + x)

	//A função Sprint converte e retorna o valor da variável em uma string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	//Para concatenar uma variável que não seja string na função Println() deve-se usar uma vírgula ","
	fmt.Println("O valor de x é ", x)

	/*A função Printf() vc pode usar uma formatação
	No exemplo abaixo conseguimos formatar o valor de `x` com 2 casas decimais
	O `%f` vc indica que o valor a imprimir é um float, se mudar para `%s` imprimiria errado ex:`O valor de x é %!s(float64=3.141516).`
	*/
	fmt.Printf("O valor de x é %.2f.", x)

	/*
		`\n` quebra uma linha (para não sair na mesma linha do print anterior)
		`%d` = tipo inteiro
		`%f` = tipo float
		`%t` = tipo boolean
		`%s` = tipo string
	*/
	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	//Também podemos utilizar o %v para fazer o binding de vários tipos de variáveis
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
