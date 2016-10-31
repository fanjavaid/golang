package main

import (
	"encoding/json"
	"fmt"
)

type Mydata struct {
	Id      string `json:"_id"`
	Gender  string
	Company string
	Email   string
	Tags    []string
	Friends []Friend
}

type Friend struct {
	Id   int
	Name string
}

func main() {
	var jsonString = `{"_id":"5816ef95d22784b23337be90","gender":"male","company":"DIGIFAD","email":"everetthobbs@digifad.com","tags":["exercitation","commodo","commodo","laborum","culpa","enim","irure"],"friends":[{"id":0,"name":"Fannie Gill"},{"id":1,"name":"Jordan Banks"},{"id":2,"name":"Cindy Velazquez"}]}`

	var jsonByte = []byte(jsonString)
	var data Mydata
	var err = json.Unmarshal(jsonByte, &data)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)
}
