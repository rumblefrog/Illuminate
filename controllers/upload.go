package controllers

import (
	"fmt"
	"net/http"
)

func UploadController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Test")

	h := r.MultipartForm.File["payload"]

	fmt.Println(*h.tmpfile)
}
