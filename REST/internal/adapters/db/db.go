package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct{}

type Order struct {
	gorm.Model
	CustomerID Customer    `json:"customer"`
	Items      []OrderItem `json:"items"`
	Total      float64     `json:"total"`
}

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  string `json:"quantity"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {

	db, openErr := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Customer{}, &Order{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
