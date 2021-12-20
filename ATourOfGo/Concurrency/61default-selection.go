package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 100)
	boom := time.After(time.Millisecond * 500)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println(".")
			time.Sleep(time.Millisecond * 50)
		}
	}
}
