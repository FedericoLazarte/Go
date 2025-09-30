package main

import "fmt"

/*
* Go solo tiene una estructura de bucle, el bucle for.
*
* El bucle for básico tiene componentes separados por punto y coma:
*
* La declaración inicial: ejecutada antes de la primera iteración.
*
* La expresión de la condición: evaluada antes de cada iteración.
*
* La declaración posterior: ejecutada al final de cada iteración.
*
* La declaración de inicio será a menudo una declaración de variable corta, y las variables
* declaradas allí son visibles solo en el ambito de la declaración for.
*
* El loop se detendrá una vez que la condición booleana evalué a falso.
*
* A diferencia de otros lenguajes como C, Java o JavaScript no hay paréntesis alrededor de los tres
* componentes de la declaración for y las llaves {} son siempre requeridas.
* */

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/*
	* La instrucción inicial y posterior son opcionales.
	* */
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	/*
	* For es el while de Go.
	* En ese punto puedes soltar los puntos y coman: El while de C se escribe for en Go.
	* */
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	/*
	* For infinito
	*
	* Si omite la condición del bucle, se repetirá para siempre, por lo que un bucle infinito se expresa de
	* forma compacta.
	* */
	// for {}
}
