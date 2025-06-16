package models

import (
	"errors"
	"slices"

	"github.com/Adrwaan/stock-api-go/config"

	"github.com/google/uuid"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Qtd  uint64 `json:"qtd"`
}

var products []Product
var logger *config.Logger = config.GetLogger()

func GetProducts() []Product {
	return products
}

func GetProductWithID(id string) (product Product, err error) {
	for _, p := range products {
		if p.ID == id {
			return p, err
		}
	}

	return product, errors.New("não foi encontrado um produto com o id informado")
}

func AddNewProduct(name string, qtd uint64) (product Product, err error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Error(err.Error())
	}

	product = Product{
		ID:   uuid.String(),
		Name: name,
		Qtd:  qtd,
	}

	products = append(products, product)

	return
}

func UpdateProduct(id, name string, qtd uint64) (product Product, err error) {
	for i, p := range products {
		if p.ID == id {
			if len(name) > 0 {
				products[i].Name = name
				products[i].Qtd = qtd
				product = products[i]

				return
			}
		}
	}

	err = errors.New("não existe um produto com este id")
	return
}

func DeleteProduct(id string) (err error) {
	for i, p := range products {
		if p.ID == id {
			products = slices.Delete(products, i, i+1)
			return
		}
	}

	err = errors.New("não existe um produto com este id")
	return
}
