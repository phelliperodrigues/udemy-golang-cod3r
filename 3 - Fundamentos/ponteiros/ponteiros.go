package main

import "fmt"

func main() {
	// Por inferência a variável `i` é um inteiro de 64 bits out seja 8 bytes
	i := 1

	// Para criar um ponteiro basta colocar do lado do tipo o `*`
	var p *int = nil
	// nil é nulo, não aponta para lugar nenhum
	// Alocamos um espaço de memória do tipo inteiro, mas esta vazia
	p = &i
	// o `&` significa "Pegue o endereços da variável `i` e guarde na variável `p` o mesmo local que a variável `i` esta apontando"
	*p++
	//quando usamos o `*` antes do ponteiro estamos desdiferenciando o ponteiro para pegar o valor
	i++
	fmt.Println(p, *p, i, &i)

	// Go não tem aritméticas de ponteiros
	// Ex: p++
}
