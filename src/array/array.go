package main

import (
	"fmt"
)

func main() {
	// Basic Array
	var names [4]string
	names[0] = "Fandi"
	names[1] = "Akhmad"
	names[2] = "Via"
	names[3] = "Hello"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Pengisian elemen array saat deklarasi
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Panjang Array yaitu ", len(fruits))
	fmt.Println("Debug nilai array : ", fruits)

	// Membuat array tanpa explicit size allocation
	var addresses = [...]string{
		"DKI Jakarta",
		"Bandung",
		"DI Yogyakarta",
	}

	fmt.Println(addresses)

	// Array multidimensi
	// array 2 rows 3 columns
	var numbers1 = [2][3]int{
		{3, 2, 4},
		{2, 6, 4},
	}

	fmt.Println(numbers1)

	// Looping array
	var worksheets = [...]string{
		"Daily",
		"Weekly",
		"Monthly",
	}

	for i := 0; i < len(worksheets); i++ {
		fmt.Println("Activating worksheet : ", worksheets[i])
	}

	// Looping array using for-range
	var excelData = [2][3]string{
		{"Jan", "10", "30"},
		{"Feb", "5", "98"},
	}

	for i, ed := range excelData {
		fmt.Println("Row -", i, " = ", ed)
	}

	// Alokasi array dengan keyword make
	var mydatas = make([]string, 2)
	mydatas[0] = "Lorem"
	mydatas[1] = "Ipsum"

	for _, data := range mydatas {
		fmt.Println(data)
	}
}
