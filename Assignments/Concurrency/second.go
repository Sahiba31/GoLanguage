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

func average(stud *Student) {
	sum := stud.sub1 + stud.sub2 + stud.sub3
	avg := float64(sum / 3.0)
	stud.average = avg
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

	for i := range temp {
		average(&temp[i])
		fmt.Println(temp[i].average)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
