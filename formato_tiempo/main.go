package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(time.Now().Format("02/01/2006"))

	fmt.Println(time.Now().Format("hoy es 2 del 1 de 2006"))

	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
}
