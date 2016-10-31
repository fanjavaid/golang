package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHello("Fandy")
	go sayHello("Wick")
	go sayHello("Jonathan Doe")

	var msg1 = <-messages
	fmt.Println(msg1)

	var msg2 = <-messages
	fmt.Println(msg2)

	var msg3 = <-messages
	fmt.Println(msg3)
}
