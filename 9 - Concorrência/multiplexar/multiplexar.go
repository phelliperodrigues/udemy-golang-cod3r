package main

import (
	"fmt"
	html "github.com/phelliperodrigues/udemy_golang_cod3r_html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// mutiplexar - misturar (mensagens) num unico canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.mercadolivre.com.br", "https://www.mercadopago.com.br"),
		html.Titulo("https://phelliperodrigues.dev", "https://patriciasoares.com.br"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
