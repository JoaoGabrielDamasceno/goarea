package main

import "fmt"

func fatorial(i int) (int, error) {
	switch {
	case i < 0:
		return -1, fmt.Errorf("Número inválido: %d", i)
	case i == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(i - 1)
		return i * fatorialAnterior, nil
	}
}

func main() {
	resultado, erro := fatorial(-1)
	fmt.Println(resultado, erro)

	resultado2, _ := fatorial(5)
	fmt.Println(resultado2)

}
