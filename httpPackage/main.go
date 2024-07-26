package main

import (
	"fmt"
	"net/http"
)

type helloWorldHandler int

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from my http server!")
}

func main() {
	fmt.Println("Server started!")
	var helloWorld helloWorldHandler
	http.ListenAndServe(":8080", helloWorld)
}
