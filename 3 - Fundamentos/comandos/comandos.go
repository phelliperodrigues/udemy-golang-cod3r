package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!\n", "Go")
}

/*
`go` => Mostra varias possibilidades do uso do go
`go help get` => O `help` mostra a documentação do comando `get` por exemplo
`go version` => Exibe a versão do go instalado
`~/go/bin/godoc -http=:6060` => Habilita a documentação offline. Acessar http://localhost:6060
	- é preciso instalar o lib
		`go install golang.org/x/tools/cmd/godoc@latest`
`go evn` => Exibe varias informaçoes de variaveis de ambiente
`go vet "arquivo.go"` => é uma forma simplificada de identificar codigo "morto", basicamente verifica se tem algo de errado com seu codigo
`go build "arquivo.go` => Constroi um executavel
`go run "arquivo.go` => Compila e executa o programa
*/
