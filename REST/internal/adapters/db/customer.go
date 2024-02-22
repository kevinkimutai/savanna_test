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
func (a Adapter) GetCustomer(customerID string, customer *domain.Customer) error {
	err := a.db.First(&customer, customerID).Error

	return err
}

func (a Adapter) DeleteCustomer(customerID string, customer *domain.Customer) error {
	err := a.db.Delete(&customer, customerID).Error

	return err
}
