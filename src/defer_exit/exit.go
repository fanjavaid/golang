package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Defer at first line")

	fmt.Println("Hello Go!")
	os.Exit(1)

	// defer tidak akan dijalankan setelah fungsi Exit dipanggil
	defer fmt.Println("Learning....")
}
