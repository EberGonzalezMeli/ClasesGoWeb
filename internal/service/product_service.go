package service

import (
	"main/internal/repository"
)

func GetProducts() []repository.Product {
	return repository.GetAll()
}

func GetProductID(id int) *repository.Product {
	produc, err := repository.GetByID(id)
	if err != nil {
		panic("error: No exite un producto con este ID")
	}
	return produc
}

func SearchByPrice(priceGt float64) []repository.Product {
	result, err := repository.SearchByPrice(priceGt)
	if err != nil {
		panic("error: No exite un producto con precio mayor al ingresado")
	}
	return result

}
