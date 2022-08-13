package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
}

type circle struct {
	r float64
}

type rectangular struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r rectangular) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func totalArea(shapes ...shapes) float64 {
	area := 0.0
	for _, shape := range shapes {
		area = area + shape.area()
	}

	return area
}

func main() {
	c := circle{4.5}
	r := rectangular{5, 6, 6, 7}
	fmt.Println("Total area: ", totalArea(c, r))
}
