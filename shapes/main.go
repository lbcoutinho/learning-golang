package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	width  float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{height: 10, base: 5}
	printArea(t)

	s := square{width: 5}
	printArea(s)
}
