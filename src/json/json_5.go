package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Grade string
}

func main() {
	var arrJSON = `[
        {"Name" : "Fandy", "Grade" : "A"},
        {"Name" : "John Doe", "Grade" : "C"}
    ]`

	var student []Student
	var err = json.Unmarshal([]byte(arrJSON), &student)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(student)
}
