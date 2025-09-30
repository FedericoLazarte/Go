package main

import "fmt"

/*
* struct
*
* Un struct es una colección de campos.
* */
type Vertex struct {
	X int
	Y int
}

/*
* Struct literal
*
* Un struct literal denota el valor de struct recién creado listando los valores de su campo.
*
* Puedes listar solo un subconjunto de campos usando la sintaxis Nombre: .(Y el orden de los campos nombrados es
* irrelevante)
*
* El prefijo especial & devuelve un puntero al valor de struct.
* */
var (
	v1 = Vertex{1, 2} // tiene tipo Vertex
	v2 = Vertex{X: 1} // Y:0 es implícito
	v3 = Vertex{}
	p2 = &Vertex{1, 2} // Tiene tipo *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	/*
	* Campos de un struct
	*
	* Los campos de un struct se acceden usando un punto
	* */
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	/*
	* Punteros a structs
	*
	* Los campos de struct se pueden acceder a través de un puntero a un struct.
	*
	* Para acceder al campo X de un struct cuando tenemos el puntero de struct p podríamos escribir (*p).X. Sin embargo,
	* esa notación es engorrosa, por lo que el lenguaje nos permite en su lugar escribir solo .X, sin la indirección
	* explícita.
	* */
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p2, v2, v3)
}
