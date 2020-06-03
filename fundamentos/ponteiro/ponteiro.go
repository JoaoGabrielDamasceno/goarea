package main

import "fmt"

func main() {
	i := 1

	//criando o ponteiro
	var p *int = nil

	//referenciando o ponteiro para o endereço da variavel i
	p = &i

	//somando o valor que ta porque tem o *
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
