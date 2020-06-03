package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	ferrari1 := ferrari{}
	ferrari1.nome = "F40"
	ferrari1.velocidadeAtual = 250
	ferrari1.turboLigado = true

	fmt.Printf("O turbo esta ligado? %v", ferrari1.turboLigado)
}
