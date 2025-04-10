package repository

import (
	"database/sql"
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		connection: db,
	}
}

func (p *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, username, email FROM userr"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.User{}, err
	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err := rows.Scan(&userObj.ID, &userObj.Username, &userObj.Email)
		if err != nil {
			return []model.User{}, err
		}
		userList = append(userList, userObj)
	}

	defer rows.Close()

	return userList, nil
}
