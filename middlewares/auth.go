package middlewares

import (
	"net/http"
)

func IsFishy(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(20000000) // 20 MB memory space for files

		key := r.FormValue("key")

		if key != "fishyisagod" {
			http.Error(w, "Key Invalid", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}
