package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go count("Sahiba", c)

	//One Method
	/*for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}*/

	//Second Method
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
		//time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
