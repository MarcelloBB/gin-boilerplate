package controller

import (
	"net/http"

	"github.com/MarcelloBB/gin-boilerplate/model"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := []model.Product{
		{
			ID:       1,
			Name:     "Product 1",
			Price:    10.99,
			Quantity: 5,
		},
	}
	c.JSON(http.StatusOK, products)
}
