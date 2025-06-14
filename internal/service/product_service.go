package service

import (
	"fmt"
	"main/internal/repository"
)

func GetProducts() []repository.Product {
	return repository.GetAll()
}

func GetProductID(id int) (*repository.Product, error) {
	produc, err := repository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error: No exite un producto con este ID")
	}
	return produc, nil
}

func SearchByPrice(priceGt float64) ([]repository.Product, error) {
	result, err := repository.SearchByPrice(priceGt)
	if err != nil {
		return nil, fmt.Errorf("error: No exite un producto con precio mayor al ingresado")
	}
	return result, nil

}

func PostProduct(newProduct *repository.Product) error {

	err := repository.CreateProduct(*newProduct)
	if err != nil {
		return fmt.Errorf("error: No se pudo crear un producto: %v", err.Error())
	}
	return nil
}
