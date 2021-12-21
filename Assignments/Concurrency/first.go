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

type Stud struct {
	index   int
	average float64
}

func average(stud Student, j int, c chan Stud) {
	sum := float64(stud.sub1 + stud.sub2 + stud.sub3)
	avg := sum / 3.0
	c <- Stud{index: j, average: avg}
}

func main() {
	start := time.Now()
	temp := make([]Student, 100)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan Stud)
	for i := range temp {
		temp[i].name = string(letters[rand.Intn(len(letters))])
		temp[i].sub1 = rand.Intn(100)
		temp[i].sub2 = rand.Intn(100)
		temp[i].sub3 = rand.Intn(100)
		//fmt.Println(i, temp[i].sub1, temp[i].sub2, temp[i].sub3)
	}

	for j, student := range temp {
		go average(student, j, c)
	}

	for _ = range temp {
		x := <-c
		temp[x.index].average = x.average
	}
	fmt.Println(temp[2].sub1, temp[2].sub2, temp[2].sub3, temp[2].average)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
