package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	dev := os.Getenv("DEVELOPMENT")
	var port int
	if dev == "1" {
		port = 8080
	} else {
		port = 5000
	}
	fmt.Println(fmt.Sprintf("Working on port: %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func main() {
	handleRequests()
}
