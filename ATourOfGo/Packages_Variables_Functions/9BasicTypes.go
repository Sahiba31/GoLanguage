package main
import "fmt"
var(
	a int = 2
	b bool = false
	c string = "Sahiba"
	d float64 = 4.5
)
func main() {
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
}