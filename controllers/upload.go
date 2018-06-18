package controllers

import (
	"fmt"
	"net/http"
)

func UploadController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Test")

	fmt.Println(r.MultipartForm)
}
