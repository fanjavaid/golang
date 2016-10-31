package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Fandi Akhmad"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodeString)
}
