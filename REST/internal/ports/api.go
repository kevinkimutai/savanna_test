package ports

import "github.com/kevinkimutai/savanna/rest/internal/application/domain"

type APIPort interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
}
