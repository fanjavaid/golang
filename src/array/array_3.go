package main

import "fmt"

func main() {
	// Basic
	var applicants [5] string
	applicants[0] = "Zainu Rochim"
	applicants[1] = "Andreas Pribadi"
	applicants[2] = "Tri Ari Kurniawati"
	applicants[3] = "Bonny Sebastian Rahardjo"
	applicants[4] = "Mohamad Ikhsan Fadzilah"

	fmt.Println("Applicant ke-1 = ", applicants[0])

	// Pengisian elemen array saat deklarasi
	var loanIds = [4] string {"bb48b4430cdf3-1501","ccd1d6867b2e2-1501","d4f7e60845771-1501"}

	// Cetak panjang array
	fmt.Println("Panjang array loanIds = ", len(loanIds))
	fmt.Println("Value loanIds ya terakhir = ", loanIds[len(loanIds) - 1])

	// Membuat array tanpa eksplisit panjangnya
	var provinces = [...] string {
		"Nanggroe Aceh Darussalam (NAD)",
		"Nusa Tenggara Barat (NTB)",
		"Nusa Tenggara Timur (NTT)",
		"Sulawesi Barat",
	}
	fmt.Println(provinces)

	fmt.Println("\nUsing loop to display array values : ")
	for i := 0; i < len(provinces); i++ {
		fmt.Println(provinces[i])
	}

	fmt.Println("\nApplicants per 2016 : ");
	for i, value := range applicants {
		fmt.Println((i+1), ". ", value)
	}

	// Membuat sebuah array dengan data loan id dan applicant name
	// multidimensional array
	var applications = [3][2] string {
		{"bb48b4430cdf3-1501", "Zainu Rochim"},
		{"ccd1d6867b2e2-1501", "Bonny Sebastian Rahardjo"},
		{"d4f7e60845771-1501", "Mohamad Ikhsan Fadzilah"},
	}
	_ = applications // Kalau variabel di atas tidak digunakan

}