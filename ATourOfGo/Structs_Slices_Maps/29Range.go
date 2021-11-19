package main
import "fmt"

// func main() {
// 	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }

func main() {
	pow := make([]int, 10)
	for i := range pow {       //If you only want the index, you can omit the second variable.
		pow[i] = 1 << i
	}
	for _, v := range pow {    //You can skip the index or value by assigning to _.
		fmt.Println(v)
	}
}