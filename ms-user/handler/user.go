package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/repository"
	"github.com/gorilla/mux"
)

var users = []entity.User{
	{Username: "dapid", Fullname: "David", Email: "dapid@example.com"},
	{Username: "joko", Fullname: "Joko Santoso", Email: "joko@example.com"},
	{Username: "budi", Fullname: "Budi Santoso", Email: "budi@example.com"},
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository()
	users := repo.FindAllUser()
	json.NewEncoder(w).Encode(users)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository()
	vars := mux.Vars(r)
	username := vars["username"]

	user, _ := repo.FindUserByUsername(username)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository()
	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := repo.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
