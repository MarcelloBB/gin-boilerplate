package controller

import (
	"net/http"

	"github.com/MarcelloBB/gin-boilerplate/model"
	"github.com/MarcelloBB/gin-boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		userUseCase: usecase,
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users := []model.User{
		{
			ID:       1,
			Username: "User 1",
		},
	}
	c.JSON(http.StatusOK, users)
}
