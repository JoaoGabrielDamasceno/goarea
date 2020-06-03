package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// sem sinal (só positivoa) ---> uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	//com sinal ---> int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)

	//float ---> float16, float32
	var x float32 = 32.12
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	//boolan
	bo := true
	fmt.Println("Bo é", bo)

	//string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamnho é ", len(s1))

	s2 := `Olá
	Meu
	Nome`

	fmt.Print(s2)
}
