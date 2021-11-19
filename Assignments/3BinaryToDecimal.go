package main
import (
	"fmt"
	//"math"
)

func BToD(num int) int{
	dec_val := 0
	//i := 0.0
	base := 1
	for n := num ; n > 0 ; n /= 10 {
		last := n % 10
		//dec_val += last * int(math.Pow(2, i))
		//i++;
		dec_val += last * base
		base *= 2
	}
	return dec_val
}

func main() {
	fmt.Println("Enter a Binary Number: ");
	var num int
	fmt.Scanln(&num)
	fmt.Println("The Decimal Number is: ", BToD(num))
}