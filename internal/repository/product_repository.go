package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

func CreateProduct(newProduct Product) error {
	err := ValidateDataPost(newProduct)
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}

	newProduct.ID = 1
	for _, producto := range Products {
		newProduct.ID = producto.ID + 1
	}

	Products = append(Products, newProduct)
	return nil
}
func ValidateDataPost(newProduct Product) error {
	formato := "02/01/2006" // formato dd/mm/yyyy

	if newProduct.Name == "" {
		return fmt.Errorf("nombre es requerido")
	}
	if newProduct.Quantity == 0 {
		return fmt.Errorf("quantity es requerido y diferente de 0")
	}
	if strings.TrimSpace(newProduct.Code_value) == "" {
		return fmt.Errorf("code_value es requerido")
	}
	for _, producto := range Products {
		if newProduct.Code_value == producto.Code_value {
			return fmt.Errorf("ya existe un Code_value con este valor")
		}
	}

	if strings.TrimSpace(newProduct.Expiration) == "" {
		return fmt.Errorf("expiration es requerido")
	}

	_, err := time.Parse(formato, newProduct.Expiration)
	if err != nil {
		return fmt.Errorf("la fecha NO tiene el formato dd/mm/yyyy")
	}

	if newProduct.Price == 0 {
		return fmt.Errorf("price es requerido y diferente de 0")
	}

	return nil
}
