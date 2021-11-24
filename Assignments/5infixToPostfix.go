package main

import "fmt"

func precedence(oper byte) int {
	if oper == '^' {
		return 3
	} else if oper == '*' || oper == '/' {
		return 2
	} else if oper == '+' || oper == '-' {
		return 1
	} else {
		return -1
	}
}

//this method converts the infix expression into postfix expression...
//Example: Infix = (a+b)*c   Postfix = ab+c*
func InfixToPostfix(str string) string {
	var res string
	var stack []byte

	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack = append(stack, str[i]) //Pushing into stack
			continue
		} else if str[i] >= 'a' && str[i] <= 'z' {
			res += string(str[i])
		} else if str[i] == ')' {
			for stack[len(stack)-1] != '(' {
				res += string(stack[len(stack)-1]) //Appending the top element
				stack = stack[:len(stack)-1]       //Pop the element
			}
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) != 0 && precedence(str[i]) <= precedence(stack[len(stack)-1]) {
				res += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, str[i])
		}
	}
	for len(stack) != 0 {
		res += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return res
}

func main() {
	var str string
	fmt.Print("Enter any String: ")
	fmt.Scanln(&str)
	fmt.Println("Postfix Expression is: ", InfixToPostfix(str))
}
