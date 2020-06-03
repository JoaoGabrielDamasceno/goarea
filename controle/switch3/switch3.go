package main

import "fmt"

//inferir qual case ele vai entrar pelo tipo do parametro que foi passado para o metodo

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"

	}
}

func main() {
	fmt.Println(tipo(3.15))
}
