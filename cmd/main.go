package main

import (
	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/MarcelloBB/gin-boilerplate/router"
)

func main() {
	config.InitializeConfig()
	router.InitRouter()
}
