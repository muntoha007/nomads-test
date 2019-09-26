package controllers

import (
	"log"
	"net/http"

	"bitbucket.org/rebelworksco/go-skeleton/libraries/api"
	"bitbucket.org/rebelworksco/go-skeleton/libraries/token"
	"bitbucket.org/rebelworksco/go-skeleton/models"
	"bitbucket.org/rebelworksco/go-skeleton/payloads/request"
	"bitbucket.org/rebelworksco/go-skeleton/payloads/response"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//Auths : struct for set Auths Dependency Injection
type Auths struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//Login : http handler for login
func (u *Auths) Login(w http.ResponseWriter, r *http.Request) error {
	var loginRequest request.LoginRequest
	err := api.Decode(r, &loginRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode auth")
	}

	uLogin := models.User{Username: loginRequest.Username}
	err = uLogin.GetByUsername(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "call login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(uLogin.Password), []byte(loginRequest.Password))
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "compare password")
	}

	token, err := token.ClaimToken(uLogin.Username)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "claim token")
	}

	var response response.TokenResponse
	response.Token = token

	return api.ResponseOK(w, response, http.StatusOK)
}
