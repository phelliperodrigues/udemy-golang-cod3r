package main

import "fmt"

/*
Para declarar uma função utilizamos a anotação `func`
Funções podem ter ou não retorno, no exemplo abaixo é uma função
sem retorno.
Funções podem ter 0..N atributos
*/
func f1() {
	fmt.Println("F1")
}

// Função recebendo parâmetros
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// Esta função não recebe nenhum parâmetro, porem retorna uma string
func f3() string {
	return "F3"
}

// Se parâmetros forem do mesmo tipo, vc pode declarar as variáveis
// e no final dizer qual tipo de todas as variáveis
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// Em Go é possível ter 0..N retornos como o exemplo abaixo
func f5() (string, string) {
	hello := "Hello" // Primeiro retorno
	world := "World" // Segundo Retorno
	return hello, world
}
func main() {
	f1()
	f2("Hello", "World")
	fmt.Println(f3())
	fmt.Println(f4("Hello", "World"))
	r5_1, r5_2 := f5()
	fmt.Printf("F5: %s %s", r5_1, r5_2)

}
