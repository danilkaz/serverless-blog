package responses

import (
	"encoding/json"
	"net/http"
)

func ResponseData(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

func ResponseInternalError(w http.ResponseWriter, err error) {
	ResponseData(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	ResponseData(w, code, map[string]string{"error": err.Error()})
}
