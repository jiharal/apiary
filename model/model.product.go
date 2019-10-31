package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type (
	ProductModel struct {
		ID        int64
		Name      string
		Price     int64
		ImageUrl  string
		CreatedAt time.Time
		CreatedBy int64
		UpdatedAt pq.NullTime
		UpdatedBy int64
	}

	ProductModelResponse struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Price     int64     `json:"price"`
		ImageUrl  string    `json:"imageurl"`
		CreatedAt time.Time `json:"created_at"`
		CreatedBy int64     `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy int64     `json:"updated_by"`
	}
)

// Convert product model into json-friendly formatted response struct (without null data type).
func (product *ProductModel) Response() ProductModelResponse {
	return ProductModelResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		ImageUrl:  product.ImageUrl,
		CreatedAt: product.CreatedAt,
		CreatedBy: product.CreatedBy,
		UpdatedAt: product.UpdatedAt.Time,
		UpdatedBy: product.UpdatedBy,
	}
}

func GetAllProduct(ctx context.Context, db *sql.DB) ([]ProductModel, error) {
	query := fmt.Sprintf(`SELECT
				id,
				name,
				price,
				imageurl,
				created_at,
				created_by,
				updated_at,
				updated_by
			FROM
				product
		`)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "model/product/list")
	}
	defer rows.Close()

	var products []ProductModel

	for rows.Next() {
		var product ProductModel

		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.ImageUrl,
			&product.CreatedAt,
			&product.CreatedBy,
			&product.UpdatedAt,
			&product.UpdatedBy,
		)
		if err != nil {
			return nil, errors.Wrap(err, "model/product/list/scan")
		}

		products = append(products, product)
	}
	return products, nil
}

func GetOneProduct(ctx context.Context, db *sql.DB, id int64) (ProductModel, error) {
	query := `
				SELECT
					id,
					name,
					price,
					imageurl,
					created_at,
					created_by,
					updated_at,
					updated_by
				FROM
					product
				WHERE
					id = $1`

	var product ProductModel

	err := db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.ImageUrl,
		&product.CreatedAt,
		&product.CreatedBy,
		&product.UpdatedAt,
		&product.UpdatedBy,
	)
	if err != nil {
		return ProductModel{}, errors.Wrap(err, "model/product/query/id")
	}

	return product, nil
}

func (product ProductModel) Insert(ctx context.Context, db *sql.DB) (ProductModel, error) {
	query := `
			INSERT INTO product (
				name,
				price,
				imageurl,
				created_by,
				created_at
			) VALUES (
				$1, $2, $3, $4, now()
			) RETURNING
				id,
				created_at`

	err := db.QueryRowContext(ctx, query,
		product.Name,
		product.Price,
		product.ImageUrl,
		product.CreatedBy,
	).Scan(
		&product.ID,
		&product.CreatedAt,
	)
	if err != nil {
		return ProductModel{}, errors.Wrap(err, "model/product/insert")
	}

	return product, nil
}

func (product ProductModel) Update(ctx context.Context, db *sql.DB) (ProductModel, error) {
	query := `
				UPDATE
					product
				SET
					name = $1,
					price = $2,
					imageurl = $3,
					updated_by = $4,
					updated_at = NOW()
				WHERE
					id = $5`

	_, err := db.ExecContext(ctx, query,
		product.Name,
		product.Price,
		product.ImageUrl,
		product.UpdatedBy,
		product.ID,
	)
	if err != nil {
		return product, errors.Wrap(err, "model/product/update")
	}

	return product, nil
}

func DeleteProductByID(ctx context.Context, db *sql.DB, id int64) error {
	query := "DELETE FROM product WHERE id = $1"

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "model/product/delete")
	}

	return nil
}
