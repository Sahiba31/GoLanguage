package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "Sahiba"
	c <- "goyal"
	//c <- "MCA"
	fmt.Println(<-c)
	fmt.Println(<-c)
}
