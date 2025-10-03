package main

import "fmt"

/*
* La interfaz vacía
*
* El tipo de interfaz que especifica cero métodos es conocida como una interfaz vacía:
*
* interface {}
*
* Una interfaz vacía puede contener valores de cualquier tipo (Cada tipo implementa por lo menos cero
* métodos).
*
* Las interfaces vacías son utilizadas opr el código que maneja valores de tipo desconocido.
*
* Por ejemplo, fmt.Print, toma cualquier número de argumentos del tipo interface{}.
* */

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hola"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
