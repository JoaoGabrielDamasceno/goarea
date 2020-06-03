package main

import "fmt"

func main() {
	x := 1
	y := 2

	//go só tem a forma posfixada
	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	// Não pode colocar esses operadores em uma comparação por exemplo
}
