package main

import (
	"fmt"
	"net/http"
)

func uploadController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprint(w, "Test")

	fmt.Println(r.PostForm["key"][0])
}
