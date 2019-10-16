package request

import (
	"bitbucket.org/rebelworksco/go-skeleton/models"
)

//NewVehicleRequest : format json request for new Vehicle
type NewVehicleRequest struct {
	BrandID     uint32 `json:"brand_id" validate:"required"`
	VehicleName string `json:"vehicle_name" validate:"required"`
}

//Transform NewVehicleRequest to Vehicle
func (u *NewVehicleRequest) Transform() *models.Vehicle {
	var vehicle models.Vehicle
	vehicle.BrandID = u.BrandID
	vehicle.VehicleName = u.VehicleName

	return &vehicle
}

//VehicleRequest : format json request for Vehicle
type VehicleRequest struct {
	ID          uint32 `json:"id,omitempty"  validate:"required"`
	BrandID     uint32 `json:"brand_id,omitempty"  validate:"required"`
	VehicleName string `json:"vehicle_name,omitempty"  validate:"required"`
}

//Transform VehicleRequest to Vehicle
func (u *VehicleRequest) Transform(vehicle *models.Vehicle) *models.Vehicle {
	if u.ID == vehicle.ID {

		vehicle.BrandID = u.BrandID

		if len(u.VehicleName) > 0 {
			vehicle.VehicleName = u.VehicleName
		}
	}
	return vehicle
}
