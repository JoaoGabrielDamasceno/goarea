package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 2.126
	b := 2
	fmt.Println(a / float64(b))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste" + string(97))

	//o correto de int para string
	fmt.Println("Teste" + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	x, _ := strconv.ParseBool("true")
	if x {
		fmt.Println("Verdadeiro")
	}
}
