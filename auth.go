package main

import (
	"net/http"
)

func isFishy(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		key, exist := r.PostForm["key"]

		if !exist {
			http.Error(w, "Key Unidentified", http.StatusBadRequest)
			return
		}

		if key[0] != "fishyisagod" {
			http.Error(w, "Key Invalid", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}
