package main

import (
	"go-auth-starter/app/config"
	"go-auth-starter/app/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load .env vars
	load_env_vars()

	e := echo.New()
	config.ServerConfig(e)

	// Setup Routes
	router.SetupRouter(e)

	// Start and log
	e.Logger.Fatal(e.Start(":3000"))
}

func load_env_vars() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
