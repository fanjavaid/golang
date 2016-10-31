package main

import "fmt"

type applicant struct {
	name  string
	email string
}

func main() {
	var applicants = []applicant{
		{name: "Fandi", email: "fanjavaid@gmail.com"},
		{name: "Akhmad", email: "akhmad.f@gmail.com"},
	}

	fmt.Println(applicants)

	for _, value := range applicants {
		fmt.Printf("\nName %s and Email is %s", value.name, value.email)
	}

	// Anonymuous struct and slice
	fmt.Println("\n")
	var users = []struct {
		username string
		password string
	}{
		{username: "fanjavaid", password: "secret"},
		{username: "eko", password: "default"},
	}

	fmt.Println(users)
}
