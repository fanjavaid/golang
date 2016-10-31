package main

import (
	"fmt"
	"time"
)

func main() {
	// menggunakan fungsi sleep
	fmt.Println("Hello Go")
	time.Sleep(time.Second * 4)
	fmt.Println("After 4 seconds")
}
