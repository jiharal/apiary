package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/jiharal/apiary/api"
)

func HandlerProductList(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	return productService.List(ctx)
}

func HandlerProductDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := StringToInt64(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "product/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return productService.Detail(ctx, id)
}

func HandlerProductCreate(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param api.ProductDataParam

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "product/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.By = api.GetContextRequesterID(ctx)
	return productService.Create(ctx, param)
}

func HandlerProductUpdate(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := StringToInt64(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "product/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	var param api.ProductDataParam
	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "product/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id
	param.By = api.GetContextRequesterID(ctx)
	return productService.Update(ctx, param)
}

func HandlerProductDelete(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := StringToInt64(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "product/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return productService.Delete(ctx, id)
}
