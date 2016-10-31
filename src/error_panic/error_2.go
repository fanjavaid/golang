// coba buat custom error di golang

package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot bet empty")
	}

	return true, nil
}

func main() {
	var name string = ""

	if valid, err := validate(name); valid {
		fmt.Println("Halo,", name)
	} else {
		// fmt.Println(err.Error())
		// use panic to show trace error
		panic(err.Error())
	}
}
