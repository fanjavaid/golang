package main

import (
	"fmt"
	"os"
	"time"
)

func timeout(duration int, ch chan<- bool) {
	time.AfterFunc(time.Duration(duration)*time.Second, func() {
		ch <- true
	})
}

func watcher(duration int, ch <-chan bool) {
	<-ch
	fmt.Println("\nTimeout! no Answer after", duration, "seconds")
	os.Exit(0)
}

func watcher2(duration int, ch <-chan bool) {
	<-ch
	fmt.Println("This is watcher 2 as a second receiver")
}

func main() {
	var data = make(chan bool)
	var duration = 5

	go timeout(duration, data)
	go watcher(duration, data)
	go watcher2(duration, data)

	var input string
	fmt.Print("What is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("Correct")
	} else {
		fmt.Println("Wrong!")
	}
}
