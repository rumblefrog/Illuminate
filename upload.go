package main

import (
	"fmt"
	"net/http"
)

func uploadController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}
