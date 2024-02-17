package server

import (
	"github.com/gofiber/fiber/v2"
)

func (a Adapter) authRouter(api fiber.Router) {
	api.Get("/login", a.Login)
	api.Get("/callback", a.Callback)
	api.Get("/user", a.User)
	api.Get("/logout", a.Logout)
}

func (a Adapter) CustomerRouter(api fiber.Router) {
	api.Post("/", a.CreateCustomer)
}
