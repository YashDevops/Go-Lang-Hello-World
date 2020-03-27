package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleRequest() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequest()
}
