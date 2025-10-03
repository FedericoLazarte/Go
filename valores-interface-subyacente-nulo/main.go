package main

import "fmt"

/*
* Valores de interfaz con valores subyacentes nulo
*
* Si valor en concreto dentro de la propia interfaz es nulo, el método será llamado con un receptor nulo.
*
* En otros lenguajes esto activaría una excepción de puntero nulo, pero en Go es común escribir métodos
* que manejen con gracia ser llamados con un receptor nulo (como el método M en este ejemplo).
*
* Tenga en cuenta que un valor de interfaz que contiene un valor concreto nulo no es considerado nulo en sí mismo.
* */

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hola"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
