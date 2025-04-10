package main

import (
	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/MarcelloBB/gin-boilerplate/db"
	"github.com/MarcelloBB/gin-boilerplate/router"
)

func main() {
	// Load config-file
	config.InitializeConfig()
	// Start Redis connection
	db.InitRedis()
	// Start api server
	router.InitRouter()
}
