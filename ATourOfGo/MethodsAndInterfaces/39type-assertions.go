package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	val, ok := i.(string)
	fmt.Println(val, ok)

	value, ok := i.(float64)
	fmt.Println(value, ok)

	f := i.(float64) //Panic
	fmt.Println(f)
}
