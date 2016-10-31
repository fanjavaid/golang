package main

import (
	"fmt"
)

func main() {
	// Membuat Map
	var chicken map[string]int
	chicken = map[string]int{}

	// Isi nilai Map
	chicken["januari"] = 40
	chicken["februari"] = 10

	// Cetak
	fmt.Println("All Chickens")
	fmt.Println(chicken)

	fmt.Println("Chicken in Januari = ", chicken["januari"])
	fmt.Println("Chicken in Mei = ", chicken["mei"]) // return 0 karena tidak ada

	// Definisikan nilai map di awal
	var sales map[int]double
	sales = map[int]double{
		0: 1500000,
		1: 4567000,
	}

	fmt.Println(sales)
}
