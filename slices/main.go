package main

import (
	"fmt"
	"strings"
)

/*
* Slices
*
* Un array tiene un tamaño fijo. Un slice, por otro lado, es una vista dinámicamente dimensionada y flexible
* de los elementos de un array. En la práctica, los slices son mucho más comunes que los arrays.
*
* El tipo []T es un slice con elementos de tipo T.
*
* Un slice se forma especificando dos índices, un límite inferior y superior, separados por dos puntos:
*
* a[low : high]
*
* Esto selecciona un rango semiabierto que incluye el primer elemento, pero excluye el último.
*
* La siguiente expresión crea un slice que incluye elementos 1 a 3 de a:
* a[1:4]
* */

func main() {
	primos := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primos[1:4]
	fmt.Println(s)

	/*
	* Los slices son como referencias a arrays
	*
	* Un slice no almacena ningún dato, solo describe una seccion de un array subyacente.
	*
	* Cambiar los elementos de un slice modifica los correspondientes elementos de su array subyacente.
	*
	* Otros slices que comparten el mismo array subyacente verán esos cambios.
	* */
	nombres := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(nombres)

	a := nombres[0:2]
	b := nombres[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(nombres)

	/*
	* Slices literal
	*
	* Un slice literal es como un array literal sin la longitud.
	*
	* Esto es un array literal:
	*
	* [3]bool{true, false, true}
	*
	* Y esto crea el mismo array que arriba y luego construye un slice que lo referencia:
	*
	* []bool{true, true, false}
	* */
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	/*
	* Slice defaults
	*
	* Al utilizar un slice, puedes omitir los límites alto o bajo para usar sus valores predeterminados.
	*
	* El default es cero para el límite inferior y la longitud del slice para el límite superior.
	*
	* Para el array
	*
	* var a[10]int
	*
	* Estas expresiones de slice son equivalentes:
	*
	* a[0:10]
	* a[:10]
	* a[0:]
	* a[:]
	* */
	s3 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s3)

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	/*
	* Longitud y capacidad de un slice
	*
	* Un slice tiene tanto una longitud como una capacidad.
	*
	* La longitud de un slice es el número de elementos que contiene.
	*
	* La capacidad de un slice es el número de elementos en el array subyacente, contando desde el primer
	* elemento en el slice.
	*
	* La longitud y la capacidad de un slice s se pueden obtener usando las expresiones len(s) y cap(s).
	*
	* Puedes extender la longitud de un slice haciendo un re-slicing del mismo, siempre que tenga capacidad suficiente.
	* Trata de cambiar una de las operaciones de slice en el programa para extenderlo mas allá de su capacidad y ver lo que sucede
	* */
	s4 := []int{2, 3, 5, 7, 11, 13}
	imprimirSlice(s4)

	// Recortar el slice para darle longitud cero.
	s4 = s4[:0]
	imprimirSlice(s4)

	// Extender su longitud
	s4 = s4[:4]
	imprimirSlice(s4)

	// Descartar los primeros dos valores.
	s4 = s4[2:]
	imprimirSlice(s4)

	/*
	* Nil slices
	*
	* El valor cero de un slice es nil.
	*
	* Un slice nulo tiene una longitud y capacidad de 0 y no tiene array subyacente.
	* */
	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nil!")
	}

	/*
	* Creando un slice con make
	*
	* Los slices pueden ser creados con la función nativa make; así es como se crean arrays de tamaño dinámico.
	*
	* La función make asigna un array inicializado con sus zero values y retorna un slice que referencia a ese array:
	*
	* a := make([]int, 5) // len(a) = 5
	*
	* Para especificar una capacidad, pasa un tercer argumento a make:
	*
	* b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	*
	* b = b[:cap(b)] // len(b) = 5, cap(b)=5
	* b = b[1:] // len(b)=4, cap(b)=4
	* */
	a2 := make([]int, 5)
	imprimirSlice2("a2", a2)

	b2 := make([]int, 0, 5)
	imprimirSlice2("b2", b2)

	c2 := b2[:2]
	imprimirSlice2("c2", c2)

	d := c2[2:5]
	imprimirSlice2("d", d)

	/*
	* Slices de slices
	*
	* Los slices pueden contener cualquier tipo, incluyendo otros slices
	* */
	// Crear un tablero de tic-tac-toe
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// Los jugdores se turnan
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/*
	* Agregando elementos a un slice
	*
	* Es común agregar nuevos elementos a un slice, por lo que Go provee la función nativa append. ver la documentación.
	*
	* func append(s []T, vs ...T) []T
	*
	* El primer parámetro s de append es un slice de tipo T, y e resto son valores T para agregar a slice.
	*
	* El valor resultante de append es un slice que contiene todos los elementos del slice original más los valores provistos.
	*
	* Si el array subyacente de s es muy pequeño para contener todos los valores provistos, se asignará un array más grande.
	* El slice resultante apuntará al nuevo array asignado.
	* */

	var s6 []int
	imprimirSlice(s6)

	// append funciona en nil slices
	s6 = append(s6, 0)
	imprimirSlice(s6)

	// El slice crece según sea necesario
	s6 = append(s6, 1)
	imprimirSlice(s6)

	// Podemos agregar más de un elemento a la vez
	s6 = append(s6, 2, 3, 4)
	imprimirSlice(s6)
}

func imprimirSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func imprimirSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
