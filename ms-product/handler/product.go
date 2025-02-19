package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/repository"
	"github.com/gorilla/mux"
)

type ProductHandler interface {
	GetProductByUserId(w http.ResponseWriter, r *http.Request)
	CreateNewProduct(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	repo repository.ProductRepository
}

func InitProductHandler(repo repository.ProductRepository) ProductHandler {
	return &productHandler{repo: repo}
}

func (h *productHandler) GetProductByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	products := h.repo.FindUserByUserId(userId)
	json.NewEncoder(w).Encode(products)
}

func (h *productHandler) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateProduct(&product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
