package router

import (
	"fmt"
	"go-auth-starter/app/config"
	"go-auth-starter/app/handlers"

	"os"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	assetPath := os.Getenv("ASSETS_PATH")

	// Serve CSS
	serveCSS(e, assetPath)
	serveFonts(e, assetPath)

	// Serve Javascript
	serveHTMX(e, assetPath)

	// Unauthed routes. User needs to be logged out
	unauthenticated := e.Group("", config.CheckUserNotAuthenticated)

	handlers.LoginHandlers(unauthenticated)
	handlers.RegisterHandlers(unauthenticated)

	// Authed routes. User needs to be logged in
	authenticated := e.Group("")
	authenticated.Use(config.CheckUserAuthentication)

	handlers.LogoutHandlers(authenticated)

	// Add workspace id to context for follow routes
	authenticated.Use(config.AddWorkspaceIdToContext)

	handlers.HomepageHandlers(authenticated)
	handlers.ProfileHandlers(authenticated)
}

func serveCSS(e *echo.Echo, assetPath string) {
	path := fmt.Sprintf("%s/css/output.css", assetPath)
	e.File("/css/output.css", path)
}

func serveFonts(e *echo.Echo, assetPath string) {
	path := fmt.Sprintf("%s/fonts/inter-4.0/inter.css", assetPath)
	e.File("/fonts/inter.css", path)
}

func serveHTMX(e *echo.Echo, assetPath string) {
	// main HTMX
	path := fmt.Sprintf("%s/js/htmx/htmx-v2.0.2.min.js", assetPath)
	e.File("/js/htmx/htmx-v2.0.2.min.js", path)

	// HTMX extensions
	path = fmt.Sprintf("%s/js/htmx/idiomorph-ext-v0.3.0.min.js", os.Getenv("ASSETS_PATH"))
	e.File("/js/htmx/idiomorph-ext-v0.3.0.min.js", path)
}
