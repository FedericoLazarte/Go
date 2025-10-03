package main

import "fmt"

/*
* Tipo switches
*
* Un tipo switch es un constructor que permite varias aserciones tipo en serie.
*
* Un tipo switch es como una declaración switch regular, pero los casos en el tipo switch
* especifican tipos (no valores), y esos valores son comparados contra otros tipos de valores
* sostenidos por el valor de interfaz dado.
*
* switch v := i.(type) {
* case T:
* acá v es tipo T
* case S:
* acá v es tipo S
* default
* }
*
* La declaracion en el tipo switch tiene la misma sintaxis que el tipo aserción i.(T), pero el tipo
* especifico T es reemlazado por a palabra clave type.
*
* Esta declaración switch comprueba si el valor de la interfaz * contiene un valor del tipo T o S. En cada caso
* T y S, la variable v será del tipo T o S respectivamente y mantienen el valor que tiene i. En el caso predeterminado
* (donde no hay ninguna coincidencia), la variable v es del mismo tipo de interfaz y valor que i.
* */

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Dos veces %v es %v\n", v, v*2)
	case string:
		fmt.Printf("%q tiene %v bytes de largo\n", v, len(v))
	default:
		fmt.Printf("No entiendo nada sobre tipos %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
