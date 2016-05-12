package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version string

func main() {
	message := fmt.Sprintf("Running version %s", GetVersion())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetVersion() string {
	return os.Getenv("VERSION")
}
