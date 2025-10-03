package main

import (
	"fmt"
	"time"
)

/*
* Errores
*
* Los programas en Go expresan el estado de error con valores de error.
*
* El tipo error es una interfaz de función incorportada similar a fmt.Stringer:
*
* type error interface {
* 	Error() string
* }
*
* (Al igual que con fmt.Stringer, el paquete fmt busca la interfaz error cuando imprime valores.)
*
* Las funciones a menudo devuelven un valor de "error", y el codigo de llamada debe manejar los errores
* probando si el error es igual a nil.
*
* i, err := strconv.Atoi("42")
* if err != nil {
* 	fmt.Printf("No se pudo ocnvertir e número %v\n", err)
* 	return
* }
* fmt.Println("Integer convertido:", i)
*
* Un error nulo denota éxito, un error no nulo denota falla.
* */

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("A las %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"No funciono",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
