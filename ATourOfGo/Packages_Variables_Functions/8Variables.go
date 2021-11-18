package main
import "fmt"

/*var a, b, c bool
func main() {
	var i int
	fmt.Println(i, a, b, c)
}*/


//Variables Initilizers
/*var i, j int = 1, 2   //you can skip the type if uh are initializing with some value.
func main() {
	var a, b, c = true, false, "Sahiba"
	fmt.Println(i, j, a, b, c)
}*/


//Short Variable Declarations
func main() {
	var i, j int = 1, 2
	k := 3
	a, b, c := true, false, "Sahiba"
	fmt.Println(i, j, k, a, b, c)
}