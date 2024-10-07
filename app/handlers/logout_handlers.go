package handlers

import (
	"go-auth-starter/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LogoutHandlers(g *echo.Group) {
	g.GET("/logout", logoutGet)
}

func logoutGet(c echo.Context) error {
	// Remove session to logout user and redirect to login
	utils.DeleteUserSession(c)
	return c.Redirect(http.StatusSeeOther, "/login")
}
