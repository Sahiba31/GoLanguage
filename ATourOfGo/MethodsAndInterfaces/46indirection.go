package main

import "fmt"

type Vertex struct {
	X, Y float64
}

//Method with pointer reciever
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Function with pointer argument
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//functions with a pointer argument must take a pointer
//while methods with pointer receivers take either a value or a pointer as the receiver
func main() {
	v := Vertex{3, 4}
	v.Scale(2)        //Sending value
	ScaleFunc(&v, 10) //Sending pointer

	p := &Vertex{4, 3}
	p.Scale(3)      //Sending pointer
	ScaleFunc(p, 8) //Sending pointer

	fmt.Println(v, p)
}
