package main

import "fmt"

func main() {
	n := 626
	if n == 626 {
		fmt.Printf("Parte verdadera\n")
	} else {
		fmt.Printf("Parte falsa\n")
	}

	// if with initialization statement
	a := 629
	if b := 626; a == b {
		fmt.Printf("Parte verdadera\n")
	} else {
		fmt.Printf("Parte falsa\n")
	}
}
