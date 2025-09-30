package main

import "fmt"

/*
* Punteros
*
* Go tiene punteros. Un puntero contiene la dirección de memoria de un valor.
*
* El tipo *T es un puntero a un valor T. Su valor cero es nil.
*
* var p *int
*
* El operador & genera un puntero a su operando.
*
* i := 42
* p = &i
*
* El operador * denota el valor subyacente del puntero.
*
* fmt.Println(*p) // lee i a través del puntero p
* *p = 21 // establece i a través del puntero p
*
* Este se conoce como "desreferencia" o "indirección".
*
* A diferencia de C, Go no tiene arimética de punteros.
* */
func main() {
	i, j := 42, 2701

	p := &i         // Apunta a la dirección de memoria de i
	fmt.Println(*p) // lee i a través del puntero
	*p = 21         // establece i a través del puntero
	fmt.Println(i)  // ve el nuevo valor de i

	p = &j         // apunta a la dirección de memoria de j
	*p = *p / 37   // divide j a través del puntero
	fmt.Println(j) // vel el nuevo valor de j
}
