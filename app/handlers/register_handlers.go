package handlers

import (
	"go-auth-starter/app/utils"
	"go-auth-starter/app/views/pages/register"
	"go-auth-starter/types/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/register", registerGet)
	g.POST("/register", registerPost)
}

func registerGet(c echo.Context) error {
	return utils.Render(c, http.StatusOK, register.Index())
}

func registerPost(c echo.Context) error {
	params := new(register.RegisterFormValues)
	if err := c.Bind(params); err != nil {
		return echo.ErrBadRequest
	}

	// process params
	params.Email = strings.ToLower(params.Email)

	// validate params
	formErrors := params.Validate()
	if len(formErrors) > 0 {
		data := register.IndexPageData{
			FormValues: *params,
			FormErrors: formErrors,
		}
		return utils.Render(c, http.StatusOK, register.IndexWith(data))
	}

	// Hash password, replacing plaintext param
	var hashErr error
	params.Password, hashErr = utils.GenerateHashedPassword(params.Password)

	if hashErr != nil {
		c.Logger().Error("Failed to hash new users password")
		return echo.ErrInternalServerError
	}

	db := utils.GetConnection(c)
	var new_user models.User

	// create transaction
	txErr := db.Transaction(func(tx *gorm.DB) error {
		// Create Workspace for user
		workspace := models.Workspace{Name: params.Workspace}
		if err := tx.Create(&workspace).Error; err != nil {
			c.Logger().Error("Failed to create workspace")
			return err
		}

		// create new user in db
		new_user = models.User{
			FirstName:         params.FirstName,
			LastName:          params.LastName,
			Email:             params.Email,
			EncryptedPassword: params.Password,
			WorkspaceID:       workspace.ID,
		}
		if err := tx.Create(&new_user).Error; err != nil {
			c.Logger().Error("Failed to create user")
			return err
		}

		return nil
	})
	if txErr != nil {
		return echo.ErrInternalServerError
	}

	// Create user session
	if err := utils.CreateUserSession(c, new_user.ID); err != nil {
		return echo.ErrInternalServerError
	}

	// redirect to root url
	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}
