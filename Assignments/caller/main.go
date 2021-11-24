package main

import (
	"fmt"

	//"github.com/Sahiba31/called"
	//"github.com/Sahiba31/called/multiplication"
	"github.com/Sahiba31/called/division"
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
	//res = called.Subtraction(first, second)
	//res = multiplication.Multiplication(first, second)
	res = division.Division(first, second)
	fmt.Println(res)
}
