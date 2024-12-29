package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	// set up configuration
	fmt.Println("Server is running on port", port)

	http.HandleFunc("/", Hello)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}