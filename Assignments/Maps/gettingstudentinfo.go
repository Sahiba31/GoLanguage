package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Students struct {
	Sdata []Student `json:"students"`
}
type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Dept string `json:"department"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("studentinfo.json")
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened studentinfo.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var candidates Students
	// we unmarshal our byteArray which contains our jsonFile's content into 'candidates' which we defined above
	json.Unmarshal(byteValue, &candidates)
	//fmt.Println(candidates)

	mp := make(map[int]Student)

	for i := 0; i < len(candidates.Sdata); i++ {
		//Created temp to store the information of one student and later on insert that into map
		temp := Student{
			Name: candidates.Sdata[i].Name,
			Id:   candidates.Sdata[i].Id,
			Dept: candidates.Sdata[i].Dept,
		}
		mp[candidates.Sdata[i].Id] = temp
	}
	// for key, value := range mp {
	// 	fmt.Println(key, value)
	// }
	var input int
	fmt.Print("Enter the Id: ")
	fmt.Scan(&input)
	_, ok := mp[input]
	if ok == false {
		fmt.Println("Key is not Valid")
	} else {
		fmt.Println("Name:", mp[input].Name, ",Department:", mp[input].Dept)
	}
}
