package main

import (
	"fmt"
)

func main() {
	var isOld bool = false

	// Sample using String datatype using doublequote or backticks
	var name string = "Fandi\nAkhmad"
	var message string = `This is html message for email.\nAttach this to your message format.`

	fmt.Println("Am i too old? ", isOld)
	fmt.Println(name)
	fmt.Println(message)

	// Menggunakan constanta
	const PHI = 3.14
	// PHI = 3 , will get Error assignment to constant
	fmt.Println("Nilai PHI : ", PHI)
}
