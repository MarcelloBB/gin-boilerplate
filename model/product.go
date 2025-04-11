package model

type Product struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: "product_1"
	Name string `json:"name"`
	// example: 50.25
	Price float64 `json:"price"`
	// example: 10
	Quantity *int `json:"quantity"`
}
