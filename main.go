package main

import (
	"log"
	"net/http"

	"github.com/ardityawahyu/lemon/handler"
	"github.com/ardityawahyu/lemon/service"
)

func main() {
	services := service.InitializeDependency()
	http.HandleFunc("/login", handler.Login(services))
	http.HandleFunc("/add_user", handler.AddUser(services))
	http.HandleFunc("/update_user", handler.UpdateUser(services))
	http.HandleFunc("/get_user", handler.GetUser(services))

	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":2233", nil))
}
