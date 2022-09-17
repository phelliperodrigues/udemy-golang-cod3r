package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	// Por padrão o switch case encontrando um retorno o processo sai do switch diferente do Java por exemplo que precisa da palavra break se não fica executando os outros cases abaixo
	// Porém caso precise executar a proxima condição é possível utilizando a palavra fallthrough, no exemplo abaixo, se a nota for 10 ele executará o case 10 e em seguida o case 9
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(7.6))
	fmt.Println(notaParaConceito(5.9))
	fmt.Println(notaParaConceito(3.2))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(-2))

}
