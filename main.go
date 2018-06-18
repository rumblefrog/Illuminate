package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RumbleFrog/Illuminate/controllers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Illuminate the world")
	})

	http.HandleFunc("/upload", controllers.UploadController)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
