package ports

import "github.com/gofiber/fiber/v2"

type AuthPort interface {
	Login(c *fiber.Ctx) error
	Callback(c *fiber.Ctx) error
}
