package main
import "fmt"

// func main() {
// 	sum := 0
// 	for i := 0 ; i < 10 ; i++ {
// 		sum += i;
// 	}
// 	fmt.Println(sum)
// }


/*func main() {
	sum := 1
	for ; sum < 1000 ; {     //init and post statements are optional
		sum += sum
	}
	fmt.Println(sum)
}*/

//For Loop Like While Loop
func main() {
	sum := 1
	for sum < 1000 {       //you can skip both the semicolons also
		sum += sum
	}
	fmt.Println(sum)
}


//Infinite Loop
func main() {
	for{
		
	}
}