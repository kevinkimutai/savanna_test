package db

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

func (a Adapter) CreateOrder(order *domain.Order) error {
	err := a.db.Create(&order).Error

	return err
}

func (a Adapter) GetOrders(order *domain.Order) error {
	err := a.db.Find(&order).Error

	return err
}
func (a Adapter) GetOrder(orderID string, order *domain.Order) error {
	err := a.db.First(&order, orderID).Error

	return err
}
