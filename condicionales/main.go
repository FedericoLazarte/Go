package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*
* La declaración if de Go son como si fueran bucles for; la expresión no necesita estar
* rodeada de paréntesis () pero las llaves {} son requeridas.
* */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
* If con una breve declaración
*
* Al igual que for, la declaración if puede comenzar con una breve declaración para ejecutar antes de la
* condición.
*
* Las variables declaradas por la instrucción solo están en el alcance hasta el final del if
*
* intenta usar v en la última declaración return
*
* if y else
*
* Las variables declaradas dentro de una declaración corta if también están disponibles dentro de cualquiera de
* los bloques else.
*
* (Ambas llamadas a pow devuelven sus resultados antes de la llamada a fmt.Println en el comienzo del main)
* */
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	/*
	* Una sentencia switch es una forma corta de escribir una secuencia de sentencias if - else. Ejecuta el
	* primer caso cuyo valor es igual a la expresión de condición.
	*
	* El switch de Go es como el de C, C++, Java, JavaScript y PHP, excepto que Go sólo ejecuta el caso seleccionado,
	* no todos los casos que le siguen. En efecto, la sentencia break que se necesita al final de cada caso en esos
	* lenguajes se proporciona automáticamente en Go. Otra diferencia importante es que los casos switch de Go no
	* necesitan ser constantes, y los valores involucrados no necesitan ser enteros.
	* */
	fmt.Print("Go se ejecuta en  ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	/*
	* Switch sin condición
	*
	* Switch sin una condición es lo mismo que switch true.
	*
	* Esta construcción puede ser una forma limpia de escribir largas cadenas if-then-else
	* */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buenos días!")
	case t.Hour() < 17:
		fmt.Println("BUenas tardes.")
	default:
		fmt.Println("Buenas noches.")

	}
}
