package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var myJSON = `{"_id":"5816ef95d22784b23337be90","gender":"male","company":"DIGIFAD","email":"everetthobbs@digifad.com","tags":["exercitation","commodo","commodo","laborum","culpa","enim","irure"],"friends":[{"id":0,"name":"Fannie Gill"},{"id":1,"name":"Jordan Banks"},{"id":2,"name":"Cindy Velazquez"}]}`

	var jsonByte = []byte(myJSON)

	var data map[string]interface{}
	// Unmarshal to map type
	var err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)
	fmt.Println(data["gender"])
}
