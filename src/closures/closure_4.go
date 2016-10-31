package main

import "fmt"

func main() {
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)

	_ = incr
	_ = incr1

	fmt.Println("counter = ", ctr())
	fmt.Println("counter 1 = ", ctr1())

	incr()
	fmt.Println("counter = ", ctr())
	fmt.Println("counter 1 = ", ctr1())
}

func counter(start int) (func() int, func()) {
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	return ctr, incr
}
