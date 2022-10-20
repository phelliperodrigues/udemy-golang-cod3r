package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"name"`
	Preco float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	p1 := produto{ID: 1, Nome: "Notebook", Preco: 1890.00, Tags: []string{"Promoção", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2, "names":"Caneta", "price": 0.99, "tags":["Papelaria", "Importado"]}`
	err := json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		return
	}
	fmt.Println(p2.Tags[1])
}
