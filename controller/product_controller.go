package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcelloBB/gin-boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		fmt.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
	}

	c.JSON(http.StatusOK, products)
}
