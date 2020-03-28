package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type test_struct struct {
	Test []int
}

func root(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Homepage")
}
func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello World")
}

func sum(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var val test_struct
	err := decoder.Decode(&val)

	if request.Body == nil {
		http.Error(response, "Please send a request body", 400)
		return
	}

	if err != nil {
		panic(err)
	}

	numbers := val.Test
	fmt.Println(numbers)
	length := len(numbers)
	result := 0
	for items := 0; items < length; items++ {
		result += numbers[items]
	}
	responseBack := {list: numbers, sum: result}
	json.NewEncoder(response).Encode(responseBack)
	fmt.Println(result)

}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", root).Methods("GET")
	myRouter.HandleFunc("/hello", hello).Methods("GET")
	myRouter.HandleFunc("/sum", sum).Methods("POST")
	address := ":8080"
	log.Println("Starting server on address", address)
	log.Fatal(http.ListenAndServe(address, myRouter))
}
func main() {
	handleRequest()
}
