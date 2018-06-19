package controllers

import (
	"fmt"
	"net/http"
)

func UploadController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Test")

	// _, _, _ := r.FormFile("payload")

	// fmt.Println(file.)
}
