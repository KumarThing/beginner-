package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/home.html"))

type PageData struct{
	Result int
	HasValue bool
	Error string
}
func main() {

http.Handle("/bin2dec/static/", 
http.StripPrefix("/bin2dec/static/",
http.FileServer(http.Dir("static"))))

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

data:= PageData {
	HasValue: false,
}
tmpl.Execute(w, data)
})

http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	inputStr := r.FormValue("input")

	if len(inputStr) == 0 {
		data := PageData{
			Error: "Please enter bionary number",
		}
		tmpl.Execute(w, data)
		return
	}

	result := 0

	for i := 0; i < len(inputStr) ; i ++ {

		
		if inputStr[i] != '0' && inputStr[i] != '1' {
			data := PageData{
				Error: "Invalid input only 0 and 1 are allowed",
			}
			tmpl.Execute(w, data)
			return
		}

		digit := int(inputStr[i]- '0')
		power := len(inputStr) - i - 1
		result += digit *(1 <<power)

	}

	data := PageData {
		Result: result,
		HasValue: true,
	
	}

	tmpl.Execute(w, data)


})
http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	data := PageData{
		
	}
	tmpl.Execute(w, data)

})



fmt.Println("Server is running in http://localhost:8080")
http.ListenAndServe(":8080", nil)

}