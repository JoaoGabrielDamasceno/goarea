package main

import "fmt"

func main() {
	fmt.Print("Linha ")
	fmt.Print("Um.")

	fmt.Println("Pula  no final.")

	x := 3.1254

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	fmt.Printf(" O valor de x é %.2f\n", x)

	a := 1
	b := 1.1657
	c := true
	d := "opa"

	fmt.Printf("\n%d %f %t %s", a, b, c, d)
	fmt.Printf("\n%v %.2v %v %v", a, b, c, d)

}
