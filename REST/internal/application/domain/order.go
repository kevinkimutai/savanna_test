package domain

import (
	"errors"
	"time"
)

type Order struct {
	ID         int         `json:"id"`
	CustomerID string      `json:"customer"`
	Items      []OrderItem `json:"items"`
	Total      float64     `json:"total"`
	CreatedAt  time.Time   `json:"created_at"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  string `json:"quantity"`
}

func NewOrder(order *Order) (*Order, error) {
	if order.CustomerID == "" || len(order.Items) == 0 {
		return order, errors.New("missing order values")
	}

	return order, nil
}
