package middlewares

import "net/http"

func IsPost(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "POST Only", http.StatusMethodNotAllowed)
	}
}
