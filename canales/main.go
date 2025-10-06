package main

import "fmt"

/*
* Canales
*
* Los canales son un conducto tipado a través del cual puede enviar y recibir los valores con el operdor
* de canal <-
*
* ch <- v // Envía v al canal ch.
* v := <-ch // Recibe desde ch, y asigna valor a v.
*
* (El flujo de los datos sigue la dirección de la flecha)
*
* Como los mapas y slices, los canales deben crerse antes de usarse:
*
* ch := maken(chan int)
*
* De manera predeterminada, enviar y recibir se bloquea hasta que el otro lado esté listo.
* Esto permite a las gorutinas sincronizarse sin bloqueos explícitos o variables condicionales.
*
* El código de ejemplo suma los números en un slice, distribuyendo el trabajo entre dos gorrutinas.
* Una vez que ambas gorutinas han completdo sus cálculos, se calcula el resultado final.
*
* */

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // envía sum a c
}

/*
* Range Close
*
* Quien envía puede cerrar un canal para indicar que no se enviarán más valores. Los receptores pueden
* comprobar si un canal ha sido cerrado asignando un segundo parámetro a la expresión de recepción.
*
* v, ok := <-ch
*
* ok es false si no hay más valores para recibir y el canal está cerrado.
*
* El loop for i := range c recibe los valores desde el canal repetidamente hasta que se cierra.
*
* Nota: Solo quien envía debería cerrar un canal, nunca el receptor. Enviar a un canal cerrado causará un panic.
*
* Otra nota: Los canales no son archivos; usualmente no necesitas cerrarlos. Cerrarlos es solo necesario cuando
* el receptor debe ser avisado de que no llegarán mas valores, por ejemplo, para terminar un loop range.
* */

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // recibido desde c
	fmt.Println(x, y, x+y)

	/*
	* Canales con Buffer
	*
	* A los canales se les puede aplicar un buffer. Proporcione la longitud del buffer como segundo argumento
	* para que make inicialice un canal con buffer:
	*
	* ch := make(chan int, 100)
	*
	* Los envíos a un canal con buffer bloquean solo cuando el buffer está lleno. Recibir bloquea cuando el buffer
	* está vacío.
	*
	* Modifica el ejemplo para sobrepasar el buffer y ver que pasa
	* */

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
}
