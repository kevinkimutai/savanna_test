package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

func (a *Adapter) CreateOrder(c *fiber.Ctx) error {
	order := new(domain.Order)
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	order, err := domain.NewOrder(order)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Create Order
	err = a.api.CreateOrder(order)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(order)

}
