package config

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Little bit of middlewares for housekeeping
func ServerConfig(e *echo.Echo) {
	e.Pre(echoMiddleware.RemoveTrailingSlash())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RateLimiter(echoMiddleware.NewRateLimiterMemoryStore(20)))
	e.Use(echoMiddleware.Secure())                     // TODO: Setup policy to protect site from XSS
	e.Use(echoMiddleware.CSRFWithConfig(csrfConfig())) // TODO: look at doing this manually?

	// TODO: Configure cors
	// e.Use(echoMiddleware.CORS())

	// Config Session Store
	ConfigSessionStore(e)

	// add database connection to request context
	e.Use(AddDbConnectionToContext)

	e.Logger.SetLevel(0)
}
