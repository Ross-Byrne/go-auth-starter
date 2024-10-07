package handlers

import (
	"go-auth-starter/app/utils"
	"go-auth-starter/app/views/pages/login"
	"go-auth-starter/types/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func LoginHandlers(g *echo.Group) {
	g.GET("/login", loginGet)
	g.POST("/login", loginPost)
}

func loginGet(c echo.Context) error {
	return utils.Render(c, http.StatusOK, login.Index())
}

func loginPost(c echo.Context) error {
	params := new(login.LoginFormValues)
	if err := c.Bind(params); err != nil {
		return echo.ErrBadRequest
	}

	// process params
	params.Email = strings.ToLower(params.Email)

	// validate params
	formErrors := params.Validate()
	if len(formErrors) > 0 {
		data := login.IndexPageData{
			FormValues: *params,
			FormErrors: formErrors,
			FlashError: "",
		}
		return utils.Render(c, http.StatusOK, login.IndexWith(data))
	}

	db := utils.GetConnection(c)

	// Find user and check password
	// If either are incorrect, give vague error
	// to protect against account guessing
	var user models.User
	if result := db.Where("email = ?", params.Email).First(&user); result.Error != nil {
		c.Logger().Errorf("Failed login attempt with email %s. Eamil not found.", params.Email)
		return utils.Render(c, http.StatusOK, login.IndexWithErrorFlash("Login failed: Invalid email or password"))
	}

	// check password is correct
	if isCorrect := utils.CompareHashAndPassword(params.Password, user.EncryptedPassword); !isCorrect {
		c.Logger().Errorf("Failed login attempt with email %s. Password incorrect", params.Email)
		return utils.Render(c, http.StatusOK, login.IndexWithErrorFlash("Login failed: Invalid email or password"))
	}

	// Create user session
	if err := utils.CreateUserSession(c, user.ID); err != nil {
		return echo.ErrInternalServerError
	}

	// redirect to root url
	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}
