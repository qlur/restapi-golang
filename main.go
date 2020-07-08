package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func citiesData(w http.ResponseWriter, r *http.Request) {
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"

	fmt.Fprintf(w, cities[1])

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/Venice", citiesData)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}
