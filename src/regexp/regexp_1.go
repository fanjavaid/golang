package main

import (
	"fmt"
	"regexp"
)

func main() {
	var fruits = "banana apple orange"
	var regex, err = regexp.Compile("[a-z]+")

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(fruits, 2)
	fmt.Println(res1)

	// Menggunakan match string
	var isMatch = regex.MatchString(fruits)
	fmt.Println(isMatch)

	// Menemukan string pertama yang cocok dengan regex
	var firstTxt = regex.FindString(fruits)
	fmt.Println(firstTxt)

	// Menemukan string index yang cocok dengan regex
	var index = regex.FindStringIndex(fruits)
	fmt.Println(index)
	fmt.Println("Text yang ada pada index", index, "yaitu", fruits[0:6])

	// replace semua string yang memenuhi kriteria regex
	var replaceTxt = regex.ReplaceAllStringFunc(fruits, func(each string) string {
		if each == "apple" {
			return "potato"
		}

		return each
	})

	fmt.Println(replaceTxt)

	//
}
