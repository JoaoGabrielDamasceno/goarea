package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID   int      `json:"id"` //maiusculo nome da variavel publico
	Nome string   `json:"nome"`
	Tags []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", []string{"Promoção", "Eltetronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id": 5, "nome": "Caneta", "tags": ["Papelaria", "Escrever"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
