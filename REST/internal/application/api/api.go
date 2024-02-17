package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna/rest/internal/application/domain"
	"github.com/kevinkimutai/savanna/rest/internal/ports"
)

type Application struct {
	db   ports.DBPort
	auth ports.AuthPort
}

func NewApplication(db ports.DBPort, auth ports.AuthPort) *Application {
	return &Application{db: db, auth: auth}
}

func (a Application) Login(c *fiber.Ctx) error {
	res := a.auth.Login(c)

	return res
}

func (a Application) Callback(c *fiber.Ctx) error {
	res := a.auth.Callback(c)

	return res
}

func (a Application) CreateCustomer(customer *domain.Customer) error {
	//a.db.CreateCustomer()
	return errors.New("Not defined Yet")

}
