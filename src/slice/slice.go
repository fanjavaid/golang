package main

import {
	"fmt"
}

func main() {
	s := make([]string, 0)
	fmt.Println(s)

	// Append some values
	s = append(s, "lorem")
	s = append(s, "ipsum", "dolor", "sit", "amet")
	fmt.Println(s)
	fmt.Println("Length of s = ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// Using slice
	// start from 2 and end < 4
	l := s[2:4]
	fmt.Println(l)

	// slice up exclude 4
	l = s[:4]
	fmt.Println(l)

	// slice up including 2
	l = s[2:]
	fmt.Println(l)
}
