package main

import "fmt"

/*
* Arrays
*
* El tipo [n]T es un array de n valores de tipo T.
*
* La expresión
*
* var a [10]int
*
* declara una variale a como un array de diez enteros.
*
* El tamaño de un array es parte de su tipo, por lo que los arrays no se pueden redimensionar. Esto parece limitante, pero no
* te preocupes; Go provee una forma conveniente de trabajar con arrays.
* */

func main() {
	var a [2]string
	a[0] = "Hola"
	a[1] = "Mundo"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primos)
}
