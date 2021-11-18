package main
import "fmt"

//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
//Stacking defers
func main() {
	fmt.Println("Counting")
	for i := 0 ; i < 10 ; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}