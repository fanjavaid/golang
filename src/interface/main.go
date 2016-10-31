package main

import "interface/listener"

func main() {
	var myButton listener.ButtonEvent
	myButton = Button{name: "My Button"}

	myButton.OnClick()
}
