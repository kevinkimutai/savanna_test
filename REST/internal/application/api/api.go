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
	sms  ports.SMSPort
}

func NewApplication(db ports.DBPort, auth ports.AuthPort, sms ports.SMSPort) *Application {
	return &Application{db: db, auth: auth, sms: sms}
}

// Auth
func (a Application) Login(c *fiber.Ctx, store *session.Store) error {
	res := a.auth.Login(c, store)

	return res
}

func (a Application) Callback(c *fiber.Ctx, store *session.Store) error {
	res := a.auth.Callback(c, store)

	return res
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

func (a Application) GetCustomer(customerID string, customer *domain.Customer) error {
	response := a.db.GetCustomer(customerID, customer)

	return response

}
func (a Application) DeleteCustomer(customerID string, customer *domain.Customer) error {
	response := a.db.DeleteCustomer(customerID, customer)

	return response
}

// Orders
func (a Application) CreateOrder(order *domain.Order) error {
	response := a.db.CreateOrder(order)

	return response

}
func (a Application) GetOrders(order *domain.Order) error {
	response := a.db.GetOrders(order)

	return response

}
func (a Application) GetOrder(orderID string, order *domain.Order) error {
	response := a.db.GetOrder(orderID, order)

	return response
}
