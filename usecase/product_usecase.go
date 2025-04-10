package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/db"
	"github.com/MarcelloBB/gin-boilerplate/model"
	"github.com/MarcelloBB/gin-boilerplate/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {
	cacheValue, err := db.GetCacheValue(db.Ctx, "products")
	if err == nil && cacheValue != "" {
		var productsResult []model.Product
		err = json.Unmarshal([]byte(cacheValue), &productsResult)

		if err != nil {
			fmt.Println("Error unmarshalling cache:", err)
		} else {
			return productsResult, nil
		}
	}

	products, err := p.repository.GetProducts()
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return []model.Product{}, err
	}

	productsSerialized, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error serializing products:", err)
	} else {
		if err := db.SetCacheValue(db.Ctx, "products", productsSerialized); err != nil {
			fmt.Printf("Error setting cache: %v\n", err)
		}
	}

	return products, nil
}
