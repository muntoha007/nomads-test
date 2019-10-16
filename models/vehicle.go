package models

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Vehicle : struct of Vehicle
type Vehicle struct {
	ID          uint32 `db:"id"`
	BrandID     uint32 `db:"brand_id"`
	VehicleName string `db:"vehicle_name"`
}

const qVehicles = `SELECT id, brand_id, vehicle_name FROM vehicles`

//List : List of Vehicle
func (u *Vehicle) List(ctx context.Context, db *sqlx.DB) ([]Vehicle, error) {
	list := []Vehicle{}
	err := db.SelectContext(ctx, &list, qVehicles)
	return list, err
}

//Get Vehicle by id
func (u *Vehicle) Get(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qVehicles+" WHERE id=?", u.ID)
}

//Create new Vehicle
func (u *Vehicle) Create(ctx context.Context, db *sqlx.DB) error {
	fmt.Printf("ID: %d\t", u.BrandID)
	const query = `
		INSERT INTO vehicles (brand_id, vehicle_name, created, updated)
		VALUES (?, ?, NOW(), NOW())
	`
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.BrandID, u.VehicleName)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint32(id)

	return nil
}

//Update Vehicle
func (u *Vehicle) Update(ctx context.Context, db *sqlx.DB) error {

	stmt, err := db.PreparexContext(ctx, `
		UPDATE vehicles 
		SET brand_id = ?,
			vehicle_name = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.BrandID, u.VehicleName, u.ID)
	return err
}

//Delete Brands
func (u *Vehicle) Delete(ctx context.Context, db *sqlx.DB) (bool, error) {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM vehicles WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
