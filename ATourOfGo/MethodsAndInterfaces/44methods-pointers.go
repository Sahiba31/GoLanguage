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

//Methods with pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f //Modifies to 30
	v.Y = v.Y * f //Modifies to 40
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
