package main

import "fmt"

func main() {
	s := make([]int, 10, 20)
	s2 := append(s, 10, 20)

	fmt.Println(s, s2)

	s[0] = 7
	fmt.Println(s, s2)
}
