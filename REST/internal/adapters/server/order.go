package server

import (
	"fmt"
	"log/slog"

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
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Create Order
	err = a.api.CreateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	orderMessage := fmt.Sprintf("Dear Kevin, Order placement successful at %v, Your Order ID: %v", order.CreatedAt, order.ID)

	//Send SMS
	//TODO:GET CUSTOMERS DETAILS & PHONE_NUMBER
	message, err := a.api.SendSMS(orderMessage, []string{"+254722670831"})
	if err != nil {
		slog.Error("SMS", "sms", err)
	}

	slog.Info("SMS", "sms", message)

	return c.Status(fiber.StatusCreated).JSON(&order)

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
		return c.Status(fiber.StatusBadRequest).SendString("missing orderID")
	}

	err := a.api.GetCustomer(orderID, order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&order)
}
