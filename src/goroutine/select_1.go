package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, el := range numbers {
		sum += el
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, el := range numbers {
		if max < el {
			max = el
		}
	}

	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 5, 6, 23, 1, 2, 5, 6, 12, 1, 90, 3122}
	fmt.Println(numbers)

	// ch1 untuk menampung nilai dari rata-rata
	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	// ch2 untuk menampung nilai tertinggi dari angka
	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg\t:%.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max\t:%d \n", max)
		}
	}
}
