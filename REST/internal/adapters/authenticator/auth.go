package authenticator

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Authenticator) Login(c *fiber.Ctx, store *session.Store) error {

	state, err := generateRandomState()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sess.Set("state", state)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	slog.Info("MYSTATE", "state", sess.Get("state"))

	authURL := a.AuthCodeURL(state)

	// Redirect with the generated state
	return c.Status(fiber.StatusTemporaryRedirect).JSON(authURL)

}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

// Handler for our callback.
func (a *Authenticator) Callback(c *fiber.Ctx, store *session.Store) error {

	// Retrieve the session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	slog.Info("Session", "store", store)
	slog.Info("State", "state", sess.Get("state"))
	// Validate the state parameter
	//state := c.Query("state")
	// if state != sess.Get("state") {
	// 	return c.Status(fiber.StatusBadRequest).SendString("Invalid state parameter.")
	// }

	// Exchange authorization code for a token
	token, err := a.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Failed to exchange an authorization code for a token.")
	}

	// Verify ID token
	idToken, err := a.VerifyIDToken(c.Context(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to verify ID Token.")
	}

	// Extract user profile from ID token claims
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Save access token and user profile in session
	sess.Set("access_token", token.AccessToken)
	sess.Set("profile", profile)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Redirect to the logged-in page
	return c.Redirect("/api/v1/auth/user", fiber.StatusTemporaryRedirect)
}
