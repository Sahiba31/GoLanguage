package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("type is int", v)
	case string:
		fmt.Println("type is string", v)
	case float64:
		fmt.Println("type is float", v)
	default:
		fmt.Println("Unknown")
	}

}

func main() {
	do(10)
	do("hello")
	do(true)
	do(95.5)
}
