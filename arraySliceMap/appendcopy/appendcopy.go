package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3}
	var slice1 []int

	//array não funciona com append
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(array1, slice1)

	slice2 := []int{5, 4, 3} //cria de tres posições então por isso só copia 3
	// slice2 := make ([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2, slice1)
}
