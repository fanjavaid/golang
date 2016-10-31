package main

import "fmt"

// first type
type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return "Hallo " + p.name
}

// second type
type applicant struct {
	name   string
	mobile string
}

func (a *applicant) printInfo() {
	fmt.Println("Name :", a.name)
	fmt.Println("Mobile :", a.mobile)
}

func main() {
	var p = person{
		name: "Fandi",
		age:  24,
	}

	fmt.Println(p.getName())

	var ap1 = &applicant{
		name:   "Fandi AKhmad",
		mobile: "085695887334",
	}

	ap1.printInfo()
}
