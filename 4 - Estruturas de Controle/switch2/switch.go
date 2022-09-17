package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//  Quando vc utiliza um switch sem nenhum valor associado ele vai procurar um case que  contenha uma expressão que seja true
	// é a mesma coisa de colocar `switch true {}`
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom Dia")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde")
	default:
		fmt.Println("Boa Noite")
	}
}
