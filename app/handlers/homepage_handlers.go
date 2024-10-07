package handlers

import (
	"go-auth-starter/app/utils"
	"go-auth-starter/app/views/pages/home"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomepageHandlers(g *echo.Group) {
	g = g.Group("")

	g.GET("", homeGet)
}

func homeGet(c echo.Context) error {
	return utils.Render(c, http.StatusOK, home.Home())
}
