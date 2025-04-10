package router

import (
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/MarcelloBB/gin-boilerplate/db"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	server := gin.Default()

	apiPort := fmt.Sprintf(":%d", config.LoadConfigIni("server", "port", 8080).(int))
	dbConnection, err := db.ConnectDB()

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}

	RegisterRoutes(server, dbConnection)
	server.Run(apiPort)

	return server
}
