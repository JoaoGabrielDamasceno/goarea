package main

import (
	"fmt"
	"math" //se colocar um m por exemplo antes da pra usar isso em vez de math no codigo
	// _ pra não usar nada
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	//forma reduzida de criar um var ---> ":="
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circuferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	g, h, i := 1, "epa", true

	fmt.Println(g, h, i)

}
