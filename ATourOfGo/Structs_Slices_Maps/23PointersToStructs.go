package main
import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	v.X = 4
	p.X = 1e9     //no need to use *p.X
	fmt.Println(v)
}