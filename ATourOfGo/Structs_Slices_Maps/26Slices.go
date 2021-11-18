package main
import "fmt"

func main() {
	primes := [6]int {1, 2, 3, 4, 5, 6}

	//s := primes[1: 4]
	var s []int = primes[1:4]
	fmt.Println(s)
	s[1] = 10
	fmt.Println(primes)    //It will also be changed... slices are like references to arrays
}


//Slice Literals - Same as array but without specifying length

// func main() {
// 	a := []int {1, 2, 3, 4}
// 	b := []bool {true, false, true}
// 	fmt.Println(a)
// 	fmt.Println(b)
// }