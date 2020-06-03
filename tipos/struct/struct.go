package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		"Lapis",
		1.79,
		0.05,
	}

	fmt.Println(produto1, produto1.precoDesconto())

	produto2 := produto{"Caneta", 1.25, 0.05}
	fmt.Println(produto2)
}
