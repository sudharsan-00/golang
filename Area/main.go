package main

import (
    "fmt"
	"math"
)

type square struct {
	lenght float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	x := math.Pi * math.Pow(c.radius, 2)
	return x
}

func (s square) area() float64 {
	x := s.lenght * s.width
	return x
}

type shape interface{
    area() float64
}

func info(s shape) float64{
    return s.area()
}

func main() {
	p1 := circle{
		radius: 4,
	}

	p2 := square{
		lenght: 5,
		width:  8,
	}

	fmt.Println(info(p1))
	fmt.Println(info(p2))
}