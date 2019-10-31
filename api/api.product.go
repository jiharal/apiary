package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/jiharal/apiary/model"
)

type (
	ProductModule struct {
		db    *sql.DB
		cache *redis.Pool
		name  string
	}

	ProductDataParam struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Price    int64  `json:"price"`
		ImageUrl string `json:"imageurl"`
		By       int64  `json:"by"`
	}
)

func NewProductModule(db *sql.DB, cache *redis.Pool) *ProductModule {
	return &ProductModule{
		db:    db,
		cache: cache,
		name:  "module/product",
	}
}

func (m ProductModule) List(ctx context.Context) ([]model.ProductModelResponse, *Error) {
	products, err := model.GetAllProduct(ctx, m.db)
	if err != nil {
		return nil, NewErrorWrap(err, m.name, "list/query",
			MessageGeneralError, http.StatusInternalServerError)
	}

	productResponse := []model.ProductModelResponse{}

	for _, product := range products {
		productResponse = append(productResponse, product.Response())
	}

	return productResponse, nil
}

func (m ProductModule) Detail(ctx context.Context, id int64) (model.ProductModelResponse, *Error) {
	product, err := model.GetOneProduct(ctx, m.db, id)
	if err != nil {
		status := http.StatusInternalServerError
		message := MessageGeneralError

		if err == sql.ErrNoRows {
			status = http.StatusNotFound
			message = http.StatusText(status)
		}

		return model.ProductModelResponse{}, NewErrorWrap(err, m.name, "detail/query",
			message, status)
	}

	return product.Response(), nil
}

func (m ProductModule) Create(ctx context.Context, param ProductDataParam) (model.ProductModelResponse, *Error) {

	product := model.ProductModel{
		Name:      param.Name,
		Price:     param.Price,
		ImageUrl:  param.ImageUrl,
		CreatedBy: param.By,
	}

	_, err := product.Insert(ctx, m.db)
	if err != nil {
		return model.ProductModelResponse{}, NewErrorWrap(err, m.name, "create",
			MessageGeneralError, http.StatusInternalServerError)
	}

	return product.Response(), nil
}

func (m ProductModule) Update(ctx context.Context, param ProductDataParam) (model.ProductModelResponse, *Error) {

	product, err := model.GetOneProduct(ctx, m.db, param.ID)
	if err != nil {
		return product.Response(), NewErrorWrap(err, m.name, "update",
			MessageGeneralError, http.StatusInternalServerError)
	}
	if param.Name != "" || param.Price != 0 || param.ImageUrl != "" {
		product.Name = param.Name
		product.Price = param.Price
		product.ImageUrl = param.ImageUrl
		product.UpdatedBy = param.By
	}

	_, err = product.Update(ctx, m.db)
	if err != nil {
		return product.Response(), NewErrorWrap(err, m.name, "update",
			MessageGeneralError, http.StatusInternalServerError)
	}

	return product.Response(), nil
}

func (m ProductModule) Delete(ctx context.Context, id int64) (interface{}, *Error) {
	err := model.DeleteProductByID(ctx, m.db, id)
	if err != nil {
		return nil, NewErrorWrap(err, m.name, "delete",
			MessageGeneralError, http.StatusInternalServerError)
	}
	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v deleted", id),
	}
	return data, nil
}
