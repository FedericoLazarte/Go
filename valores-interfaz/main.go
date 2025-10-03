package main

import (
	"fmt"
	"math"
)

/*
* Valores de interfaz
*
* Debajo del capo, los valores de interfaz se pueden considerar mucho una tupa de un valor y un tipo de
* contacto:
*
* (valor, tipo)
*
* Un valor de interfaz contiene un valor de un tipo concreto subyacente específico.
*
* Llamar a un método en un valor de interfaz ejecuta el método del mismo nombre en su tipo subyacente.
* */

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hola"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
