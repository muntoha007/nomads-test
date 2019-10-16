package response

import (
	"bitbucket.org/rebelworksco/go-skeleton/models"
)

//BrandResponse : format json response for brand
type BrandResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

//Transform from Brand model to Brand response
func (u *BrandResponse) Transform(brand *models.Brand) {
	u.ID = brand.ID
	u.Name = brand.Name
}
