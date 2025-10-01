package main

import (
	"fmt"
	"math"
)

/*
* Una función puede tener  cero o más argumentos.
*
* La función add toma dos parámetros de tipo int.
* */
func add(x int, y int) int {
	return x + y
}

/*
* Cuando dos o más parámetros de función con nombre consecutivos comparten el mismo tipo,
* se puede omitir el tipo de todos menos del último.
* */
func add2(x, y int) int {
	return x + y
}

/*
* Una función puede devolver cualquier número de resultados.
* */
func swap(x, y string) (string, string) {
	return y, x
}

/*
* Valores de retorno con nombre
*
* Los valores de retorno de Go pueden ser nombrados. Si es así, se tratan como variables definidas
* en la parte superior de la función.
*
* Estos nombres deben utilizarse para documentar el significado de los valores devueltos.
*
* Una instrucción return sin argumentos devuelve los valores de retorno nombrados. Esto se conoce como
* naked return (retorno desnudo).
*
* Las declaraciones de retorno desnudas deben usarse solo en funciones cortas, como en el ejemplo de abajo.
*
* En funciones más largas puedes dañar la legibilidad.
* */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(20, 10))

	a, b := swap("hola", "mundo")
	fmt.Println(a, b)

	fmt.Println(split(17))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(calcular((hypot)))
	fmt.Println(calcular(math.Pow))
}

/*
* Valores de función
*
* Las funciones son valores también. Pueden ser pasadas como argumentos a funciones y retornarlas desde funciones.
*
* Los valores de función pueden ser usados como argumentos de funciones y valores de retorno.
* */
func calcular(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
