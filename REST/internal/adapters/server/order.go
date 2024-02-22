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

func (a *Adapter) GetOrders(c *fiber.Ctx) error {
	order := new(domain.Order)

	err := a.api.GetOrders(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&order)
}

func (a *Adapter) GetOrder(c *fiber.Ctx) error {
	order := new(domain.Customer)

	orderID := c.Query("customerId")
	if orderID == "" {
		c.Status(fiber.StatusBadRequest).SendString("missing orderID")
	}

	err := a.api.GetCustomer(orderID, order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&order)
}
