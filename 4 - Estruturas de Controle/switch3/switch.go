package main

import (
	"fmt"
	"time"
)

/*
A gente vai fazer testes em cima de um parâmetro genérico, a gente vai passar um parâmetro de um tipo genérico e
vai tentar inferir quais são os 7 - Sistema de Tipos de variáveis que foram passadas e a partir do tipo do dado que foi passado
se for um inteiro ele vai entrar em um case, e se for string ele vai entrar em outro.
Se for uma função ele vai entrar case, e assim a gente vai conseguir a partir de um switch selecionar os 7 - Sistema de Tipos do dado
que foi passado como parâmetro da função
*/

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}
