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
type Vehicles struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of Vehicles
func (u *Vehicles) List(w http.ResponseWriter, r *http.Request) error {
	var vehicles models.Vehicle
	list, err := vehicles.List(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting Vehicles list")
	}

	var listResponse []*response.VehicleResponse
	for _, vehicle := range list {
		var vehicleResponse response.VehicleResponse
		vehicleResponse.Transform(&vehicle)
		listResponse = append(listResponse, &vehicleResponse)
	}

	return api.ResponseOK(w, listResponse, http.StatusOK)
}

//View : http handler for retrieve vehicle by id
func (u *Vehicles) View(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting")
	}

	var vehicle models.Vehicle
	vehicle.ID = uint32(id)
	err = vehicle.Get(r.Context(), u.Db)

	if err == sql.ErrNoRows {
		u.Log.Printf("ERROR : %+v", err)
		return api.NotFoundError(err, "")
	}

	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Vehicle")
	}

	var response response.VehicleResponse
	response.Transform(&vehicle)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Create : http handler for create new brand
func (u *Vehicles) Create(w http.ResponseWriter, r *http.Request) error {
	var vehicleRequest request.NewVehicleRequest
	err := api.Decode(r, &vehicleRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode vehicle")
	}

	vehicle := vehicleRequest.Transform()
	err = vehicle.Create(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Create Vehicle")
	}

	var response response.VehicleResponse
	response.Transform(vehicle)
	return api.ResponseOK(w, response, http.StatusCreated)
}

//Update : http handler for update vehicles by id
func (u *Vehicles) Update(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var vehicle models.Vehicle
	vehicle.ID = uint32(id)
	err = vehicle.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Role")
	}

	var vehicleRequest request.VehicleRequest
	err = api.Decode(r, &vehicleRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Decode Role")
	}

	if vehicleRequest.ID <= 0 {
		vehicleRequest.ID = vehicle.ID
	}
	vehicleUpdate := vehicleRequest.Transform(&vehicle)
	err = vehicleUpdate.Update(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Update Vehicle")
	}

	var response response.VehicleResponse
	response.Transform(vehicleUpdate)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Delete : http handler for delete vehicle by id
func (u *Vehicles) Delete(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var vehicle models.Vehicle
	vehicle.ID = uint32(id)
	err = vehicle.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get role")
	}

	isDelete, err := vehicle.Delete(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Delete vehicle")
	}

	if !isDelete {
		err = errors.New("error delete vehicle")
		u.Log.Printf("ERROR : %+v", err)
		return err
	}

	return api.ResponseOK(w, nil, http.StatusNoContent)
}
