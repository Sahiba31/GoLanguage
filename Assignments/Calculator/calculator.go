package main
import (
	"fmt"
	"sample.com/events/addition"
	"sample.com/events/subtraction"
	"sample.com/events/multiplication"
	"sample.com/events/division"
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
		res = addition.Addition(first, second)
	case "-":
		res = subtraction.Subtraction(first, second)
	case "*":
		res = multiplication.Multiplication(first, second)
	case "/":
		if second == 0 {
			for second == 0 {
				fmt.Println("Second Number cant be Zero")
				fmt.Print("Enter Second Number Again: ")
				fmt.Scanln(&second)
			}
		}
		res = division.Division(first, second)
	default:
		fmt.Println("Enter any Valid Operator")
	}
	fmt.Printf("%d %s %d = %d\n", first, op, second, res)
}