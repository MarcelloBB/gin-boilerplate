package model

// swagger:model
type User struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement json:"id"`
	// example: johndoe
	Username string `json:"username"`
	// example: foo@bar.com
	Email string `json:"email"`
}
