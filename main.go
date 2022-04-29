package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page!")
	fmt.Println("Call Endpoint / ")
}

func handleReq() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleReq()
}
