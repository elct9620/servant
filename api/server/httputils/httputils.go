package httputils

import (
	"encoding/json"
	"net/http"
)

type ApiFunc = func(w http.ResponseWriter, r *http.Request) error

func WriteJson(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	enc := json.NewEncoder(w)
	return enc.Encode(data)
}
