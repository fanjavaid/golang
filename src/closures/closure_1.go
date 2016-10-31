package main

import "fmt"

func main() {
	// named function
	printMessage("Hello Go!")

	// anonymous function declared and called
	func (message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	// get anonymous function and call it
	printFunc := getPrintMessage()
	printFunc("Hello, World!")
}

func printMessage(message string) {
	fmt.Println(message)
}

// buat sebuah fungsi
// dengan kembalian berupa fungsi
func getPrintMessage() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}