package router

import (
	"encoding/json"
	"net/http"

	"github.com/jiharal/apiary/api"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, *api.Error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errs []string

	// Ignore error from form parsing as it's insignificant.
	r.ParseForm()

	data, err := fn(w, r)
	if err != nil {
		logger.Err.WithError(err.Err).Println("Serve error.")
		errs = append(errs, err.Error())
		w.WriteHeader(err.StatusCode)
		resp := api.Response{
			Status: http.StatusText(err.StatusCode),
			Data:   data,
			BaseResponse: api.BaseResponse{
				Errors: errs,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			logger.Err.WithError(err).Println("Encode response error.")
			return
		}
	} else {
		resp := api.Response{
			Status: http.StatusText(200),
			Data:   data,
			BaseResponse: api.BaseResponse{
				Errors: errs,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			logger.Err.WithError(err).Println("Encode response error.")
			return
		}
	}
}
