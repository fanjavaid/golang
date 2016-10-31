package main

import "fmt"

func main() {
	calculate(20,"celcius","fahrenheit")
}

func calculate(temperature float32, from string, to string) {
	var finalTemp float32 = 0;

	log := func() {
		fmt.Printf("Convert %.2f from %s to %s = %.2f", temperature, from, to, finalTemp)
	}

	convert := func() {
		if (from == "celcius" && to == "fahrenheit") {
			finalTemp = temperature * 1.8 + 32
		}

		log()
	}

	convert()
}
