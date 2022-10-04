package main

import "fmt"

// Quando temos uma função com nome de 'init()' o go executa antes de tudo, até mesmo antes da função 'main()'
func init() {
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("Main...")
}
