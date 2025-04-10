package routes

import (
	"github.com/MarcelloBB/gin-boilerplate/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", handler.GetUsers)
	}
}
