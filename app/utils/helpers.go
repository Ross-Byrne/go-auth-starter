// Helpers for more complex or reusable logic in request handlers

package utils

import (
	"context"
	"errors"
	"fmt"
	"go-auth-starter/database"
	"go-auth-starter/types/contextkey"
	"go-auth-starter/types/models"
	"log"
	"net/mail"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const LOCATION_KEY string = "location"
const CSRF_TOKEN_KEY string = "csrf"

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// set location for frontend
	ctx := SetLocation(c.Request().Context(), c.Request().URL.Path)

	// Pass CSRF token to temple context
	if token := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); token != "" {
		ctx = context.WithValue(ctx, CSRF_TOKEN_KEY, token)
	}

	return t.Render(ctx, c.Response().Writer)
}

func SetLocation(ctx context.Context, location string) context.Context {
	return context.WithValue(ctx, LOCATION_KEY, location)
}

func GetLocation(ctx context.Context) string {
	if location, ok := ctx.Value(LOCATION_KEY).(string); ok {
		return location
	}
	return ""
}

func GetCsrfTokenHeader(ctx context.Context) string {
	if token, ok := ctx.Value(CSRF_TOKEN_KEY).(string); ok {
		return fmt.Sprintf(`{ "X-CSRF-Token": %q }`, token)
	}
	return ""
}

func GetConnection(c echo.Context) *gorm.DB {
	return c.Get(contextkey.DB).(*gorm.DB)
}

func GetWorkspaceId(c echo.Context) uint {
	return c.Get(contextkey.WORKSPACE_ID).(uint)
}

// Create User session. Returns error
func CreateUserSession(c echo.Context, user_id uint) error {
	sess, err := session.Get("session", c)
	if err != nil {
		c.Logger().Errorf("Error creating new session: %s\n", err.Error())
		return err
	}

	sess.Values["user_id"] = user_id
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Errorf("Error saving session: %s\n", err.Error())
		return err
	}

	c.Logger().Infof("New Session: %s\n", sess.Values)
	return nil
}

// Delete user session
func DeleteUserSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		c.Logger().Errorf("Error getting session: %s\n", err.Error())
		return err
	}

	// set to -1 to delete cookie
	sess.Options.MaxAge = -1

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Errorf("Error saving session: %s\n", err.Error())
		return err
	}
	return nil
}

// Check if user has a session
func RequestHasSession(c echo.Context) bool {
	sess, err := session.Get("session", c)
	if err != nil || sess.IsNew {
		c.Logger().Info("No session, redirecting to login")
		return false
	}

	return true
}

// get user id from session
func GetUserIdFromSession(c echo.Context) (uint, error) {
	sess, err := session.Get("session", c)
	if err != nil || sess.IsNew {
		c.Logger().Error("Cannot get user id from session, No session")
		return 0, err
	}

	// get user id from session
	userID := sess.Values["user_id"].(uint)
	if userID < 1 {
		c.Logger().Error("User id from session not valid")
		return 0, errors.New("User id from session not valid")
	}

	return userID, nil
}

func IsUserEmailAvailable(email string) (bool, error) {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Println("Failed to verify email.")
		return false, err
	}
	// convert to lower case to avoid issues
	email = strings.ToLower(email)

	// count number of entities with email
	var count int64
	db.Model(&models.User{}).Where(&models.User{Email: email}).Count(&count)

	// if none found with email, it's available
	return count == 0, nil
}

// validation function for emails
func IsValidEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}
