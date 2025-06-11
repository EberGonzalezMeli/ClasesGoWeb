package handler

import (
	"encoding/json"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RequestBodyProduct struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data    *struct {
		Name         string  `json:"name"`
		Quantity     int     `json:"quantity"`
		Code_value   string  `json:"code_value"`
		Is_published bool    `json:"is_published"`
		Expiration   string  `json:"expiration"`
		Price        float64 `json:"price"`
	} `json:"data"`
	Error bool `json:"error"`
}

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
	products, err := service.SearchByPrice(priceGt)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(products)

}
