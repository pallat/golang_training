package main

import (
	"fmt"
	"math"
)

type VertexV2 struct {
	Vertex
	Z int
}

func (v *VertexV2) ZValue() int {
	return v.Z
}

func (v *VertexV2) Scale(f float64) {
	v.X = (v.X * f) * 2
	v.Y = (v.Y * f) * 2
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexV2{Vertex{X: 3, Y: 4}, 10}
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v.ZValue())
}
