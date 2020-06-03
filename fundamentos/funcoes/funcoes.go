package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}

/*
Se criar um outro programa na pasta que só tem a função main
consegue usar essas funções mas precisa ir no terminal -> pasta -> fo run *.go
*/
