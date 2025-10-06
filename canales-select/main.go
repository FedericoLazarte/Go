package main

import "fmt"
import "time"

/*
* Select
*
* La declaración select permite que una gorutina espere a multiles operaciones de comunicación.
*
* Un select bloquea hasta que uno de sus cases pueda ejecutarse, entonces lo ejecuta. Escoge un case
* al azar si varios de ellos están listos.
* */

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	/*
	* Selección por Default
	*
	* El case default en un select se ejecuta si no hay otro case listo.
	*
	* Utilice un case default para intentar enviar o recibir sin bloquear.
	*
	* select {
	* case i := <- c:
	* // use i
	* default:
	* 	// recibir de c bloquearía
	* }
	* */

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
