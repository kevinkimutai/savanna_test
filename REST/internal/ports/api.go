package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

type APIPort interface {
	//Customer
	CreateCustomer(customer *domain.Customer) error
	GetCustomers(customer *domain.Customer) error
	GetCustomer(customerID string, customer *domain.Customer) error
	DeleteCustomer(customerID string, customer *domain.Customer) error

	//Auth
	Login(fiber *fiber.Ctx, store *session.Store) error
	Callback(fiber *fiber.Ctx, store *session.Store) error

	//Orders
	CreateOrder(order *domain.Order) error
	GetOrders(order *domain.Order) error
	GetOrder(orderID string, order *domain.Order) error

	//SMS
	SendSMS(msg string, phoneNumbers []string) (string, error)
}
