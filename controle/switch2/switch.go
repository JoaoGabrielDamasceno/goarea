package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //procura alguma case que seja true, como se tivesse passando true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Bom tarde!")
	default:
		fmt.Println("Boa noite!")

	}
}
