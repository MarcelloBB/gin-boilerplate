package usecase

import (
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
	return p.repository.GetProducts()
	// if err != nil {
	// 	return []model.Product{}, err
	// }
	// return products, nil
}
