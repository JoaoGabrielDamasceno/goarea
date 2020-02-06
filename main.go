package goarea

import "math"

//Pi é valor numérico
const Pi = 3.1416

// Circ calcula area do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func _Triang(base, altura float64) float64 {
	return (base * altura) / 2
}
