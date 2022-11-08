package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// interfaces are implemented implicitly
// we do not have to explicitly mention that it does so
// example `something implements interface``
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a)
}

func main() {
	var a Abser
	f := MyFloat(float64(-25))
	v := Vertex{3, 4}

	a = f
	describe(a)
	fmt.Println(a.Abs())

	a = &v
	describe(a)
	fmt.Println(a.Abs())
}
