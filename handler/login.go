package handler

import (
	"net/http"

	"github.com/ardityawahyu/lemon/service"
)

func Login(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: call login service here
		w.WriteHeader(http.StatusOK)
	}
}
