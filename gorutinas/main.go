package main

import (
	"fmt"
	"time"
)

/*
* Gorutinas
*
* Una gorutina es un hilo ligero manejado por el runtime de Go.
*
* go f(x, y, z)
*
* Inicia una nueva gorutina
*
* f(x, y, z)
*
* La evaluación de f, x, y, z ocurre en la gorutina en curso y la ejecución de f ocurre en la nueva
* gorutina.
*
* Las gorutinas se ejecutan en el mismo espacio de direcciones, así que el acceso a memoria compartida
* debe estar sincronizada. El paquete sync provee algunas primitivas útiles para ello, aunque no las necesitará
* mucho en Go ya que hay otras primitivas. (ver canales)
* */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("mundo")
	say("hola")
}
