package usecase

import "github.com/MarcelloBB/gin-boilerplate/model"

type UserUseCase struct {
	// UserRepository repository.UserRepository
}

func NewUserUseCase() UserUseCase {
	return UserUseCase{}
}

func (p *UserUseCase) GetUsers() ([]model.User, error) {
	return []model.User{}, nil
}
