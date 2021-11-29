package main

import "fmt"

func findPrevGreater(arr []int, n int) {
	var stack []int
	fmt.Printf("%d ", -1)
	stack = append(stack, arr[0])

	for i := 1; i < n; i++ {
		top := stack[len(stack)-1]
		for len(stack) > 0 && top < arr[i] {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top = stack[len(stack)-1]
			}
		}
		if len(stack) == 0 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", stack[len(stack)-1])
		}
		stack = append(stack, arr[i])
	}
}

func main() {
	var n int
	fmt.Print("Enter Size of Array: ")
	fmt.Scanln(&n)
	//var arr []int
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		//var x int
		//fmt.Scan(&x)
		//arr = append(arr, x)
		fmt.Scan(&arr[i])
	}
	findPrevGreater(arr, n)
}
