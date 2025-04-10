package router

import (
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	apiPort := fmt.Sprintf(":%d", config.LoadConfigIni("server", "port", 8080).(int))

	r := gin.Default()
	RegisterRoutes(r)
	r.Run(apiPort)

	return r
}
