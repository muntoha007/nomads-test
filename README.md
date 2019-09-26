# Go Skeleton
- GO Skeleton is skeleton for create API (Backend) in Rebelworks 
- To help rapid development 
- To standardize the coding

## Get Started
- copy all files on [go-skeleton](https://bitbucket.org/rebelworksco/go-skeleton/src) to your project directory
- cp .env.example .env
- edit .env with your environment
- go mod init your_git_source_project
- edit all text "bitbucket.org/rebelworksco/go-skeleton" in go import with your_git_source_project
- go run cmd/main.go migrate
- go run cmd/main.go seed
- go run cmd/main.go scan-access
- go test
- go run main.go

## Life Cycle
Request -> Middleware -> Controller -> Models -> Response

## Directory Structure
```
-> cmd 
-> controllers
-> libraries 
-> middleware
-> models 
-> payloads
	-> request 
	-> response
-> routing
-> schema
```

## Routing
- Open routing/route.go
- In the API function, add your routing
```
    // Products Routing
	{
		products := controllers.Products{Db: db, Log: log}
		app.Handle(http.MethodGet, "/products", products.List)
        app.Handle(http.MethodPost, "/products", products.Create)
        app.Handle(http.MethodGet, "/products/{id}", products.View)
        app.Handle(http.MethodPut, "/products/{id}", products.Update)
        app.Handle(http.MethodDelete, "/products/{id}", products.Delete)
	}
```

## Migration
- Open schema migrate.go
- Add array migrations with your products table structure
```
    {
		Version:     6,
		Description: "Add products",
		Script: `
CREATE TABLE products (
	id   BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	name         varchar(255) NOT NULL UNIQUE,
	created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);`,
	},
	
```
- go run cmd/main.go migrate

## Controller
- Create new file controllers/products.go to handle all products routing
```
package controllers

import (
	...
)

// Products : struct for set Products Dependency Injection
type Products struct {
	Db  *sqlx.DB
	Log *log.Logger
}

// List : http handler for returning list of products
func (u *Products) List(w http.ResponseWriter, r *http.Request) error {
	var product models.Product
	list, err := product.List(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting products list")
	}

	var listResponse []*response.ProductResponse
	for _, product := range list {
		var productResponse response.ProductResponse
		productResponse.Transform(&product)
		listResponse = append(listResponse, &productResponse)
	}

	return api.ResponseOK(w, listResponse, http.StatusOK)
}

// View : http handler for retrieve product by id
func (u *Products) View(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting")
	}

	var product models.Product
	product.ID = uint64(id)
	err = product.Get(r.Context(), u.Db)

	if err == sql.ErrNoRows {
		u.Log.Printf("ERROR : %+v", err)
		return api.NotFoundError(err, "")
	}

	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get product")
	}

	var response response.ProductResponse
	response.Transform(&product)
	return api.ResponseOK(w, response, http.StatusOK)
}

// Create : http handler for create new product
func (u *Products) Create(w http.ResponseWriter, r *http.Request) error {
	var productRequest request.NewProductRequest
	err := api.Decode(r, &productRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode product")
	}

	product := productRequest.Transform()
	err = product.Create(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Create product")
	}

	var response response.ProductResponse
	response.Transform(product)
	return api.ResponseOK(w, response, http.StatusCreated)
}

// Update : http handler for update product by id
func (u *Products) Update(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var product models.Product
	product.ID = uint64(id)
	err = product.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get product")
	}

	var productRequest request.ProductRequest
	err = api.Decode(r, &productRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Decode product")
	}

	if productRequest.ID <= 0 {
		productRequest.ID = product.ID
	}
	productUpdate := productRequest.Transform(&product)
	err = productUpdate.Update(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Update product")
	}

	var response response.ProductResponse
	response.Transform(productUpdate)
	return api.ResponseOK(w, response, http.StatusOK)
}

// Delete : http handler for delete product by id
func (u *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var product models.Product
	product.ID = uint64(id)
	err = product.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get product")
	}

	err = product.Delete(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Delete product")
	}

	return api.ResponseOK(w, nil, http.StatusNoContent)
}


```

## Model
- Create new file models/product.go
```
package models

import (
	...
)

// Product : struct of Product
type Product struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}

const qProducts = `SELECT id, name FROM products`

// List of products
func (u *Product) List(ctx context.Context, db *sqlx.DB) ([]Product, error) {
	list := []Product{}
	err := db.SelectContext(ctx, &list, qProducts)
	return list, err
}

// Get product by id
func (u *Product) Get(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qProducts+" WHERE id=?", u.ID)
}

// Create new product
func (u *Product) Create(ctx context.Context, db *sqlx.DB) error {
	const query = `
		INSERT INTO products (name, created)
		VALUES (?, NOW())
	`
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

    defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u.Name)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint64(id)

	return nil
}

// Update product
func (u *Product) Update(ctx context.Context, db *sqlx.DB) error {

	stmt, err := db.PreparexContext(ctx, `
		UPDATE products 
		SET name = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

    defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.Name, u.ID)
	return err
}

// Delete product
func (u *Product) Delete(ctx context.Context, db *sqlx.DB) error {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM products WHERE id = ?`)
	if err != nil {
		return err
	}

    defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.ID)
	return err
}

```

## Payload Request
- Create new file payloads/request/product_request.go
```
package request

import (
	...
)

// NewProductRequest : format json request for new product
type NewProductRequest struct {
	Name string `json:"name"`
}

// Transform NewProductRequest to Product
func (u *NewProductRequest) Transform() *models.Product {
	var product models.Product
	product.Name = u.Name

	return &product
}

// ProductRequest : format json request for product
type ProductRequest struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Transform ProductRequest to Product
func (u *ProductRequest) Transform(product *models.Product) *models.Product {
	if u.ID == product.ID {
		if len(u.Name) > 0 {
			product.Name = u.Name
		}
	}
	return product
}

``` 

### Validation
- Each request need validation
- For detail example please read [validator example](https://github.com/go-playground/validator/tree/v9.29.1/_examples)
- Open payloads/request/product_request.go
- Add validation tag to NewProductRequest
```
type NewProductRequest struct {
	Name string `json:"name" validate:"required"`
}
``` 
- Add validation tag to ProductRequest
```
type ProductRequest struct {
	ID   uint64 `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty"`
}
```

## Payload Response
- Create new file payloads/response/product_response.go
```
package response

import (
	...
)

// ProductResponse : format json response for product
type ProductResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// Transform from Product model to Product response
func (u *ProductResponse) Transform(product *models.Product) {
	u.ID = product.ID
	u.Name = product.Name
}

```