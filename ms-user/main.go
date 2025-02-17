package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/config"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	route := mux.NewRouter()

	route.HandleFunc("/", handler.Hello)
	route.HandleFunc("/user/list", handler.GetUserList).Methods("GET")
	route.HandleFunc("/user/{username}", handler.GetUserByUsername).Methods("GET")
	route.HandleFunc("/user", handler.RegisterNewUser).Methods("POST")

	fmt.Println("Server is listening on port 8081...")

	err = http.ListenAndServe(":8081", route)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
