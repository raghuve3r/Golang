package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Printf("%0.3f\n", z.area())
}
func main() {
	s := Square{4}
	c := Circle{3}
	info(s)
	info(c)
}
