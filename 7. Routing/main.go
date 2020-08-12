package main

import (
	"fmt"
	"net/http"
)

func ActionUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Auth Berhasil")
}

func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/users", ActionUsers)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":8002"
	server.Handler = mux

	fmt.Println("server started")
	server.ListenAndServe()
}
