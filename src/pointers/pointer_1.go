package main

import "fmt"

func main() {
	i := 5

	zeroVal(i)
	zeroPtr(&i)

	// Ada variable
	var age int = 24
	var agePointer *int = &age

	fmt.Println("Value variable age = ", age)
	fmt.Println("Alamat memory variable age = ", &age)

	fmt.Println("Value pointer agePointer = ", *agePointer)
	fmt.Println("Alamat memory pointer agePointer = ", agePointer)

	// gue coba ganti value nih di pointer
	*agePointer = 25
	fmt.Println("Value variable age = ", age)
}

func zeroVal(val int) {
	fmt.Println(val)
}

func zeroPtr(iptr *int) {
	fmt.Println(*iptr)
}
