// This is custom goose binary with sqlite3 support only.

package main

import (
	"context"
	"flag"
	_ "go-auth-starter/database/migrations"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const MIN_ARGS int = 1

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	// load .env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env vars
	var (
		dir       = os.Getenv("GOOSE_MIGRATION_DIR")
		db_string = os.Getenv("GOOSE_DBSTRING")
		driver    = os.Getenv("GOOSE_DRIVER")
	)

	flags.Parse(os.Args[MIN_ARGS:])
	args := flags.Args()

	if len(args) < MIN_ARGS {
		flags.Usage()
		return
	}

	command := args[0]

	db, err := goose.OpenDBWithDriver(driver, db_string)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > MIN_ARGS {
		arguments = append(arguments, args[MIN_ARGS:]...)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
