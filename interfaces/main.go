package main

import (
	"fmt"
	"math"
)

/*
* Interface
*
* Un interface type se define como un conjunto de firmas de métodos.
*
* Un valor del tipo interfaz puede contener valor que se implemente en sos métodos.
*
* NOTA: Hay un error en el ejemplo de a = v. Vertex (e tipo de valor) no implementa Abser porqe el método Abs esta definido
* solo en *Vertex (el tipo de puntero).
* */

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implementa Abser
	a = &v // a *Vertex implementa Abser

	// En la siguiente línea, v es de tipo Vertex (no *Vertex)
	// Y no implementa Abser
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
