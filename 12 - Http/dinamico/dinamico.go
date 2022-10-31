package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(writer http.ResponseWriter, _ *http.Request) {
	horaAtual := time.Now().Format("02/01/2006 03:04:05")
	_, err := fmt.Fprintf(writer, "<h1> Hora certa: %s</h1>", horaAtual)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hora-certa", horaCerta)
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
