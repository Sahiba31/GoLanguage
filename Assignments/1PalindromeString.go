package main
import "fmt"


//O(N) Time and O(1) Space
func CheckPalin(s string) bool {
	low := 0
	high := len(s) - 1
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

//O(N) Time and O(N) Space
/*func CheckPalin(s string) bool {
	res := []byte {}
	for i := len(s) - 1 ; i >= 0 ; i-- {
		res = append(res, s[i])    //append operation is constant time operation
	}
	return s == string(res)
}*/

func main() {
	//s := "abcba"
	fmt.Println("Enter Any String:")
	var s string
	fmt.Scanln(&s)
	fmt.Println(CheckPalin(s))
}