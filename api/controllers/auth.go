package controllers

import (
	"encoding/json"
	"net/http"
	"simplerestapi/api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

	if loginRequest.Username == "admin" && loginRequest.Password == "123" {
		token := "fake-jwt-token-123456"

		json.NewEncoder(w).Encode(models.LoginResponse{
			Status:  "success",
			Message: "Login successful",
			Token:   token,
		})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(models.LoginResponse{
		Status:  "error",
		Message: "Invalid username or password",
	})
}
