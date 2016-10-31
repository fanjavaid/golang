package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("This statement run after 4 seconds")
		ch <- true
	})

	fmt.Println("start")
	fmt.Println(<-ch)
	fmt.Println("Finish")
}
