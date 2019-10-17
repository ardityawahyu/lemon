package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ardityawahyu/lemon/repo"
	"github.com/ardityawahyu/lemon/service"
)

func GetUser(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			idString := r.FormValue("id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			um := s.User.GetUser(id)
			if reflect.DeepEqual(um, repo.UserModel{}) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(um)
			w.Write(reqBodyBytes.Bytes())
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func AddUser(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			um := repo.UserModel{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
				Email:    r.FormValue("email"),
			}

			id := s.User.AddUser(um)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%d", id)))
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func UpdateUser(s service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			idString := r.FormValue("ID")
			id, err := strconv.Atoi(idString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			um := repo.UserModel{
				ID:       id,
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
				Email:    r.FormValue("email"),
			}

			if s.User.UpdateUser(um) {
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
