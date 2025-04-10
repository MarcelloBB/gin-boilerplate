package router

import (
	"github.com/MarcelloBB/gin-boilerplate/routes"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run()

	return r
}
