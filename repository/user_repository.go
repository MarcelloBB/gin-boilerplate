package repository

import (
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		connection: db,
	}
}

func (p *UserRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}
	err := p.connection.Find(&users).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.User{}, err
	}

	return users, nil
}
