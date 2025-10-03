package main

import "fmt"

/*
* Tipo aserciones
*
* Una aserción de tipo proporciona acceso al valor concreto subyacente de un valor de interfaz.
*
* t := i.(T)
*
* Esta declaración firma que el valor de la interfaz i sostiene el tipo concreto T y asigna
* el valor subyacente T a la variable t.
*
* Si i no contiene a T, la declaración provocará panic.
*
* Para probar si un valor de interfaz contiene un tipo especifico, una aserción de tipo puede devolver dos
* valores: el valor subyacente y un valor booleano que informa si la afirmación tuvo éxito.
*
* t, ok := i.(T)
*
* Si i contiene a T, entonces t será el valor subyacente y ok será verdadero.
*
* Si no, ok será falso y t será el valor cero del tipo T, y ningún panic ocurrira.
*
* Tenga en cuenta la similitud entre esta sintaxis y la de leer un map.
* */

func main() {
	var i interface{} = "hola"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
