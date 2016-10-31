package main

import "fmt"

// Nested struct
type applications struct {
	applicant struct {
		name  string
		email string
	}
	loanId string
}

func main() {

	var app1 = applications{
		loanId: "B1101",
	}
	app1.applicant.name = "Fandi"
	app1.applicant.email = "example@gmail.com"

	fmt.Println(app1)
}
