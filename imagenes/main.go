package main

import (
	"fmt"
	"image"
)

/*
* El paquete image define la interfaz Image:
*
* pacage image
*
* type Image interface {
* 	ColorModel() color.Model
* 	Bounds() Rectangle
* 	At(x, y int) color.Color
* }
*
* Nota: el valor devuelto Rectángulo del método BOunds es en realidad in image.Rectangle, ya que
* la declaración está dentro del paquete image.
*
* Los tipos color.Color y color.Model también son interfaces, pero lo ignoraremos usando las
* implementaciones predefinidas color.RGBA y color.RGBAModel. Estas interfaces y tipos están
* especificados por image/paquete color
* */

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
