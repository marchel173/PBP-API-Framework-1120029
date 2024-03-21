package controllers

import (
	"encoding/json"
	"net/http"

	m "github.com/Martini/models"
)

func PrintError(status int, message string, w http.ResponseWriter) {
	var errResponse m.ErrorResponse
	errResponse.Status = status
	errResponse.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errResponse)
}

func PrintSuccess(status int, message string, w http.ResponseWriter) {
	var sucResponse m.ErrorResponse
	sucResponse.Status = status
	sucResponse.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sucResponse)
}
