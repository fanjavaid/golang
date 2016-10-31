package main

import "fmt"

// fungsi defer dijalankan setelah akhir block dari sebuah fungsi
// jika ada beberapa defer maka, defer akan dijalankan dari defer terakhir ke defer teratas
func main() {
	defer fmt.Println("Go Programming Language") // 3
	defer fmt.Println("Loading...")
	fmt.Println("Defer and Exit")   // 1
	defer fmt.Println("Welcome to") // 2
}
