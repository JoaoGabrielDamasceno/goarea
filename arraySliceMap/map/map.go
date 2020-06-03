//elemento chave valor
package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123] = "João"
	aprovados[225] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%v - (CPF %v)\n", nome, cpf)
	}

	delete(aprovados, 123)
	fmt.Println(aprovados[225])

}
