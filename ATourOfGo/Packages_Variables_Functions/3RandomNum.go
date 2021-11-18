package main
import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random Number is:", rand.Intn(200))
	fmt.Println(rand.Intn(200))
    fmt.Println(rand.Intn(200))
    rand.Seed(time.Now().UTC().UnixNano())
	//Seed is used to generate random numbers everytime
	fmt.Println(rand.Intn(200))
	fmt.Println(rand.Intn(200))
	fmt.Println(rand.Intn(200))
}