package main

import "fmt"

func main() {
	// for, C-like
	n := 10
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", i)
	}

	// for, while-like
	n2 := 10
	i := 0
	for i < n2 {
		fmt.Printf("%d\n", i)
		i++
	}

	// for-each like
	a := [4]int{626, 625, 624}

	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
}
