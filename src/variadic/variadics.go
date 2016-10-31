package main

import "fmt"

func calculate (numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	
	return avg
}

func printValue (values ...string) {
	for i, value := range values {
		fmt.Println(i,". ", value)
	}
}

func main() {
	var avg = calculate(4, 5, 8, 9, 1, 2, 10)
	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
	
	fmt.Println(msg)

	var datas = []string {
		"Hello", 
		"Guys",
		"I",
		"Learn",
		"Go Lang",
	}

	printValue(datas...)
}
