package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type rectangle struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func(s square) area() float64{
	return s.side*s.side
}

func(r rectangle)area()float64{
	return r.length*r.width
}
type shape interface{
	area() float64
}
func info(s shape)float64{
	return s.area()
}

func main(){
	c:=circle{
		2.5,
	}
	fmt.Printf("Area of cirlce: %.2f\n",info(c))
	s:=square{
		4,
	}
	fmt.Printf("Area of square: %.2f\n",info(s))
	r:=rectangle{
		2.8,
		4.5,
	}
	fmt.Printf("Area of rectangle: %.2f\n",info(r))
}
