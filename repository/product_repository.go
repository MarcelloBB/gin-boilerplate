package repository

import (
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	products := []model.Product{}
	err := p.connection.Debug().Find(&products).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	fmt.Println("Products found:", products)

	return products, nil
}
