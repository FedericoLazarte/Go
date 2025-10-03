package main

import "fmt"

/*
* Interfaces implícitas
*
* Un tipo implementa una interfaz implementando sus métodos. No hay una declaración explícita de intenciones,
* ni la palabra clave "implements".
*
* Las interfaces implícitas desacoplan la definición de una interfaz de su implementación, que luego podría aparecer
* en cualquier paquete sin arreglo previo.
* */

type I interface {
	M()
}

type T struct {
	S string
}

// Este método significa que el tipo T implementa la interface I,
// Pero no necesitamos declarar explícitamente lo que hace.

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hola"}
	i.M()
}
