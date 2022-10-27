package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	if total < 0 {
		fmt.Println("Sem cobertura de teste")
	}
	media := total / float64(len(numeros))
	mediaArrendondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendondada
}

/*
Roda teste e exibe cobertura
go test -cover

Roda teste e gera relatorio
go test cover -func=resultado.out

Exibe resultado relatorio no terminal
go test -coverprofile=resultado.out

Exibe resultado relatorio em uma pagina html
go tool cover -html=resultado.out
*/
