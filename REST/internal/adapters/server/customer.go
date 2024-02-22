package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
)

func (a *Adapter) CreateCustomer(c *fiber.Ctx) error {
	customer := new(domain.Customer)
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	customer, err := domain.NewCustomer(customer)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Create Customer
	err = a.api.CreateCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&customer)

}

func (a *Adapter) GetCustomers(c *fiber.Ctx) error {
	customer := new(domain.Customer)

	err := a.api.GetCustomers(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&customer)

}

func (a *Adapter) GetCustomer(c *fiber.Ctx) error {
	customer := new(domain.Customer)

	customerID := c.Query("customerId")
	if customerID == "" {
		c.Status(fiber.StatusBadRequest).SendString("missing customerID")
	}

	err := a.api.GetCustomer(customerID, customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&customer)

}

func (a *Adapter) DeleteCustomer(c *fiber.Ctx) error {
	customer := new(domain.Customer)

	customerID := c.Query("customerId")
	if customerID == "" {
		c.Status(fiber.StatusBadRequest).SendString("missing customerID")
	}

	err := a.api.DeleteCustomer(customerID, customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(204).SendString("OK")
}
