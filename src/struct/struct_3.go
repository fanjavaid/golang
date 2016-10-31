package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// create anonymous struct
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{name: "Fandi Akhmad", age: 24}
	s1.grade = 90

	fmt.Println("Name:", s1.person.name)
	fmt.Println("Age:", s1.person.age)
	fmt.Println("Grade:", s1.grade)

	var s2 = struct {
		loanId string
		name   string
	}{
		loanId: "bb48b4430cdf3-1501",
		name:   "Fandi Akhmad",
	}

	fmt.Println(s2)
}
