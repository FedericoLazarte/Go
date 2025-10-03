package main

import "fmt"

/*
*	Valores de interfaz nulo
*
*	Un valor de interfaz nulo no contiene ni un valor ni tipo en concreto.
*
*	Llamar a un método en una interfaz nula es un error de tiempo de ejecución porque
*	no hay ningún tipo dentro de la tupla de la interfaz ara indicar que método concreto llamar.
* */

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("%v, %T", i, i)
}
