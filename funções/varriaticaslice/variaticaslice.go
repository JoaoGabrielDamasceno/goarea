package main

import "fmt"

func imprimirAprovados(array ...string) {
	fmt.Println("Aprovados")
	for i, valor := range array {
		fmt.Printf("%v) %v\n", i, valor)
	}
}

func main() {
	aprovados := []string{"João", "Maria", "Ricardo"}
	imprimirAprovados(aprovados...)
}
