package main

import "fmt"

/*
* La instrucción var declara una lista de variables, como en las listas de argumentos de función
* el tipo es el último en la declaración.
*
* Una sentencia var, puede estar a nivel de paquete o de función. Ambos se muestran en el ejemplo de abajo
*
* Una sentencia var, puede estar a nivel de paquete o de función. Ambos se muestran en el ejemplo de abajo.
* */
var c, python, java bool

/*
* Una declaración var puede incluir inicializadores, uno por variable.
*
* Si hay un inicializador, se puede omitir el tipo; la variable tomará el tipo inicializador.*/
var i2, j2 int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	c2, python2, java2 := true, false, "no!"
	fmt.Println(i2, j2, c2, python2, java2)

	/*
	* Sólo dentro de una función, la declaración de asignación corta := se puede usar en lugar de una
	* declaración var con tipo implícito.
	*
	* Fuera de una función, cada declaración comienza con una palabra clave (var, func, etc), por lo que
	* la construcción := no está disponible en ese ámbito superior.*/
	k := 3
	c3, python3, java3 := true, false, "!no"
	fmt.Println(k, c3, python3, java3)

	/*
	* Valores cero
	*
	* Las variables declaradas sin un valor inicial explícito reciben su valor cero (zero value) por defecto
	*
	* 0 para tipos numéricos
	* false para tipos booleanos, y
	* "" para cadenas de texto
	* */
	var i4 int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i4, f, b, s)
}
