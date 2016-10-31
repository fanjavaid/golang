package main 

import "fmt" 

func main() {
	var names[4] string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var members = [2] string {"fandi", "akhmad"}

	fmt.Println("Length = ", len(members));

	// Inisialisasi Array dengan gaya vertikal
	var fruits = [4] string {
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println(fruits)

	// Inisialisasi array tanpa panjang
	var applicants = [...] string {"edo", "sinta", "prabu", "maulana"}
	fmt.Println(applicants)

	// Membuat array multidimensi
	var numbers = [2][3] int {{1,5,6}, {6,5,3}}
	fmt.Println(numbers);
	fmt.Println(len(numbers))

	// Loop data array
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Loop 1 : %s\n", fruits[i])
	}

	// Loop dengan for range
	for key, fruit := range fruits {
		fmt.Printf("Loop Key %d : %s\n", key, fruit)
	}

	// Penggunaan variabel underscore
	var mydatas = [4] string {"Document 1", "Document 2", "Document 3", "Document 4"}
	for _, data := range mydatas {
		fmt.Println(data)
	}

	// Alokasi elemen array menggunakan keyword `make`
	// Merujuk ke dokumentasi Go, make digunakan untuk inisialisasi SLICE
	var buah = make([] string, 2)
	buah[0] = "Apel"
	buah[1] = "Jeruk"

	fmt.Println(buah)
}