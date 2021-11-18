package main
import "fmt"

type Vertex struct {
	X int
	Y int         //X, Y int
}

/*func main() {
	fmt.Println(Vertex{1, 2})
}*/

//Struct Fields
func main() {
	var v Vertex
	fmt.Println(v)
	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Println(v1.X, v1.Y)
}