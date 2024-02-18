package server

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (a Adapter) Login(c *fiber.Ctx) error {

	response := a.api.Login(c)

	return response
}
func (a Adapter) Callback(c *fiber.Ctx) error {
	response := a.api.Callback(c)

	return response
}

func (a Adapter) User(c *fiber.Ctx) error {
	store := session.New()

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Retrieve the user's profile from the session
	profile := sess.Get("profile")

	// Do something with the user's profile
	if profile == nil {
		return c.Status(fiber.StatusNotFound).SendString("User profile not found in session")
	}

	// Return the user's profile as a response
	return c.JSON(profile)

}

func (a Adapter) Logout(c *fiber.Ctx) error {
	// Parse logout URL
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Determine scheme (http or https)
	scheme := "http"
	if c.Protocol() == "https" {
		scheme = "https"
	}

	// Parse returnTo URL
	returnTo, err := url.Parse(scheme + "://" + c.Hostname())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Add returnTo and client_id parameters
	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	// Redirect to logout URL
	return c.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}

func IsAuthenticated(c *fiber.Ctx) error {
	// Check if the user is authenticated
	if c.Locals("profile") == nil {
		// Redirect to the homepage if the user is not authenticated
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorised.")
	}

	// Call the next middleware or handler
	return c.Next()
}
