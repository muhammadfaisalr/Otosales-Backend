package rest

import (
	"Otosales/models"
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, msg string, d interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	response := model.BaseResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: msg,
		Data:    d,
	}

	return json.NewEncoder(w).Encode(response)
}
