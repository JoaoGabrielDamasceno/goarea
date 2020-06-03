package main

import "fmt"

func imprimeResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com ", nota)
	} else {
		fmt.Println("Reprovado com ", nota)
	}
}
func main() {
	imprimeResultado(7.5)
	imprimeResultado(6)
}
