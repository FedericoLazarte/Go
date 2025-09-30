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

/*
* Inferencia de tipos
*
* Al declarar una variable sin especificar un tipo explícito (ya sea utilizando la sintaxis de expresión := 0 var =),
* el tipo de variable se deduce del valor del lado derecho.
*
* Cuando se escribe el lado derecho de la declaración, la nueva variable es del mismo tipo:
*
* var i int
* j := i // j es int
*
* Pero cuando el lado derecho contiene una constante numérica sin tipo, la nueva riable puede ser int, float64 o
* complex128 dependiendo de la precisión de la constante:
*
* i := 42 // int
* f := 3.142 // float64
* g := 0.867 + 0.5i // complex128
* */

/** constantes
*
* Las constantes se declarando como variables, pero con la palabra clave const.
*
* Las constantes pueden ser caracteres, cadenas de texto, booleanos o valores numéricos.
*
* Las constantes no se pueden declarar usando la sintaxis :=
* */
const Pi = 3.14

/*
* Constantes numéricas
*
* Las costantes numéricas son valores de alta precisión.
*
* Una constante sin tipo toma el tipo que necesita su contexto.
*
*
* */
const (
	// Crear un número enorme desplazando 1 bit a la izquierda 100 lugares.
	// Es decir, el número binario que es 1 seguido de 100 ceros.
	Big = 1 << 100
	// Moverlo a la derecha de nuevo 99 lugares, por lo que terminamos con: 1<<1 o 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

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

	fmt.Println("Constante:", Pi)

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
