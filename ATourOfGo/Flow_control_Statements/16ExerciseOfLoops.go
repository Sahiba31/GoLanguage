package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	temp := z
	count := 0
	for i := 0 ; true && i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		if temp == z {
			count++
		} else {
			temp = z
			count = 0
		}
		if count == 3 {
			break
		}
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}