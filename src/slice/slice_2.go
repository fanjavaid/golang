package main

import "fmt"

func main() {
	// []T
	var affiliators = [] string {
		"Noni Noviyanti Letarie Sitinjak",
		"Artie Perwita Sari",
		"Nasrul Bima Setiaji",
		"Lengky Angling Kusuma Wardhana",
	}

	fmt.Println("Slice values : ");
	fmt.Println(affiliators)

	fmt.Println("\nGet Length using len = ", len(affiliators))
	fmt.Println("Get Length using cap = ", cap(affiliators))

	// Slice operations
	var top2Affiliators = affiliators[0:2]
	fmt.Println(top2Affiliators)

	// Create slice using make
	var applicantEmails = make([] string, 3)
	applicantEmails[0] = "slank_jurustandur@yahoo.com"
	applicantEmails[1] = "fransischaputri.sfp@gmail.com"
	applicantEmails[2] = "mrasyidin21@gmail.com"

	fmt.Println("Before Append = ", applicantEmails)

	// Gue bisa append nih
	applicantEmails = append(applicantEmails, "muhammadnafiah1612@gmail.com")
	fmt.Println("After Append = ", applicantEmails)

	// Slice juga bisa di copy dari satu slice ke yang lain
	var tempApplicantEmails = make([] string, len(applicantEmails))
	copy(tempApplicantEmails, applicantEmails)

	fmt.Println("Values of temporary applicant emails is = ", tempApplicantEmails)

	// Waktunya demo operator "slice" pada slice ??
	var loanAmounts = [] float32 {
		4320000.0,
		5000000.0,
		2378000.0,
		1879000.0,
		1340000.0,
	}
	temp := loanAmounts[0:3] // cuma ambil potongan dari 0 - 2
	fmt.Println(temp)
	
	temp2 := loanAmounts[1:4] // cuma ambil potongan dari 1 - 3
	fmt.Println(temp2)

	temp3 := loanAmounts[3:5] // cuma ambil potongan 3 - 4
	fmt.Println(temp3)
}