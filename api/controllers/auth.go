package controllers

import (
	"encoding/json"
	"net/http"
	"simplerestapi/api/models"
	"simplerestapi/api/services"
)

var authService = services.NewAuthService()

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

	token, err := authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Token:   token,
	})
}
