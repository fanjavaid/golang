package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Fandi Akhmad",
			"Message": "Go Programming Language!",
		}

		var t, err = template.ParseFiles("web/template.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		t.Execute(w, data)
	})

	fmt.Println("starting server")
	http.ListenAndServe(":8080", nil)
}
