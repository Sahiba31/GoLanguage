package main
import "fmt"

func CheckPalin(num int) bool {
	n := num
	temp := 0
	for n != 0 {   				//for n := num ; n > 0 ; n /= 10 
		last := n % 10
		temp = temp * 10 + last
		n /= 10
	}
	return num == temp
}

func main() {
	fmt.Println("Enter any Number: ")
	var num int
	fmt.Scanln(&num)
	fmt.Println(CheckPalin(num))
}