package main

import "math"

//Com maiusculo é público tanto dentro e fora do arquivo (Pacote todo)
//Com minusculo é privado só "funciona" dentro do arquivo

// Pontos representa um ponto do plano cartesiano
type Pontos struct {
	x float64
	y float64
}

func catetos(a, b Pontos) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia calcula a distancia entre dois pontos
func Distancia(a, b Pontos) (dist float64) {
	cx, cy := catetos(a, b)
	dist = math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
	return
}
