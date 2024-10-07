package handlers

import (
	"go-auth-starter/app/utils"
	"go-auth-starter/app/views/pages/profile"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProfileHandlers(g *echo.Group) {
	g.GET("/profile", profileGet)
}

func profileGet(c echo.Context) error {
	return utils.Render(c, http.StatusOK, profile.Profile())
}
