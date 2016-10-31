package main

import "fmt"

type Button struct {
	name string
}

func (b Button) OnClick() {
	fmt.Println("Button clicked")
}
