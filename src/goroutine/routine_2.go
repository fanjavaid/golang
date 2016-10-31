package main

import "fmt"

func printLoop(text string) {
	for i := 0; i < 150; i++ {
		fmt.Println(text, "ke-", i)
	}
}

func main() {
	fmt.Println("Normal Call : ")
	printLoop("Go routine!")

	fmt.Println("Call using go routine")
	go printLoop("Real Go routine!")

	// call anonymous function using go routine
	go func(text string) {
		fmt.Println(text)
	}("welcome!")

	var input string
	fmt.Scanln(&input)
}
