package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			r, err := regexp.Compile("<title>(.*?)</title>")
			if err != nil {
				fmt.Println(err)
			}
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://phelliperodrigues.dev", "https://patriciasoares.com.br")
	t2 := titulo("https://www.mercadolivre.com.br", "https://www.mercadopago.com.br")

	fmt.Println("Primeiro: ", <-t1, "<|>", <-t2)
	fmt.Println("Segundo: ", <-t1, "<|>", <-t2)
}
