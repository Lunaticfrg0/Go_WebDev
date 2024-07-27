package main

import (
	"io"
	"net/http"
)

func initialPageHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		NotFoundPageHandler(res, req, http.StatusNotFound)
		return
	}
	io.WriteString(res, "Welcome to base site")
}
func AboutPageHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to about page")
}
func InfoPageHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Info page")
}
func NotFoundPageHandler(res http.ResponseWriter, req *http.Request, status int) {
	res.WriteHeader(status)
	if status == http.StatusNotFound {
		io.WriteString(res, "404 Not found page")
	}
}

func main() {
	http.HandleFunc("/", initialPageHandler)
	//Strict just about
	http.HandleFunc("/about", AboutPageHandler)
	//Any child get under info/ Ex: info/prueba
	http.HandleFunc("/info/", InfoPageHandler)

	http.ListenAndServe(":8080", nil)

}
