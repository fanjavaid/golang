package main

import "fmt"

func main() {
	// Inisialisasi slice
	// dengan nilai default tiap element = 0
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Isi nilai tiap elemen seperti mengisi nilai pada Array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len mengembalikan panjang dari slice
	fmt.Println("len:", len(s))

	// Keyword append merupakan perintah basic dari slice
	// yang membuat slice lebih full feature daripada array
	// append mengembalikan nilai slice yang baru yg sudah ditambahkan dengan data
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("after append:", s)

	// Slice juga support untuk di copy ke slice yang lain
	// Walaupun dengan panjang yang berbeda? Ya
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:",c)

	// Slice juga mendukung operasi slice atau memotong
	// dengan syntax data_slice[index:panjang slice maksimum]
	// [a b c d e]
	l := c[1:2]
	fmt.Println("sl1:", l)

	// Jika pemotongan tidak dimasukkan parameter pertama, maka
	// Pemotongan akan dilakukan sepanjang slice dikurangi 1, artinya mengambil indexnya
	m := c[:5] 
	fmt.Println("sl2:", m)

	// Jika pemotongan tidak dimasukkan parameter kedua, maka
	// pemotongan akan dilakukan sepanjang slice termasuk index yang didefinisikan di parameter
	n := c[2:]
	fmt.Println("sl3:", n)
}
