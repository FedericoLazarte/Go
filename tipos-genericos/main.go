package main

/*
*	Tipos genéricos
*
*	Además de funciones genéricas. Go también soporta tipos genéricos. Un tipo puede parametrizarse con un
*	parámetro de tipo, lo que podría ser útil para implementar estructurs de datos genéricas.
*
*	Este ejemplo se muestra una declaración de tipo simple para una lista de un solo slice enlace que contiene
*	cualquier tipo de valor.
* */

// List representa una lista enlazada individualmente que contiene
// valore de cualquier tipo
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

}
