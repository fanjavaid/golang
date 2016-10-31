package main

import (
	"fmt"
	"time"
)

func main() {
	// konversi dari string ke time
	var myTime = "2016-10-20 10:00:34"
	var layoutFormat = "2006-01-02 15:04:05"

	var date time.Time

	date, _ = time.Parse(layoutFormat, myTime)
	fmt.Println(myTime, " -> ", date.String())
}
