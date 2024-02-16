package api

import (
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
	"github.com/kevinkimutai/savanna/rest/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	a.db.CreateCustomer()

}
