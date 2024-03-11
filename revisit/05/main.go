package main

import (
	"fmt"
	"math"
)

func main() {
	shapes := []shape{
		circle{2.2},
		square{1.5},
		rectangle{5, 6},
		triangle{5, 7},
	}
	for _, v := range shapes {
		v.area()
	}
}

type circle struct {
	radius float64
}

type square struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

type triangle struct {
	base   float64
	height float64
}

func (c circle) area() {
	fmt.Printf("%.2f\n", math.Pi*c.radius*c.radius)
}

func (s square) area() {
	fmt.Printf("%.2f\n", s.radius*s.radius)
}

func (r rectangle) area() {
	fmt.Printf("%.2f\n", r.length*r.width)
}

func (t triangle) area() {
	fmt.Printf("%.2f\n", 0.5*(t.base*t.height))
}

type shape interface {
	area()
}
