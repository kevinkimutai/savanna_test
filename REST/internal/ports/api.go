package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

type APIPort interface {
	CreateCustomer(customer *domain.Customer) error
	Login(fiber *fiber.Ctx, store *session.Store) error
	Callback(fiber *fiber.Ctx, store *session.Store) error
	CreateOrder(order *domain.Order) error
}
