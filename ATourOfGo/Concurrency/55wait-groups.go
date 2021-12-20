package main

import (
	"fmt"
	"sync"
	"time"
)

func count(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("Sahiba")
		wg.Done()
	}()

	wg.Wait()
}
