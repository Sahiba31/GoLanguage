package main
import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int {2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[ : 0]
	printSlice(s)

	s = s[ : 4]
	printSlice(s)

	s = s[2 : ]
	printSlice(s)

	//We can also create slices using built-in make function
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
	
	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}