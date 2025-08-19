package main 

import "fmt" 

/*
	Interfaces:
	- Definen comportamientos (conjuntos de métodos)
	- Cualquier tipo que implemente los métodos automáticamente satisface la interfaz
	(no se declara explícitamente como en Java o C#).

	A tener en cuenta:
	- Go usa interfaces implícitas -> muy flexible.
	- interface{} es la interfaz vacía (puede contener cualquier cosa)
*/

type Lector interface {
	Leer() string
}

type Persona struct {
	Nombre string
}

func (p Persona) Leer() string {
	return "Hola, soy " + p.Nombre
}

func saludar(l Lector) {
	fmt.Println(l.Leer())
}

func main() {
	p := Persona{"Ana"}
	saludar(p) // "Hola, soy Ana"
}