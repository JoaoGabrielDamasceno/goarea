package main

import "fmt"

type carro interface {
	ligarTurbo()
}

type ferrari struct {
	nome            string
	velocidadeAtual int
	turbo           bool
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	var f1 ferrari = ferrari{
		nome:            "F40",
		velocidadeAtual: 0,
		turbo:           false,
	}

	f1.ligarTurbo()
	fmt.Println(f1.turbo)

	var f2 carro = &ferrari{"F50", 0, false}
	f2.ligarTurbo()
	fmt.Println(f2)

}
