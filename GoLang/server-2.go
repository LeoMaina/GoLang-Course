package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is the home page")
	})
	log.Println("Starting server....")
	http.ListenAndServe(":8080", nil)
}
