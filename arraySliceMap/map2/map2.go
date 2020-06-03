package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João Silva":  6.25,
		"Ana Cardoso": 7.25,
	}

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
