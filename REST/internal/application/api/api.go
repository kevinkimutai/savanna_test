package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
	"github.com/kevinkimutai/savanna/rest/internal/ports"
)

type Application struct {
	db   ports.DBPort
	auth ports.AuthPort
}

func NewApplication(db ports.DBPort, auth ports.AuthPort) *Application {
	return &Application{db: db, auth: auth}
}

func (a Application) Login(c *fiber.Ctx, store *session.Store) error {
	res := a.auth.Login(c, store)

	return res
}

func (a Application) Callback(c *fiber.Ctx, store *session.Store) error {
	res := a.auth.Callback(c, store)

	return res
}

func (a Application) CreateOrder(order *domain.Order) error {

	response := a.db.CreateOrder(order)

	return response

}

// Customers
func (a Application) CreateCustomer(customer *domain.Customer) error {
	response := a.db.CreateCustomer(customer)

	return response

}

func (a Application) GetCustomers(customer *domain.Customer) error {
	response := a.db.GetCustomers(customer)

	return response

}
