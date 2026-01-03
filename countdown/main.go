package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tmpl.Execute(w, nil)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
