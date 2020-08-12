package main

import (
	"fmt"
	"net/http"
)

func ActionUsers(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	fmt.Println("Auth Berhasil")
}

func main() {
	http.HandleFunc("/users", ActionUsers)

	server := new(http.Server)
	server.Addr = ":8002"

	fmt.Println("server started")
	server.ListenAndServe()
}
