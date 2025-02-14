package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dapidsurya/cakap-aws-internship/model/entity"
	"github.com/gorilla/mux"
)

var users = []entity.User{
	{Username: "dapid", Fullname: "David", Email: "dapid@example.com"},
	{Username: "joko", Fullname: "Joko Santoso", Email: "joko@example.com"},
	{Username: "budi", Fullname: "Budi Santoso", Email: "budi@example.com"},
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	for _, user := range users {
		if user.Username == username {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}
