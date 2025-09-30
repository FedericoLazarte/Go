package main

import (
	"fmt"
	"math/cmplx"
)

/*
* Los tipos básicos en Go son:
*
* bool
*
* string
*
* int int8 int16 int32 int64
* uint uint8 uint16 uint32 uint64 uintptr
*
* byte // alias para uint8
*
* rune // aias par int32, representa un punto de código Unicode
*
* float32 float64
*
* complex64 complex128
*
* Las declaraciones de variables pueden "factorizarse" en bloques, como las declaraciones de importación.
*
* Los tipos int, uint y uintptr suelen tener 32 bits y 64 bits de longid en sistemas de 32 y 64 bits.
*
* Cuando se necesita un valor entero, debe usarse int a menos que tenga una razón específica para usar un tipo
* entero de tamaño o sin signo.
* */

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("TIpo: %T Valor: %v\n", z, z)
}
