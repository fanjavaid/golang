package main

import "fmt"

type member struct {
	name string
	dob  string
}

func (m *member) SetName(pName string) {
	m.name = pName
}

func (m *member) GetName() string {
	return m.name
}

func main() {
	// var myMember = &member{}
	var myMember member
	myMember.SetName("Fandi Akhmad")
	// myMember.name = "Fandi"
	fmt.Println(myMember.name)

	// var reflectValue = reflect.ValueOf(myMember)
	// var method = reflectValue.MethodByName("SetName")
	// method.Call([]reflect.Value{
	// 	reflect.ValueOf("Wick"),
	// })
	//
	// fmt.Println(*myMember)
}
