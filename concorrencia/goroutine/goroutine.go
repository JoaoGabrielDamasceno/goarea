package main

import (
	"fmt"
	"time"
)

func fale(pessoa, fala string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%v-> %v (vez %v)\n", pessoa, fala, qtd)
	}
}

func main() {
	go fale("Maria", "Bom dia!", 10)
	fale("João", "Boa noite!", 5)
}
