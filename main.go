///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var address = ":8080"

type test_struct struct {
	Number []int
}

type responeJson struct {
	Number []int `json:"number"`
	Sum    int   `json:"sum"`
}

type logLevel struct {
	Loglevel string `json:loglevel`
}

var LOGGER = logrus.Logger{}

// function for root uri
func root(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Homepage")
}

// function to print hello
func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello World")

}

// function to send sum
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

	numbers := val.Number
	length := len(numbers)
	result := 0
	for items := 0; items < length; items++ {
		result += numbers[items]
	}
	finalResponse := responeJson{
		Number: numbers,
		Sum:    result,
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(finalResponse)
}

//Function to change log level
func loglevel(response http.ResponseWriter, request *http.Request) {
	var logtemp logLevel
	decoder := json.NewDecoder(request.Body).Decode(&logtemp)
	if decoder != nil {
		panic(decoder)
	}
	fmt.Print(logtemp.Loglevel)
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", root).Methods("GET")
	myRouter.HandleFunc("/hello", hello).Methods("GET")
	myRouter.HandleFunc("/sum", sum).Methods("POST")
	myRouter.HandleFunc("/loglevel", loglevel).Methods("POST")
	logrus.Infof("Successfully initialized Server At port " + address)
	logrus.Errorln(http.ListenAndServe(address, myRouter))
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
	handleRequest()
}
