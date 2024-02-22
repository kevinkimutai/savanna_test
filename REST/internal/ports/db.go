package ports

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

type DBPort interface {
	CreateOrder(order *domain.Order) error

	//Customers
	CreateCustomer(customer *domain.Customer) error
	GetCustomers(customer *domain.Customer) error
	GetCustomer(customer *domain.Customer) error
}
