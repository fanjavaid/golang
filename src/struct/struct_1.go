package main

import "fmt"

func main() {
	type applicant struct {
		loanId string
		name   string
	}

	var fandy = applicant{}
	fandy.loanId = "B1101"
	fandy.name = "Fandi Akhmad"

	var via = applicant{"B1102", "Via P."}

	fmt.Println(fandy)
	fmt.Println(via.loanId)

	// Struct dan pointer
	var objPointer *applicant = &fandy
	objPointer.name = "Akhmad"
	fmt.Println("objPointer = ", objPointer)
}
