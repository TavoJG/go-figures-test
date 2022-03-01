package figuras

import "fmt"

type Figura interface {
	area() float32
	perimetro() float32
}

func Medidas(g Figura) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}
