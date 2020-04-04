package controllers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Error{
		Error: err.Error(),
	})
}
