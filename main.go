package main

import
	"github.com/ardityawahyu/lemon/handler"
	"github.com/ardityawahyu/lemon/service"

func main() {
	services := service.InitializeDependency()
	http.HandleFunc("/login", handler.login)

	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":2233", nil))
​
}