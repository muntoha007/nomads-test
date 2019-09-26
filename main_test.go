package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	apiTest "bitbucket.org/rebelworksco/go-skeleton/controllers/tests"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/auth"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/config"
	"bitbucket.org/rebelworksco/go-skeleton/routing"
	"bitbucket.org/rebelworksco/go-skeleton/schema"
	"bitbucket.org/rebelworksco/go-skeleton/tests"
)

var token string

func TestMain(t *testing.T) {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	db, teardown := tests.NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	if err := auth.ScanAccess(db); err != nil {
		t.Fatal(err)
	}

	log := log.New(os.Stderr, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// unit test for user
	//user := unitTest.User{Db: db}
	//t.Run("UnitUsersList", user.List)
	//t.Run("UnitUsersCrud", user.Crud)

	// api test for auths
	auths := apiTest.Auths{App: routing.API(db, log)}
	t.Run("ApiLogin", auths.Login)
	token = auths.Token

	// api test for users
	users := apiTest.Users{App: routing.API(db, log), Token: token}
	t.Run("ApiUsersList", users.List)
	t.Run("APiUsersCrud", users.Crud)
}
