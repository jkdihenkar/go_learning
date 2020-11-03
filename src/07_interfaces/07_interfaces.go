package main

import (
	"fmt"
	"math"
)

/*
	interfaces and clasic shapes example
*/

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{radius: 14}
	r := Rectangle{height: 12, width: 15}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}

/*

[jd@jdpc go_learning]$ go run 07_interfaces/07_interfaces.go
615.7521601035994
180

*/
