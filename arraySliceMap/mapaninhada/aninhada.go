package main

import "fmt"

func main() {
	letras := map[string]map[string]float64{
		"G": {
			"Giulio":  654.25,
			"Giovane": 578.52,
		},
		"J": {
			"João": 588.00,
			"Juca": 847.20,
		},
	}

	for letra, valor := range letras {
		fmt.Println(letra, valor)
	}
}
