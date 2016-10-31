package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	age   int
	person
}

func main() {
	var s1 = student{}
	s1.name = "Fandi"
	s1.age = 25
	s1.person.age = 0
	s1.grade = 98

	fmt.Println(s1)
	fmt.Println(s1.person.name)
	fmt.Println(s1.person.age)
	fmt.Println(s1.age)

	var p2 = person{name: "M.Iqbal"}
	var s2 = student{age: 24, grade: 90, person: p2}
	fmt.Println(s2)
}
