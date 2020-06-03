package main

import "fmt"

type luxuoso interface {
	baliza()
}

type esportivo interface {
	turbo()
}

type esportivoLuxuoso interface {
	luxuoso
	esportivo
}

type bmw struct{}

func (b bmw) baliza() {
	fmt.Println("Baliza...")
}

func (b bmw) turbo() {
	fmt.Println("Turbo...")
}

func main() {
	var carro esportivoLuxuoso = bmw{}
	carro.baliza()
	carro.turbo()
}
