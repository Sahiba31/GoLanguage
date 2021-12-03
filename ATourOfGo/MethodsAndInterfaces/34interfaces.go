package main

import (
	"fmt"
	"math"
)

//Creating an interface
type MyInterface interface {
	Abs() float64
}

//Creating our custom type
type MyFloat float64

//Creating structure
type Vertex struct {
	X, Y float64
}

//Implementing method of MyInterface interface
//Method with reciever value of MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Method with pointer reciver of structure
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//Creating variable of interface
	var a MyInterface
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements MyInterface
	a = &v // a *Vertex implements MyInterface
	//a = v     v is a Vertex (not *Vertex) and does NOT implement MyInterface.
	fmt.Println(a.Abs())
}
