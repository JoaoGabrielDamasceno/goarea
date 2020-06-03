package main

import "fmt"

type item struct {
	id    int
	qtd   int
	preco float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, valor := range p.itens {
		total += valor.preco * float64(valor.qtd)
	}
	return total
}

func main() {
	var pedido1 pedido
	pedido1 =
		pedido{
			userID: 10,
			itens: []item{
				item{1, 2, 12.10},
				item{2, 1, 23},
				item{11, 100, 3.13},
			},
		}

	fmt.Println(pedido1.valorTotal())
}
