package main

import "fmt"

// O Desafio é refatorar o exercício de if-elseif para utilizar um switch sem valor como implementado no exercício switch2
func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(8.1))
	fmt.Println(notaParaConceito(5.8))
	fmt.Println(notaParaConceito(4.9))
	fmt.Println(notaParaConceito(0))
}
