package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("Soma", a+b)
	fmt.Println("Subtração", a-b)
	fmt.Println("Multiplicação", a*b)
	fmt.Println("Divisão", a/b)
	fmt.Println("Módulo", a%b)

	//bit a bit
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Potencia =>", math.Pow(c, d))

}
