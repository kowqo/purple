package request

import (
	"net/http"
	"purple/pkg/resp"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.JSON(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	IsValid(body)
	if err != nil {
		resp.JSON(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}
