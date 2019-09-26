package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/auth"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/config"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/database"
	"bitbucket.org/rebelworksco/go-skeleton/schema"
	"github.com/pkg/errors"
)

func main() {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	if err := run(); err != nil {
		log.Printf("error: shutting down: %s", err)
		os.Exit(1)
	}
}

func run() error {

	flag.Parse()

	// =========================================================================
	// Start Database

	dbx, err := database.Openx()
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer dbx.Close()

	switch flag.Arg(0) {
	case "migrate":
		db, err := database.Open()
		if err != nil {
			return errors.Wrap(err, "connecting to db")
		}
		defer db.Close()
		if err := schema.Migrate(db); err != nil {
			return errors.Wrap(err, "applying migrations")
		}
		fmt.Println("Migrations complete")

	case "seed":
		if err := schema.Seed(dbx); err != nil {
			return errors.Wrap(err, "seeding database")
		}
		fmt.Println("Seed data complete")
	case "scan-access":
		if err := auth.ScanAccess(dbx); err != nil {
			return errors.Wrap(err, "scan access")
		}
		fmt.Println("Scan access complete")
	}

	return nil
}
