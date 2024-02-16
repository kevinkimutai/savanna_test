package server

import (
	"github.com/gofiber/fiber/v2"
)

func (a Adapter) CustomerRouter(api fiber.Router) {
	api.Post("/", a.CreateCustomer)
}
