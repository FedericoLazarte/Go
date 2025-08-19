package main

import "fmt"

func main() {
	/*
		Arrays:
		- Colección de elementos de longitud fija y del mismo tipo.
		- La longitud forma parte del tipo -> [3]int != [4]int
		- No cambian de tamaño después de declarados.
	*/

	var numeros = [3]int{10, 20, 30}
	fmt.Println(numeros[0]) // 10

	/*
		Slices
		- Más usados que los arrays
		- Son vistas dinámicas sobre arrays -> tamaño flexible.
		- Tienen: longitud (len) y capacidad (cap).

		A tener en cuenta:
		- Cuando se hace append, si se supera la capacidad, Go crea un nuevo array interno.
		- Pueden compartir memoria con el array original (ojo con modificaciones)
	*/

	nombres := []string{"Ana", "Luis"}
	nombres = append(nombres, "Marta")
	fmt.Println(nombres, len(nombres), cap(nombres))

	/*
		Maps:
		- Colecciones clave -> valor
		- Muy útiles para búsquedas rápidas.

		A tener en cuenta:
		- Las claves deben ser un tipo comparable (string, int, etc).
		- Si accedés a una clave inexistente, devuelve el valor cero del tipo (ej: 0 para int)
		- Podés verificar existencia con el segundo valor:
	*/

	edades := map[string]int{
		"Ana": 20,
		"Luis": 30,
	}
	fmt.Println(edades["Ana"]) // 3o
	
	edades["Marta"] = 25 // Agregar
	delete(edades, "Luis") // Eliminar

	edad, ok := edades["Pedro"]
	if !ok {
		fmt.Println("Pedro no está en el mapa")
	} else {
		fmt.Println(edad)
	}

	/*
		Structs:
		- Agrupan datos de diferentes tipos bajo un mismo nombre.
		- Son como "objetos sin métodos" (aunque luego se les puede asociar métodos).

		A tener en cuenta:
		- Se usan muchísimo en Go para modelar entidades.
		- Son por valor (copias), pero podés usar punteros si querés compartir memoria.
	*/
	type Persona struct {
		Nombre string
		Edad int
	}

	p := Persona{"Ana", 20}
	fmt.Println(p.Nombre, p.Edad)

	/*
		Pointers:
		- Guardan direcciones de memoria
		- Permiten modificar datos sin copiarlos

		A tener en cuenta:
		- Go no tiene aritmética de punteros como C.
		- Son esenciales cuando pasás structs grandes a funciones para no copiarlos.
	*/
	x := 10
	puntero := &x // puntero a x
	*puntero = 20 // cambia x
	fmt.Println(puntero)
	fmt.Println(x) // 20
}
