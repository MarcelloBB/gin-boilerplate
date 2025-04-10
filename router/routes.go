package router

import (
	"github.com/MarcelloBB/gin-boilerplate/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controller.GetUsers)
		api.GET("/products", controller.GetProducts)
	}
}
