package main

import "fmt"

func valoraprovado(numero int) int {
	defer fmt.Println("Fim!")

	if numero > 1000 {
		fmt.Println("Valor alto!")
		return numero
	}

	fmt.Println("Valor Baixo!")
	return numero
}

func main() {
	fmt.Println(valoraprovado(5000))
}
