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
	prodcts := service.GetProducts()
	json.NewEncoder(w).Encode(prodcts)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, err := service.GetProductID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)

}

func GetSearchByPriceHandler(w http.ResponseWriter, r *http.Request) {

	priceGtStr := r.URL.Query().Get("priceGt")
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		http.Error(w, "priceGt tiene un formato no valido", http.StatusBadRequest)
		return
	}
	products, err := service.SearchByPrice(priceGt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(products)

}

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product *repository.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Formato invalido", http.StatusBadRequest)
		return
	}

	if err := service.PostProduct(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
