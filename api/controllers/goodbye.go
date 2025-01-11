package controllers

import (
	"encoding/json"
	"net/http"
)

func Goodbye(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Goodbye World"}
	json.NewEncoder(w).Encode(response)
}
