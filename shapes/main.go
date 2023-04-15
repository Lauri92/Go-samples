package main

import (
	"fmt"
	"math"
)

type shape interface {
	getShapeArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}
type circle struct {
	radius float64
}

func main() {
	t := triangle{base: 5, height: 10}
	s := square{sideLength: 10}
	c := circle{radius: 7}

	printArea(t)
	printArea(s)
	printArea(c)
}

func (s square) getShapeArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getShapeArea() float64 {
	return t.base * t.height / 2
}

func (c circle) getShapeArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func printArea(s shape) {
	fmt.Println(s.getShapeArea())
}
