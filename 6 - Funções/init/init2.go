package main

import "fmt"

// Rodando no terminal `go run *.go` ele executa os 2 arquivos init.go e ini2.go mesmo que o init2.go não tenha a função main()
func init() {
	fmt.Println("Inicializando2")
}
