package main

import (
	"fmt"
	"sort"
)

func sortArray(arr []string, n int) {
	//map of int (as key) and array of string (as values)
	mp := make(map[int][]string)
	var vec []int //array for storing all keys so that we can sort them later easily
	for i := 0; i < n; i++ {
		length := len(arr[i])
		_, ok := mp[length]
		if ok == false {
			vec = append(vec, length)
		}
		mp[length] = append(mp[length], arr[i])
	}
	sort.Ints(vec) //sorting for length
	for i := 0; i < len(vec); i++ {
		sort.Strings(mp[vec[i]]) //sorting the inside strings according to alphabets
		for j := 0; j < len(mp[vec[i]]); j++ {
			fmt.Printf("%s ", mp[vec[i]][j])
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Enter the size of array of strings: ")
	fmt.Scan(&n)
	fmt.Print("Enter the strings: ")
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sortArray(arr, n)
}
