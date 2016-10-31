package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Gender string
	Phone  string
}

func main() {
	var persons = []Person{
		{
			Name:   "Fandi",
			Gender: "Male",
			Phone:  "085695887334",
		},
		{
			Name:   "John Doe",
			Gender: "Male",
			Phone:  "081280085422",
		},
	}

	var jsonData, err = json.Marshal(persons)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(jsonData))
}
