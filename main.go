package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page!")
	fmt.Println("Call Endpoint / ")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Call Endpoint /articles")
	json.NewEncoder(w).Encode(Articles)
}

func handleReq() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "John Doe", Desc: "John Doe Descriptions ...", Content: "John Doe Contents..."},
		Article{Title: "John Doe 1", Desc: "John Doe Descriptions 2 ...", Content: "John Doe Contents 2 ..."},
	}
	handleReq()
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:content`
}

var Articles []Article
