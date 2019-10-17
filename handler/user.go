package handler

import (
	"net/http"

	"github.com/ardityawahyu/lemon/service"
)

func AddUser(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: call user add service here
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateUser(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: call user update service here
		w.WriteHeader(http.StatusOK)
	}
}
