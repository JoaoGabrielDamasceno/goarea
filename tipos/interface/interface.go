package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%v - R$ %v", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func main() {
	var coisa imprimivel = pessoa{"ROBERTO", "SILVA"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça", 179.99}
	imprimir(p2)
}
