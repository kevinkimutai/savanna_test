package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthPort interface {
	Login(c *fiber.Ctx, store *session.Store) error
	Callback(c *fiber.Ctx, store *session.Store) error
}
