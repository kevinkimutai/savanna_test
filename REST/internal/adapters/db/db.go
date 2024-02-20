package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uint   `gorm:"primaryKey;uniqueIndex"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []Order
}

type Order struct {
	ID          uint     `gorm:"primaryKey;uniqueIndex"`
	CustomerID  uint     //`gorm:"uniqueIndex"`
	Customer    Customer `gorm:"foreignKey:CustomerID"`
	OrderNumber string
	OrderDate   time.Time
	Items       []Product `gorm:"many2many:order_products;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Total       float64 `json:"total"`
}

type Product struct {
	ID          uint   `gorm:"primaryKey;uniqueIndex"`
	ProductName string `json:"product_name"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItem struct {
	ProductID uint   `json:"product_id"`
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

	err := db.AutoMigrate(&Customer{}, &Order{}, &Product{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
