package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)

	//varios slices apontando para um mesmo array
	//slice não é um array! é o pedaço de um array

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)
}
