package handler

import (
	"encoding/json"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	prodcts := repository.GetAll()
	json.NewEncoder(w).Encode(prodcts)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, err := repository.GetByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)

}

func GetSearchByPriceHandler(w http.ResponseWriter, r *http.Request) {

	priceGtStr := r.URL.Query().Get("priceGt")
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		http.Error(w, "valor invalido para priceGt", http.StatusBadRequest)
		return
	}
	products := service.SearchByPrice(priceGt)
	json.NewEncoder(w).Encode(products)

}
