package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    int
	Name  string
	Price float64
	Tags  []string
}

func main() {
	var jsonString = `{"id":1,"name":"A green door","price":12.50,"tags":["home","green"]}`
	var jsonData = []byte(jsonString)

	var prod Product
	var err = json.Unmarshal(jsonData, &prod)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Unmarshal JSON result : ")
	fmt.Println(prod)
}
