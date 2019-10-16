package routing

import (
	"log"
	"net/http"

	"bitbucket.org/rebelworksco/go-skeleton/controllers"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/api"
	"bitbucket.org/rebelworksco/go-skeleton/middleware"
	"github.com/jmoiron/sqlx"
)

func mid(db *sqlx.DB, log *log.Logger) []api.Middleware {
	var mw []api.Middleware
	mw = append(mw, middleware.Errors(db, log))
	mw = append(mw, middleware.Auths(db, log, []string{"/login", "/health"}))

	return mw
}

//API : handler api
func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := api.NewApp(db, log, mid(db, log)...)

	// Health Routing
	{
		check := controllers.Checks{Db: db}
		app.Handle(http.MethodGet, "/health", check.Health)
	}

	// Auth Routing
	{
		auth := controllers.Auths{Db: db, Log: log}
		app.Handle(http.MethodPost, "/login", auth.Login)
	}

	// Users Routing
	{
		users := controllers.Users{Db: db, Log: log}
		app.Handle(http.MethodGet, "/users", users.List)
		app.Handle(http.MethodGet, "/users/{id}", users.View)
		app.Handle(http.MethodPost, "/users", users.Create)
		app.Handle(http.MethodPut, "/users/{id}", users.Update)
		app.Handle(http.MethodDelete, "/users/{id}", users.Delete)
	}

	// Roles Routing
	{
		roles := controllers.Roles{Db: db, Log: log}
		app.Handle(http.MethodGet, "/roles", roles.List)
		app.Handle(http.MethodGet, "/roles/{id}", roles.View)
		app.Handle(http.MethodPost, "/roles", roles.Create)
		app.Handle(http.MethodPut, "/roles/{id}", roles.Update)
		app.Handle(http.MethodDelete, "/roles/{id}", roles.Delete)
		app.Handle(http.MethodPost, "/roles/{id}/access/{access_id}", roles.Grant)
		app.Handle(http.MethodDelete, "/roles/{id}/access/{access_id}", roles.Revoke)
	}

	// Access Routing
	{
		access := controllers.Access{Db: db, Log: log}
		app.Handle(http.MethodGet, "/access", access.List)
	}

	// Brands Routing
	{
		brands := controllers.Brands{Db: db, Log: log}
		app.Handle(http.MethodGet, "/brands", brands.List)
		app.Handle(http.MethodGet, "/brands/{id}", brands.View)
		app.Handle(http.MethodPost, "/brands", brands.Create)
		app.Handle(http.MethodPut, "/brands/{id}", brands.Update)
		app.Handle(http.MethodDelete, "/brands/{id}", brands.Delete)
	}

	// Vehicles Routing
	{
		vehicles := controllers.Vehicles{Db: db, Log: log}
		app.Handle(http.MethodGet, "/vehicles", vehicles.List)
		app.Handle(http.MethodGet, "/vehicles/{id}", vehicles.View)
		app.Handle(http.MethodPost, "/vehicles", vehicles.Create)
		app.Handle(http.MethodPut, "/vehicles/{id}", vehicles.Update)
		app.Handle(http.MethodDelete, "/vehicles/{id}", vehicles.Delete)
	}

	return app
}
