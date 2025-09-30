package main

/*
*	Cada programa de Go se compone de paquetes.
*
*	Los programas en Go comienzan a ejecutarse en el paquete main que siempre debe
*	existir como punto de entrada para el flujo del programa.
*
*	El programa de abajo, está usando los paquetes con rutas de importación "fmt" y "math/rand".
*
*	Por convención, el nombre del paquete es el mismo que el último elemento de la ruta de importación. Por ejemplo:
*	el paquete "math/rand" comprende archivos que comienzan con la declaración package rand. En cambio "math" viene siendo parte 
*	de la ruta donde se encuentra el paquete y a su vez también puede ser un paquete que comprende archivos que comienzan con la 
*	declaración package math.
* */

/*
* Esta forma de escribir los import es buena prácica. Se puede esecribir una por una, pero no sería lo mejor.
* */
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Mi número favorito es", rand.Intn(10))

	/*
	* En go, un nombre se exporta si comienza con una letra mayúscula. Por ejemplo, Pizza es un nombre exportado, al igual que Pi
	* que se exporta desde el paquete math.
	*
	* Al importar un paquete, solo se puede hacer referencia a sus nombres exportados. Cualquier nombre "no exportado" no es accesible
	* desde fuera del paquete.
	* */
	fmt.Println(math.Pi)
}
