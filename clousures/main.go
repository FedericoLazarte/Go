package main

import "fmt"

/*
* Clousures
*
* Las funciones de Go pueden ser clousures. Un clousure es un valor de función que referencia variables
* fuera de su cuerpo. La función puede acceder y asignar a las variables referenciadas; en este sentido la
* función está "ligada" a las variables.
*
* Por ejemplo, la función sumador returna un clousure. Cada clousure está ligado a su propia variable sum.
* */
func sumador() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := sumador(), sumador()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
