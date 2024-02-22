package db

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

func (a Adapter) CreateOrder(order *domain.Order) error {

	err := a.db.Create(&order).Error

	return err
}
