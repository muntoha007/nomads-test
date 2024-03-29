package models

import (
	"context"
	"database/sql"
	"errors"

	"bitbucket.org/rebelworksco/go-skeleton/libraries/token"
	"github.com/jmoiron/sqlx"
)

//Access : struct of Access
type Access struct {
	ID       uint32        `db:"id"`
	ParentID sql.NullInt64 `db:"parent_id"`
	Name     string        `db:"name"`
	Alias    string        `db:"alias"`
}

const qAccess = `SELECT id, parent_id, name, alias FROM access`

//List : List of access
func (u *Access) List(ctx context.Context, tx *sqlx.Tx) ([]Access, error) {
	list := []Access{}
	err := tx.SelectContext(ctx, &list, qAccess)
	return list, err
}

//GetByName : get access by name
func (u *Access) GetByName(ctx context.Context, tx *sqlx.Tx) error {
	return tx.GetContext(ctx, u, qAccess+" WHERE name=?", u.Name)
}

//GetByAlias : get access by alias
func (u *Access) GetByAlias(ctx context.Context, tx *sqlx.Tx) error {
	return tx.GetContext(ctx, u, qAccess+" WHERE alias=?", u.Alias)
}

//Get : get access by id
func (u *Access) Get(ctx context.Context, tx *sqlx.Tx) error {
	return tx.GetContext(ctx, u, qAccess+" WHERE id=?", u.ID)
}

//Create new Access
func (u *Access) Create(ctx context.Context, tx *sqlx.Tx) error {
	const query = `
		INSERT INTO access (parent_id, name, alias, created)
		VALUES (?, ?, ?, NOW())
	`
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.ParentID, u.Name, u.Alias)
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

//Delete : delete user
func (u *Access) Delete(ctx context.Context, tx *sqlx.Tx) (bool, error) {
	stmt, err := tx.PreparexContext(ctx, `DELETE FROM access WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetIDs : get array of access id
func (u *Access) GetIDs(ctx context.Context, db *sqlx.DB) ([]uint32, error) {
	var access []uint32

	rows, err := db.QueryContext(ctx, "SELECT id FROM access WHERE name != 'root'")
	if err != nil {
		return access, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			return access, err
		}
		access = append(access, id)
	}

	return access, rows.Err()
}

// IsAUth for check user authorization
func (u *Access) IsAuth(ctx context.Context, db *sqlx.DB, tokenparam interface{}, controller string, route string) (bool, error) {
	query := `
	SELECT true
	FROM users
	JOIN roles_users ON users.id = roles_users.user_id
	JOIN roles ON roles_users.role_id = roles.id
	JOIN access_roles ON roles.id = access_roles.role_id
	JOIN access ON access_roles.access_id = access.id
	WHERE (access.name = 'root' OR access.name = ? OR access.name = ?)
	AND users.id = ?
	`
	var isAuth bool
	var err error

	if tokenparam == nil {
		return isAuth, errors.New("Bad request for token")
	}

	isValid, username := token.ValidateToken(tokenparam.(string))
	if !isValid {
		return isAuth, errors.New("Bad request for invalid token")
	}

	user := User{Username: username}
	err = user.GetByUsername(ctx, db)
	if err != nil {
		return isAuth, err
	}

	err = db.QueryRowContext(ctx, query, controller, route, user.ID).Scan(&isAuth)

	return isAuth, err
}
