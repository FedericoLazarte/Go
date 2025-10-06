package main

import "fmt"

/*
* Parámetros de tipo
*
* Las funciones en GO pueden ser escritas para trabajar con múltiples tipos usando parámetros de tipo. Los parámetros
* de tipo de una función aparecen entre paréntesis, antes de los de la función.
*
* func Index[T comparable](s []T, x T)
*
* Esta declaración significa que s es un slice de cualquier tipo T que cumpla la restricción incorporada
* comparable. x es también un valor del mismo tipo.
*
* comparable es una restricción útil que permite utilizar operadores == y != sobre valores del tipo.
* En este ejemplo, la utilizamos para comparar un valor con todos los elementos del slice hasta encontrar
* una coincidencia. Esta función Index funciona para cualquier tipo que admita la operación.
* */

// Index devuelve el índice de x en s, o -1 si no se encuentra.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v y x son del tipo T, que tiene la restricción
		// comparable, por lo que podemos usar == aquí.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index funciona en un slice de int
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index funciona también en un sliec de enteros
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hola"))
}
