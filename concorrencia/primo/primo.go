package main

import "fmt"

func isPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

}
