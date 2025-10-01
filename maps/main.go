package main

import "fmt"

type Vertex struct {
	Lag, Long float64
}

/*
* Un map asocia claves a valores.
*
* El valor cero de un map es nil. Un map nulo no tiene claves, ni se pueden agregar claves.
*
* La función make retorna un map del tipo dado, inicializado y listo para usar.
* */
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["bell labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["bell labs"])

	/*
	*	map literales
	*
	*	los maps literales son como los struct literales, pero las claves son requeridas.
	* */
	m2 := map[string]Vertex{
		"bell labs": {
			40.68433, -74.39967,
		},
		"google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	/*
	*	map literales continuos
	*
	*	si el tipo de nivel superior es solo un nombre de tipo, puedes omitirlo de los elementos del literal.
	* */
	m3 := map[string]Vertex{
		"bell labs": {40.1232, -39.1232},
		"google":    {12.3134, 9.12321},
	}
	fmt.Println(m3)

	/*
	*	Mutando Maps
	*
	*	Insertar o actualizar un elemento en el map:
	*	m[key] = elem
	*
	*	Obtener un elemento
	*	elem = m[key]
	*
	*	Eliminar un elemento
	*	delete(m, key)
	*
	*	Prueba que una cave esté presente con una signación de dos valores
	*	elem, ok = m[key]
	*
	*	Si key está e m, entonces ok es true
	*
	*	Si key no está en m, entonces elem es el valor cero para el tipo de elemento del map.
	* */
	m4 := make(map[string]int)

	m4["respuesta"] = 42
	fmt.Println("el valor:", m4["respuesta"])

	m4["respuesta"] = 48
	fmt.Println("el valor:", m4["respuesta"])

	delete(m4, "respuesta")
	fmt.Println("el valor:", m4["respuesta"])

	v, ok := m4["respuesta"]
	fmt.Println("El valor:", v, "¿Presente?", ok)
}
