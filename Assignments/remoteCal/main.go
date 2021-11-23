package main

import (
	"fmt"

	"github.com/Sahiba31/calculations"
)

func main() {
	var op string
	var first, second int
	fmt.Print("Please Enter the first number: ")
	fmt.Scanln(&first)
	fmt.Print("Please Enter the second number: ")
	fmt.Scanln(&second)
	fmt.Print("Enter an Operator(+, -, *, /): ")
	fmt.Scanln(&op)
	res := 0
	switch op {
	case "+":
		res = calculations.Addition(first, second)
	case "-":
		res = calculations.Subtraction(first, second)
	case "*":
		res = calculations.Multiplication(first, second)
	case "/":
		if second == 0 {
			for second == 0 {
				fmt.Println("Second Number cant be Zero")
				fmt.Print("Enter Second Number Again: ")
				fmt.Scanln(&second)
			}
		}
		res = calculations.Division(first, second)
	default:
		fmt.Println("Enter any Valid Operator")
	}
	fmt.Printf("%d %s %d = %d\n", first, op, second, res)
}
