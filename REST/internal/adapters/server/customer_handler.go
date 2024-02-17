package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

func (a Adapter) CreateCustomer(c *fiber.Ctx) error {
	customer := new(domain.Customer)
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	customer, err := domain.NewCustomer(customer)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Create Customer
	return a.api.CreateCustomer(customer)

}
