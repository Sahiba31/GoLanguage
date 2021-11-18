package main
import (
	"fmt"
	"time"
)


//Switch Evaluation Order
/*func main() {
	fmt.Println("When's Saturday!!!")
	today := time.Now().Weekday()
	fmt.Println("Today is:",today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far")
	}
}*/

//Switch with no condition
func main() {
	today := time.Now()
	switch {
	case today.Hour() < 12:
		fmt.Println("Good Morning")
	case today.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}