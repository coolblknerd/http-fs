package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "data")
}
func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
