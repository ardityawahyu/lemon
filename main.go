package main

import
	"handler"

func main() {
	http.HandleFunc("/login", handler.login)

	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":2233", nil))
​
}