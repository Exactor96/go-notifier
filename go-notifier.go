package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	dev := os.Getenv("DEVELOPMENT")
	var port int
	if dev == "1" {
		port = 8080
	} else {
		envport := os.Getenv("PORT")
		port, _ = strconv.Atoi(envport)
	}
	fmt.Println(fmt.Sprintf("Working on port: %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func main() {
	handleRequests()
}
