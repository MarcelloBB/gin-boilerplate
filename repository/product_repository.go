package repository

import (
	"database/sql"
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price, quantity FROM product"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObt model.Product

	for rows.Next() {
		err := rows.Scan(&productObt.ID, &productObt.Name, &productObt.Price, &productObt.Quantity)
		if err != nil {
			return []model.Product{}, err
		}
		productList = append(productList, productObt)
	}

	defer rows.Close()

	return productList, nil
}
