package main

import (
	"fmt"
	"runtime"
)

// buat sebuah fungsi untuk pengiriman data
// menggunakan go routine
// chan<- menandakan jika fungsi ini hanya digunakan sebagai channel pengirim
func sendMessages(message chan<- string) {
	for i := 0; i < 15; i++ {
		message <- fmt.Sprintf("Data - %d", i)
	}
	close(message)
}

// buat sebuah fungsi untuk menerima data dari channel
// <- chan
func getMessages(message <-chan string) {
	for msg := range message {
		fmt.Println(msg)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	// buat sebuah channel
	var msgChannel = make(chan string)

	// jalankan fungsi pengiriman data dengan goroutine
	go sendMessages(msgChannel)

	// cetak pesan
	getMessages(msgChannel)
}
