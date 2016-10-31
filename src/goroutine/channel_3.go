package main

import (
	"fmt"
	"runtime"
)

func printMsg(msg chan string) {
	fmt.Println(<-msg)
}

func main() {
	runtime.GOMAXPROCS(4)

	var chMsg = make(chan string)

	for _, each := range []string{"john", "fandy", "wick"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			chMsg <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMsg(chMsg)
	}
}
