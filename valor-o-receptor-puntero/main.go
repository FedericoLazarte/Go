package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
* Eligiendo un valor o receptor de puntero.
*
* Hay dos razones para usar un receptor de puntero.
*
* La segunda es para evitar copiar el valor en cada método llamado. Esto puede ser más eficiente si el receptor
* es un struct largo, por ejemplo.
*
* En este ejempo Scale y Abs tienen el tipo de receptor *Vertex, aunque el método Abs no necesita modificar
* su receptor.
*
* En general, todos los métodos de un tipo determinado deben tener un valor o un puntero receptor, pero no una mezcla
* de ambos. (ver Interface)
* */

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Antes de la función Scale: %+v, Abs: %v\n ", v, v.Abs())
	v.Scale(5)
	fmt.Printf("Después de la función Scale :%+v, Abs: %v\n", v, v.Abs())
}
