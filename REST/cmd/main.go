package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/savanna/rest/internal/adapters/authenticator"
	"github.com/kevinkimutai/savanna/rest/internal/adapters/db"
	"github.com/kevinkimutai/savanna/rest/internal/adapters/server"
	"github.com/kevinkimutai/savanna/rest/internal/application/api"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {

		log.Fatal("error loading .env file", err)
	}

	//ENV variables
	DBURL := os.Getenv("POSTGRES_URL")
	PORT := os.Getenv("PORT")

	dbAdapter, err := db.NewAdapter(DBURL)
	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	application := api.NewApplication(dbAdapter, auth)

	server := server.NewAdapter(application, PORT)

	//Start Server
	server.Run()

}
