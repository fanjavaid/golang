package main

import (
	"fmt"
)

func main() {
	var age = 23

	if age >= 60 {
		fmt.Println("Too old")
	} else if age >= 30 {
		fmt.Println("Productive Person")
	} else {
		fmt.Println("Hmmm Young enough...")
	}

	fmt.Println("-----------------------------")

	var point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	fmt.Println("-----------------------------")

	var nilai = 6
	switch nilai {
		case 8 :
			fmt.Println("Bagus!")
		case 7 :
			fmt.Println("Cukup")
		default :
			fmt.Println("Kurang!")
	}

	fmt.Println("-----------------------------")

	var promo = "weekend"
	switch promo {
		case "ramadhan" :
			fmt.Println("Promo 50%")
		case "lebaran" :
			fmt.Println("Promo 70%")
		case "weekend", "holiday" : // Multiple condition
			fmt.Println("Promo 25%")
		default :
			fmt.Println("Promo 0%")	
	}

	fmt.Println("-----GAYA IF ELSE------")

	var nilaiUjian = 90
	switch {
		case nilaiUjian > 80 : 
			fmt.Println("A")
		case nilaiUjian > 70 : 
			fmt.Println("B")
		case nilaiUjian > 60 : 
			fmt.Println("C")
		default : 
			fmt.Println("D")
	}

	fmt.Println("-----FALLTRHOUGH------")

	var myPoint = 6
	switch {
		case myPoint == 8 :
			fmt.Println("perfect")
		case (myPoint < 8) && (myPoint > 3) :
			fmt.Println("awesome")
			fallthrough 
		case point < 5 :
			fmt.Println("You need to learn more")
		default :
			{
				fmt.Println("Not bad")
				fmt.Println("You need to learn more")
			}
	}

}
