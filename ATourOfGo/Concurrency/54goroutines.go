package main

import (
	"fmt"
	"time"
)

func count(thing string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(thing)
	}
}

func main() {

	//First
	//Prints Sahiba 5 times then goyal 5 times (Normal functions)
	/*count("Sahiba")
	count("goyal")*/

	//Second
	//It'll say go and run this function in the background and then continue to the next line immediately
	//and this creates what is called a go routine and that go routine will run concurrently.
	/*go count("Sahiba")
	count("goyal")*/

	//Third
	//Doesnt print anything beacuse main goroutine will terminate
	go count("Sahiba")
	go count("goyal")

	//But if add this statement then it will run for 2 seconds after that again terminate
	//time.Sleep(2 * time.Second)

	//This is a blocking call it'll wait for user input, when we press enter then it will terminate.
	fmt.Scanln()
}
