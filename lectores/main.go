package main

import (
	"fmt"
	"io"
	"strings"
)

/*
* Lectores
*
* El paquete io específica la interfaz io.Reader, que representa e final de lectura de un flujo de datos.
*
* La librería estándar de Go contiene muchas implementaciones de esta interfaz, incluidos archivos, conexiones
* de red, compresores, cifrados y otros.
*
* La interfaz io.Reader tiene un método Read:
*
* func (T) Read(b []byte) (n int, err error)
*
* Read llena el segmento de bytes dado con datos y devuelve el número de bytes poblado y un valor de error.
* Devuelve un error io.EOF cuando la transmisión termina.
*
* El código de ejemplo crea un strings.Reader y consume su salida de 8 bytes a la vez
* */

func main() {
	r := strings.NewReader("Hola, lector!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
