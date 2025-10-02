package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
* Métodos
*
* Go no tiene clases. Sin embargo, puedes definir métodos en tipos.
*
* Un método es una función con un argumento receptor especial.
*
* El receptor aparece en su propio argumento entre la palabra clave func y el nombre del método.
*
* En este ejemplo, el método Abs tiene un receptor de tipo Vertex llamado v.
* */
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
* Los métodos son funciones
*
* Recuerda: un metodo es solo una función con un argumento receptor.
*
* Aquí Abs escrito como una función regular sin ningún cambio en su funcionalidad
* */
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
* Métodos continuos
*
* También, puedes declarar un método en tipos que son estructura.
*
* En este ejemplo, vemos un tipo numérico MyFloat con un método Abs.
*
* Solo puedes declarar un método con un receptor cuyo tipo este definido en el mismo paquete como el método. No
* puedes declarar un método con un receptor cuyo tipo este definido en otro paquete (que incluye los tipos integrados
* como int)
*
* */
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
* Receptores de puntero
*
* Puedes declarar métodos con receptores de puntero.
*
* Esto significa que el tipo de receptor tiene la sintaxis literal *T para algún tipo T. (Además T no puede ser un
* puntero con *int)
*
* Por ejemplo, acá el método Scale se define en *Vertex.
*
* Los métodos con receptores de punteros pueden modificar el valor al que el receptor apunta (como Scale lo hace aquí)
* Dado que los métodos a menudo necesitan modificar su receptor, los receptores de punteros son más comunes que los
* receptores de valores.
*
* Con un receptor de valor, el método Scale opera en una copia del valor Vertex original. (Este comportamiento
* es igual para cualquier otro argumento de función). El método Scale debe tener un puntero receptor para cambiar
* el valor de Vertex declarando en la función main.
* */
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
}
