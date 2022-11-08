package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(n float64) {
	v.X = v.X * n
	v.Y = v.Y * n
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Printf("absolute v = %f", v.Abs())
}
