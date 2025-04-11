package model

// swagger:model
type Client struct {
	// example: 1
	ID int `gorm:"primaryKey" json:"id"`
	// example: company1
	Code string `json:"code"`
}
