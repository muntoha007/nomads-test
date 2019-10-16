package request

import (
	"bitbucket.org/rebelworksco/go-skeleton/models"
)

//NewBrandRequest : format json request for new Brand
type NewBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

//Transform NewBrandRequest to Brand
func (u *NewBrandRequest) Transform() *models.Brand {
	var brand models.Brand
	brand.Name = u.Name

	return &brand
}

//BrandRequest : format json request for Brand
type BrandRequest struct {
	ID   uint32 `json:"id,omitempty"  validate:"required"`
	Name string `json:"name,omitempty"  validate:"required"`
}

//Transform BrandRequest to Brand
func (u *BrandRequest) Transform(brand *models.Brand) *models.Brand {
	if u.ID == brand.ID {
		if len(u.Name) > 0 {
			brand.Name = u.Name
		}
	}
	return brand
}
