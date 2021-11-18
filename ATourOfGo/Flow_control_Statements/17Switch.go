package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "ios":
		fmt.Println("IOS")
	case "darwin":
		fmt.Println("Darwin")
	default:
		fmt.Println("Any other")
	}
}