package main

import (
	"fmt"
	"math"
)

/*
* La expresión T(v) convierte el valor v en T.
*
* Algunas conversiones numéricas
*
* A diferencia de C, en Go la asignación entre elementos de diferentes tipos requiere una conversión
* explicita.
* */

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
