package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println(reflectValue.Type())
	fmt.Println(reflectValue.Int())

	// Pengaksesan value dg interface
	fmt.Println(reflectValue.Interface())
}
