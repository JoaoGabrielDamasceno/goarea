package main

import "fmt"

func f1() {
	fmt.Println("--F1--")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2:  %v %v\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) //retorna uma string formatada
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Bom", "dia")

	r3, r4 := f3(), f4("Boa", "Tarde")
	fmt.Println(r3)
	fmt.Println(r4)

	r5, _ := f5()
	fmt.Println(r5)
}
