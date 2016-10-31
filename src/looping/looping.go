package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(`*`)
		}
		fmt.Println("")
	}
	for i := 10; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(`*`)
		}
		fmt.Println("")
	}

	var i = 0
	for i < 5 {
		fmt.Println("Contoh menggunakan looping")
		i++
	}

	// Unlimited for
	// but can interupt using break
	var a = 0
	for {
		fmt.Println("Angka ", a)

		a++

		if a == 5 {
			break
		}
	}

	// using continue and break
	for y := 0; y < 10; y++ {
		if y == 3 {
			continue
		}

		if y == 6 {
			break
		}

		fmt.Println("Angka ke ", y)
	}

	// break using label
outerloop:
	for z := 0; z < 10; z++ {
		for f := 0; f < 3; f++ {
			if z == 3 {
				break outerloop
			}

			fmt.Println("Matriks [", z, "][", f, "]")
		}
	}
}

// 12345

// 1
// 12
// 123
// 1234
// 12345

// i
// .....
// 0
// 1
// 2

// 0
// 0 1
// 0 1 2
