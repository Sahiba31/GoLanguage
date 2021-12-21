package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Student struct {
	name    string
	id      int
	sub1    int
	sub2    int
	sub3    int
	average float64
}

func average(stud []Student, c chan float64) {
	for i := 0; i < 100; i++ {
		sum := stud[i].sub1 + stud[i].sub2 + stud[i].sub3
		avg := float64(sum / 3.0)
		stud[i].average = avg
		c <- stud[i].average
	}
	close(c)
}

func main() {
	start := time.Now()
	temp := make([]Student, 100)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range temp {
		temp[i].name = string(letters[rand.Intn(len(letters))])
		temp[i].sub1 = rand.Intn(100)
		temp[i].sub2 = rand.Intn(100)
		temp[i].sub3 = rand.Intn(100)
	}

	c := make(chan float64)
	go average(temp, c)
	for avg := range c {
		fmt.Println(avg)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
