package main

import (
	"html/template"
	"log"
	"net/http"
)

type formHandler int

func (m formHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("eescano-Key", "write custom header demo")

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d formHandler
	http.ListenAndServe(":8080", d)
}
