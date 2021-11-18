package main
import "fmt"

func add(x int , y int) int {
	return x + y;
}

func sub(x, y int) int {
	return x - y;
}

func main(){
	fmt.Println(add(20, 40))
	fmt.Println(sub(20, 40))
}