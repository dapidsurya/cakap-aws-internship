package main

import (
	"fmt"
	"net/http"

	"github.com/dapidsurya/cakap-aws-internship/handler"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Cakap x AWS Academy!")
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", hello)
	route.HandleFunc("/user/list", handler.GetUserList).Methods("GET")
	route.HandleFunc("/user/{username}", handler.GetUserByUsername).Methods("GET")

	fmt.Println("Server is listening on port 8081...")

	err := http.ListenAndServe(":8081", route)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
