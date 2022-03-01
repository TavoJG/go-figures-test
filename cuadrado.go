package figuras

type Cuadrado struct {
	Ancho  float32
	Altura float32
}

func (c *Cuadrado) area() float32 {
	return c.Ancho * c.Altura
}

func (c *Cuadrado) perimetro() float32 {
	return 2*c.Ancho + 2*c.Altura
}
