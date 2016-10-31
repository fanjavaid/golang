package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 interface{} = &person{name: "Fandy", age: 24}
	var name = p1.(*person).name

	fmt.Println(name)
}
