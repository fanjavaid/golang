package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2016, 5, 20, 10, 00, 43, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
}
