package response

import (
	"bitbucket.org/rebelworksco/go-skeleton/models"
)

//VehicleResponse : format json response for vehicle
type VehicleResponse struct {
	ID          uint32 `json:"id"`
	BrandID     uint32 `json:"brand_id"`
	VehicleName string `json:"vehicle_name"`
}

//Transform from Vehicle model to Vehicle response
func (u *VehicleResponse) Transform(vehicle *models.Vehicle) {
	u.ID = vehicle.ID
	u.BrandID = vehicle.BrandID
	u.VehicleName = vehicle.VehicleName
}
