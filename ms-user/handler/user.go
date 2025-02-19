package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/repository"
	"github.com/gorilla/mux"
)

type UserHandler interface {
	GetUserList(w http.ResponseWriter, r *http.Request)
	GetUserByUsername(w http.ResponseWriter, r *http.Request)
	RegisterNewUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	repo repository.UserRepository
}

func InitUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{repo: repo}
}

func (h *userHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users := h.repo.FindAllUser()
	json.NewEncoder(w).Encode(users)
}

func (h *userHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, _ := h.repo.FindUserByUsername(username)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
