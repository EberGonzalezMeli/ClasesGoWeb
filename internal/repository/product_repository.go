package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var Products []Product

func ReadJson(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Products)
	if err != nil {
		return err
	}
	log.Println("Productos cargados: ", len(Products))
	return nil
}

func GetAll() []Product {
	return Products
}

func GetByID(id int) (*Product, error) {
	for _, item := range Products {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("error: No exite un producto con este ID")
}

func SearchByPrice(priceGt float64) ([]Product, error) {
	var result []Product
	for _, p := range Products {
		if p.Price > priceGt {
			result = append(result, p)
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, fmt.Errorf("error: No exite un producto con precio mayor a %v", priceGt)
}
