package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Illuminate the world")
	})

	http.HandleFunc("/upload", uploadController)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
