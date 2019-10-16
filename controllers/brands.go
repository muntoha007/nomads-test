package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/rebelworksco/go-skeleton/libraries/api"
	"bitbucket.org/rebelworksco/go-skeleton/models"
	"bitbucket.org/rebelworksco/go-skeleton/payloads/request"
	"bitbucket.org/rebelworksco/go-skeleton/payloads/response"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Brands : struct for set Driver Dependency Injection
type Brands struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of brands
func (u *Brands) List(w http.ResponseWriter, r *http.Request) error {
	var brand models.Brand
	list, err := brand.List(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting brands list")
	}

	var listResponse []*response.BrandResponse
	for _, brand := range list {
		var brandResponse response.BrandResponse
		brandResponse.Transform(&brand)
		listResponse = append(listResponse, &brandResponse)
	}

	return api.ResponseOK(w, listResponse, http.StatusOK)
}

//View : http handler for retrieve brand by id
func (u *Brands) View(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting")
	}

	var brand models.Brand
	brand.ID = uint32(id)
	err = brand.Get(r.Context(), u.Db)

	if err == sql.ErrNoRows {
		u.Log.Printf("ERROR : %+v", err)
		return api.NotFoundError(err, "")
	}

	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Brand")
	}

	var response response.BrandResponse
	response.Transform(&brand)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Create : http handler for create new brand
func (u *Brands) Create(w http.ResponseWriter, r *http.Request) error {
	var brandRequest request.NewBrandRequest
	err := api.Decode(r, &brandRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode brand")
	}

	brand := brandRequest.Transform()
	err = brand.Create(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Create Brand")
	}

	var response response.BrandResponse
	response.Transform(brand)
	return api.ResponseOK(w, response, http.StatusCreated)
}

//Update : http handler for update brand by id
func (u *Brands) Update(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var brand models.Brand
	brand.ID = uint32(id)
	err = brand.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Role")
	}

	var brandRequest request.BrandRequest
	err = api.Decode(r, &brandRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Decode Role")
	}

	if brandRequest.ID <= 0 {
		brandRequest.ID = brand.ID
	}
	brandUpdate := brandRequest.Transform(&brand)
	err = brandUpdate.Update(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Update brand")
	}

	var response response.BrandResponse
	response.Transform(brandUpdate)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Delete : http handler for delete brand by id
func (u *Brands) Delete(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var brand models.Brand
	brand.ID = uint32(id)
	err = brand.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get role")
	}

	isDelete, err := brand.Delete(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Delete brand")
	}

	if !isDelete {
		err = errors.New("error delete brand")
		u.Log.Printf("ERROR : %+v", err)
		return err
	}

	return api.ResponseOK(w, nil, http.StatusNoContent)
}
