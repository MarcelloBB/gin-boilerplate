package controller

import (
	"net/http"

	"github.com/MarcelloBB/gin-boilerplate/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []model.User{
		{
			ID:       1,
			Username: "user1",
			Email:    "L0lOw@example.com",
		},
	}
	c.JSON(http.StatusOK, users)
}
