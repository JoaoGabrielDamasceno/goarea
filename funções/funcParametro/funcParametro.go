package main

import "fmt"

func mult(p1, p2 int) int {
	return p1 * p2
}

func exec(funcao func(int, int) int, x, y int) int {
	return mult(x, y)
}

func main() {
	fmt.Println(exec(mult, 3, 2))
}
