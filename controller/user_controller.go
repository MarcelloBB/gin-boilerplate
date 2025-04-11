package controller

import (
	"net/http"

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

// GetUsers godoc
// @Summary      List users
// @Description  Returns a list of users
// @Tags         user
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      500  {object}  map[string]string
// @Router       /user [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
