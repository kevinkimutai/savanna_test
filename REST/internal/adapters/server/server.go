package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinkimutai/savanna/rest/internal/ports"
)

type Adapter struct {
	port string
	api  ports.APIPort
}

func NewAdapter(api ports.APIPort, port string) *Adapter {
	return &Adapter{port: port}
}

func (a Adapter) Run() {
	app := fiber.New()

	//LOGGER MIDDLEWARE
	app.Use(logger.New())

	//API ROUTES
	app.Route("/api/v1/customer", a.CustomerRouter)

	app.Listen(":" + a.port)

}
