package router

import (
	"database/sql"

	"github.com/MarcelloBB/gin-boilerplate/controller"
	"github.com/MarcelloBB/gin-boilerplate/repository"
	"github.com/MarcelloBB/gin-boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConnection *sql.DB) {
	// Initialize repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Initialize use cases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	UserUseCase := usecase.NewUserUseCase()

	// Initialize controllers
	ProductController := controller.NewProductController(ProductUseCase)
	UserController := controller.NewUserController(UserUseCase)

	api := r.Group("/api")
	{
		api.GET("/users", UserController.GetUsers)
		api.GET("/products", ProductController.GetProducts)
	}
}
