package main

import (
	"fmt"
	"math"
)

func postfixEval(str string) int {
	var stack []int
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			//fmt.Println((str[i]))
			stack = append(stack, int(str[i]-'0'))
		} else {
			oper1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			oper2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res := 0
			switch str[i] {
			case '+':
				res = oper2 + oper1
			case '-':
				res = oper2 - oper1
			case '*':
				res = oper2 * oper1
			case '/':
				res = oper2 / oper1
			case '^':
				res = int(math.Pow(float64(oper2), float64(oper1)))
			default:
				fmt.Println("Enter Valid Operator")
			}
			stack = append(stack, res)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	var str string
	fmt.Print("Enter any String: ")
	fmt.Scanln(&str)
	fmt.Println("Result is: ", postfixEval(str))
}
