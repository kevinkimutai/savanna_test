package db

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

func (a Adapter) CreateCustomer(customer *domain.Customer) error {
	err := a.db.Create(&customer).Error

	return err
}

func (a Adapter) GetCustomers(customer *domain.Customer) error {
	err := a.db.Find(&customer).Error

	return err
}
