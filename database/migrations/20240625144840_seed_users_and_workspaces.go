package migrations

import (
	"context"
	"database/sql"
	"go-auth-starter/database"
	"go-auth-starter/types/models"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upSeedUsersAndWorkspaces, downSeedUsersAndWorkspaces)
}

func upSeedUsersAndWorkspaces(ctx context.Context, tx *sql.Tx) error {
	db, err := database.ConnectToDB()
	if err != nil {
		return err
	}

	txErr := db.Transaction(func(tx *gorm.DB) error {
		// Create default workspace
		if err := tx.Create(&models.Workspace{Name: "Default"}).Error; err != nil {
			return err
		}

		var workspace models.Workspace
		if err := tx.First(&workspace).Error; err != nil {
			return err
		}

		data := []models.User{
			{Email: "u1@test.com", FirstName: "u1", LastName: "u1", EncryptedPassword: "password", WorkspaceID: workspace.ID},
			{Email: "u2@test.com", FirstName: "u2", LastName: "u2", EncryptedPassword: "password", WorkspaceID: workspace.ID},
			{Email: "u3@test.com", FirstName: "u3", LastName: "u3", EncryptedPassword: "password", WorkspaceID: workspace.ID},
			{Email: "u4@test.com", FirstName: "u4", LastName: "u4", EncryptedPassword: "password", WorkspaceID: workspace.ID},
			{Email: "u5@test.com", FirstName: "u5", LastName: "u5", EncryptedPassword: "password", WorkspaceID: workspace.ID},
		}

		// create test users
		for _, user := range data {
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return txErr
}

func downSeedUsersAndWorkspaces(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
