package main

import "fmt"

func fatorial(i uint) uint {
	switch {
	case i == 0:
		return 1
	default:
		fatorialAnterior := fatorial(i - 1)
		return i * fatorialAnterior
	}
}

func main() {

	resultado2 := fatorial(5)
	fmt.Println(resultado2)

}
