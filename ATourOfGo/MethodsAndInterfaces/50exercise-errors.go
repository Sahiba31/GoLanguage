package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	temp := z
	count := 0
	for i := 0; true && i < 10; i++ {
		z -= (z*z - x) / (2 * z)
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
