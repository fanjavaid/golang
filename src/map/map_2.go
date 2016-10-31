package main 

import "fmt"

func main() {
	var billings map[string]int
	billings = map[string] int {
		"82180085422" : 50000,
		"85695887334" : 100000,
	}

	fmt.Println("Data Billings : ")
	fmt.Println(billings)

	fmt.Println("Data Billing untuk ANI 82180085422 = Rp ", billings["82180085422"])

	var applicants = map[string] int {
		"bb48b4430cdf3-1501" : 2500000,
		"ccd1d6867b2e2-1501" : 4706000,
	}

	fmt.Println(applicants)

	// Pake fungsi delete nya maps
	delete(applicants, "ccd1d6867b2e2-1501")

	fmt.Println(applicants)
}