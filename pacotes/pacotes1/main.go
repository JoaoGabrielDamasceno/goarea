package main

import "fmt"

func main() {
	p1 := Pontos{2.0, 2.0}
	p2 := Pontos{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
