package server

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Adapter) AuthRouter(api fiber.Router) {
	api.Get("/login", a.Login)
	api.Get("/callback", a.Callback)
	api.Get("/user", a.User)
	api.Get("/logout", a.Logout)
}

func (a *Adapter) CustomerRouter(api fiber.Router) {
	api.Post("/", IsAuthenticated(a.store), a.CreateCustomer)
	api.Get("/", IsAuthenticated(a.store), a.GetCustomers)
	api.Get("/:customerId", IsAuthenticated(a.store), a.GetCustomer)
	api.Delete("/:customerId", IsAuthenticated(a.store), a.DeleteCustomer)
}

func (a *Adapter) OrderRouter(api fiber.Router) {
	api.Post("/", a.CreateOrder)
	api.Get("/", a.GetOrders)
	api.Get("/:orderId", a.GetOrder)

}
