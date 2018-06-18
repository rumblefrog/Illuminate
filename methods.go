package main

import "net/http"

func isPost(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "POST Only", http.StatusMethodNotAllowed)
	}
}
