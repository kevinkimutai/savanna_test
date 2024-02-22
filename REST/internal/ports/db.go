package ports

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

type DBPort interface {
	CreateOrder(order *domain.Order) error
	GetOrders(order *domain.Order) error
	GetOrder(orderID string, order *domain.Order) error

	//Customers
	CreateCustomer(customer *domain.Customer) error
	GetCustomers(customer *domain.Customer) error
	GetCustomer(customerID string, customer *domain.Customer) error
	DeleteCustomer(customerID string, customer *domain.Customer) error
}
