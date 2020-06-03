package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 45, 4, 5} //quem conta a quantidade de elementos é o compilador

	for i, numero := range numeros {
		fmt.Printf("%v %v\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Printf("%v\n", num)
	}
}
