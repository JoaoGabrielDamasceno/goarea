package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p *pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	var pessoa1 pessoa
	pessoa1 = pessoa{
		nome:      "João",
		sobrenome: "Damasceno",
	}

	fmt.Println(pessoa1.getNomeCompleto())

	completo := "Vinicius Caralho"
	pessoa1.setNomeCompleto(completo)
	fmt.Println(pessoa1.getNomeCompleto())

}
