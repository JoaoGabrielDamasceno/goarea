package main

import (
	"fmt"
	"time"
)

//Channel (canal) -  é a forma de comunicação entre goroutines
//channel é um tipo

func multDoisTresQuatro(n int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * n

	time.Sleep(time.Second)
	c <- 3 * n

	time.Sleep(5 * time.Second)
	c <- 3 * n

}

func main() {
	c := make(chan int)

	go multDoisTresQuatro(5, c)
	a, b, tres := <-c, <-c, <-c

	fmt.Println(a, b, tres)
}
