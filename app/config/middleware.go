package config

import (
	"go-auth-starter/app/utils"
	"go-auth-starter/database"
	"go-auth-starter/types/contextkey"
	"go-auth-starter/types/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wader/gormstore/v2"
)

// Check if user is authenticated
func CheckUserAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// if no session, redirect to login
		if !utils.RequestHasSession(c) {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// otherwise, all good keep going
		return next(c)
	}
}

// check if user is Authenticated and redriect to root
// if trying to get to unauthed pages
func CheckUserNotAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// if user has sesson, redirect to root
		if utils.RequestHasSession(c) {
			return c.Redirect(http.StatusSeeOther, "/")
		}

		// otherwise, all good keep going
		return next(c)
	}
}

// configures the session store using gorm db sessions
func ConfigSessionStore(e *echo.Echo) {
	session_secret_key := os.Getenv("SESSION_SECRET_KEY")
	isProduction, parseErr := strconv.ParseBool(os.Getenv("PRODUCTION"))
	if parseErr != nil {
		(*e).Logger.Fatal("Config Session Store: Failed to parse production env")
		return
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		(*e).Logger.Fatal("Config Session Store: Failed to connect to DB")
		return
	}

	session_store := gormstore.New(db, []byte(session_secret_key))

	// set store options
	session_store.SessionOpts.Path = "/"
	session_store.SessionOpts.MaxAge = 60 * 60 * 48 // max age is 2 days. 60 seconds x 60 minutes x 48 hours
	session_store.SessionOpts.HttpOnly = true
	session_store.SessionOpts.Secure = isProduction // Only true in Production
	session_store.SessionOpts.SameSite = http.SameSiteStrictMode

	// use store
	e.Use(session.Middleware(session_store))

	// db cleanup every hour
	// close channel to stop cleanup
	cleanupChan := make(chan struct{})
	go session_store.PeriodicCleanup(1*time.Hour, cleanupChan)
}

func AddDbConnectionToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get connect to db for request
		db, err := database.ConnectToDB()
		if err != nil {
			c.Logger().Error("Failed to add db connection to request context")
			return echo.ErrInternalServerError
		}

		// add connection to context
		c.Set(contextkey.DB, db)

		return next(c)
	}
}

func AddWorkspaceIdToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get connect to db for request
		db := utils.GetConnection(c)

		// get user id from session
		userID, err := utils.GetUserIdFromSession(c)
		if err != nil {
			utils.DeleteUserSession(c)
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// get workspace id from logged in user
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.Logger().Error("Failed to find user from session user id")
			utils.DeleteUserSession(c)
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// add workspace id to context
		c.Set(contextkey.WORKSPACE_ID, user.WorkspaceID)

		return next(c)
	}
}

func csrfConfig() middleware.CSRFConfig {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION"))
	if err != nil {
		panic("Failed to read env vars while setting up CSRF Config")
	}

	return middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookieSecure:   isProduction, // Only true in Production
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}
}
