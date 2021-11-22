package main
import "fmt"

func isMatchingPair(first, second byte) bool {
	if first == '(' && second == ')' {
		return true
	} else if first == '[' && second == ']' {
		return true
	} else if first == '{' && second == '}' {
		return true
	}
	return false
}

//This function checks that the string is balanced or not
func isBalanced(str string) bool {
	var stack []byte
	for i := 0 ; i < len(str) ; i++ {
		if str[i] == '(' || str[i] == '{' || str[i] == '[' {
			stack = append(stack, str[i])
			continue
		} else {
			if len(stack) == 0 {
				//fmt.Println("Not Balanced")
				return false
			} else if isMatchingPair(stack[len(stack) - 1], str[i]) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 { 
		return false
	}
	return true
}


func main() {
	var str string
	fmt.Println("Enter any string of parenthesis to check: ")
	fmt.Scanln(&str)
	fmt.Println(isBalanced(str))
}