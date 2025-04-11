package router

import (
	"database/sql"

	"github.com/MarcelloBB/gin-boilerplate/controller"
	"github.com/MarcelloBB/gin-boilerplate/docs"
	"github.com/MarcelloBB/gin-boilerplate/repository"
	"github.com/MarcelloBB/gin-boilerplate/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, dbConnection *sql.DB) {
	// Initialize repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	UserRepository := repository.NewUserRepository(dbConnection)

	// Initialize use cases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	UserUseCase := usecase.NewUserUseCase(UserRepository)

	// Initialize controllers
	ProductController := controller.NewProductController(ProductUseCase)
	UserController := controller.NewUserController(UserUseCase)

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// @BasePath /api/v1
	api := r.Group(basePath)
	{
		api.GET("/user", UserController.GetUsers)
		api.GET("/product", ProductController.GetProducts)
	}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
