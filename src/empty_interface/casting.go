package main

import "fmt"

func main() {
	var number interface{}

	number = 5
	total := number.(int) * 10
	fmt.Println(total)
}
