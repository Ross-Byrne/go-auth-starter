package migrations

import (
	"context"
	"database/sql"
	"go-auth-starter/database"
	"go-auth-starter/types/models"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	return database.TryAutoMigrate(&models.User{}, &models.Workspace{})
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	return database.TryDropTables(&models.User{}, &models.Workspace{})
}
