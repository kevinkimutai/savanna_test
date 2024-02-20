package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kevinkimutai/savanna/rest/internal/ports"
)

type Adapter struct {
	port  string
	api   ports.APIPort
	store *session.Store
}

func NewAdapter(api ports.APIPort, port string) *Adapter {
	return &Adapter{api: api, port: port, store: session.New()}
}

func (a *Adapter) GetSessionStore() *session.Store {
	return a.store
}

func (a *Adapter) Run() {
	app := fiber.New()
	//LOGGER MIDDLEWARE
	app.Use(logger.New())

	// Define routes
	app.Route("/api/v1/auth", a.AuthRouter)
	app.Route("/api/v1/customer", a.CustomerRouter)
	app.Route("/api/v1/order", a.OrderRouter)

	app.Listen(":" + a.port)

}
