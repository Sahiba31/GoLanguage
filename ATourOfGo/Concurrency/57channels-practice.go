package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	c <- sum
}

func main() {
	arr := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
