package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := nAleatorio(); i > 5 { // definido apenas dentro do if / também pode ser feito no switch
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
