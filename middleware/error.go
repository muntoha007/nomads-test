package middleware

import (
	"log"
	"net/http"

	"bitbucket.org/rebelworksco/go-skeleton/libraries/api"
	"github.com/jmoiron/sqlx"
)

func Errors(db *sqlx.DB, log *log.Logger) api.Middleware {
	fn := func(before api.Handler) api.Handler {
		h := func(w http.ResponseWriter, r *http.Request) error {

			err := before(w, r)

			if err != nil {
				// Log the error.
				log.Printf("ERROR : %+v", err)

				// Response to the error.
				if err := api.ResponseError(w, err); err != nil {
					return err
				}
			}

			// Return nil to indicate the error has been handled.
			return nil
		}

		return h
	}

	return fn
}
