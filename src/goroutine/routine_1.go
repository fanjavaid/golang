package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	go print(5, "Hallo")
	print(5, "Apa kabar?")

	var input string
	fmt.Scanln(&input)
}
