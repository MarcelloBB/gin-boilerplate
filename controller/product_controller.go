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

// GetProducts godoc
// @Summary      List products
// @Description  Returns a list of products
// @Tags         product
// @Produce      json
// @Success      200  {array}   model.Product
// @Failure      500  {object}  map[string]string
// @Router       /product [get]
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		fmt.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
	}

	c.JSON(http.StatusOK, products)
}
