package main

import "fmt"

func main() {
	outer("Fandy A.")
}

func outer(name string) {
	text := "Modified by " + name

	foo := func() {
		fmt.Println(text) // foo has access to text variable
	}

	// call foo() closure
	foo()
}