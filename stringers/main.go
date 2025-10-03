package main

import (
	"fmt"
)

/*
* Stringers
*
* Una de las interfaces más ubicuas es Stringer definida por el paquete fmt.
*
* type Stringer interface {
* 	String() string
* }
*
* un Stringer es un tipo que puede describirse a sí mismo como un string. El paquete fmt (y muchos otros) buscan
* esta interfaz para imprimir valores.
* */

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v años)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
