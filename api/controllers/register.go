package controllers

import (
	"encoding/json"
	"net/http"
	"simplerestapi/api/models"
	"simplerestapi/api/services"
)

var userService = services.NewUserService()

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var registerRequest models.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.RegisterUserResponse{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

	json.NewEncoder(w).Encode(models.RegisterUserResponse{
		Status:  "success",
		Message: "Register successful",
	})
}
