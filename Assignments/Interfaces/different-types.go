package main

import (
	"fmt"
	"reflect"
)

func do(i interface{}) {
	v := reflect.ValueOf(i)
	//fmt.Println(v)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("Type is an integer:", v)
	case reflect.Float64:
		fmt.Println("Type is a float:", v)
	case reflect.String:
		fmt.Println("Type is a string:", v)
	case reflect.Bool:
		fmt.Println("Type is a boolean:", v)
	case reflect.Map:
		fmt.Println("Type is a map:", v)
	case reflect.Slice:
		fmt.Printf("Type is slice: len = %d, %v\n", v.Len(), v)
	default:
		fmt.Println("Type is unknown: ", v)
	}
}

func main() {
	do(12)
	do(13.5)
	do("Hello")
	do(true)
	do(map[int]string{1: "Sahiba", 2: "Anu"})
	do([]int{1, 2, 3})
}
