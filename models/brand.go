package models

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//Brand : struct of Brand
type Brand struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
}

const qBrands = `SELECT id, name FROM brands`

//List : List of Brand
func (u *Brand) List(ctx context.Context, db *sqlx.DB) ([]Brand, error) {
	list := []Brand{}
	err := db.SelectContext(ctx, &list, qBrands)
	return list, err
}

//Get Brand by id
func (u *Brand) Get(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qBrands+" WHERE id=?", u.ID)
}

//Create new Brand
func (u *Brand) Create(ctx context.Context, db *sqlx.DB) error {
	const query = `
		INSERT INTO brands (name, created)
		VALUES (?, NOW())
	`
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.Name)
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

//Update Brands
func (u *Brand) Update(ctx context.Context, db *sqlx.DB) error {

	stmt, err := db.PreparexContext(ctx, `
		UPDATE brands 
		SET name = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Name, u.ID)
	return err
}

//Delete Brands
func (u *Brand) Delete(ctx context.Context, db *sqlx.DB) (bool, error) {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM brands WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
