package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
* Aquí vemos los métodos Abs y Scae reescritos como funciones.
*
* Nuevamenente, prueba eliminando el signo * de la función Scale.
* */
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
