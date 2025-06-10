package main

import (
	"log"
	"main/internal/handler"
	"main/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := repository.ReadJson("/Users/ebgonzalez/Documents/BootCamp/BasicoWeb/Ejercicios Clase 2/docs/db/products.json")
	if err != nil {
		log.Fatal("No se pudo leer productos:", err)
	}

	rt := chi.NewRouter()

	rt.Get("/products", handler.GetProductsHandler)
	rt.Get("/products/{id}", handler.GetProductByIDHandler)
	rt.Get("/products/search", handler.GetSearchByPriceHandler)

	log.Println("Servidor escuchando en :8080")
	http.ListenAndServe(":8080", rt)
}
