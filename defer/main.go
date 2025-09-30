package main

import "fmt"

/*
* defer
*
* Una instrucción defer, aplaza la ejecución de una función hasta el retorno de la función
* circundante (que la contiene).
*
* Los argumentos de la llamada diferida se evalúan inmediatamente, pero la llamada a la función
* no se ejecuta hasta que la función que la rodea retorna.
* */
func main() {
	defer fmt.Println("mundo")

	fmt.Println("hola")

	/*
	* Apilando defers
	*
	* Las llamadas a funciones diferidas se introducen en una pila. Cuando una función retorna, sus llamadas
	* diferidas se ejecutan en orden de último en entrar, primero en salir.
	*
	* Para aprender mas de la declaración defer leer defer panic y recover en la documentación.
	* */
	fmt.Println("Contando")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("listo")
}
