package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

type APIPort interface {
	CreateCustomer(customer *domain.Customer) error
	Login(fiber *fiber.Ctx) error
	Callback(fiber *fiber.Ctx) error
}
