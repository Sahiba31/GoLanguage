package main
import (
	"fmt"
	"math"
)

func main(){
	i := 42
	j := float64(i)
	k := uint(j)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(i, j, k, x, y, f, z)
}