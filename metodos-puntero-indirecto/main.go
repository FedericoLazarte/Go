package main

import (
	"fmt"
	"math"
)

/*
* Métodos y puntero indirecto.
*
* Comparando los anteriores dos programas, paquete punteros-funciones, podrías haber notado
* que funciones con un argumento de puntero deben tomar un puntero.
*
* var v Vertex
* ScaleFun(v, 5) // Error de compilación
* ScaleFun(&v, 5) // OK
*
* Mientras que los métodos con receptores de puntero toman o un valor o un puntero como el receptor cuando ellos
* son llamados:
*
* var v Vertex
* v.Scale(5) // OK
* p := &v
* p.Scale(10) // OK
*
* Para la declaración v.Scale(5), a pesar de que v es un valor y no un puntero, el método receptor de puntero es llamado
* automáticamente. Esto es porque, como conveniencia Go interpreta la declaración v.Scale(5) como (&v).Scale(5),
* ya que el método Scale tiene un puntero receptor.
* */

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFun(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
* Lo equivalente ocurre en sentido inverso.
*
* Las funciones que toman un argumento de valor deben tomar un valor de ese tipo específico:
*
* var v Vertex
* fmt.Println(AbsFunc(v)) // ok
* fmt.Println(AbsFunc(&v)) // Error de compilación
*
* Mientas que los métodos con receptores de valor toman un valor o un puntero como el receptor cuando son llamadas:
*
* var v Vertex
* fmt.Println(v.Abs()) // OK
* p := &v
* fmt.Println(p.Abs()) // OK
*
* En este caso, el método llama a p.Abs() es interpretado como (*p).Abs()
* */

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFun(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFun(p, 8)

	fmt.Println(v, p)

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println(AbsFunc(v2))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))
}
