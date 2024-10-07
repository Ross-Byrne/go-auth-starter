package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	db_string := os.Getenv("GOOSE_DBSTRING")
	db, err := gorm.Open(sqlite.Open(db_string), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	return db, nil
}

// get gorm migrator for goose migrations
func Migrator() (gorm.Migrator, error) {
	db, err := ConnectToDB()
	if err != nil {
		log.Println("Migrator: Failed to connect to database")
		return nil, err
	}

	return db.Migrator(), nil
}

// helps for migrations
func TryAutoMigrate(models ...interface{}) error {
	migrator, err := Migrator()
	if err != nil {
		return err
	}
	return migrator.AutoMigrate(models...)
}

func TryDropTables(models ...interface{}) error {
	migrator, err := Migrator()
	if err != nil {
		return err
	}

	return migrator.DropTable(models...)
}
