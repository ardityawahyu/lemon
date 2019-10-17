package handler

import (
	"net/http"

	"github.com/ardityawahyu/lemon/service"
)

func Login(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")

			if s.Login.Login(username, password) {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}
}
