package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	//Interface values with nil underlying values
	i = t
	descibe(i)
	i.M()

	i = &T{"Hello"}
	descibe(i)
	i.M()

	i = F(math.Pi)
	descibe(i)
	i.M()
}

func descibe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
