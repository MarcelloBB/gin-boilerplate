package usecase

import (
	"github.com/MarcelloBB/gin-boilerplate/model"
	"github.com/MarcelloBB/gin-boilerplate/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (p *UserUseCase) GetUsers() ([]model.User, error) {
	return p.repository.GetUsers()
}
